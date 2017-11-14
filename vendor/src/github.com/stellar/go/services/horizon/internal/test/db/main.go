// Package db provides helpers to connect to test databases.  It has no
// internal dependencies on horizon and so should be able to be imported by
// any horizon package.
package db

import (
	"testing"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	// pq enables postgres support
	_ "github.com/lib/pq"
	db "github.com/stellar/go/support/db/dbtest"
)

var (
	coreDB		*sqlx.DB
	coreUrl 	*string
	horizonDB 	*sqlx.DB
	horizonUrl 	*string
)

// Horizon returns a connection to the horizon test database
func Horizon(t *testing.T) *sqlx.DB {
	if horizonDB != nil {
		return horizonDB
	}
	postgres := db.Postgres(t)
	horizonUrl = &postgres.DSN
	horizonDB = postgres.Open()
	return horizonDB
}

// HorizonURL returns the database connection the url any test
// use when connecting to the history/horizon database
func HorizonURL() string {
	if horizonUrl == nil  {
		log.Panic(fmt.Errorf("Horizon not initialized"))
	}
	return *horizonUrl
}

// StellarCore returns a connection to the stellar core test database
func StellarCore(t *testing.T) *sqlx.DB {
	if coreDB != nil {
		return coreDB
	}
	postgres := db.Postgres(t)
	coreUrl = &postgres.DSN
	coreDB = postgres.Open()
	return coreDB
}

// StellarCoreURL returns the database connection the url any test
// use when connecting to the stellar-core database
func StellarCoreURL() string {
	if coreUrl == nil  {
		log.Panic(fmt.Errorf("StellarCore not initialized"))
	}
	return *coreUrl
}
