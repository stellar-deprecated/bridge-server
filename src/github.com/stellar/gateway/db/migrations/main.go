package migrations

import (
	"github.com/Sirupsen/logrus"
	"github.com/jmoiron/sqlx"
	"github.com/rubenv/sql-migrate"
)

// go-bindata -ignore .+\.go$ -pkg migrations -o bindata.go ./mysql ./postgres

type MigrationManager struct {
	db     *sqlx.DB
	dbType string
	log    *logrus.Entry
}

func NewMigrationManager(dbType string, dbUrl string) (m MigrationManager, err error) {
	m.db, err = sqlx.Connect(dbType, dbUrl)
	m.dbType = dbType
	m.log = logrus.WithFields(logrus.Fields{
		"service": "MigrationManager",
	})
	return
}

func (m MigrationManager) MigrateUp() {
	source := m.getAssetMigrationSource()
	n, err := migrate.Exec(m.db.DB, m.dbType, source, migrate.Up)
	if err != nil {
		m.log.Print("Error migrating: ", err)
	}
	m.log.Printf("Applied %d migrations!", n)
}

func (m MigrationManager) getAssetMigrationSource() (source *migrate.AssetMigrationSource) {
	source = &migrate.AssetMigrationSource{
		Asset:    Asset,
		AssetDir: AssetDir,
		Dir:      m.dbType,
	}
	return
}
