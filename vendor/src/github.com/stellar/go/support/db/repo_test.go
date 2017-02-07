package db

import (
	"testing"

	"github.com/stellar/go/support/db/dbtest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRepo(t *testing.T) {
	db := dbtest.Postgres(t).Load(testSchema)
	defer db.Close()

	assert := assert.New(t)
	require := require.New(t)
	repo := &Repo{DB: db.Open()}
	defer repo.DB.Close()

	assert.Equal("postgres", repo.Dialect())

	var count int
	err := repo.GetRaw(&count, "SELECT COUNT(*) FROM people")
	assert.NoError(err)
	assert.Equal(3, count)

	var names []string
	err = repo.SelectRaw(&names, "SELECT name FROM people")
	assert.NoError(err)
	assert.Len(names, 3)

	ret, err := repo.ExecRaw("DELETE FROM people")
	assert.NoError(err)
	deleted, err := ret.RowsAffected()
	assert.NoError(err)
	assert.Equal(int64(3), deleted)

	// Test args (NOTE: there is a simple escaped arg to ensure no error is raised
	// during execution)
	db.Load(testSchema)
	var name string
	err = repo.GetRaw(
		&name,
		"SELECT name FROM people WHERE hunger_level = ? AND name != '??'",
		1000000,
	)
	assert.NoError(err)
	assert.Equal("scott", name)

	// Test NoRows
	err = repo.GetRaw(
		&name,
		"SELECT name FROM people WHERE hunger_level = ?",
		1234,
	)
	assert.True(repo.NoRows(err))

	// Test transactions
	db.Load(testSchema)
	require.NoError(repo.Begin(), "begin failed")
	err = repo.GetRaw(&count, "SELECT COUNT(*) FROM people")
	assert.NoError(err)
	assert.Equal(3, count)
	_, err = repo.ExecRaw("DELETE FROM people")
	assert.NoError(err)
	err = repo.GetRaw(&count, "SELECT COUNT(*) FROM people")
	assert.NoError(err)
	assert.Equal(0, count, "people did not appear deleted inside transaction")
	assert.NoError(repo.Rollback(), "rollback failed")

	// Ensure commit works
	require.NoError(repo.Begin(), "begin failed")
	repo.ExecRaw("DELETE FROM people")
	assert.NoError(repo.Commit(), "commit failed")
	err = repo.GetRaw(&count, "SELECT COUNT(*) FROM people")
	assert.NoError(err)
	assert.Equal(0, count)

	// ensure that selecting into a populated slice clears the slice first
	db.Load(testSchema)
	require.Len(names, 3, "ids slice was not preloaded with data")
	err = repo.SelectRaw(&names, "SELECT name FROM people limit 2")
	assert.NoError(err)
	assert.Len(names, 2)

	// Test ReplacePlaceholders
	out, err := repo.ReplacePlaceholders("? = ? = ? = ??")
	if assert.NoError(err) {
		assert.Equal("$1 = $2 = $3 = ?", out)
	}
}

const testSchema = `
CREATE TABLE  IF NOT EXISTS people (
    name character varying NOT NULL,
    hunger_level integer NOT NULL
);
DELETE FROM people;
INSERT INTO people (name, hunger_level) VALUES ('scott', 1000000);
INSERT INTO people (name, hunger_level) VALUES ('jed', 10);
INSERT INTO people (name, hunger_level) VALUES ('bartek', 10);
`
