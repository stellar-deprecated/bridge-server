package dbtest

import (
	"fmt"
	"os/exec"

	_ "github.com/lib/pq"
	"github.com/stellar/go/support/errors"
)

// Postgres provisions a new, blank database with a random name on the localhost
// of the running process.  It assumes that you have postgres running on the
// default port, have the command line postgres tools installed, and that the
// current user has access to the server.  It panics on the event of a failure.
func Postgres() *DB {
	var result DB
	name := randomName()
	result.Dialect = "postgres"
	result.DSN = fmt.Sprintf("postgres://localhost/%s?sslmode=disable", name)

	// create the db
	err := exec.Command("createdb", name).Run()
	if err != nil {
		err = errors.Wrap(err, "createdb failed")
		panic(err)
	}

	result.closer = func() {
		err := exec.Command("dropdb", name).Run()
		if err != nil {
			err = errors.Wrap(err, "dropdb failed")
			panic(err)
		}
	}

	return &result
}
