package entities

import (
	"time"
)

type ReceivedPayment struct {
	exists      bool
	Id          *int64    `db:"id"`
	OperationId string    `db:"operation_id"`
	ProcessedAt time.Time `db:"processed_at"`
	PagingToken string    `db:"paging_token"`
	Status      string    `db:"status"`
}

func (e *ReceivedPayment) GetId() *int64 {
	if e.Id == nil {
		return nil
	}
	newId := *e.Id
	return &newId
}

func (e *ReceivedPayment) SetId(id int64) {
	e.Id = &id
}

func (e *ReceivedPayment) IsNew() bool {
	return !e.exists
}

func (e *ReceivedPayment) SetExists() {
	e.exists = true
}
