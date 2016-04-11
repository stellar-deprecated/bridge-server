package entities

import (
	"time"
)

type SentTransactionStatus string

const (
	SentTransactionStatusSending SentTransactionStatus = "sending"
	SentTransactionStatusSuccess SentTransactionStatus = "success"
	SentTransactionStatusFailure SentTransactionStatus = "failure"
)

type SentTransaction struct {
	exists        bool
	Id            *int64                `db:"id"`
	TransactionId string                `db:"transaction_id"`
	Status        SentTransactionStatus `db:"status"` // sending/success/failure
	Source        string                `db:"source"`
	SubmittedAt   time.Time             `db:"submitted_at"`
	SucceededAt   *time.Time            `db:"succeeded_at"`
	Ledger        *uint64               `db:"ledger"`
	EnvelopeXdr   string                `db:"envelope_xdr"`
	ResultXdr     *string               `db:"result_xdr"`
}

func (e *SentTransaction) GetId() *int64 {
	if e.Id == nil {
		return nil
	}
	newId := *e.Id
	return &newId
}

func (e *SentTransaction) SetId(id int64) {
	e.Id = &id
}

func (e *SentTransaction) IsNew() bool {
	return !e.exists
}

func (e *SentTransaction) SetExists() {
	e.exists = true
}

func (e *SentTransaction) MarkSucceeded(ledger uint64) {
	e.Status = SentTransactionStatusSuccess
	e.Ledger = &ledger
	now := time.Now()
	e.SucceededAt = &now
}

func (e *SentTransaction) MarkFailed(resultXdr string) {
	e.Status = SentTransactionStatusFailure
	e.ResultXdr = &resultXdr
}
