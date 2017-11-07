package history

import (
	"fmt"
	"math"

	sq "github.com/Masterminds/squirrel"
	"github.com/stellar/go/services/horizon/internal/db2"
)

// PagingToken returns a cursor for this trade
func (r *Trade) PagingToken() string {
	return fmt.Sprintf("%d-%d", r.HistoryOperationID, r.Order)
}

// Trades provides a helper to filter rows from the `history_trades` table
// with pre-defined filters.  See `TradesQ` methods for the available filters.
func (q *Q) Trades() *TradesQ {
	return &TradesQ{
		parent: q,
		sql:    selectTrade,
	}
}

// TradesForAssetPair provides a helper to filter rows from the `history_trades` table
// with the base filter of a specific asset pair.  See `TradesQ` methods for further available filters.
func (q *Q) TradesForAssetPair(baseAssetId int64, counterAssetId int64) *TradesQ {
	var sql sq.SelectBuilder
	if baseAssetId < counterAssetId {
		sql = selectTrade.Where(sq.Eq{"base_asset_id": baseAssetId, "counter_asset_id": counterAssetId})
	} else {
		sql = selectReverseTrade.Where(sq.Eq{"base_asset_id": counterAssetId, "counter_asset_id": baseAssetId})
	}
	return &TradesQ{
		parent: q,
		sql:    sql,
	}
}

// Page specifies the paging constraints for the query being built by `q`.
func (q *TradesQ) Page(page db2.PageQuery) *TradesQ {
	if q.Err != nil {
		return q
	}

	op, idx, err := page.CursorInt64Pair(db2.DefaultPairSep)
	if err != nil {
		q.Err = err
		return q
	}

	// constrain the second portion of the cursor pair to 32-bits
	if idx > math.MaxInt32 {
		idx = math.MaxInt32
	}

	switch page.Order {
	case "asc":
		q.sql = q.sql.
			Where(`(
					 htrd.history_operation_id > ?
				OR (
							htrd.history_operation_id = ?
					AND htrd.order > ?
				))`, op, op, idx).
			OrderBy("htrd.history_operation_id asc, htrd.order asc")
	case "desc":
		q.sql = q.sql.
			Where(`(
					 htrd.history_operation_id < ?
				OR (
							htrd.history_operation_id = ?
					AND htrd.order < ?
				))`, op, op, idx).
			OrderBy("htrd.history_operation_id desc, htrd.order desc")
	}

	q.sql = q.sql.Limit(page.Limit)
	return q
}

// Select loads the results of the query specified by `q` into `dest`.
func (q *TradesQ) Select(dest interface{}) error {
	if q.Err != nil {
		return q.Err
	}

	q.Err = q.parent.Select(dest, q.sql)
	return q.Err
}

var selectTrade = sq.Select(
	"history_operation_id",
	"htrd.\"order\"",
	"htrd.ledger_closed_at",
	"htrd.offer_id",
	"base_accounts.address as base_account",
	"base_assets.asset_type as base_asset_type",
	"base_assets.asset_code as base_asset_code",
	"base_assets.asset_issuer as base_asset_issuer",
	"htrd.base_amount",
	"counter_accounts.address as counter_account",
	"counter_assets.asset_type as counter_asset_type",
	"counter_assets.asset_code as counter_asset_code",
	"counter_assets.asset_issuer as counter_asset_issuer",
	"htrd.counter_amount",
	"htrd.base_is_seller",
).From("history_trades htrd").
	Join("history_accounts base_accounts ON base_account_id = base_accounts.id").
	Join("history_accounts counter_accounts ON counter_account_id = counter_accounts.id").
	Join("history_assets base_assets ON base_asset_id = base_assets.id").
	Join("history_assets counter_assets ON counter_asset_id = counter_assets.id")

var selectReverseTrade = sq.Select(
	"history_operation_id",
	"htrd.\"order\"",
	"htrd.ledger_closed_at",
	"htrd.offer_id",
	"counter_accounts.address as base_account",
	"counter_assets.asset_type as base_asset_type",
	"counter_assets.asset_code as base_asset_code",
	"counter_assets.asset_issuer as base_asset_issuer",
	"htrd.base_amount",
	"base_accounts.address as counter_account",
	"base_assets.asset_type as counter_asset_type",
	"base_assets.asset_code as counter_asset_code",
	"base_assets.asset_issuer as counter_asset_issuer",
	"htrd.counter_amount",
	"NOT(htrd.base_is_seller) as base_is_seller",
).From("history_trades htrd").
	Join("history_accounts base_accounts ON base_account_id = base_accounts.id").
	Join("history_accounts counter_accounts ON counter_account_id = counter_accounts.id").
	Join("history_assets base_assets ON base_asset_id = base_assets.id").
	Join("history_assets counter_assets ON counter_asset_id = counter_assets.id")