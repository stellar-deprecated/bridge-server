// Package build provides a builder system for constructing various xdr
// structures used by the stellar network.
//
// At the core of this package is the *Builder and *Mutator types.  A Builder
// object (ex. PaymentBuilder, TransactionBuilder) contain an underlying xdr
// struct that is being iteratively built by having zero or more Mutator structs
// applied to it. See ExampleTransactionBuilder in main_test.go for an example.
package build

import (
	"github.com/stellar/go-stellar-base/network"
	"github.com/stellar/go-stellar-base/xdr"
)

const (
	// MemoTextMaxLength represents the maximum number of bytes a valid memo of
	// type "MEMO_TEXT" can be.
	MemoTextMaxLength = 28
)

var (
	// PublicNetwork is a mutator that configures the transaction for submission
	// to the main public stellar network.
	PublicNetwork = Network{network.PublicNetworkPassphrase}

	// TestNetwork is a mutator that configures the transaction for submission
	// to the test stellar network (often called testnet).
	TestNetwork = Network{network.TestNetworkPassphrase}

	// DefaultNetwork is a mutator that configures the transaction for submission
	// to the default stellar network.  Integrators may change this value to
	// another `Network` mutator if they would like to effect the default in a
	// process-global manner.
	DefaultNetwork = TestNetwork
)

// AllowTrustAsset is a mutator capable of setting the asset on
// an operations that have one.
type AllowTrustAsset struct {
	Code string
}

// Authorize is a mutator capable of setting the `authorize` flag
type Authorize struct {
	Value bool
}

// CreditAmount is a mutator that configures a payment to be using credit
// asset and have the amount provided.
type CreditAmount struct {
	Code string
	Issuer string
	Amount string
}

// Defaults is a mutator that sets defaults
type Defaults struct{}

// Destination is a mutator capable of setting the destination on
// an operations that have one.
type Destination struct {
	AddressOrSeed string
}

// MemoHash is a mutator that sets a memo on the mutated transaction of type
// MEMO_HASH.
type MemoHash struct {
	Value xdr.Hash
}

// MemoID is a mutator that sets a memo on the mutated transaction of type
// MEMO_ID.
type MemoID struct {
	Value uint64
}

// MemoReturn is a mutator that sets a memo on the mutated transaction of type
// MEMO_RETURN.
type MemoReturn struct {
	Value xdr.Hash
}

// MemoText is a mutator that sets a memo on the mutated transaction of type
// MEMO_TEXT.
type MemoText struct {
	Value string
}

// NativeAmount is a mutator that configures a payment to be using native
// currency and have the amount provided.
type NativeAmount struct {
	Amount string
}

// Sequence is a mutator that sets the sequence number on a transaction
type Sequence struct {
	Sequence uint64
}

// Sign is a mutator that contributes a signature of the provided envelope's
// transaction with the configured key
type Sign struct {
	Seed string
}

// SourceAccount is a mutator capable of setting the source account on
// an xdr.Operation and an xdr.Transaction
type SourceAccount struct {
	AddressOrSeed string
}

// Trustor is a mutator capable of setting the trustor on
// allow_trust operation.
type Trustor struct {
	Address string
}

// Network establishes the stellar network that a transaction should apply to.
// This modifier influences how a transaction is hashed for the purposes of signature generation.
type Network struct {
	Passphrase string
}

// ID returns the network ID derived from this struct's Passphrase
func (n *Network) ID() [32]byte {
	return network.ID(n.Passphrase)
}
