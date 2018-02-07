package resource

import (
	"github.com/stellar/go/amount"
	"github.com/stellar/go/services/horizon/internal/db2/assets"
	"github.com/stellar/go/services/horizon/internal/render/hal"
	"github.com/stellar/go/xdr"
	"golang.org/x/net/context"
)

// Populate fills out the details
func (res *AssetStat) Populate(
	ctx context.Context,
	row assets.AssetStatsR,
) (err error) {

	res.Asset.Type = row.Type
	res.Asset.Code = row.Code
	res.Asset.Issuer = row.Issuer
	res.Amount = amount.StringFromInt64(row.Amount)
	res.NumAccounts = row.NumAccounts
	res.Flags = AccountFlags{
		(row.Flags & int8(xdr.AccountFlagsAuthRequiredFlag)) != 0,
		(row.Flags & int8(xdr.AccountFlagsAuthRevocableFlag)) != 0,
	}
	res.PT = row.SortKey

	res.Links.Toml = hal.NewLink(row.Toml)
	return
}

// PagingToken implementation for hal.Pageable
func (res AssetStat) PagingToken() string {
	return res.PT
}
