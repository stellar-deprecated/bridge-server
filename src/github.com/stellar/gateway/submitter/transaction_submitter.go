package submitter

import (
	"encoding/hex"
	"fmt"
	"strconv"
	"sync"
	"time"

	"encoding/json"

	"github.com/Sirupsen/logrus"
	"github.com/stellar/gateway/db"
	"github.com/stellar/gateway/db/entities"
	"github.com/stellar/go/build"
	"github.com/stellar/go/clients/horizon"
	"github.com/stellar/go/keypair"
	"github.com/stellar/go/network"
	"github.com/stellar/go/support/errors"
	"github.com/stellar/go/xdr"
)

// TransactionSubmitterInterface helps mocking TransactionSubmitter
type TransactionSubmitterInterface interface {
	SubmitTransaction(seed string, operation, memo interface{}) (response horizon.TransactionSuccess, err error)
	SignAndSubmitRawTransaction(seed string, tx *xdr.Transaction) (response horizon.TransactionSuccess, err error)
}

// TransactionSubmitter submits transactions to Stellar Network
type TransactionSubmitter struct {
	Horizon       *horizon.Client
	Accounts      map[string]*Account // seed => *Account
	EntityManager db.EntityManagerInterface
	Network       build.Network
	log           *logrus.Entry
	now           func() time.Time
}

// Account represents account used to signing and sending transactions
type Account struct {
	Keypair        keypair.KP
	Seed           string
	SequenceNumber uint64
	Mutex          sync.Mutex
}

// NewTransactionSubmitter creates a new TransactionSubmitter
func NewTransactionSubmitter(
	horizon *horizon.Client,
	entityManager db.EntityManagerInterface,
	networkPassphrase string,
	now func() time.Time,
) (ts TransactionSubmitter) {
	ts.Horizon = horizon
	ts.EntityManager = entityManager
	ts.Accounts = make(map[string]*Account)
	ts.Network = build.Network{networkPassphrase}
	ts.log = logrus.WithFields(logrus.Fields{
		"service": "TransactionSubmitter",
	})
	ts.now = now
	return
}

// LoadAccount loads currect state of Stellar account
func (ts *TransactionSubmitter) LoadAccount(seed string) (account *Account, err error) {
	account = &Account{}
	account.Keypair, err = keypair.Parse(seed)
	if err != nil {
		ts.log.Print("Invalid seed")
		return
	}

	accountResponse, err := ts.Horizon.LoadAccount(account.Keypair.Address())
	if err != nil {
		return
	}

	account.Seed = seed
	account.SequenceNumber, err = strconv.ParseUint(accountResponse.Sequence, 10, 64)
	return
}

// InitAccount loads an account and returns error if it fails
func (ts *TransactionSubmitter) InitAccount(seed string) (err error) {
	_, err = ts.GetAccount(seed)
	return
}

// GetAccount returns an account by a given seed
func (ts *TransactionSubmitter) GetAccount(seed string) (account *Account, err error) {
	account, exist := ts.Accounts[seed]
	if !exist {
		account, err = ts.LoadAccount(seed)
		ts.Accounts[seed] = account
	}
	return
}

// SignAndSubmitRawTransaction will:
// - update sequence number of the transaction to the current one,
// - sign it,
// - submit it to the network.
func (ts *TransactionSubmitter) SignAndSubmitRawTransaction(seed string, tx *xdr.Transaction) (response Response, err error) {
	account, err := ts.GetAccount(seed)
	if err != nil {
		return
	}

	tx.SeqNum = account.nextSequence()

	hash, err := network.HashTransaction(tx, ts.Network.Passphrase)
	if err != nil {
		err = errors.Wrap(err, "hashing failed")
		return
	}

	sig, err := account.Keypair.SignDecorated(hash[:])
	if err != nil {
		err = errors.Wrap(err, "signing failed")
		return
	}

	envelopeXdr := xdr.TransactionEnvelope{
		Tx:         *tx,
		Signatures: []xdr.DecoratedSignature{sig},
	}

	txeB64, err := xdr.MarshalBase64(envelopeXdr)
	if err != nil {
		err = errors.Wrap(err, "marshal envelope failed")
		return
	}

	sentTransaction := &entities.SentTransaction{
		TransactionID: hex.EncodeToString(hash[:]),
		Status:        entities.SentTransactionStatusSending,
		Source:        account.Keypair.Address(),
		SubmittedAt:   ts.now(),
		EnvelopeXdr:   txeB64,
	}
	err = ts.EntityManager.Persist(sentTransaction)
	if err != nil {
		err = errors.Wrap(err, "initial save tx failed")
		return
	}

	hresp, err := ts.Horizon.SubmitTransaction(txeB64)

	if err != nil {

		// NOTE: on any error, always reload the account state from horizon.  Yes,
		// this causes more loads against horizon than is strictly necessary, but it
		// greatly simplifies to code flow until we can take the time to refactor
		// this whole module.
		//
		// BUG(scott): a failure here will cause the currently processing
		// transaction to have no status recorded.
		rerr := account.resetSequence(ts.Horizon)
		if rerr != nil {
			err = errors.Wrap(err, "failed to reset account sequence after tx failure")
			return
		}

		switch cerr := errors.Cause(err).(type) {
		case *horizon.Error:
			var result string
			xdri, ok := cerr.Problem.Extras["result_xdr"]
			if ok {
				err := json.Unmarshal(xdri, &result)
				if err != nil {
					result = fmt.Sprintf("<error: %s>", err)
				}
			} else {
				result = "<empty>"
			}

			sentTransaction.MarkFailed(result)
		default:
			err = errors.Wrap(err, "submit transaction failed")
			return
		}
	} else {
		sentTransaction.MarkSucceeded(hresp.Ledger)
	}

	err = ts.EntityManager.Persist(sentTransaction)
	if err != nil {
		err = errors.Wrap(err, "save tx status failed")
		return
	}

	err = response.Populate(&hresp)
	if err != nil {
		err = errors.Wrap(err, "populate response failed")
		return
	}

	return
}

// SubmitTransaction builds and submits transaction to Stellar network
func (ts *TransactionSubmitter) SubmitTransaction(
	seed string,
	operation interface{},
	memo interface{},
) (response Response, err error) {
	account, err := ts.GetAccount(seed)
	if err != nil {
		return
	}

	operationMutator, ok := operation.(build.TransactionMutator)
	if !ok {
		ts.log.Error("Cannot cast operationMutator to build.TransactionMutator")
		err = errors.New("Cannot cast operationMutator to build.TransactionMutator")
		return
	}

	mutators := []build.TransactionMutator{
		build.SourceAccount{account.Seed},
		ts.Network,
		operationMutator,
	}

	if memo != nil {
		memoMutator, ok := memo.(build.TransactionMutator)
		if !ok {
			ts.log.Error("Cannot cast memo to build.TransactionMutator")
			err = errors.New("Cannot cast memo to build.TransactionMutator")
			return
		}
		mutators = append(mutators, memoMutator)
	}

	txBuilder := build.Transaction(mutators...)

	return ts.SignAndSubmitRawTransaction(seed, txBuilder.TX)
}

// BuildTransaction is used in compliance server. The sequence number in built transaction will be equal 0!
func BuildTransaction(accountID, networkPassphrase string, operation, memo interface{}) (transaction *xdr.Transaction, err error) {
	operationMutator, ok := operation.(build.TransactionMutator)
	if !ok {
		err = errors.New("Cannot cast operationMutator to build.TransactionMutator")
		return
	}

	mutators := []build.TransactionMutator{
		build.SourceAccount{accountID},
		build.Sequence{0},
		build.Network{networkPassphrase},
		operationMutator,
	}

	if memo != nil {
		memoMutator, ok := memo.(build.TransactionMutator)
		if !ok {
			err = errors.New("Cannot cast memo to build.TransactionMutator")
			return
		}
		mutators = append(mutators, memoMutator)
	}

	txBuilder := build.Transaction(mutators...)

	return txBuilder.TX, txBuilder.Err
}
