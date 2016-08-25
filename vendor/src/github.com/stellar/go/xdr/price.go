package xdr

import (
	"math/big"
)

// String returns a string represenation of `p`
func (p *Price) String() string {
	return big.NewRat(int64(p.N), int64(p.D)).FloatString(7)
}
