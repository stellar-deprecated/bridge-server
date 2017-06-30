package bridge

import (
	"github.com/stellar/gateway/protocols"
	b "github.com/stellar/go/build"
)

// AllowTrustOperationBody represents allow_trust operation
type AllowTrustOperationBody struct {
	Source    *string
	AssetCode string `json:"asset_code"`
	Trustor   string
	Authorize bool
}

// ToTransactionMutator returns stellar/go TransactionMutator
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
		return protocols.NewInvalidParameterError("asset_code", op.AssetCode, "Asset code is invalid")
	}

	if !protocols.IsValidAccountID(op.Trustor) {
		return protocols.NewInvalidParameterError("trustor", op.Trustor, "Trustor must be a public key (starting with `G`).")
	}

	if op.Source != nil && !protocols.IsValidAccountID(*op.Source) {
		return protocols.NewInvalidParameterError("source", *op.Source, "Source must be a public key (starting with `G`).")
	}

	return nil
}
