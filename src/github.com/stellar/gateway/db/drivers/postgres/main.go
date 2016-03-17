package postgres

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/rubenv/sql-migrate"
	"github.com/stellar/gateway/db/entities"
)

type PostgresDriver struct {
	database *sqlx.DB
}

func (d *PostgresDriver) Init(url string) (err error) {
	d.database, err = sqlx.Connect("postgres", url)
	return
}

// go-bindata -ignore .+\.go$ -pkg postgres -o bindata.go ./migrations
func (d *PostgresDriver) MigrateUp(component string) (migrationsApplied int, err error) {
	source := d.getAssetMigrationSource()
	migrationsApplied, err = migrate.Exec(d.database.DB, "postgres", source, migrate.Up)
	return
}

func (d *PostgresDriver) GetAuthorizedTransactionByMemo(memo string) (*entities.AuthorizedTransaction, error) {
	var authorizedTransaction entities.AuthorizedTransaction
	err := d.database.Get(&authorizedTransaction, "SELECT * FROM AuthorizedTransaction WHERE memo = :memo;", memo)
	if err != nil {
		return nil, err
	}
	return &authorizedTransaction, nil
}

func (d *PostgresDriver) GetLastReceivedPayment() (*entities.ReceivedPayment, error) {
	var receivedPayment entities.ReceivedPayment
	err := d.database.Get(&receivedPayment, "SELECT * FROM ReceivedPayment ORDER BY id DESC LIMIT 1")
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, nil
		} else {
			return nil, err
		}
	}
	return &receivedPayment, nil
}

func (d *PostgresDriver) Insert(object entities.Entity) (id int64, err error) {
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

		if tag == "id" && object.GetId() == nil {
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
		if object.GetId() == nil {
			return 0, fmt.Errorf("Not autoincrement but ID nil")
		}
		id = *object.GetId()
	}

	if err == nil {
		object.SetId(id)
		object.SetExists()
	}

	return
}

func (d *PostgresDriver) Update(object entities.Entity) (err error) {
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

func (d *PostgresDriver) getAssetMigrationSource() (source *migrate.AssetMigrationSource) {
	source = &migrate.AssetMigrationSource{
		Asset:    Asset,
		AssetDir: AssetDir,
		Dir:      "migrations",
	}
	return
}
