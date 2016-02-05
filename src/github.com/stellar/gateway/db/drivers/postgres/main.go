package postgres

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/rubenv/sql-migrate"
	"github.com/stellar/gateway/db"
)

type PostgresDriver struct {
	database *sqlx.DB
}

func (d *PostgresDriver) Init(url string) (err error) {
	d.database, err = sqlx.Connect("postgres", url)
	return
}

// go-bindata -ignore .+\.go$ -pkg postgres -o bindata.go ./migrations
func (d *PostgresDriver) MigrateUp() (migrationsApplied int, err error) {
	source := d.getAssetMigrationSource()
	migrationsApplied, err = migrate.Exec(d.database.DB, "postgres", source, migrate.Up)
	return
}

func (d *PostgresDriver) InsertReceivedPayment(object *db.ReceivedPayment) (id int64, err error) {
	query := `
	INSERT INTO ReceivedPayment
		(operation_id, processed_at, paging_token, status)
	VALUES
		(:operation_id, :processed_at, :paging_token, :status)
	RETURNING id`
	id, err = d.insert(query, object)
	return
}

func (d *PostgresDriver) UpdateReceivedPayment(object *db.ReceivedPayment) (err error) {
	query := `
	UPDATE ReceivedPayment SET
		operation_id = :operation_id,
		processed_at = :processed_at,
		paging_token = :paging_token,
		status = :status
	WHERE
		id = :id`
	err = d.update(query, object)
	return
}

func (d *PostgresDriver) GetLastReceivedPayment() (*db.ReceivedPayment, error) {
	var receivedPayment db.ReceivedPayment
	err := d.database.Get(&receivedPayment, "SELECT * FROM ReceivedPayment ORDER BY id DESC LIMIT 1")
	if err != nil {
		return nil, err
	}
	return &receivedPayment, nil
}

func (d *PostgresDriver) InsertSentTransaction(object *db.SentTransaction) (id int64, err error) {
	query := `
	INSERT INTO SentTransaction
		(status, source, submitted_at, succeeded_at, ledger, envelope_xdr, result_xdr)
	VALUES
		(:status, :source, :submitted_at, :succeeded_at, :ledger, :envelope_xdr, :result_xdr)
	RETURNING id`
	id, err = d.insert(query, object)
	return
}

func (d *PostgresDriver) UpdateSentTransaction(object *db.SentTransaction) (err error) {
	query := `
	UPDATE SentTransaction SET
		status = :status,
		source = :source,
		submitted_at = :submitted_at,
		succeeded_at = :succeeded_at,
		ledger = :ledger,
		envelope_xdr = :envelope_xdr,
		result_xdr = :result_xdr
	WHERE
		id = :id`
	err = d.update(query, object)
	return
}

func (d *PostgresDriver) insert(query string, object interface{}) (id int64, err error) {
	// TODO cache prepared statement
	stmt, err := d.database.PrepareNamed(query)
	if err != nil {
		return
	}

	err = stmt.Get(&id, object)
	if err != nil {
		return
	}
	return
}

func (d *PostgresDriver) update(query string, object interface{}) (err error) {
	_, err = d.database.NamedExec(query, object)
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
