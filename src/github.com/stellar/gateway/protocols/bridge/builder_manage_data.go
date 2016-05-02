package bridge

import (
	"encoding/base64"

	"github.com/stellar/gateway/protocols"
	b "github.com/stellar/go-stellar-base/build"
)

// ManageDataOperationBody represents manage_data operation
type ManageDataOperationBody struct {
	Source *string
	Name   string
	Data   string
}

// ToTransactionMutator returns go-stellar-base TransactionMutator
func (op ManageDataOperationBody) ToTransactionMutator() b.TransactionMutator {
	var builder b.ManageDataBuilder

	if op.Data == "" {
		builder = b.ClearData(op.Name)
	} else {
		// This is validated in Validate()
		data, _ := base64.StdEncoding.DecodeString(op.Data)
		builder = b.SetData(op.Name, data)
	}

	if op.Source != nil {
		builder.Mutate(b.SourceAccount{*op.Source})
	}

	return builder
}

// Validate validates if operation body is valid.
func (op ManageDataOperationBody) Validate() error {
	if len(op.Name) > 64 {
		return protocols.NewInvalidParameterError("name", op.Name)
	}

	data, err := base64.StdEncoding.DecodeString(op.Data)
	if err != nil || len(data) > 64 {
		return protocols.NewInvalidParameterError("data", op.Data)
	}

	if op.Source != nil && !protocols.IsValidAccountID(*op.Source) {
		return protocols.NewInvalidParameterError("source", *op.Source)
	}

	return nil
}
