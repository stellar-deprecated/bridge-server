package db

import (
	"fmt"

	"github.com/Sirupsen/logrus"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

type EntityManagerInterface interface {
	Persist(object Entity) (err error)
}

type EntityManager struct {
	db  *sqlx.DB
	log *logrus.Entry
}

func NewEntityManager(dbType string, url string) (em EntityManager, err error) {
	em.db, err = sqlx.Connect(dbType, url)
	em.log = logrus.WithFields(logrus.Fields{
		"service": "EntityManager",
	})
	return
}

func (em *EntityManager) Persist(object Entity) (err error) {
	objectType := fmt.Sprintf("%T", object)
	var query string

	if object.GetId() != nil {
		// Update
		query, err = GetUpdateQuery(objectType)
	} else {
		// Insert
		query, err = GetInsertQuery(objectType)
	}

	if err != nil {
		return
	}

	result, err := em.db.NamedExec(query, object)
	if err != nil {
		return
	}

	if object.GetId() == nil {
		var id int64
		// Just inserted a new object - set it's ID
		id, err = result.LastInsertId()
		if err != nil {
			return
		}
		object.SetId(id)
	}
	return
}
