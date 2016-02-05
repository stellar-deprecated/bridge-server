package db

import (
	"fmt"

	"github.com/Sirupsen/logrus"
)

type EntityManagerInterface interface {
	Persist(object Entity) (err error)
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

func (em EntityManager) Persist(object Entity) (err error) {
	objectType := fmt.Sprintf("%T", object)

	if object.GetId() != nil {
		// Update
		switch objectType {
		case "*db.ReceivedPayment":
			err = em.driver.UpdateReceivedPayment(object.(*ReceivedPayment))
		case "*db.SentTransaction":
			err = em.driver.UpdateSentTransaction(object.(*SentTransaction))
		default:
			err = fmt.Errorf("Unknown object: %s (must be a pointer)", objectType)
		}
	} else {
		// Insert
		var id int64
		switch objectType {
		case "*db.ReceivedPayment":
			id, err = em.driver.InsertReceivedPayment(object.(*ReceivedPayment))
		case "*db.SentTransaction":
			id, err = em.driver.InsertSentTransaction(object.(*SentTransaction))
		default:
			err = fmt.Errorf("Unknown object: %s (must be a pointer)", objectType)
		}

		if err == nil {
			object.SetId(id)
		}
	}
	return
}
