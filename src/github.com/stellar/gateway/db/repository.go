package db

import (
	"github.com/Sirupsen/logrus"
	"github.com/stellar/gateway/db/entities"
	"github.com/stellar/go/support/db"
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
	repo *db.Repo
	log  *logrus.Entry
}

// NewRepository creates a new Repository using driver
func NewRepository(driver Driver) (r Repository) {
	r.repo = &db.Repo{DB: driver.DB()}
	r.log = logrus.WithFields(logrus.Fields{
		"service": "Repository",
	})
	return
}

// GetLastCursorValue returns last cursor value from a DB
func (r Repository) GetLastCursorValue() (cursor *string, err error) {
	receivedPayment, err := r.getLastReceivedPayment()
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

	var found entities.AuthorizedTransaction

	err := r.repo.GetRaw(
		&found,
		"SELECT * FROM AuthorizedTransaction WHERE memo = ?",
		memo,
	)

	if r.repo.NoRows(err) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &found, nil
}

// GetAllowedFiByDomain returns allowed FI by a domain
func (r Repository) GetAllowedFiByDomain(domain string) (*entities.AllowedFi, error) {

	var found entities.AllowedFi

	err := r.repo.GetRaw(
		&found,
		"SELECT * FROM AllowedFI WHERE domain = ?",
		domain,
	)

	if r.repo.NoRows(err) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &found, nil
}

// GetAllowedUserByDomainAndUserID returns allowed user by domain and userID
func (r Repository) GetAllowedUserByDomainAndUserID(domain, userID string) (*entities.AllowedUser, error) {

	var found entities.AllowedUser

	err := r.repo.GetRaw(
		&found,
		"SELECT * FROM AllowedUser WHERE fi_domain = ? AND user_id = ?",
		domain,
		userID,
	)

	if r.repo.NoRows(err) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &found, nil
}

// GetReceivedPaymentByID returns received payment by id
func (r Repository) GetReceivedPaymentByID(id int64) (*entities.ReceivedPayment, error) {

	var found entities.ReceivedPayment

	err := r.repo.GetRaw(
		&found,
		"SELECT * FROM ReceivedPayment WHERE id = ?",
		id,
	)

	if r.repo.NoRows(err) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &found, nil
}

// getLastReceivedPayment returns the last received payment
func (r Repository) getLastReceivedPayment() (*entities.ReceivedPayment, error) {
	var receivedPayment entities.ReceivedPayment
	err := r.repo.GetRaw(&receivedPayment, "SELECT * FROM ReceivedPayment ORDER BY id DESC LIMIT 1")

	if r.repo.NoRows(err) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &receivedPayment, nil
}
