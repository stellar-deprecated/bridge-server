package history

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/stellar/go/xdr"
)

func (q *Q) GetAssetByID(dest interface{}, id int64) (err error) {
	sql := sq.Select("id", "asset_type", "asset_code", "asset_issuer").From("history_assets").Limit(1).Where(sq.Eq{"id": id})
	err = q.Get(dest, sql)
	return
}

func (q *Q) GetAssetID(asset xdr.Asset) (id int64, err error) {

	var (
		assetType   string
		assetCode   string
		assetIssuer string
	)

	err = asset.Extract(&assetType, &assetCode, &assetIssuer)
	if err != nil {
		return
	}

	sql := sq.Select("id").From("history_assets").Limit(1).Where(sq.Eq{
		"asset_type":   assetType,
		"asset_code":   assetCode,
		"asset_issuer": assetIssuer})

	err = q.Get(&id, sql)
	return
}
