package bridge

import (
	"github.com/stellar/gateway/protocols"
	b "github.com/stellar/go-stellar-base/build"
)

// AllowTrustOperationBody represents create_account operation
type AllowTrustOperationBody struct {
	Source    *string
	AssetCode string `json:"asset_code"`
	Trustor   string
	Authorize bool
}

// ToTransactionMutator returns go-stellar-base TransactionMutator
func (op AllowTrustOperationBody) ToTransactionMutator() b.TransactionMutator {
	mutators := []interface{}{
		b.AllowTrustAsset{op.AssetCode},
		b.Trustor{op.Trustor},
		b.Authorize{op.Authorize},
	}

	if op.Source != nil {
		mutators = append(mutators, b.SourceAccount{*op.Source})
	}

	return b.AllowTrust(mutators...)
}

// Validate validates if operation body is valid.
func (op AllowTrustOperationBody) Validate() error {
	if !protocols.IsValidAssetCode(op.AssetCode) {
		return protocols.NewInvalidParameterError("asset_code", op.AssetCode)
	}

	if !protocols.IsValidAccountID(op.Trustor) {
		return protocols.NewInvalidParameterError("trustor", op.Trustor)
	}

	if op.Source != nil && !protocols.IsValidAccountID(*op.Source) {
		return protocols.NewInvalidParameterError("source", *op.Source)
	}

	return nil
}
