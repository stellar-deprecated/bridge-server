package protocols

import (
	"github.com/stellar/go-stellar-base/amount"
	"github.com/stellar/go-stellar-base/keypair"
)

// IsValidAccountID returns true if account ID is valid
func IsValidAccountID(accountID string) bool {
	_, err := keypair.Parse(accountID)
	if err != nil {
		return false
	}
	return true
}

// IsValidAssetCode returns true if asset code is valid
func IsValidAssetCode(code string) bool {
	if len(code) < 1 || len(code) > 12 {
		return false
	}
	return true
}

// IsValidAmount returns true if amount is valid
func IsValidAmount(a string) bool {
	_, err := amount.Parse(a)
	if err != nil {
		return false
	}
	return true
}
