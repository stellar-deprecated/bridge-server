package submitter

import (
	"bytes"
	"encoding/hex"
	"errors"
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/stellar/gateway/db"
	"github.com/stellar/gateway/db/entities"
	"github.com/stellar/gateway/horizon"
	"github.com/stellar/go/build"
	"github.com/stellar/go/hash"
	"github.com/stellar/go/keypair"
	"github.com/stellar/go/xdr"
)

// TransactionSubmitterInterface helps mocking TransactionSubmitter
type TransactionSubmitterInterface interface {
	SubmitTransaction(paymentID *string, seed string, operation, memo interface{}) (response horizon.SubmitTransactionResponse, err error)
	SignAndSubmitRawTransaction(paymentID *string, seed string, tx *xdr.Transaction) (response horizon.SubmitTransactionResponse, err error)
}

// TransactionSubmitter submits transactions to Stellar Network
type TransactionSubmitter struct {
	Horizon       horizon.HorizonInterface
	Accounts      map[string]*Account // seed => *Account
	AccountsMutex sync.Mutex
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
	horizon horizon.HorizonInterface,
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

// LoadAccount loads current state of Stellar account and creates a map entry if it didn't exist
func (ts *TransactionSubmitter) LoadAccount(seed string) (*Account, error) {
	ts.AccountsMutex.Lock()

	account, exist := ts.Accounts[seed]
	if exist {
		ts.AccountsMutex.Unlock()
		return account, nil
	}

	kp, err := keypair.Parse(seed)
	if err != nil {
		ts.log.Print("Invalid seed")
		ts.AccountsMutex.Unlock()
		return nil, err
	}

	ts.Accounts[seed] = &Account{
		Seed:    seed,
		Keypair: kp,
	}
	ts.AccountsMutex.Unlock()

	ts.Accounts[seed].Mutex.Lock()
	defer ts.Accounts[seed].Mutex.Unlock()

	accountResponse, err := ts.Horizon.LoadAccount(ts.Accounts[seed].Keypair.Address())
	if err != nil {
		return nil, err
	}

	ts.Accounts[seed].SequenceNumber, err = strconv.ParseUint(accountResponse.SequenceNumber, 10, 64)
	if err != nil {
		return nil, err
	}

	return ts.Accounts[seed], nil
}

// InitAccount loads an account and returns error if it fails
func (ts *TransactionSubmitter) InitAccount(seed string) (err error) {
	_, err = ts.LoadAccount(seed)
	return
}

// SignAndSubmitRawTransaction will:
// - update sequence number of the transaction to the current one,
// - sign it,
// - submit it to the network.
func (ts *TransactionSubmitter) SignAndSubmitRawTransaction(paymentID *string, seed string, tx *xdr.Transaction) (response horizon.SubmitTransactionResponse, err error) {
	account, err := ts.LoadAccount(seed)
	if err != nil {
		return
	}

	account.Mutex.Lock()
	account.SequenceNumber++
	tx.SeqNum = xdr.SequenceNumber(account.SequenceNumber)
	account.Mutex.Unlock()

	hash, err := TransactionHash(tx, ts.Network.Passphrase)
	if err != nil {
		ts.log.Print("Error calculating transaction hash")
		return
	}

	sig, err := account.Keypair.SignDecorated(hash[:])
	if err != nil {
		ts.log.Print("Error signing a transaction")
		return
	}

	envelopeXdr := xdr.TransactionEnvelope{
		Tx:         *tx,
		Signatures: []xdr.DecoratedSignature{sig},
	}

	txeB64, err := xdr.MarshalBase64(envelopeXdr)
	if err != nil {
		ts.log.Print("Cannot encode transaction envelope")
		return
	}

	transactionHashBytes, err := TransactionHash(tx, ts.Network.Passphrase)
	if err != nil {
		ts.log.WithFields(logrus.Fields{"err": err}).Warn("Error calculating tx hash")
		return
	}

	sentTransaction := &entities.SentTransaction{
		PaymentID:     paymentID,
		TransactionID: hex.EncodeToString(transactionHashBytes[:]),
		Status:        entities.SentTransactionStatusSending,
		Source:        account.Keypair.Address(),
		SubmittedAt:   ts.now(),
		EnvelopeXdr:   txeB64,
	}
	err = ts.EntityManager.Persist(sentTransaction)
	if err != nil {
		return
	}

	response, err = ts.Horizon.SubmitTransaction(txeB64)
	if err != nil {
		ts.log.Error("Error submitting transaction ", err)
		return
	}

	if response.Ledger != nil {
		sentTransaction.MarkSucceeded(*response.Ledger)
	} else {
		var result string
		if response.Extras != nil {
			result = response.Extras.ResultXdr
		} else {
			result = "<empty>"
		}
		sentTransaction.MarkFailed(result)
	}
	err = ts.EntityManager.Persist(sentTransaction)
	if err != nil {
		return
	}

	// Sync sequence number
	if response.Extras != nil && response.Extras.ResultXdr == "AAAAAAAAAAD////7AAAAAA==" {
		account.Mutex.Lock()
		ts.log.Print("Syncing sequence number for ", account.Keypair.Address())
		accountResponse, err2 := ts.Horizon.LoadAccount(account.Keypair.Address())
		if err2 != nil {
			ts.log.Error("Error updating sequence number ", err)
		} else {
			account.SequenceNumber, _ = strconv.ParseUint(accountResponse.SequenceNumber, 10, 64)
		}
		account.Mutex.Unlock()
	}
	return
}

// SubmitTransaction builds and submits transaction to Stellar network
func (ts *TransactionSubmitter) SubmitTransaction(paymentID *string, seed string, operation, memo interface{}) (response horizon.SubmitTransactionResponse, err error) {
	account, err := ts.LoadAccount(seed)
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

	if txBuilder.Err != nil {
		err = txBuilder.Err
		return
	}

	return ts.SignAndSubmitRawTransaction(paymentID, seed, txBuilder.TX)
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

// TransactionHash returns transaction hash for a given Transaction based on the network
func TransactionHash(tx *xdr.Transaction, networkPassphrase string) ([32]byte, error) {
	var txBytes bytes.Buffer

	_, err := fmt.Fprintf(&txBytes, "%s", hash.Hash([]byte(networkPassphrase)))
	if err != nil {
		return [32]byte{}, err
	}

	_, err = xdr.Marshal(&txBytes, xdr.EnvelopeTypeEnvelopeTypeTx)
	if err != nil {
		return [32]byte{}, err
	}

	_, err = xdr.Marshal(&txBytes, tx)
	if err != nil {
		return [32]byte{}, err
	}

	return hash.Hash(txBytes.Bytes()), nil
}
