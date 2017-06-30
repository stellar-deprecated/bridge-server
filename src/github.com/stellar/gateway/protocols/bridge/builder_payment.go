package bridge

import (
	"github.com/stellar/gateway/protocols"
	b "github.com/stellar/go/build"
)

// PaymentOperationBody represents payment operation
type PaymentOperationBody struct {
	Source      *string
	Destination string
	Amount      string
	Asset       protocols.Asset
}

// ToTransactionMutator returns go-stellar-base TransactionMutator
func (op PaymentOperationBody) ToTransactionMutator() b.TransactionMutator {
	mutators := []interface{}{
		b.Destination{op.Destination},
	}

	if op.Asset.Code != "" && op.Asset.Issuer != "" {
		mutators = append(
			mutators,
			b.CreditAmount{op.Asset.Code, op.Asset.Issuer, op.Amount},
		)
	} else {
		mutators = append(
			mutators,
			b.NativeAmount{op.Amount},
		)
	}

	if op.Source != nil {
		mutators = append(mutators, b.SourceAccount{*op.Source})
	}

	return b.Payment(mutators...)
}

// Validate validates if operation body is valid.
func (op PaymentOperationBody) Validate() error {
	if !protocols.IsValidAccountID(op.Destination) {
		return protocols.NewInvalidParameterError("destination", op.Destination, "Destination must be a public key (starting with `G`).")
	}

	if !protocols.IsValidAmount(op.Amount) {
		return protocols.NewInvalidParameterError("amount", op.Amount, "Invalid amount.")
	}

	if !op.Asset.Validate() {
		return protocols.NewInvalidParameterError("asset", op.Asset.String(), "Invalid asset.")
	}

	if op.Source != nil && !protocols.IsValidAccountID(*op.Source) {
		return protocols.NewInvalidParameterError("source", *op.Source, "Source must be a public key (starting with `G`).")
	}

	return nil
}
