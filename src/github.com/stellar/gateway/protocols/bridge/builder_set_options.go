package bridge

import (
	"github.com/stellar/gateway/protocols"
	b "github.com/stellar/go/build"
)

// SetOptionsOperationBody represents set_options operation
type SetOptionsOperationBody struct {
	Source          *string
	InflationDest   *string           `json:"inflation_dest"`
	SetFlags        *[]int            `json:"set_flags"`
	ClearFlags      *[]int            `json:"clear_flags"`
	MasterWeight    *uint32           `json:"master_weight"`
	LowThreshold    *uint32           `json:"low_threshold"`
	MediumThreshold *uint32           `json:"medium_threshold"`
	HighThreshold   *uint32           `json:"high_threshold"`
	HomeDomain      *string           `json:"home_domain"`
	Signer          *SetOptionsSigner `json:"signer"`
}

// SetOptionsSigner is a struct that representing signer in SetOptions operation body
type SetOptionsSigner struct {
	PublicKey string `json:"public_key"`
	Weight    uint32 `json:"weight"`
}

// ToTransactionMutator returns go-stellar-base TransactionMutator
func (op SetOptionsOperationBody) ToTransactionMutator() b.TransactionMutator {
	var mutators []interface{}

	if op.InflationDest != nil {
		mutators = append(mutators, b.InflationDest(*op.InflationDest))
	}

	if op.SetFlags != nil {
		for _, flag := range *op.SetFlags {
			mutators = append(mutators, b.SetFlag(flag))
		}
	}

	if op.ClearFlags != nil {
		for _, flag := range *op.ClearFlags {
			mutators = append(mutators, b.ClearFlag(flag))
		}
	}

	if op.MasterWeight != nil {
		mutators = append(mutators, b.MasterWeight(*op.MasterWeight))
	}

	if op.LowThreshold != nil {
		mutators = append(mutators, b.SetLowThreshold(*op.LowThreshold))
	}

	if op.MediumThreshold != nil {
		mutators = append(mutators, b.SetMediumThreshold(*op.MediumThreshold))
	}

	if op.HighThreshold != nil {
		mutators = append(mutators, b.SetHighThreshold(*op.HighThreshold))
	}

	if op.HomeDomain != nil {
		mutators = append(mutators, b.HomeDomain(*op.HomeDomain))
	}

	if op.Signer != nil {
		mutators = append(mutators, b.Signer{
			Address: op.Signer.PublicKey,
			Weight:  op.Signer.Weight,
		})
	}

	return b.SetOptions(mutators...)
}

// Validate validates if operation body is valid.
func (op SetOptionsOperationBody) Validate() error {
	if op.InflationDest != nil && !protocols.IsValidAccountID(*op.InflationDest) {
		return protocols.NewInvalidParameterError("inflation_dest", *op.InflationDest)
	}

	if op.Signer != nil {
		if !protocols.IsValidAccountID(op.Signer.PublicKey) {
			return protocols.NewInvalidParameterError("signer.public_key", op.Signer.PublicKey)
		}
	}

	if op.Source != nil && !protocols.IsValidAccountID(*op.Source) {
		return protocols.NewInvalidParameterError("source", *op.Source)
	}

	return nil
}
