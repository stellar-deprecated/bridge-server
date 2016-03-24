package build

import (
	"bytes"
	"encoding/hex"
	"errors"
	"fmt"

	"github.com/stellar/go-stellar-base/hash"
	"github.com/stellar/go-stellar-base/xdr"
)

// Transaction groups the creation of a new TransactionBuilder with a call
// to Mutate.
func Transaction(muts ...TransactionMutator) (result *TransactionBuilder) {
	result = &TransactionBuilder{}
	result.Mutate(muts...)
	result.Mutate(Defaults{})
	return
}

// TransactionMutator is a interface that wraps the
// MutateTransaction operation.  types may implement this interface to
// specify how they modify an xdr.Transaction object
type TransactionMutator interface {
	MutateTransaction(*TransactionBuilder) error
}

// TransactionBuilder represents a Transaction that is being constructed.
type TransactionBuilder struct {
	TX        *xdr.Transaction
	NetworkID [32]byte
	Err       error
}

// Mutate applies the provided TransactionMutators to this builder's transaction
func (b *TransactionBuilder) Mutate(muts ...TransactionMutator) {
	if b.TX == nil {
		b.TX = &xdr.Transaction{}
	}

	for _, m := range muts {
		err := m.MutateTransaction(b)
		if err != nil {
			b.Err = err
			return
		}
	}
}

// Hash returns the hash of this builder's transaction.
func (b *TransactionBuilder) Hash() ([32]byte, error) {
	var txBytes bytes.Buffer

	_, err := fmt.Fprintf(&txBytes, "%s", b.NetworkID)
	if err != nil {
		return [32]byte{}, err
	}

	_, err = xdr.Marshal(&txBytes, xdr.EnvelopeTypeEnvelopeTypeTx)
	if err != nil {
		return [32]byte{}, err
	}

	_, err = xdr.Marshal(&txBytes, b.TX)
	if err != nil {
		return [32]byte{}, err
	}

	return hash.Hash(txBytes.Bytes()), nil
}

// HashHex returns the hex-encoded hash of this builder's transaction
func (b *TransactionBuilder) HashHex() (string, error) {
	hash, err := b.Hash()
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(hash[:]), nil
}

// Sign returns an new TransactionEnvelopeBuilder using this builder's
// transaction as the basis and with signatures of that transaction from the
// provided Signers.
func (b *TransactionBuilder) Sign(signers ...string) (result TransactionEnvelopeBuilder) {
	result.Mutate(b)

	for _, s := range signers {
		result.Mutate(Sign{s})
	}

	return
}

// ------------------------------------------------------------
//
//   Mutator implementations
//
// ------------------------------------------------------------

// MutateTransaction for AllowTrustBuilder causes the underylying AllowTrustOp
// to be added to the operation list for the provided transaction
func (m AllowTrustBuilder) MutateTransaction(o *TransactionBuilder) error {
	if m.Err != nil {
		return m.Err
	}

	m.O.Body, m.Err = xdr.NewOperationBody(xdr.OperationTypeAllowTrust, m.AT)
	o.TX.Operations = append(o.TX.Operations, m.O)
	return m.Err
}

// MutateTransaction for CreateAccountBuilder causes the underylying
// CreateAccountOp to be added to the operation list for the provided
// transaction
func (m CreateAccountBuilder) MutateTransaction(o *TransactionBuilder) error {
	if m.Err != nil {
		return m.Err
	}

	m.O.Body, m.Err = xdr.NewOperationBody(xdr.OperationTypeCreateAccount, m.CA)
	o.TX.Operations = append(o.TX.Operations, m.O)
	return m.Err
}

// MutateTransaction for Defaults sets reasonable defaults on the transaction being built
func (m Defaults) MutateTransaction(o *TransactionBuilder) error {

	if o.TX.Fee == 0 {
		o.TX.Fee = xdr.Uint32(100 * len(o.TX.Operations))
	}

	if o.NetworkID == [32]byte{} {
		o.NetworkID = DefaultNetwork.ID()
	}
	return nil
}

// MutateTransaction for MemoHash sets the memo.
func (m MemoHash) MutateTransaction(o *TransactionBuilder) (err error) {
	o.TX.Memo, err = xdr.NewMemo(xdr.MemoTypeMemoHash, m.Value)
	return
}

// MutateTransaction for MemoID sets the memo.
func (m MemoID) MutateTransaction(o *TransactionBuilder) (err error) {
	o.TX.Memo, err = xdr.NewMemo(xdr.MemoTypeMemoId, xdr.Uint64(m.Value))
	return
}

// MutateTransaction for MemoReturn sets the memo.
func (m MemoReturn) MutateTransaction(o *TransactionBuilder) (err error) {
	o.TX.Memo, err = xdr.NewMemo(xdr.MemoTypeMemoReturn, m.Value)
	return
}

// MutateTransaction for MemoText sets the memo.
func (m MemoText) MutateTransaction(o *TransactionBuilder) (err error) {

	if len([]byte(m.Value)) > MemoTextMaxLength {
		err = errors.New("Memo too long; over 28 bytes")
		return
	}

	o.TX.Memo, err = xdr.NewMemo(xdr.MemoTypeMemoText, m.Value)
	return
}

// MutateTransaction for Network sets the Network ID to use when signing this transaction
func (m Network) MutateTransaction(o *TransactionBuilder) error {
	o.NetworkID = m.ID()
	return nil
}

// MutateTransaction for PaymentBuilder causes the underylying PaymentOp
// to be added to the operation list for the provided transaction
func (m PathPaymentBuilder) MutateTransaction(o *TransactionBuilder) error {
	if m.Err != nil {
		return m.Err
	}

	m.O.Body, m.Err = xdr.NewOperationBody(xdr.OperationTypePathPayment, m.P)
	o.TX.Operations = append(o.TX.Operations, m.O)
	return m.Err
}

// MutateTransaction for PaymentBuilder causes the underylying PaymentOp
// to be added to the operation list for the provided transaction
func (m PaymentBuilder) MutateTransaction(o *TransactionBuilder) error {
	if m.Err != nil {
		return m.Err
	}

	m.O.Body, m.Err = xdr.NewOperationBody(xdr.OperationTypePayment, m.P)
	o.TX.Operations = append(o.TX.Operations, m.O)
	return m.Err
}

// MutateTransaction for Sequence sets the SeqNum on the transaction.
func (m Sequence) MutateTransaction(o *TransactionBuilder) error {
	o.TX.SeqNum = xdr.SequenceNumber(m.Sequence)
	return nil
}

// MutateTransaction for SourceAccount sets the transaction's SourceAccount
// to the pubilic key for the address provided
func (m SourceAccount) MutateTransaction(o *TransactionBuilder) error {
	return setAccountId(m.AddressOrSeed, &o.TX.SourceAccount)
}
