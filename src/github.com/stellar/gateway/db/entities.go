package db

import (
	"time"
)

type Entity interface {
	GetId() *int64
	SetId(int64)
}

type ReceivedPayment struct {
	Id          *int64    `db:"id"`
	OperationId string    `db:"operation_id"`
	ProcessedAt time.Time `db:"processed_at"`
	PagingToken string    `db:"paging_token"`
	Status      string    `db:"status"`
}

type SentTransaction struct {
	Id            *int64     `db:"id"`
	Status        string     `db:"status"` // sending/success/failure
	Source        string     `db:"source"`
	SubmittedAt   time.Time  `db:"submitted_at"`
	SucceededAt   *time.Time `db:"succeeded_at"`
	OperationType string     `db:"operation_type"`
	Ledger        *uint64    `db:"ledger"`
	EnvelopeXdr   string     `db:"envelope_xdr"`
	ResultXdr     *string    `db:"result_xdr"`
}

func (rp *ReceivedPayment) GetId() *int64 {
	return rp.Id
}

func (rp *ReceivedPayment) SetId(id int64) {
	rp.Id = &id
}

func (st *SentTransaction) GetId() *int64 {
	return st.Id
}

func (st *SentTransaction) SetId(id int64) {
	st.Id = &id
}

func (st *SentTransaction) MarkSucceeded(ledger uint64) {
	st.Status = "success"
	st.Ledger = &ledger
	now := time.Now()
	st.SucceededAt = &now
}

func (st *SentTransaction) MarkFailed(resultXdr string) {
	st.Status = "failure"
	st.ResultXdr = &resultXdr
}
