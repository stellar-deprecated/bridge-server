// +build cgo

package dbtest

import (
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/require"
)

func TestSqlite(t *testing.T) {
	tdb := Sqlite(t)
	t.Log(tdb.DSN)

	db, err := sqlx.Open("sqlite3", tdb.DSN)
	require.NoError(t, err)
	_, err = db.Exec("SELECT 1")
	require.NoError(t, err)

	db.Close()
	tdb.Close()
	db, err = sqlx.Open("sqlite", tdb.DSN)
	require.Error(t, err)

	tdb.Close()
}
