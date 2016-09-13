package submitter

import (
	"strconv"

	"github.com/pkg/errors"
	"github.com/stellar/go/clients/horizon"
	"github.com/stellar/go/xdr"
)

// Address returns the accounts address
func (a *Account) Address() string {
	return a.Keypair.Address()
}

// nextSequence increments and returns the next sequence number to use for the
// account
func (a *Account) nextSequence() xdr.SequenceNumber {
	a.Mutex.Lock()
	a.SequenceNumber++
	result := xdr.SequenceNumber(a.SequenceNumber)
	a.Mutex.Unlock()

	return result
}

//
func (a *Account) resetSequence(h *horizon.Client) error {
	a.Mutex.Lock()
	defer a.Mutex.Unlock()

	resp, err := h.LoadAccount(a.Address())
	if err != nil {
		return errors.Wrap(err, "load account failed")
	}

	a.SequenceNumber, err = strconv.ParseUint(resp.Sequence, 10, 64)
	if err != nil {
		return errors.Wrap(err, "parse sequence failed")
	}

	return nil
}
