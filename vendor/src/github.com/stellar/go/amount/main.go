// Package amount provides utilities for converting numbers to/from
// the format used internally to stellar-core.
//
// stellar-core represents asset "amounts" as 64-bit integers, but to enable
// fractional units of an asset, horizon, the client-libraries and other built
// on top of stellar-core use a convention, encoding amounts as a string of
// decimal digits with up to seven digits of precision in the fractional
// portion. For example, an amount shown as "101.001" in horizon would be
// represented in stellar-core as 1010010000.
package amount

import (
	"fmt"
	"math/big"
	"strconv"

	"github.com/stellar/go/xdr"
)

// One is the value of one whole unit of currency.  Stellar uses 7 fixed digits
// for fractional values, thus One is 10 million (10^7)
const One = 10000000

// MustParse is the panicking version of Parse
func MustParse(v string) xdr.Int64 {
	ret, err := Parse(v)
	if err != nil {
		panic(err)
	}
	return ret
}

// Parse parses the provided as a stellar "amount", i.e. A 64-bit signed integer
// that represents a decimal number with 7 digits of significance in the
// fractional portion of the number.
func Parse(v string) (xdr.Int64, error) {
	var f, o, r big.Rat

	_, ok := f.SetString(v)
	if !ok {
		return xdr.Int64(0), fmt.Errorf("cannot parse amount: %s", v)
	}

	o.SetInt64(One)
	r.Mul(&f, &o)

	is := r.FloatString(0)
	i, err := strconv.ParseInt(is, 10, 64)
	if err != nil {
		return xdr.Int64(0), err
	}
	return xdr.Int64(i), nil
}

// String returns an "amount string" from the provided raw xdr.Int64 value `v`.
func String(v xdr.Int64) string {
	return StringFromInt64(int64(v))
}

// StringFromInt64 returns an "amount string" from the provided raw int64 value `v`.
func StringFromInt64(v int64) string {
	var f, o, r big.Rat
	f.SetInt64(v)
	o.SetInt64(One)
	r.Quo(&f, &o)
	return r.FloatString(7)
}
