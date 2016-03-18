package db

import (
	"github.com/Sirupsen/logrus"
	"github.com/stellar/gateway/db/entities"
)

type EntityManagerInterface interface {
	Persist(object entities.Entity) (err error)
}

type EntityManager struct {
	driver Driver
	log    *logrus.Entry
}

func NewEntityManager(driver Driver) (em EntityManager) {
	em.driver = driver
	em.log = logrus.WithFields(logrus.Fields{
		"service": "EntityManager",
	})
	return
}

// Persists an object in DB.
//
// If `object.IsNew()` equals true object will be inserted.
// Otherwise, it will found using `object.GetId()` and updated.
func (em EntityManager) Persist(object entities.Entity) (err error) {
	if object.IsNew() {
		_, err = em.driver.Insert(object)
	} else {
		err = em.driver.Update(object)
	}
	return
}
