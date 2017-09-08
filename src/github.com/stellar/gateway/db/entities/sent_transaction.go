package entities

import (
	"database/sql/driver"
	"errors"
	"time"
)

// SentTransactionStatus type represents sent transaction status
type SentTransactionStatus string

// Scan implements database/sql.Scanner interface
func (s *SentTransactionStatus) Scan(src interface{}) error {
	value, ok := src.([]byte)
	if !ok {
		return errors.New("Cannot convert value to SentTransactionStatus")
	}
	*s = SentTransactionStatus(value)
	return nil
}

// Value implements driver.Valuer
func (status SentTransactionStatus) Value() (driver.Value, error) {
	return driver.Value(string(status)), nil
}

var _ driver.Valuer = SentTransactionStatus("")

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
	ID            *int64                `db:"id" json:"id"`
	TransactionID string                `db:"transaction_id" json:"transaction_id"`
	Status        SentTransactionStatus `db:"status" json:"status"` // sending/success/failure
	Source        string                `db:"source" json:"source"`
	SubmittedAt   time.Time             `db:"submitted_at" json:"submitted_at"`
	SucceededAt   *time.Time            `db:"succeeded_at" json:"succeeded_at"`
	Ledger        *uint64               `db:"ledger" json:"ledger"`
	EnvelopeXdr   string                `db:"envelope_xdr" json:"envelope_xdr"`
	ResultXdr     *string               `db:"result_xdr" json:"result_xdr"`
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
func (e *SentTransaction) MarkSucceeded(ledger uint64) {
	e.Status = SentTransactionStatusSuccess
	e.Ledger = &ledger
	now := time.Now()
	e.SucceededAt = &now
}

// MarkFailed marks transaction as failed
func (e *SentTransaction) MarkFailed(resultXdr string) {
	e.Status = SentTransactionStatusFailure
	e.ResultXdr = &resultXdr
}
