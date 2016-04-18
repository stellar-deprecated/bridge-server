package entities

import (
	"time"
)

// ReceivedPayment represents payment received by the gateway server
type ReceivedPayment struct {
	exists      bool
	ID          *int64    `db:"id"`
	OperationID string    `db:"operation_id"`
	ProcessedAt time.Time `db:"processed_at"`
	PagingToken string    `db:"paging_token"`
	Status      string    `db:"status"`
}

// GetID returns ID of the entity
func (e *ReceivedPayment) GetID() *int64 {
	if e.ID == nil {
		return nil
	}
	newID := *e.ID
	return &newID
}

// SetID sets ID of the entity
func (e *ReceivedPayment) SetID(id int64) {
	e.ID = &id
}

// IsNew returns true if the entity has not been persisted yet
func (e *ReceivedPayment) IsNew() bool {
	return !e.exists
}

// SetExists sets entity as persisted
func (e *ReceivedPayment) SetExists() {
	e.exists = true
}
