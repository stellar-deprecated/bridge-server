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

type MockEntityManager struct {
	mock.Mock
}

func (m *MockEntityManager) Delete(object entities.Entity) (err error) {
	a := m.Called(object)
	return a.Error(0)
}

func (m *MockEntityManager) Persist(object entities.Entity) (err error) {
	a := m.Called(object)
	return a.Error(0)
}

type MockFederationResolver struct {
	mock.Mock
}

func (m *MockFederationResolver) Resolve(address string) (response federation.Response, stellarToml stellartoml.StellarToml, err error) {
	a := m.Called(address)
	return a.Get(0).(federation.Response), a.Get(1).(stellartoml.StellarToml), a.Error(2)
}

func (m *MockFederationResolver) GetDestination(federationUrl, address string) (response federation.Response, err error) {
	a := m.Called(federationUrl, address)
	return a.Get(0).(federation.Response), a.Error(1)
}

type MockHttpClient struct {
	mock.Mock
}

func (m *MockHttpClient) PostForm(url string, data url.Values) (resp *http.Response, err error) {
	a := m.Called(url, data)
	return a.Get(0).(*http.Response), a.Error(1)
}

type MockHorizon struct {
	mock.Mock
}

func (m *MockHorizon) LoadAccount(accountId string) (response horizon.AccountResponse, err error) {
	a := m.Called(accountId)
	return a.Get(0).(horizon.AccountResponse), a.Error(1)
}

func (m *MockHorizon) LoadMemo(p *horizon.PaymentResponse) (err error) {
	a := m.Called(p)
	return a.Error(0)
}

func (m *MockHorizon) StreamPayments(accountId string, cursor *string, onPaymentHandler horizon.PaymentHandler) (err error) {
	a := m.Called(accountId, cursor, onPaymentHandler)
	return a.Error(0)
}

func (m *MockHorizon) SubmitTransaction(txeBase64 string) (response horizon.SubmitTransactionResponse, err error) {
	a := m.Called(txeBase64)
	return a.Get(0).(horizon.SubmitTransactionResponse), a.Error(1)
}

type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) GetLastCursorValue() (cursor *string, err error) {
	a := m.Called()
	return a.Get(0).(*string), a.Error(1)
}

func (m *MockRepository) GetAuthorizedTransactionByMemo(memo string) (*entities.AuthorizedTransaction, error) {
	a := m.Called(memo)
	return a.Get(0).(*entities.AuthorizedTransaction), a.Error(1)
}

func (m *MockRepository) GetAllowedFiByDomain(domain string) (*entities.AllowedFi, error) {
	a := m.Called(domain)
	return a.Get(0).(*entities.AllowedFi), a.Error(1)
}

func (m *MockRepository) GetAllowedUserByDomainAndUserId(domain, userId string) (*entities.AllowedUser, error) {
	a := m.Called(domain, userId)
	return a.Get(0).(*entities.AllowedUser), a.Error(1)
}

type MockStellartomlResolver struct {
	mock.Mock
}

func (m *MockStellartomlResolver) GetStellarToml(domain string) (stellarToml stellartoml.StellarToml, err error) {
	a := m.Called(domain)
	return a.Get(0).(stellartoml.StellarToml), a.Error(1)
}

func (m *MockStellartomlResolver) GetStellarTomlByAddress(address string) (stellarToml stellartoml.StellarToml, err error) {
	a := m.Called(address)
	return a.Get(0).(stellartoml.StellarToml), a.Error(1)
}

type MockTransactionSubmitter struct {
	mock.Mock
}

func (ts *MockTransactionSubmitter) SubmitTransaction(seed string, operation, memo interface{}) (response horizon.SubmitTransactionResponse, err error) {
	a := ts.Called(seed, operation, memo)
	return a.Get(0).(horizon.SubmitTransactionResponse), a.Error(1)
}

func (ts *MockTransactionSubmitter) SignAndSubmitRawTransaction(seed string, tx *xdr.Transaction) (response horizon.SubmitTransactionResponse, err error) {
	a := ts.Called(seed, tx)
	return a.Get(0).(horizon.SubmitTransactionResponse), a.Error(1)
}

var PredefinedTime time.Time

func Now() time.Time {
	return PredefinedTime
}
