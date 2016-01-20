package db

import (
	"github.com/Sirupsen/logrus"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

type RepositoryInterface interface {
	GetLastCursorValue() (cursor *string, err error)
}

type Repository struct {
	db  *sqlx.DB
	log *logrus.Entry
}

func NewRepository(dbType string, url string) (r Repository, err error) {
	r.db, err = sqlx.Connect(dbType, url)
	r.log = logrus.WithFields(logrus.Fields{
		"service": "Repository",
	})
	return
}

func (r Repository) GetLastCursorValue() (cursor *string, err error) {
	var receivedPayment ReceivedPayment
	err = r.db.Get(&receivedPayment, "SELECT * FROM `ReceivedPayment` ORDER BY id DESC LIMIT 1")
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, nil
		} else {
			return nil, err
		}
	}
	return &receivedPayment.PagingToken, nil
}
