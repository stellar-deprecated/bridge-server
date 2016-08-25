// Package db is the base package for database access at stellar.  It primarily
// exposes Repo which is a lightweight wrapper around a *sqlx.DB that provides
// utility methods (See the repo tests for examples).
//
// In addition to the query methods, this package provides query logging and
// stateful transaction management.
package db

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/stellar/go/support/errors"
	"golang.org/x/net/context"

	// Enable mysql
	_ "github.com/go-sql-driver/mysql"
	// Enable postgres
	_ "github.com/lib/pq"
)

// Conn represents a connection to a single database.
type Conn interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Get(dest interface{}, query string, args ...interface{}) error
	Rebind(sql string) string
	Queryx(query string, args ...interface{}) (*sqlx.Rows, error)
	Select(dest interface{}, query string, args ...interface{}) error
}

// Repo provides helper methods for making queries against `DB` and provides
// utilities such as automatic query logging and transaction management.  NOTE:
// A Repo is designed to be lightweight and temporarily lived (usually request
// scoped) which is one reason it is acceptable for it to store a context.
type Repo struct {
	// DB is the database connection that queries should be executed against.
	DB *sqlx.DB

	// Ctx is the optional context in which the repo is operating under.
	Ctx context.Context

	tx *sqlx.Tx
}

// Open the postgres database at `url` and returns a new *Repo using it.
func Open(dialect, url string) (*Repo, error) {
	db, err := sqlx.Connect(dialect, url)
	if err != nil {
		return nil, errors.Wrap(err, "connect failed")
	}

	return &Repo{DB: db}, nil
}

// Wrap wraps a bare *sql.DB (from the database/sql stdlib package) in a
// *db.Repo instance.
func Wrap(base *sql.DB, dialect string) *Repo {
	return &Repo{DB: sqlx.NewDb(base, dialect)}
}

// ensure various types conform to Conn interface
var _ Conn = (*sqlx.Tx)(nil)
var _ Conn = (*sqlx.DB)(nil)
