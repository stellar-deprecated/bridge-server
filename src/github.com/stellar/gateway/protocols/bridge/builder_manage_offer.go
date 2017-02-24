package bridge

import (
	"strconv"

	"github.com/stellar/gateway/protocols"
	b "github.com/stellar/go/build"
)

// ManageOfferOperationBody represents manage_offer operation
type ManageOfferOperationBody struct {
	PassiveOffer bool `json:"-"`
	Source       *string
	Selling      protocols.Asset
	Buying       protocols.Asset
	Amount       string
	Price        string
	OfferID      *string `json:"offer_id"`
}

// uint64

// ToTransactionMutator returns go-stellar-base TransactionMutator
func (op ManageOfferOperationBody) ToTransactionMutator() b.TransactionMutator {
	mutators := []interface{}{
		b.Amount(op.Amount),
		b.Rate{
			Selling: op.Selling.ToBaseAsset(),
			Buying:  op.Buying.ToBaseAsset(),
			Price:   b.Price(op.Price),
		},
	}

	if op.OfferID != nil {
		// Validated in Validate()
		u, _ := strconv.ParseUint(*op.OfferID, 10, 64)
		mutators = append(mutators, b.OfferID(u))
	}

	if op.Source != nil {
		mutators = append(mutators, b.SourceAccount{*op.Source})
	}

	return b.ManageOffer(op.PassiveOffer, mutators...)
}

// Validate validates if operation body is valid.
func (op ManageOfferOperationBody) Validate() error {
	if op.OfferID != nil {
		_, err := strconv.ParseUint(*op.OfferID, 10, 64)
		if err != nil {
			return protocols.NewInvalidParameterError("offer_id", *op.OfferID)
		}
	}

	if op.Source != nil && !protocols.IsValidAccountID(*op.Source) {
		return protocols.NewInvalidParameterError("source", *op.Source)
	}

	return nil
}
