package postgres

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"

	"github.com/jmoiron/sqlx"
	// To load pq driver
	_ "github.com/lib/pq"
	migrate "github.com/rubenv/sql-migrate"
	"github.com/stellar/gateway/db/entities"
)

//go:generate go-bindata -ignore .+\.go$ -pkg postgres -o bindata.go ./migrations_gateway ./migrations_compliance

// Driver implements Driver interface using Postgres connection
type Driver struct {
	database *sqlx.DB
}

// Init initializes DB connection
func (d *Driver) Init(url string) (err error) {
	d.database, err = sqlx.Connect("postgres", url)
	return
}

func (d *Driver) DB() *sqlx.DB {
	return d.database
}

// MigrateUp migrates DB using migrate files
func (d *Driver) MigrateUp(component string) (migrationsApplied int, err error) {
	source := d.getAssetMigrationSource(component)
	migrationsApplied, err = migrate.Exec(d.database.DB, "postgres", source, migrate.Up)
	return
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
	case *entities.AllowedFi:
		err = stmt.Get(&id, object)
	case *entities.AllowedUser:
		err = stmt.Get(&id, object)
	case *entities.AuthData:
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
	case *entities.AllowedFi:
		_, err = d.database.NamedExec(query, object)
	case *entities.AllowedUser:
		_, err = d.database.NamedExec(query, object)
	case *entities.AuthData:
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
	_, tableName, err := getTypeData(object)

	if err != nil {
		return
	}

	query := "DELETE FROM " + tableName + " WHERE id = :id;"
	_, err = d.database.NamedExec(query, object)

	return
}

// GetOne returns a single entity based on a seach conditions
func (d *Driver) GetOne(object entities.Entity, where string, params ...interface{}) (entities.Entity, error) {
	_, tableName, err := getTypeData(object)
	if err != nil {
		return nil, err
	}

	sql := "SELECT * FROM " + tableName + " WHERE " + where + " LIMIT 1;"
	sql = sqlx.Rebind(sqlx.DOLLAR, sql)

	err = d.database.Get(object, sql, params...)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, nil
		}
		return nil, err
	}
	object.SetExists() // Mark this entity as existing
	return object, err
}

// GetMany returns many entities
func (d *Driver) GetMany(slice interface{}, where, order, offset, limit *string, params ...interface{}) (err error) {
	_, tableName, err := getTypeData(slice)
	if err != nil {
		return
	}

	var query bytes.Buffer

	query.WriteString("SELECT * FROM " + tableName)

	if where != nil {
		query.WriteString(" WHERE " + *where)
	}

	if order != nil {
		query.WriteString(" ORDER BY " + *order)
	}

	if offset != nil {
		query.WriteString(" OFFSET " + *offset)
	}

	if limit != nil {
		query.WriteString(" LIMIT " + *limit)
	}

	query.WriteString(";")

	switch slice := slice.(type) {
	case *[]*entities.ReceivedPayment:
		err = d.database.Select(slice, query.String(), params...)
		tmp := *slice
		for i := range tmp {
			tmp[i].SetExists()
		}
		slice = &tmp
	case *[]*entities.SentTransaction:
		err = d.database.Select(slice, query.String(), params...)
		tmp := *slice
		for i := range tmp {
			tmp[i].SetExists()
		}
		slice = &tmp
	}

	if err != nil && err.Error() == "sql: no rows in result set" {
		return nil
	}

	return
}

func getTypeData(object interface{}) (typeValue reflect.Type, tableName string, err error) {
	switch object := object.(type) {
	case *entities.AuthorizedTransaction:
		typeValue = reflect.TypeOf(*object)
		tableName = "AuthorizedTransaction"
	case *entities.AllowedFi:
		typeValue = reflect.TypeOf(*object)
		tableName = "AllowedFI"
	case *entities.AllowedUser:
		typeValue = reflect.TypeOf(*object)
		tableName = "AllowedUser"
	case *entities.AuthData:
		typeValue = reflect.TypeOf(*object)
		tableName = "AuthData"
	case *entities.SentTransaction:
		typeValue = reflect.TypeOf(*object)
		tableName = "SentTransaction"
	case *entities.ReceivedPayment:
		typeValue = reflect.TypeOf(*object)
		tableName = "ReceivedPayment"
	case *[]*entities.SentTransaction:
		tableName = "SentTransaction"
	case *[]*entities.ReceivedPayment:
		tableName = "ReceivedPayment"
	default:
		return typeValue, tableName, fmt.Errorf("Unknown entity type: %T", object)
	}
	return
}

func (d *Driver) getAssetMigrationSource(component string) (source *migrate.AssetMigrationSource) {
	source = &migrate.AssetMigrationSource{
		Asset:    Asset,
		AssetDir: AssetDir,
		Dir:      "migrations_" + component,
	}
	return
}
