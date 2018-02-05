package bitcoin

import (
	"math/big"

	"github.com/stellar/go/services/bifrost/common"
)

func (t Transaction) ValueToStellar() string {
	valueSat := new(big.Int).SetInt64(t.ValueSat)
	valueBtc := new(big.Rat).Quo(new(big.Rat).SetInt(valueSat), satInBtc)
	return valueBtc.FloatString(common.StellarAmountPrecision)
}
