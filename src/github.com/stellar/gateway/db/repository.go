package db

import (
	"github.com/Sirupsen/logrus"
	"github.com/stellar/gateway/db/entities"
)

type RepositoryInterface interface {
	GetLastCursorValue() (cursor *string, err error)
	GetAuthorizedTransactionByMemo(memo string) (*entities.AuthorizedTransaction, error)
	GetAllowedFiByDomain(domain string) (*entities.AllowedFi, error)
	GetAllowedUserByDomainAndUserId(domain, userId string) (*entities.AllowedUser, error)
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

func (r Repository) GetAuthorizedTransactionByMemo(memo string) (*entities.AuthorizedTransaction, error) {
	authorizedTransaction, err := r.driver.GetOne(&entities.AuthorizedTransaction{}, "memo = ?", memo)
	if authorizedTransaction == nil {
		return nil, err
	} else {
		return authorizedTransaction.(*entities.AuthorizedTransaction), err
	}
}

func (r Repository) GetAllowedFiByDomain(domain string) (*entities.AllowedFi, error) {
	allowedFi, err := r.driver.GetOne(&entities.AllowedFi{}, "domain = ?", domain)
	if allowedFi == nil {
		return nil, err
	} else {
		return allowedFi.(*entities.AllowedFi), err
	}
}

func (r Repository) GetAllowedUserByDomainAndUserId(domain, userId string) (*entities.AllowedUser, error) {
	allowedUser, err := r.driver.GetOne(&entities.AllowedUser{}, "fi_domain = ? && user_id = ?", domain, userId)
	if allowedUser == nil {
		return nil, err
	} else {
		return allowedUser.(*entities.AllowedUser), err
	}
}
