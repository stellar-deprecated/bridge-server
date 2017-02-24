package bridge

import (
	"strconv"

	"github.com/stellar/gateway/protocols"
	b "github.com/stellar/go/build"
)

// PathPaymentOperationBody represents path_payment operation
type PathPaymentOperationBody struct {
	Source *string

	SendMax   string          `json:"send_max"`
	SendAsset protocols.Asset `json:"send_asset"`

	Destination       string
	DestinationAmount string          `json:"destination_amount"`
	DestinationAsset  protocols.Asset `json:"destination_asset"`

	Path []protocols.Asset
}

// ToTransactionMutator returns go-stellar-base TransactionMutator
func (op PathPaymentOperationBody) ToTransactionMutator() b.TransactionMutator {
	var path []b.Asset
	for _, pathAsset := range op.Path {
		path = append(path, pathAsset.ToBaseAsset())
	}

	mutators := []interface{}{
		b.Destination{op.Destination},
		b.PayWithPath{
			Asset:     op.SendAsset.ToBaseAsset(),
			MaxAmount: op.SendMax,
			Path:      path,
		},
	}

	if op.DestinationAsset.Code != "" && op.DestinationAsset.Issuer != "" {
		mutators = append(
			mutators,
			b.CreditAmount{
				op.DestinationAsset.Code,
				op.DestinationAsset.Issuer,
				op.DestinationAmount,
			},
		)
	} else {
		mutators = append(
			mutators,
			b.NativeAmount{op.DestinationAmount},
		)
	}

	if op.Source != nil {
		mutators = append(mutators, b.SourceAccount{*op.Source})
	}

	return b.Payment(mutators...)
}

// Validate validates if operation body is valid.
func (op PathPaymentOperationBody) Validate() error {
	if !protocols.IsValidAccountID(op.Destination) {
		return protocols.NewInvalidParameterError("destination", op.Destination)
	}

	if !protocols.IsValidAmount(op.SendMax) {
		return protocols.NewInvalidParameterError("send_max", op.SendMax)
	}

	if !protocols.IsValidAmount(op.DestinationAmount) {
		return protocols.NewInvalidParameterError("destination_amount", op.DestinationAmount)
	}

	if !op.SendAsset.Validate() {
		return protocols.NewInvalidParameterError("send_asset", op.SendAsset.String())
	}

	if !op.DestinationAsset.Validate() {
		return protocols.NewInvalidParameterError("destination_asset", op.DestinationAsset.String())
	}

	if op.Source != nil && !protocols.IsValidAccountID(*op.Source) {
		return protocols.NewInvalidParameterError("source", *op.Source)
	}

	for i, asset := range op.Path {
		if !asset.Validate() {
			return protocols.NewInvalidParameterError("path["+strconv.Itoa(i)+"]", asset.String())
		}
	}

	return nil
}
