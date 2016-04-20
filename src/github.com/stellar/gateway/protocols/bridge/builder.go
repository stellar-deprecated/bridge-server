package bridge

import (
	"encoding/json"

	"github.com/stellar/gateway/protocols"
	b "github.com/stellar/go-stellar-base/build"
)

// OperationType is the type of operation
type OperationType string

const (
	// OperationTypeCreateAccount represents create_account operation
	OperationTypeCreateAccount OperationType = "create_account"
	// OperationTypePayment represents payment operation
	OperationTypePayment OperationType = "payment"
	// OperationTypePathPayment represents path_payment operation
	OperationTypePathPayment OperationType = "path_payment"
	// OperationTypeManageOffer represents manage_offer operation
	OperationTypeManageOffer OperationType = "manage_offer"
	// OperationTypeCreatePassiveOffer represents create_passive_offer operation
	OperationTypeCreatePassiveOffer OperationType = "create_passive_offer"
	// OperationTypeSetOptions represents set_options operation
	OperationTypeSetOptions OperationType = "set_options"
	// OperationTypeChangeTrust represents change_trust operation
	OperationTypeChangeTrust OperationType = "change_trust"
	// OperationTypeAllowTrust represents allow_trust operation
	OperationTypeAllowTrust OperationType = "allow_trust"
	// OperationTypeAccountMerge represents account_merge operation
	OperationTypeAccountMerge OperationType = "account_merge"
	// OperationTypeInflation represents inflation operation
	OperationTypeInflation OperationType = "inflation"
	// OperationTypeManageData represents manage_data operation
	OperationTypeManageData OperationType = "manage_data"
)

// BuilderRequest represents request made to /builder endpoint of bridge server
type BuilderRequest struct {
	Source         string
	SequenceNumber string `json:"sequence_number"`
	Operations     []Operation
	Signers        []string
}

// Operation contains
type Operation struct {
	Type OperationType
	Data json.RawMessage // delay parsing until we know operation type
}

// OperationData interface is a common interface for builder operations
type OperationData interface {
	ToTransactionMutator() b.TransactionMutator
}

// CreateAccountOperationData represents create_account operation
type CreateAccountOperationData struct {
	Source          *string
	Destination     string
	StartingBalance string `json:"starting_balance"`
}

// ToTransactionMutator returns go-stellar-base TransactionMutator
func (op CreateAccountOperationData) ToTransactionMutator() b.TransactionMutator {
	mutators := []interface{}{
		b.Destination{op.Destination},
		b.NativeAmount{op.StartingBalance},
	}

	if op.Source != nil {
		mutators = append(mutators, b.SourceAccount{*op.Source})
	}

	return b.CreateAccount(mutators...)
}

// PaymentOperationData represents payment operation
type PaymentOperationData struct {
	Source      *string
	Destination string
	Amount      string
	AssetCode   *string `json:"asset_code"`
	AssetIssuer *string `json:"asset_issuer"`
}

// ToTransactionMutator returns go-stellar-base TransactionMutator
func (op PaymentOperationData) ToTransactionMutator() b.TransactionMutator {
	mutators := []interface{}{
		b.Destination{op.Destination},
	}

	if op.AssetCode != nil && op.AssetIssuer != nil {
		mutators = append(
			mutators,
			b.CreditAmount{*op.AssetCode, *op.AssetIssuer, op.Amount},
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

// BuilderResponse represents response returned by /builder endpoint of bridge server
type BuilderResponse struct {
	protocols.SuccessResponse
	TransactionEnvelope string `json:"transaction_envelope"`
}

// Marshal marshals BuilderResponse
func (response *BuilderResponse) Marshal() []byte {
	json, _ := json.MarshalIndent(response, "", "  ")
	return json
}
