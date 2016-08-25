// +build cgo

package db

import (
	_ "github.com/mattn/go-sqlite3"
)

// This file simply includes the sqlite3 driver when in a cgo environment,
// enabling it for use when using the db package
var _ = 0
