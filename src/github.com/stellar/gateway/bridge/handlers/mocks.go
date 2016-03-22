package handlers

import (
	"github.com/stretchr/testify/mock"
)

type MockAddressResolverHelper struct {
	mock.Mock
}

// func (arh *MockAddressResolverHelper) GetStellarToml(domain string) (stellarToml StellarToml, err error) {
// 	a := arh.Called(domain)
// 	return a.Get(0).(StellarToml), a.Error(1)
// }

// func (arh *MockAddressResolverHelper) GetDestination(federationUrl, address string) (destination StellarDestination, err error) {
// 	a := arh.Called(federationUrl, address)
// 	return a.Get(0).(StellarDestination), a.Error(1)
// }
