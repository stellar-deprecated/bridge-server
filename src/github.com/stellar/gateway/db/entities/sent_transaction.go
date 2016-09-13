package entities

import (
	"time"
)

// SentTransactionStatus type represents sent transaction status
type SentTransactionStatus string

const (
	// SentTransactionStatusSending is a status indicating that transaction is sending
	SentTransactionStatusSending SentTransactionStatus = "sending"
	// SentTransactionStatusSuccess is a status indicating that transaction has been successfully sent
	SentTransactionStatusSuccess SentTransactionStatus = "success"
	// SentTransactionStatusFailure is a status indicating that there has been an error while sending a transaction
	SentTransactionStatusFailure SentTransactionStatus = "failure"
)

// SentTransaction represents transaction sent by the gateway server
type SentTransaction struct {
	exists        bool
	ID            *int64                `db:"id"`
	TransactionID string                `db:"transaction_id"`
	Status        SentTransactionStatus `db:"status"` // sending/success/failure
	Source        string                `db:"source"`
	SubmittedAt   time.Time             `db:"submitted_at"`
	SucceededAt   *time.Time            `db:"succeeded_at"`
	Ledger        *uint64               `db:"ledger"`
	EnvelopeXdr   string                `db:"envelope_xdr"`
	ResultXdr     *string               `db:"result_xdr"`
}

// GetID returns ID of the entity
func (e *SentTransaction) GetID() *int64 {
	if e.ID == nil {
		return nil
	}
	newID := *e.ID
	return &newID
}

// SetID sets ID of the entity
func (e *SentTransaction) SetID(id int64) {
	e.ID = &id
}

// IsNew returns true if the entity has not been persisted yet
func (e *SentTransaction) IsNew() bool {
	return !e.exists
}

// SetExists sets entity as persisted
func (e *SentTransaction) SetExists() {
	e.exists = true
}

// MarkSucceeded marks transaction as succeeded
func (e *SentTransaction) MarkSucceeded(ledger int32) {
	e.Status = SentTransactionStatusSuccess
	l := uint64(ledger)
	e.Ledger = &l
	now := time.Now()
	e.SucceededAt = &now
}

// MarkFailed marks transaction as failed
func (e *SentTransaction) MarkFailed(resultXdr string) {
	e.Status = SentTransactionStatusFailure
	e.ResultXdr = &resultXdr
}
