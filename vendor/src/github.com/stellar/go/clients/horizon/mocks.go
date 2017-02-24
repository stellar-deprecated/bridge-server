package horizon

import "github.com/stretchr/testify/mock"

// MockClient is a mockable horizon client.
type MockClient struct {
	mock.Mock
}

// LoadAccount is a mocking a method
func (m *MockClient) LoadAccount(accountID string) (Account, error) {
	a := m.Called(accountID)
	return a.Get(0).(Account), a.Error(1)
}

// LoadMemo is a mocking a method
func (m *MockClient) LoadMemo(p *Payment) error {
	a := m.Called(p)
	return a.Error(0)
}

// StreamPayments is a mocking a method
func (m *MockClient) StreamPayments(accountID string, cursor *string, handler PaymentHandler) error {
	a := m.Called(accountID, cursor, handler)
	return a.Error(0)
}

// StreamTransactions is a mocking a method
func (m *MockClient) StreamTransactions(accountID string, cursor *string, handler TransactionHandler) error {
	a := m.Called(accountID, cursor, handler)
	return a.Error(0)
}

// SubmitTransaction is a mocking a method
func (m *MockClient) SubmitTransaction(txeBase64 string) (TransactionSuccess, error) {
	a := m.Called(txeBase64)
	return a.Get(0).(TransactionSuccess), a.Error(1)
}
