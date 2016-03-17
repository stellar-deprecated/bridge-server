package db

import (
	"github.com/Sirupsen/logrus"
)

type RepositoryInterface interface {
	GetLastCursorValue() (cursor *string, err error)
}

type Repository struct {
	driver Driver
	log    *logrus.Entry
}

func NewRepository(driver Driver) (r Repository) {
	r.driver = driver
	r.log = logrus.WithFields(logrus.Fields{
		"service": "Repository",
	})
	return
}

func (r Repository) GetLastCursorValue() (cursor *string, err error) {
	receivedPayment, err := r.driver.GetLastReceivedPayment()
	if err != nil {
		return nil, err
	} else if receivedPayment == nil {
		return nil, nil
	} else {
		return &receivedPayment.PagingToken, nil
	}
}
