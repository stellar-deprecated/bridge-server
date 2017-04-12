package db

import (
	"testing"

	"github.com/stellar/go/support/db/dbtest"
	"github.com/stretchr/testify/assert"
)

func TestGetTable(t *testing.T) {
	db := dbtest.Postgres(t).Load(testSchema)
	defer db.Close()
	sess := &Session{DB: db.Open()}
	defer sess.DB.Close()

	tbl := sess.GetTable("users")
	if assert.NotNil(t, tbl) {
		assert.Equal(t, "users", tbl.Name)
		assert.Equal(t, sess, tbl.Session)
	}

}
