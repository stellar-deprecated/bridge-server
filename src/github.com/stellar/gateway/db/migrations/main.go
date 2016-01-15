package migrations

import (
	"github.com/jmoiron/sqlx"
	"github.com/rubenv/sql-migrate"
	"github.com/Sirupsen/logrus"
)

// go-bindata -ignore .+\.go$ -pkg migrations -o bindata.go .

type MigrationManager struct {
	db     *sqlx.DB
	dbType string
	log    *logrus.Entry
}

func NewMigrationManager(dbType string, url string) (m MigrationManager, err error) {
	var params string
	switch dbType {
	case "mysql":
		params = "?parseTime=true"
	}

	url = url+params
	m.db, err = sqlx.Connect(dbType, url)
	m.dbType = dbType
	m.log = logrus.WithFields(logrus.Fields{
		"service": "MigrationManager",
	})
	return
}

func (m MigrationManager) MigrateUp() {
	source := getAssetMigrationSource()
	n, err := migrate.Exec(m.db.DB, m.dbType, source, migrate.Up)
	if err != nil {
		m.log.Print("Error migrating: ", err);
	}
	m.log.Printf("Applied %d migrations!", n)
}

func getAssetMigrationSource() (source *migrate.AssetMigrationSource) {
	source = &migrate.AssetMigrationSource{
		Asset:    Asset,
		AssetDir: AssetDir,
		Dir:      "",
	}
	return
}
