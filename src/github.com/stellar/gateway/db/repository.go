package db

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/stellar/gateway/db/entities"
	"github.com/stellar/go/support/db"
)

// RepositoryInterface helps mocking Repository
type RepositoryInterface interface {
	GetLastCursorValue() (cursor *string, err error)
	GetAuthorizedTransactionByMemo(memo string) (*entities.AuthorizedTransaction, error)
	GetSentTransactionByPaymentID(paymentID string) (*entities.SentTransaction, error)
	GetAllowedFiByDomain(domain string) (*entities.AllowedFi, error)
	GetAllowedUserByDomainAndUserID(domain, userID string) (*entities.AllowedUser, error)
	GetAuthData(requestID string) (*entities.AuthData, error)
	GetReceivedPaymentByOperationID(operationID int64) (*entities.ReceivedPayment, error)
	GetReceivedPayments(page, limit int) ([]*entities.ReceivedPayment, error)
	GetSentTransactions(page, limit int) ([]*entities.SentTransaction, error)
}

// Repository helps getting data from DB
type Repository struct {
	driver Driver
	repo   *db.Session
	log    *logrus.Entry
}

// NewRepository creates a new Repository using driver
func NewRepository(driver Driver) (r Repository) {
	r.driver = driver
	r.repo = &db.Session{DB: driver.DB()}
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

	found.SetExists()
	return &found, nil
}

// GetSentTransactionByPaymentID returns sent transaction searching by payment ID
func (r Repository) GetSentTransactionByPaymentID(paymentID string) (*entities.SentTransaction, error) {

	var found entities.SentTransaction

	err := r.repo.GetRaw(
		&found,
		"SELECT * FROM SentTransaction WHERE payment_id = ?",
		paymentID,
	)

	if r.repo.NoRows(err) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	found.SetExists()
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

	found.SetExists()
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

	found.SetExists()
	return &found, nil
}

// GetReceivedPaymentByOperationID returns received payment by operation_id
func (r Repository) GetReceivedPaymentByOperationID(operationID int64) (*entities.ReceivedPayment, error) {

	var found entities.ReceivedPayment

	err := r.repo.GetRaw(
		&found,
		"SELECT * FROM ReceivedPayment WHERE operation_id = ?",
		operationID,
	)

	if r.repo.NoRows(err) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	found.SetExists()
	return &found, nil
}

// GetAuthData returns received auth data by request_id
func (r Repository) GetAuthData(requestID string) (*entities.AuthData, error) {
	var found entities.AuthData

	err := r.repo.GetRaw(
		&found,
		"SELECT * FROM AuthData WHERE request_id = ?",
		requestID,
	)

	if r.repo.NoRows(err) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	found.SetExists()
	return &found, nil
}

// GetReceivedPayments returns received payments
func (r Repository) GetReceivedPayments(page, limit int) ([]*entities.ReceivedPayment, error) {
	payments := []*entities.ReceivedPayment{}

	if page == 0 {
		page = 1
	}

	offset := (page - 1) * limit

	limitQuery := fmt.Sprintf("%d", limit)
	offsetQuery := fmt.Sprintf("%d", offset)
	orderQuery := "id desc"

	err := r.driver.GetMany(&payments, nil, &orderQuery, &offsetQuery, &limitQuery)
	return payments, err
}

// GetSentTransactions returns received payments
func (r Repository) GetSentTransactions(page, limit int) ([]*entities.SentTransaction, error) {
	transactions := []*entities.SentTransaction{}

	if page == 0 {
		page = 1
	}

	offset := (page - 1) * limit

	limitQuery := fmt.Sprintf("%d", limit)
	offsetQuery := fmt.Sprintf("%d", offset)
	orderQuery := "id desc"

	err := r.driver.GetMany(&transactions, nil, &orderQuery, &offsetQuery, &limitQuery)
	return transactions, err
}

// getLastReceivedPayment returns the last received payment
func (r Repository) getLastReceivedPayment() (*entities.ReceivedPayment, error) {
	var receivedPayment entities.ReceivedPayment
	// DO NOT use `processed_at` as payment can be reprocessed. Reprocessing will update `processed_at`
	// value but not `id`.
	err := r.repo.GetRaw(&receivedPayment, "SELECT * FROM ReceivedPayment ORDER BY id DESC LIMIT 1")

	if r.repo.NoRows(err) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	receivedPayment.SetExists()
	return &receivedPayment, nil
}
