package mysql

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/rubenv/sql-migrate"
	"github.com/stellar/gateway/db"
)

type MysqlDriver struct {
	database *sqlx.DB
}

func (d *MysqlDriver) Init(url string) (err error) {
	d.database, err = sqlx.Connect("mysql", url)
	return
}

// go-bindata -ignore .+\.go$ -pkg mysql -o bindata.go ./migrations
func (d *MysqlDriver) MigrateUp() (migrationsApplied int, err error) {
	source := d.getAssetMigrationSource()
	migrationsApplied, err = migrate.Exec(d.database.DB, "mysql", source, migrate.Up)
	return
}

func (d *MysqlDriver) InsertReceivedPayment(object *db.ReceivedPayment) (id int64, err error) {
	query := `
	INSERT INTO ReceivedPayment
		(operation_id, processed_at, paging_token, status)
	VALUES
		(:operation_id, :processed_at, :paging_token, :status)`
	id, err = d.insert(query, object)
	return
}

func (d *MysqlDriver) UpdateReceivedPayment(object *db.ReceivedPayment) (err error) {
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

func (d *MysqlDriver) GetLastReceivedPayment() (*db.ReceivedPayment, error) {
	var receivedPayment db.ReceivedPayment
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

func (d *MysqlDriver) InsertSentTransaction(object *db.SentTransaction) (id int64, err error) {
	query := `
	INSERT INTO SentTransaction
		(status, source, submitted_at, succeeded_at, ledger, envelope_xdr, result_xdr)
	VALUES
		(:status, :source, :submitted_at, :succeeded_at, :ledger, :envelope_xdr, :result_xdr)`
	id, err = d.insert(query, object)
	return
}

func (d *MysqlDriver) UpdateSentTransaction(object *db.SentTransaction) (err error) {
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

func (d *MysqlDriver) insert(query string, object interface{}) (id int64, err error) {
	result, err := d.database.NamedExec(query, object)
	if err != nil {
		return
	}

	id, err = result.LastInsertId()
	return
}

func (d *MysqlDriver) update(query string, object interface{}) (err error) {
	_, err = d.database.NamedExec(query, object)
	return
}

func (d *MysqlDriver) getAssetMigrationSource() (source *migrate.AssetMigrationSource) {
	source = &migrate.AssetMigrationSource{
		Asset:    Asset,
		AssetDir: AssetDir,
		Dir:      "migrations",
	}
	return
}
