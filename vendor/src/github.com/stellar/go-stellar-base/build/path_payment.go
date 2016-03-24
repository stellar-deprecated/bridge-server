package build

import (
	"errors"
	
	"github.com/stellar/go-stellar-base/amount"
	"github.com/stellar/go-stellar-base/xdr"
)

// PathPayment groups the creation of a new PathPaymentBuilder with a call to Mutate.
func PathPayment(muts ...interface{}) (result PathPaymentBuilder) {
	result.Mutate(muts...)
	return
}

// PathPaymentMutator is a interface that wraps the
// MutatePathPayment operation.  types may implement this interface to
// specify how they modify an xdr.PathPaymentOp object
type PathPaymentMutator interface {
	MutatePathPayment(*xdr.PathPaymentOp) error
}

// PathPaymentBuilder represents a transaction that is being built.
type PathPaymentBuilder struct {
	O   xdr.Operation
	P   xdr.PathPaymentOp
	Err error
}

// Mutate applies the provided mutators to this builder's payment or operation.
func (b *PathPaymentBuilder) Mutate(muts ...interface{}) {
	for _, m := range muts {
		var err error
		switch mut := m.(type) {
		case PathPaymentMutator:
			err = mut.MutatePathPayment(&b.P)
		case OperationMutator:
			err = mut.MutateOperation(&b.O)
		default:
			err = errors.New("Mutator type not allowed")
		}

		if err != nil {
			b.Err = err
			return
		}
	}
}

// MutatePayment for Destination sets the PathPaymentOp's Destination field
func (m Destination) MutatePathPayment(o *xdr.PathPaymentOp) error {
	return setAccountId(m.AddressOrSeed, &o.Destination)
}

// MutatePayment for PathDestination sets the PathPaymentOp's DestinationAsset and
// DestinationAmount fields
func (m PathDestination) MutatePathPayment(o *xdr.PathPaymentOp) (err error) {
	switch {
	case m.Asset.Native:
		o.DestAsset, err = xdr.NewAsset(xdr.AssetTypeAssetTypeNative, nil)
	case !m.Asset.Native:
		o.DestAsset, err = createAlphaNumAsset(m.Asset.Code, m.Asset.Issuer)
	default:
		err = errors.New("Unknown Asset type")
	}

	if err != nil {
		return
	}

	o.DestAmount, err = amount.Parse(m.Amount)
	return
}

// MutatePayment for PathSend sets the PathPaymentOp's SendAsset and
// SendMax fields
func (m PathSend) MutatePathPayment(o *xdr.PathPaymentOp) (err error) {
	switch {
	case m.Asset.Native:
		o.SendAsset, err = xdr.NewAsset(xdr.AssetTypeAssetTypeNative, nil)
	case !m.Asset.Native:
		o.SendAsset, err = createAlphaNumAsset(m.Asset.Code, m.Asset.Issuer)
	default:
		err = errors.New("Unknown Asset type")
	}

	if err != nil {
		return
	}

	o.SendMax, err = amount.Parse(m.MaxAmount)
	return
}

// MutatePayment for Path sets the PathPaymentOp's Path field
func (m Path) MutatePathPayment(o *xdr.PathPaymentOp) (err error) {
	var path []xdr.Asset
	var xdrAsset xdr.Asset

	for _, asset := range m.Assets {
		switch {
		case asset.Native:
			xdrAsset, err = xdr.NewAsset(xdr.AssetTypeAssetTypeNative, nil)
			path = append(path, xdrAsset)
		case !asset.Native:
			xdrAsset, err = createAlphaNumAsset(asset.Code, asset.Issuer)
			path = append(path, xdrAsset)
		default:
			err = errors.New("Unknown Asset type")
		}

		if err != nil {
			return err
		}
	}

	o.Path = path
	return
}
