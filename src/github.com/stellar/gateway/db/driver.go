package db

import (
	"github.com/stellar/gateway/db/entities"
)

type Driver interface {
	Init(url string) (err error)
	MigrateUp() (migrationsApplied int, err error)

	Insert(object entities.Entity) (id int64, err error)
	Update(object entities.Entity) (err error)

	GetLastReceivedPayment() (*entities.ReceivedPayment, error)
}
