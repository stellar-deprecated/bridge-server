package gateway

import (
	"errors"
	"sync"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/stellar/gateway/db"
	"github.com/stellar/gateway/horizon"
	"github.com/stellar/go-stellar-base/build"
	"github.com/stellar/go-stellar-base/keypair"
)

type TransactionSubmitter struct {
	Horizon       *horizon.Horizon
	Accounts      map[string]*Account // seed => *Account
	EntityManager *db.EntityManager
	log           *logrus.Entry
}

type Account struct {
	Keypair        keypair.KP
	Seed           string
	SequenceNumber uint64
	Mutex          sync.Mutex
}

func NewTransactionSubmitter(horizon *horizon.Horizon, entityManager *db.EntityManager) (ts TransactionSubmitter) {
	ts.Horizon = horizon
	ts.EntityManager = entityManager
	ts.Accounts = make(map[string]*Account)
	ts.log = logrus.WithFields(logrus.Fields{
		"service": "TransactionSubmitter",
	})
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
	account.SequenceNumber = accountResponse.SequenceNumber
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

func (ts *TransactionSubmitter) SubmitTransaction(seed string, operation interface{}) (response horizon.SubmitTransactionResponse, err error) {
	account, err := ts.GetAccount(seed)
	if err != nil {
		return
	}

	var sequenceNumber uint64

	account.Mutex.Lock()
	account.SequenceNumber++
	sequenceNumber = account.SequenceNumber
	account.Mutex.Unlock()

	mutator, ok := operation.(build.TransactionMutator)
	if !ok {
		ts.log.Error("Cannot cast to build.TransactionMutator")
		err = errors.New("Cannot cast to build.TransactionMutator")
		return
	}

	tx := build.Transaction(
		build.SourceAccount{account.Seed},
		build.Sequence{sequenceNumber},
		mutator,
	)

	txe := tx.Sign(seed)
	txeB64, err := txe.Base64()

	if err != nil {
		ts.log.Error("Cannot encode transaction envelope ", err)
		return
	}

	sentTransaction := &db.SentTransaction{
		Status:      "sending",
		Source:      account.Keypair.Address(),
		SubmittedAt: time.Now(),
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
		sentTransaction.MarkFailed(response.Extras.ResultXdr)
	}
	err = ts.EntityManager.Persist(sentTransaction)
	if err != nil {
		return
	}

	// Sync sequence number
	if response.Errors != nil && response.Errors.TransactionErrorCode == "transaction_bad_seq" {
		account.Mutex.Lock()
		ts.log.Print("Syncing sequence number for ", account.Keypair.Address())
		accountResponse, _ := ts.Horizon.LoadAccount(account.Keypair.Address())
		account.SequenceNumber = accountResponse.SequenceNumber
		account.Mutex.Unlock()
	}

	return
}
