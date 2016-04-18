package entities

import (
	"time"
)

// AuthorizedTransaction represents authorized transaction
type AuthorizedTransaction struct {
	exists         bool
	ID             *int64    `db:"id"`
	TransactionID  string    `db:"transaction_id"`
	Memo           string    `db:"memo"`
	TransactionXdr string    `db:"transaction_xdr"`
	AuthorizedAt   time.Time `db:"authorized_at"`
	Data           string    `db:"data"`
}

// GetID returns ID of the entity
func (e *AuthorizedTransaction) GetID() *int64 {
	return e.ID
}

// SetID sets ID of the entity
func (e *AuthorizedTransaction) SetID(id int64) {
	e.ID = &id
}

// IsNew returns true if the entity has not been persisted yet
func (e *AuthorizedTransaction) IsNew() bool {
	return !e.exists
}

// SetExists sets entity as persisted
func (e *AuthorizedTransaction) SetExists() {
	e.exists = true
}
