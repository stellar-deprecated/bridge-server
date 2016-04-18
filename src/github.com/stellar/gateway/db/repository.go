package db

import (
	"github.com/Sirupsen/logrus"
	"github.com/stellar/gateway/db/entities"
)

// RepositoryInterface helps mocking Repository
type RepositoryInterface interface {
	GetLastCursorValue() (cursor *string, err error)
	GetAuthorizedTransactionByMemo(memo string) (*entities.AuthorizedTransaction, error)
	GetAllowedFiByDomain(domain string) (*entities.AllowedFi, error)
	GetAllowedUserByDomainAndUserID(domain, userID string) (*entities.AllowedUser, error)
	GetReceivedPaymentByID(id int64) (*entities.ReceivedPayment, error)
}

// Repository helps getting data from DB
type Repository struct {
	driver Driver
	log    *logrus.Entry
}

// NewRepository creates a new Repository using driver
func NewRepository(driver Driver) (r Repository) {
	r.driver = driver
	r.log = logrus.WithFields(logrus.Fields{
		"service": "Repository",
	})
	return
}

// GetLastCursorValue returns last cursor value from a DB
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

// GetAuthorizedTransactionByMemo returns authorized transaction searching by memo
func (r Repository) GetAuthorizedTransactionByMemo(memo string) (*entities.AuthorizedTransaction, error) {
	authorizedTransaction, err := r.driver.GetOne(&entities.AuthorizedTransaction{}, "memo = ?", memo)
	if authorizedTransaction == nil {
		return nil, err
	}
	return authorizedTransaction.(*entities.AuthorizedTransaction), err
}

// GetAllowedFiByDomain returns allowed FI by a domain
func (r Repository) GetAllowedFiByDomain(domain string) (*entities.AllowedFi, error) {
	allowedFi, err := r.driver.GetOne(&entities.AllowedFi{}, "domain = ?", domain)
	if allowedFi == nil {
		return nil, err
	}
	return allowedFi.(*entities.AllowedFi), err
}

// GetAllowedUserByDomainAndUserID returns allowed user by domain and userID
func (r Repository) GetAllowedUserByDomainAndUserID(domain, userID string) (*entities.AllowedUser, error) {
	allowedUser, err := r.driver.GetOne(&entities.AllowedUser{}, "fi_domain = ? && user_id = ?", domain, userID)
	if allowedUser == nil {
		return nil, err
	}
	return allowedUser.(*entities.AllowedUser), err
}

// GetReceivedPaymentByID returns received payment by id
func (r Repository) GetReceivedPaymentByID(id int64) (*entities.ReceivedPayment, error) {
	receivedPayment, err := r.driver.GetOne(&entities.ReceivedPayment{}, "id = ?", id)
	if receivedPayment == nil {
		return nil, err
	}
	return receivedPayment.(*entities.ReceivedPayment), err
}
