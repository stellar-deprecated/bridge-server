package mocks

import (
	"time"

	"github.com/stellar/gateway/db/entities"
	"github.com/stellar/gateway/horizon"
	"github.com/stretchr/testify/mock"
)

type MockEntityManager struct {
	mock.Mock
}

func (m *MockEntityManager) Persist(object entities.Entity) (err error) {
	a := m.Called(object)
	return a.Error(0)
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

type MockTransactionSubmitter struct {
	mock.Mock
}

func (ts *MockTransactionSubmitter) SubmitTransaction(seed string, operation, memo interface{}) (response horizon.SubmitTransactionResponse, err error) {
	a := ts.Called(seed, operation, memo)
	return a.Get(0).(horizon.SubmitTransactionResponse), a.Error(1)
}

var PredefinedTime time.Time

func Now() time.Time {
	return PredefinedTime
}
