package mocks

import (
	"net/http"
	"net/url"
	"time"

	"github.com/stellar/gateway/db/entities"
	"github.com/stellar/gateway/horizon"
	"github.com/stellar/gateway/protocols/federation"
	"github.com/stellar/gateway/protocols/stellartoml"
	"github.com/stellar/go-stellar-base/xdr"
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

// MockFederationResolver ...
type MockFederationResolver struct {
	mock.Mock
}

// Resolve is a mocking a method
func (m *MockFederationResolver) Resolve(address string) (response federation.Response, stellarToml stellartoml.StellarToml, err error) {
	a := m.Called(address)
	return a.Get(0).(federation.Response), a.Get(1).(stellartoml.StellarToml), a.Error(2)
}

// GetDestination is a mocking a method
func (m *MockFederationResolver) GetDestination(federationURL, address string) (response federation.Response, err error) {
	a := m.Called(federationURL, address)
	return a.Get(0).(federation.Response), a.Error(1)
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

// Do is a mocking a method
func (m *MockHTTPClient) Do(req *http.Request) (resp *http.Response, err error) {
	a := m.Called(req)
	return a.Get(0).(*http.Response), a.Error(1)
}

// MockHorizon ...
type MockHorizon struct {
	mock.Mock
}

// LoadAccount is a mocking a method
func (m *MockHorizon) LoadAccount(accountID string) (response horizon.AccountResponse, err error) {
	a := m.Called(accountID)
	return a.Get(0).(horizon.AccountResponse), a.Error(1)
}

// LoadMemo is a mocking a method
func (m *MockHorizon) LoadMemo(p *horizon.PaymentResponse) (err error) {
	a := m.Called(p)
	return a.Error(0)
}

// StreamPayments is a mocking a method
func (m *MockHorizon) StreamPayments(accountID string, cursor *string, onPaymentHandler horizon.PaymentHandler) (err error) {
	a := m.Called(accountID, cursor, onPaymentHandler)
	return a.Error(0)
}

// SubmitTransaction is a mocking a method
func (m *MockHorizon) SubmitTransaction(txeBase64 string) (response horizon.SubmitTransactionResponse, err error) {
	a := m.Called(txeBase64)
	return a.Get(0).(horizon.SubmitTransactionResponse), a.Error(1)
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

// MockStellartomlResolver ...
type MockStellartomlResolver struct {
	mock.Mock
}

// GetStellarToml is a mocking a method
func (m *MockStellartomlResolver) GetStellarToml(domain string) (stellarToml stellartoml.StellarToml, err error) {
	a := m.Called(domain)
	return a.Get(0).(stellartoml.StellarToml), a.Error(1)
}

// GetStellarTomlByAddress is a mocking a method
func (m *MockStellartomlResolver) GetStellarTomlByAddress(address string) (stellarToml stellartoml.StellarToml, err error) {
	a := m.Called(address)
	return a.Get(0).(stellartoml.StellarToml), a.Error(1)
}

// MockTransactionSubmitter ...
type MockTransactionSubmitter struct {
	mock.Mock
}

// SubmitTransaction is a mocking a method
func (ts *MockTransactionSubmitter) SubmitTransaction(seed string, operation, memo interface{}) (response horizon.SubmitTransactionResponse, err error) {
	a := ts.Called(seed, operation, memo)
	return a.Get(0).(horizon.SubmitTransactionResponse), a.Error(1)
}

// SignAndSubmitRawTransaction is a mocking a method
func (ts *MockTransactionSubmitter) SignAndSubmitRawTransaction(seed string, tx *xdr.Transaction) (response horizon.SubmitTransactionResponse, err error) {
	a := ts.Called(seed, tx)
	return a.Get(0).(horizon.SubmitTransactionResponse), a.Error(1)
}

// PredefinedTime is a time.Time object that will be returned by Now() function
var PredefinedTime time.Time

// Now is a mocking a method
func Now() time.Time {
	return PredefinedTime
}
