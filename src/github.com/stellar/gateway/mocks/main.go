package mocks

import (
	"net/http"
	"net/url"
	"time"

	"github.com/stellar/gateway/db"
	"github.com/stellar/gateway/db/entities"
	"github.com/stellar/gateway/horizon"
	"github.com/stellar/go/clients/stellartoml"
	fproto "github.com/stellar/go/protocols/federation"
	"github.com/stellar/go/xdr"
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

var _ db.EntityManagerInterface = &MockEntityManager{}

// MockFederationResolver ...
type MockFederationResolver struct {
	mock.Mock
}

// LookupByAddress is a mocking a method
func (m *MockFederationResolver) LookupByAddress(addy string) (*fproto.NameResponse, error) {
	a := m.Called(addy)
	return a.Get(0).(*fproto.NameResponse), a.Error(1)
}

// LookupByAccountID is a mocking a method
func (m *MockFederationResolver) LookupByAccountID(aid string) (*fproto.IDResponse, error) {
	a := m.Called(aid)
	return a.Get(0).(*fproto.IDResponse), a.Error(1)
}

// ForwardRequest is a mocking a method
func (m *MockFederationResolver) ForwardRequest(domain string, fields url.Values) (*fproto.NameResponse, error) {
	a := m.Called(domain, fields)
	return a.Get(0).(*fproto.NameResponse), a.Error(1)
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

// Get is a mocking a method
func (m *MockHTTPClient) Get(url string) (resp *http.Response, err error) {
	a := m.Called(url)
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

// LoadOperation is a mocking a method
func (m *MockHorizon) LoadOperation(operationID string) (response horizon.PaymentResponse, err error) {
	a := m.Called(operationID)
	return a.Get(0).(horizon.PaymentResponse), a.Error(1)
}

// LoadMemo is a mocking a method
func (m *MockHorizon) LoadMemo(p *horizon.PaymentResponse) (err error) {
	a := m.Called(p)
	return a.Error(0)
}

// LoadAccountMergeAmount is a mocking a method
func (m *MockHorizon) LoadAccountMergeAmount(p *horizon.PaymentResponse) (err error) {
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

// GetReceivedPaymentByOperationID is a mocking a method
func (m *MockRepository) GetReceivedPaymentByOperationID(operationID int64) (*entities.ReceivedPayment, error) {
	a := m.Called(operationID)
	if a.Get(0) == nil {
		return nil, a.Error(1)
	}
	return a.Get(0).(*entities.ReceivedPayment), a.Error(1)
}

// GetAuthData is a mocking a method
func (m *MockRepository) GetAuthData(id string) (*entities.AuthData, error) {
	a := m.Called(id)
	if a.Get(0) == nil {
		return nil, a.Error(1)
	}
	return a.Get(0).(*entities.AuthData), a.Error(1)
}

func (m *MockRepository) GetReceivedPayments(page, limit int) ([]*entities.ReceivedPayment, error) {
	a := m.Called(page, limit)
	if a.Get(0) == nil {
		return nil, a.Error(1)
	}
	return a.Get(0).([]*entities.ReceivedPayment), a.Error(1)
}

func (m *MockRepository) GetSentTransactions(page, limit int) ([]*entities.SentTransaction, error) {
	a := m.Called(page, limit)
	if a.Get(0) == nil {
		return nil, a.Error(1)
	}
	return a.Get(0).([]*entities.SentTransaction), a.Error(1)
}

func (m *MockRepository) GetSentTransactionByPaymentID(paymentID string) (*entities.SentTransaction, error) {
	a := m.Called(paymentID)
	if a.Get(0) == nil {
		return nil, a.Error(1)
	}
	return a.Get(0).(*entities.SentTransaction), a.Error(1)
}

var _ db.RepositoryInterface = &MockRepository{}

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

func (m *MockStellartomlResolver) GetStellarToml(domain string) (resp *stellartoml.Response, err error) {
	a := m.Called(domain)
	return a.Get(0).(*stellartoml.Response), a.Error(1)
}

// GetStellarTomlByAddress is a mocking a method
func (m *MockStellartomlResolver) GetStellarTomlByAddress(addy string) (*stellartoml.Response, error) {
	a := m.Called(addy)
	return a.Get(0).(*stellartoml.Response), a.Error(1)
}

// MockTransactionSubmitter ...
type MockTransactionSubmitter struct {
	mock.Mock
}

// SubmitTransaction is a mocking a method
func (ts *MockTransactionSubmitter) SubmitTransaction(paymentID *string, seed string, operation, memo interface{}) (response horizon.SubmitTransactionResponse, err error) {
	a := ts.Called(paymentID, seed, operation, memo)
	return a.Get(0).(horizon.SubmitTransactionResponse), a.Error(1)
}

// SignAndSubmitRawTransaction is a mocking a method
func (ts *MockTransactionSubmitter) SignAndSubmitRawTransaction(paymentID *string, seed string, tx *xdr.Transaction) (response horizon.SubmitTransactionResponse, err error) {
	a := ts.Called(paymentID, seed, tx)
	return a.Get(0).(horizon.SubmitTransactionResponse), a.Error(1)
}

// PredefinedTime is a time.Time object that will be returned by Now() function
var PredefinedTime time.Time

// Now is a mocking a method
func Now() time.Time {
	return PredefinedTime
}
