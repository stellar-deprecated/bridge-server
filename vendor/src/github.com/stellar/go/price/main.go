package price

import (
	"errors"
	"fmt"
	"math"
	"math/big"

	"github.com/stellar/go/xdr"
)

// Parse  calculates and returns the best rational approximation of the given real number price.
func Parse(v string) (xdr.Price, error) {
	return continuedFraction(v)
}

// continuedFraction calculates and returns the best rational approximation of the given real number.
func continuedFraction(price string) (xdrPrice xdr.Price, err error) {
	number := &big.Rat{}
	maxInt32 := &big.Rat{}
	zero := &big.Rat{}
	one := &big.Rat{}

	_, ok := number.SetString(price)
	if !ok {
		return xdrPrice, fmt.Errorf("cannot parse price: %s", price)
	}

	maxInt32.SetInt64(int64(math.MaxInt32))
	zero.SetInt64(int64(0))
	one.SetInt64(int64(1))

	fractions := [][2]*big.Rat{
		{zero, one},
		{one, zero},
	}

	i := 2
	for {
		if number.Cmp(maxInt32) == 1 {
			break
		}

		f := &big.Rat{}
		h := &big.Rat{}
		k := &big.Rat{}

		a := floor(number)
		f.Sub(number, a)
		h.Mul(a, fractions[i-1][0])
		h.Add(h, fractions[i-2][0])
		k.Mul(a, fractions[i-1][1])
		k.Add(k, fractions[i-2][1])

		if h.Cmp(maxInt32) == 1 || k.Cmp(maxInt32) == 1 {
			break
		}

		fractions = append(fractions, [2]*big.Rat{h, k})
		if f.Cmp(zero) == 0 {
			break
		}
		number.Quo(one, f)
		i++
	}

	n, d := fractions[len(fractions)-1][0], fractions[len(fractions)-1][1]

	if n.Cmp(zero) == 0 || d.Cmp(zero) == 0 {
		return xdrPrice, errors.New("Couldn't find approximation")
	}

	return xdr.Price{
		N: xdr.Int32(n.Num().Int64()),
		D: xdr.Int32(d.Num().Int64()),
	}, nil
}

func floor(n *big.Rat) *big.Rat {
	f := &big.Rat{}
	z := new(big.Int)
	z.Div(n.Num(), n.Denom())
	f.SetInt(z)
	return f
}
