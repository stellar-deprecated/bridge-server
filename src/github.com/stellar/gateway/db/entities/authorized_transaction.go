package entities

import (
	"time"
)

type AuthorizedTransaction struct {
	exists         bool
	Id             *int64    `db:"id"`
	TransactionId  string    `db:"transaction_id"`
	Memo           string    `db:"memo"`
	TransactionXdr string    `db:"transaction_xdr"`
	AuthorizedAt   time.Time `db:"authorized_at"`
	Data           string    `db:"data"`
}

func (e *AuthorizedTransaction) GetId() *int64 {
	return e.Id
}

func (e *AuthorizedTransaction) SetId(id int64) {
	e.Id = &id
}

func (e *AuthorizedTransaction) IsNew() bool {
	return !e.exists
}

func (e *AuthorizedTransaction) SetExists() {
	e.exists = true
}
