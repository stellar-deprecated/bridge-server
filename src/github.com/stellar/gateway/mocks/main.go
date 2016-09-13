package mocks

import (
	"net/http"
	"net/url"
	"time"

	"github.com/stellar/gateway/db/entities"
	"github.com/stellar/go-stellar-base/xdr"
	"github.com/stellar/go/clients/horizon"
	"github.com/stretchr/testify/mock"
)

// MockEntityManager ...
type MockEntityManager struct {
	mock.Mock
}

// Delete is a mocking a method
func (m *MockEntityManager) Delete(object entities.Entity) (err error) {
	a := m.Called(object)
	return a.Error(0)
}

// Persist is a mocking a method
func (m *MockEntityManager) Persist(object entities.Entity) (err error) {
	a := m.Called(object)
	return a.Error(0)
}

// MockHTTPClient ...
type MockHTTPClient struct {
	mock.Mock
}

// PostForm is a mocking a method
func (m *MockHTTPClient) PostForm(url string, data url.Values) (resp *http.Response, err error) {
	a := m.Called(url, data)
	return a.Get(0).(*http.Response), a.Error(1)
}

// MockRepository ...
type MockRepository struct {
	mock.Mock
}

// GetLastCursorValue is a mocking a method
func (m *MockRepository) GetLastCursorValue() (cursor *string, err error) {
	a := m.Called()
	return a.Get(0).(*string), a.Error(1)
}

// GetAuthorizedTransactionByMemo is a mocking a method
func (m *MockRepository) GetAuthorizedTransactionByMemo(memo string) (*entities.AuthorizedTransaction, error) {
	a := m.Called(memo)
	if a.Get(0) == nil {
		return nil, a.Error(1)
	}
	return a.Get(0).(*entities.AuthorizedTransaction), a.Error(1)
}

// GetAllowedFiByDomain is a mocking a method
func (m *MockRepository) GetAllowedFiByDomain(domain string) (*entities.AllowedFi, error) {
	a := m.Called(domain)
	if a.Get(0) == nil {
		return nil, a.Error(1)
	}
	return a.Get(0).(*entities.AllowedFi), a.Error(1)
}

// GetAllowedUserByDomainAndUserID is a mocking a method
func (m *MockRepository) GetAllowedUserByDomainAndUserID(domain, userID string) (*entities.AllowedUser, error) {
	a := m.Called(domain, userID)
	if a.Get(0) == nil {
		return nil, a.Error(1)
	}
	return a.Get(0).(*entities.AllowedUser), a.Error(1)
}

// GetReceivedPaymentByID is a mocking a method
func (m *MockRepository) GetReceivedPaymentByID(id int64) (*entities.ReceivedPayment, error) {
	a := m.Called(id)
	if a.Get(0) == nil {
		return nil, a.Error(1)
	}
	return a.Get(0).(*entities.ReceivedPayment), a.Error(1)
}

// MockSignerVerifier ...
type MockSignerVerifier struct {
	mock.Mock
}

// Sign is a mocking a method
func (m *MockSignerVerifier) Sign(secretSeed string, message []byte) (string, error) {
	a := m.Called(secretSeed, message)
	return a.String(0), a.Error(1)
}

// Verify is a mocking a method
func (m *MockSignerVerifier) Verify(publicKey string, message, signature []byte) error {
	a := m.Called(publicKey, message, signature)
	return a.Error(0)
}

// MockTransactionSubmitter ...
type MockTransactionSubmitter struct {
	mock.Mock
}

// SubmitTransaction is a mocking a method
func (ts *MockTransactionSubmitter) SubmitTransaction(seed string, operation, memo interface{}) (response horizon.TransactionSuccess, err error) {
	a := ts.Called(seed, operation, memo)
	return a.Get(0).(horizon.TransactionSuccess), a.Error(1)
}

// SignAndSubmitRawTransaction is a mocking a method
func (ts *MockTransactionSubmitter) SignAndSubmitRawTransaction(seed string, tx *xdr.Transaction) (response horizon.TransactionSuccess, err error) {
	a := ts.Called(seed, tx)
	return a.Get(0).(horizon.TransactionSuccess), a.Error(1)
}

// PredefinedTime is a time.Time object that will be returned by Now() function
var PredefinedTime time.Time

// Now is a mocking a method
func Now() time.Time {
	return PredefinedTime
}
