package postgres

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/jmoiron/sqlx"
	// To load pq driver
	_ "github.com/lib/pq"
	"github.com/rubenv/sql-migrate"
	"github.com/stellar/gateway/db/entities"
)

// Driver implements Driver interface using Postgres connection
type Driver struct {
	database *sqlx.DB
}

// Init initializes DB connection
func (d *Driver) Init(url string) (err error) {
	d.database, err = sqlx.Connect("postgres", url)
	return
}

// MigrateUp migrates DB using migrate files
// go-bindata -ignore .+\.go$ -pkg postgres -o bindata.go ./migrations
func (d *Driver) MigrateUp(component string) (migrationsApplied int, err error) {
	source := d.getAssetMigrationSource()
	migrationsApplied, err = migrate.Exec(d.database.DB, "postgres", source, migrate.Up)
	return
}

// GetAuthorizedTransactionByMemo returns authorized transaction by memo
func (d *Driver) GetAuthorizedTransactionByMemo(memo string) (*entities.AuthorizedTransaction, error) {
	var authorizedTransaction entities.AuthorizedTransaction
	err := d.database.Get(&authorizedTransaction, "SELECT * FROM AuthorizedTransaction WHERE memo = :memo;", memo)
	if err != nil {
		return nil, err
	}
	return &authorizedTransaction, nil
}

// GetLastReceivedPayment returns the last received payment
func (d *Driver) GetLastReceivedPayment() (*entities.ReceivedPayment, error) {
	var receivedPayment entities.ReceivedPayment
	err := d.database.Get(&receivedPayment, "SELECT * FROM ReceivedPayment ORDER BY id DESC LIMIT 1")
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, nil
		}
		return nil, err
	}
	return &receivedPayment, nil
}

// Insert inserts the entity to a DB
func (d *Driver) Insert(object entities.Entity) (id int64, err error) {
	value, tableName, err := getTypeData(object)

	if err != nil {
		return 0, err
	}

	fieldsCount := value.NumField()
	var fieldNames []string
	var fieldValues []string

	for i := 0; i < fieldsCount; i++ {
		field := value.Field(i)
		tag := field.Tag.Get("db")
		if tag == "" {
			continue
		}

		if tag == "id" && object.GetID() == nil {
			// To handle error:
			// null value in column "id" violates not-null constraint
			continue
		}

		fieldNames = append(fieldNames, tag)
		fieldValues = append(fieldValues, ":"+tag)
	}

	query := "INSERT INTO " + tableName + " (" + strings.Join(fieldNames, ", ") + ") VALUES (" + strings.Join(fieldValues, ", ") + ") RETURNING id;"

	// TODO cache prepared statement
	stmt, err := d.database.PrepareNamed(query)
	if err != nil {
		return
	}

	switch object := object.(type) {
	case *entities.AuthorizedTransaction:
		err = stmt.Get(&id, object)
	case *entities.SentTransaction:
		err = stmt.Get(&id, object)
	case *entities.ReceivedPayment:
		err = stmt.Get(&id, object)
	}

	if err != nil {
		return
	}

	if id == 0 {
		// Not autoincrement
		if object.GetID() == nil {
			return 0, fmt.Errorf("Not autoincrement but ID nil")
		}
		id = *object.GetID()
	}

	if err == nil {
		object.SetID(id)
		object.SetExists()
	}

	return
}

// Update updates the entity to a DB
func (d *Driver) Update(object entities.Entity) (err error) {
	value, tableName, err := getTypeData(object)

	if err != nil {
		return err
	}

	fieldsCount := value.NumField()

	query := "UPDATE " + tableName + " SET "
	var fields []string

	for i := 0; i < fieldsCount; i++ {
		field := value.Field(i)
		if field.Tag.Get("db") == "id" || field.Tag.Get("db") == "" {
			continue
		}
		fields = append(fields, field.Tag.Get("db")+" = :"+field.Tag.Get("db"))
	}

	query += strings.Join(fields, ", ") + " WHERE id = :id;"

	switch object := object.(type) {
	case *entities.AuthorizedTransaction:
		_, err = d.database.NamedExec(query, object)
	case *entities.SentTransaction:
		_, err = d.database.NamedExec(query, object)
	case *entities.ReceivedPayment:
		_, err = d.database.NamedExec(query, object)
	}

	return
}

// Delete delets the entity from a DB
func (d *Driver) Delete(object entities.Entity) (err error) {
	err = errors.New("Not implemented yet")
	return
}

// GetOne returns a single entity based on a seach conditions
func (d *Driver) GetOne(object entities.Entity, where string, params ...interface{}) (entities.Entity, error) {
	return nil, errors.New("Not implemented yet")
}

func getTypeData(object interface{}) (typeValue reflect.Type, tableName string, err error) {
	switch object := object.(type) {
	case *entities.AuthorizedTransaction:
		typeValue = reflect.TypeOf(*object)
		tableName = "AuthorizedTransaction"
	case *entities.SentTransaction:
		typeValue = reflect.TypeOf(*object)
		tableName = "SentTransaction"
	case *entities.ReceivedPayment:
		typeValue = reflect.TypeOf(*object)
		tableName = "ReceivedPayment"
	default:
		return typeValue, tableName, fmt.Errorf("Unknown entity type: %T", object)
	}
	return
}

func (d *Driver) getAssetMigrationSource() (source *migrate.AssetMigrationSource) {
	source = &migrate.AssetMigrationSource{
		Asset:    Asset,
		AssetDir: AssetDir,
		Dir:      "migrations",
	}
	return
}
