package dbtest

import (
	"fmt"
	"os/exec"

	_ "github.com/go-sql-driver/mysql"
	"github.com/stellar/go/support/errors"
)

// Mysql provisions a new, blank database with a random name on the localhost of
// the running process.  It assumes that you have mysql running and that the
// root user has access with no password.  It panics on
// the event of a failure.
func Mysql() *DB {
	var result DB
	name := randomName()
	result.Dialect = "mysql"
	result.DSN = fmt.Sprintf("root@/%s", name)

	// create the db
	err := exec.Command("mysql", "-e", fmt.Sprintf("CREATE DATABASE %s;", name)).Run()
	if err != nil {
		err = errors.Wrap(err, "createdb failed")
		panic(err)
	}

	result.closer = func() {
		err := exec.Command("mysql", "-e", fmt.Sprintf("DROP DATABASE %s;", name)).Run()
		if err != nil {
			err = errors.Wrap(err, "dropdb failed")
			panic(err)
		}
	}

	return &result
}
