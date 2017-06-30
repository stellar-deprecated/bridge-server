package protocols

import (
	"github.com/stellar/go/amount"
	"github.com/stellar/go/keypair"
)

// IsValidAccountID returns true if account ID is valid
func IsValidAccountID(accountID string) bool {
	_, err := keypair.Parse(accountID)
	if err != nil {
		return false
	}

	if accountID[0] != 'G' {
		return false
	}

	return true
}

// IsValidSecret returns true if secret is valid
func IsValidSecret(secret string) bool {
	_, err := keypair.Parse(secret)
	if err != nil {
		return false
	}

	if secret[0] != 'S' {
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
