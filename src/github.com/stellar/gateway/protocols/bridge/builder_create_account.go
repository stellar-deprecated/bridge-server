package bridge

import (
	"github.com/stellar/gateway/protocols"
	b "github.com/stellar/go-stellar-base/build"
)

// CreateAccountOperationBody represents create_account operation
type CreateAccountOperationBody struct {
	Source          *string
	Destination     string
	StartingBalance string `json:"starting_balance"`
}

// ToTransactionMutator returns go-stellar-base TransactionMutator
func (op CreateAccountOperationBody) ToTransactionMutator() b.TransactionMutator {
	mutators := []interface{}{
		b.Destination{op.Destination},
		b.NativeAmount{op.StartingBalance},
	}

	if op.Source != nil {
		mutators = append(mutators, b.SourceAccount{*op.Source})
	}

	return b.CreateAccount(mutators...)
}

// Validate validates if operation body is valid.
func (op CreateAccountOperationBody) Validate() error {
	if !protocols.IsValidAccountID(op.Destination) {
		return protocols.NewInvalidParameterError("destination", op.Destination)
	}

	if !protocols.IsValidAmount(op.StartingBalance) {
		return protocols.NewInvalidParameterError("starting_balance", op.StartingBalance)
	}

	if op.Source != nil && !protocols.IsValidAccountID(*op.Source) {
		return protocols.NewInvalidParameterError("source", *op.Source)
	}

	return nil
}
