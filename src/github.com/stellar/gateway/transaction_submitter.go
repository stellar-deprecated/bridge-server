package gateway

import (
	"errors"
	"log"

	"github.com/stellar/gateway/horizon"
	"github.com/stellar/go-stellar-base/build"
	"github.com/stellar/go-stellar-base/keypair"
)

type TransactionSubmitter struct {
	Horizon  *horizon.Horizon
	Accounts map[string]Account // seed => Account
}

type Account struct {
	Keypair        keypair.KP
	Seed           string
	SequenceNumber uint64
}

func NewTransactionSubmitter(horizon *horizon.Horizon) (ts TransactionSubmitter) {
	ts.Horizon = horizon
	ts.Accounts = make(map[string]Account)
	return
}

func (ts *TransactionSubmitter) LoadAccount(seed string) (account Account, err error) {
	account.Keypair, err = keypair.Parse(seed)
	if err != nil {
		log.Print("Invalid seed")
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

func (ts *TransactionSubmitter) GetAccount(seed string) (account Account, err error) {
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

	mutator, ok := operation.(build.TransactionMutator)
	if !ok {
		log.Print("Cannot cast to build.TransactionMutator")
		err = errors.New("Cannot cast to build.TransactionMutator")
		return
	}

	tx := build.Transaction(
		build.SourceAccount{account.Seed},
		build.Sequence{account.SequenceNumber + 1},
		mutator,
	)

	txe := tx.Sign(seed)
	txeB64, err := txe.Base64()

	if err != nil {
		log.Print("Cannot encode transaction envelope ", err)
		return
	}

	response, err = ts.Horizon.SubmitTransaction(txeB64)
	if err != nil {
		log.Print("Error submitting transaction ", err)
		// Sync sequence number
		// TODO do it only if it's BAD_SEQ error
		accountResponse, _ := ts.Horizon.LoadAccount(account.Keypair.Address())
		account.SequenceNumber = accountResponse.SequenceNumber
	} else {
		account.SequenceNumber++
	}

	return
}
