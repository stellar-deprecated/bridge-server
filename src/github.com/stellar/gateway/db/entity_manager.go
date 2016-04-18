package db

import (
	"github.com/Sirupsen/logrus"
	"github.com/stellar/gateway/db/entities"
)

// EntityManagerInterface allows mocking EntityManager
type EntityManagerInterface interface {
	Delete(object entities.Entity) (err error)
	Persist(object entities.Entity) error
}

// EntityManager is responsible for persisting object to DB
type EntityManager struct {
	driver Driver
	log    *logrus.Entry
}

// NewEntityManager creates a new EntityManager using driver
func NewEntityManager(driver Driver) (em EntityManager) {
	em.driver = driver
	em.log = logrus.WithFields(logrus.Fields{
		"service": "EntityManager",
	})
	return
}

// Delete an object from DB.
func (em EntityManager) Delete(object entities.Entity) error {
	return em.driver.Delete(object)
}

// Persist persists an object in DB.
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
