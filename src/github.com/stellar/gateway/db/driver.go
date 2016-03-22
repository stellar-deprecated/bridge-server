package db

import (
	"github.com/stellar/gateway/db/entities"
)

type Driver interface {
	Init(url string) (err error)
	MigrateUp(component string) (migrationsApplied int, err error)

	Insert(object entities.Entity) (id int64, err error)
	Update(object entities.Entity) (err error)
	Delete(object entities.Entity) (err error)

	GetOne(object entities.Entity, where string, params ...interface{}) (entities.Entity, error)
	GetLastReceivedPayment() (*entities.ReceivedPayment, error)
}
