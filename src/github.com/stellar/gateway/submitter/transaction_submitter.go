package submitter

import (
	"errors"
	"strconv"
	"sync"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/stellar/gateway/db"
	"github.com/stellar/gateway/db/entities"
	"github.com/stellar/gateway/horizon"
	"github.com/stellar/go-stellar-base/build"
	"github.com/stellar/go-stellar-base/keypair"
)

type TransactionSubmitterInterface interface {
	SubmitTransaction(seed string, operation, memo interface{}) (response horizon.SubmitTransactionResponse, err error)
}

type TransactionSubmitter struct {
	Horizon       horizon.HorizonInterface
	Accounts      map[string]*Account // seed => *Account
	EntityManager db.EntityManagerInterface
	Network       build.Network
	log           *logrus.Entry
	now           func() time.Time
}

type Account struct {
	Keypair        keypair.KP
	Seed           string
	SequenceNumber uint64
	Mutex          sync.Mutex
}

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
	account.SequenceNumber, err = strconv.ParseUint(accountResponse.SequenceNumber, 10, 64)
	return
}

func (ts *TransactionSubmitter) InitAccount(seed string) (err error) {
	_, err = ts.GetAccount(seed)
	return
}

func (ts *TransactionSubmitter) GetAccount(seed string) (account *Account, err error) {
	account, exist := ts.Accounts[seed]
	if !exist {
		account, err = ts.LoadAccount(seed)
		ts.Accounts[seed] = account
	}
	return
}

func (ts *TransactionSubmitter) SubmitTransaction(seed string, operation, memo interface{}) (response horizon.SubmitTransactionResponse, err error) {
	account, err := ts.GetAccount(seed)
	if err != nil {
		return
	}

	var sequenceNumber uint64

	account.Mutex.Lock()
	account.SequenceNumber++
	sequenceNumber = account.SequenceNumber
	account.Mutex.Unlock()

	operationMutator, ok := operation.(build.TransactionMutator)
	if !ok {
		ts.log.Error("Cannot cast operationMutator to build.TransactionMutator")
		err = errors.New("Cannot cast operationMutator to build.TransactionMutator")
		return
	}

	mutators := []build.TransactionMutator{
		build.SourceAccount{account.Seed},
		build.Sequence{sequenceNumber},
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

	tx := build.Transaction(mutators...)

	txe := tx.Sign(seed)
	txeB64, err := txe.Base64()

	if err != nil {
		ts.log.Error("Cannot encode transaction envelope ", err)
		return
	}

	sentTransaction := &entities.SentTransaction{
		Status:      entities.SENT_TRANSACTION_STATUS_SENDING,
		Source:      account.Keypair.Address(),
		SubmittedAt: ts.now(),
		EnvelopeXdr: txeB64,
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
	if response.Error != nil && response.Error.Code == "transaction_bad_seq" {
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
