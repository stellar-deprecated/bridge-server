// +build cgo

package dbtest

import (
	"io/ioutil"
	"os"

	_ "github.com/mattn/go-sqlite3"
	"github.com/pkg/errors"
)

// Sqlite provisions a new, blank database sqlite database.  It panics on the event of a failure.
func Sqlite() *DB {
	var result DB

	tmpfile, err := ioutil.TempFile("", "test.sqlite")
	if err != nil {
		err = errors.Wrap(err, "create temp file failed")
		panic(err)
	}

	tmpfile.Close()
	err = os.Remove(tmpfile.Name())

	if err != nil {
		err = errors.Wrap(err, "remove first temp file failed")
		panic(err)
	}

	result.Dialect = "sqlite3"
	result.DSN = tmpfile.Name()
	result.closer = func() {
		err := os.Remove(tmpfile.Name())
		if err != nil {
			err = errors.Wrap(err, "remove db file failed")
			panic(err)
		}
	}

	return &result
}
