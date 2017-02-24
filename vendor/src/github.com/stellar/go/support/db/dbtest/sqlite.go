// +build cgo

package dbtest

import (
	"io/ioutil"
	"os"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/require"
)

// Sqlite provisions a new, blank database sqlite database.  It panics on the
// event of a failure.
func Sqlite(t *testing.T) *DB {
	var result DB

	tmpfile, err := ioutil.TempFile("", "test.sqlite")
	require.NoError(t, err)

	tmpfile.Close()
	err = os.Remove(tmpfile.Name())
	require.NoError(t, err)

	result.Dialect = "sqlite3"
	result.DSN = tmpfile.Name()
	result.t = t

	result.closer = func() {
		err := os.Remove(tmpfile.Name())
		require.NoError(t, err)
	}

	return &result
}
