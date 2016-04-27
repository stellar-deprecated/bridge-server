package bridge

import (
	"github.com/stellar/gateway/protocols"
	b "github.com/stellar/go-stellar-base/build"
)

// ManageOfferOperationBody represents manage_offer operation
type ManageOfferOperationBody struct {
	Source  *string
	Selling protocols.Asset
	Buying  protocols.Asset
	Amount  string
	Price   string
	OfferID *uint64 `json:"offer_id"`
}

// ToTransactionMutator returns go-stellar-base TransactionMutator
func (op ManageOfferOperationBody) ToTransactionMutator() b.TransactionMutator {
	mutators := []interface{}{
		b.Amount(op.Amount),
		b.Rate{
			Selling: op.Selling.ToBaseAsset(),
			Buying:  op.Buying.ToBaseAsset(),
			Price:   op.Price,
		},
	}

	if op.OfferID != nil {
		mutators = append(mutators, b.OfferID{*op.OfferID})
	}

	if op.Source != nil {
		mutators = append(mutators, b.SourceAccount{*op.Source})
	}

	return b.ManageOffer(mutators...)
}

// Validate validates if operation body is valid.
func (op ManageOfferOperationBody) Validate() error {
	if op.Source != nil && !protocols.IsValidAccountID(*op.Source) {
		return protocols.NewInvalidParameterError("source", *op.Source)
	}

	return nil
}
