package db

import (
	"github.com/jmoiron/sqlx"
	"github.com/stellar/gateway/db/entities"
)

// Driver interface allows mocking database driver
type Driver interface {
	Init(url string) (err error)
	DB() *sqlx.DB
	MigrateUp(component string) (migrationsApplied int, err error)

	Insert(object entities.Entity) (id int64, err error)
	Update(object entities.Entity) (err error)
	Delete(object entities.Entity) (err error)

	GetOne(object entities.Entity, where string, params ...interface{}) (entities.Entity, error)
	GetMany(slice interface{}, where, order, limit *string, params ...interface{}) (err error)
}
