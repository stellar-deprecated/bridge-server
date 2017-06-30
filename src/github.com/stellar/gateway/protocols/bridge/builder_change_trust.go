package bridge

import (
	"github.com/stellar/gateway/protocols"
	b "github.com/stellar/go/build"
)

// ChangeTrustOperationBody represents change_trust operation
type ChangeTrustOperationBody struct {
	Source *string
	Asset  protocols.Asset
	// nil means max limit
	Limit *string
}

// ToTransactionMutator returns go-stellar-base TransactionMutator
func (op ChangeTrustOperationBody) ToTransactionMutator() b.TransactionMutator {
	mutators := []interface{}{
		op.Asset.ToBaseAsset(),
	}

	if op.Limit == nil {
		// Set MaxLimit
		mutators = append(mutators, b.MaxLimit)
	} else {
		mutators = append(mutators, b.Limit(*op.Limit))
	}

	if op.Source != nil {
		mutators = append(mutators, b.SourceAccount{*op.Source})
	}

	return b.ChangeTrust(mutators...)
}

// Validate validates if operation body is valid.
func (op ChangeTrustOperationBody) Validate() error {
	if !op.Asset.Validate() {
		return protocols.NewInvalidParameterError("asset", op.Asset.String(), "Asset is invalid.")
	}

	if op.Limit != nil {
		if !protocols.IsValidAmount(*op.Limit) {
			return protocols.NewInvalidParameterError("limit", *op.Limit, "Limit is not a valid amount.")
		}
	}

	if op.Source != nil && !protocols.IsValidAccountID(*op.Source) {
		return protocols.NewInvalidParameterError("source", *op.Source, "Source must be a public key (starting with `G`).")
	}

	return nil
}
