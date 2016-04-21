package bridge

import (
	"encoding/json"
	"strconv"

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

// Process parses operations and creates OperationBody object for each operation
func (r BuilderRequest) Process() error {
	var err error
	for i, operation := range r.Operations {
		var operationBody OperationBody

		switch operation.Type {
		case OperationTypeCreateAccount:
			var createAccount CreateAccountOperationBody
			err = json.Unmarshal(operation.RawBody, &createAccount)
			operationBody = createAccount
		case OperationTypePayment:
			var payment PaymentOperationBody
			err = json.Unmarshal(operation.RawBody, &payment)
			operationBody = payment
		case OperationTypePathPayment:
			var pathPayment PathPaymentOperationBody
			err = json.Unmarshal(operation.RawBody, &pathPayment)
			operationBody = pathPayment
		case OperationTypeChangeTrust:
			var changeTrust ChangeTrustOperationBody
			err = json.Unmarshal(operation.RawBody, &changeTrust)
			operationBody = changeTrust
		case OperationTypeAllowTrust:
			var allowTrust AllowTrustOperationBody
			err = json.Unmarshal(operation.RawBody, &allowTrust)
			operationBody = allowTrust
		default:
			return protocols.NewInvalidParameterError("operations["+strconv.Itoa(i)+"][type]", string(operation.Type))
		}

		if err != nil {
			return protocols.NewInvalidParameterError("operations["+strconv.Itoa(i)+"][body]", "")
		}

		r.Operations[i].Body = operationBody
	}

	return nil
}

// Validate validates if the request is correct.
func (r BuilderRequest) Validate() error {
	if !protocols.IsValidAccountID(r.Source) {
		return protocols.NewInvalidParameterError("source", r.Source)
	}

	for _, operation := range r.Operations {
		err := operation.Body.Validate()
		if err != nil {
			return err
		}
	}

	return nil
}

// Operation struct contains operation type and body
type Operation struct {
	Type    OperationType
	RawBody json.RawMessage `json:"body"` // Delay parsing until we know operation type
	Body    OperationBody   `json:"-"`    // Created during processing stage
}

// OperationBody interface is a common interface for builder operations
type OperationBody interface {
	ToTransactionMutator() b.TransactionMutator
	Validate() error
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
