package bridge

import (
	"github.com/stellar/gateway/protocols"
	b "github.com/stellar/go-stellar-base/build"
)

// InflationOperationBody represents inflation operation
type InflationOperationBody struct {
	Source *string
}

// ToTransactionMutator returns go-stellar-base TransactionMutator
func (op InflationOperationBody) ToTransactionMutator() b.TransactionMutator {
	var mutators []interface{}

	if op.Source != nil {
		mutators = append(mutators, b.SourceAccount{*op.Source})
	}

	return b.Inflation(mutators...)
}

// Validate validates if operation body is valid.
func (op InflationOperationBody) Validate() error {
	if op.Source != nil && !protocols.IsValidAccountID(*op.Source) {
		return protocols.NewInvalidParameterError("source", *op.Source)
	}

	return nil
}
