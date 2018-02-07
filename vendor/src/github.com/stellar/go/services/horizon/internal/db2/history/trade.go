package history

import (
	"fmt"
	"math"

	sq "github.com/Masterminds/squirrel"
	"github.com/stellar/go/services/horizon/internal/db2"
	"github.com/stellar/go/support/errors"
	"github.com/stellar/go/support/time"
	"github.com/stellar/go/xdr"
)

// PagingToken returns a cursor for this trade
func (r *Trade) PagingToken() string {
	return fmt.Sprintf("%d-%d", r.HistoryOperationID, r.Order)
}

// Trades provides a helper to filter rows from the `history_trades` table
// with pre-defined filters.  See `TradesQ` methods for the available filters.
func (q *Q) Trades() *TradesQ {
	trades := &TradesQ{
		parent: q,
		sql:    selectTrade,
	}
	return trades.JoinAccounts().JoinAssets()
}

// ReverseTrades provides a helper to filter rows from the `history_trades` table
// with pre-defined filters and reversed base/counter.  See `TradesQ` methods for the available filters.
func (q *Q) ReverseTrades() *TradesQ {
	trades := &TradesQ{
		parent: q,
		sql:    selectReverseTrade,
	}
	return trades.JoinAccounts().JoinAssets()
}

// TradesForAssetPair provides a helper to filter rows from the `history_trades` table
// with the base filter of a specific asset pair.  See `TradesQ` methods for further available filters.
func (q *Q) TradesForAssetPair(baseAssetId int64, counterAssetId int64) *TradesQ {
	orderPreserved, baseAssetId, counterAssetId := getCanonicalAssetOrder(baseAssetId, counterAssetId)
	var trades *TradesQ
	if orderPreserved {
		trades = q.Trades()
	} else {
		trades = q.ReverseTrades()
	}
	return trades.forAssetPair(baseAssetId, counterAssetId)
}

// ForOffer filters the query results by the offer id.
func (q *TradesQ) ForOffer(id int64) *TradesQ {
	q.sql = q.sql.Where("htrd.offer_id = ?", id)
	return q
}

//Filter by asset pair. This function is private to ensure that correct order and proper select statement are coupled
func (q *TradesQ) forAssetPair(baseAssetId int64, counterAssetId int64) *TradesQ {
	q.sql = q.sql.Where(sq.Eq{"base_asset_id": baseAssetId, "counter_asset_id": counterAssetId})
	return q
}

func (q *TradesQ) OrderBy(order string) *TradesQ {
	q.sql = q.sql.OrderBy("ledger_closed_at " + order)
	return q
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

func (q *TradesQ) JoinAccounts() *TradesQ {
	q.sql = q.sql.
		Join("history_accounts base_accounts ON base_account_id = base_accounts.id").
		Join("history_accounts counter_accounts ON counter_account_id = counter_accounts.id")
	return q
}

func (q *TradesQ) JoinAssets() *TradesQ {
	q.sql = q.sql.
		Join("history_assets base_assets ON base_asset_id = base_assets.id").
		Join("history_assets counter_assets ON counter_asset_id = counter_assets.id")
	return q
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
).From("history_trades htrd")

var selectReverseTrade = sq.Select(
	"history_operation_id",
	"htrd.\"order\"",
	"htrd.ledger_closed_at",
	"htrd.offer_id",
	"counter_accounts.address as base_account",
	"counter_assets.asset_type as base_asset_type",
	"counter_assets.asset_code as base_asset_code",
	"counter_assets.asset_issuer as base_asset_issuer",
	"htrd.counter_amount as base_amount",
	"base_accounts.address as counter_account",
	"base_assets.asset_type as counter_asset_type",
	"base_assets.asset_code as counter_asset_code",
	"base_assets.asset_issuer as counter_asset_issuer",
	"htrd.base_amount as counter_amount",
	"NOT(htrd.base_is_seller) as base_is_seller",
).From("history_trades htrd")

var tradesInsert = sq.Insert("history_trades").Columns(
	"history_operation_id",
	"\"order\"",
	"ledger_closed_at",
	"offer_id",
	"base_account_id",
	"base_asset_id",
	"base_amount",
	"counter_account_id",
	"counter_asset_id",
	"counter_amount",
	"base_is_seller",
)

// Trade records a trade into the history_trades table
func (q *Q) InsertTrade(
	opid int64,
	order int32,
	buyer xdr.AccountId,
	trade xdr.ClaimOfferAtom,
	ledgerClosedAt time.Millis,
) error {
	sellerAccountId, err := q.GetCreateAccountID(trade.SellerId)
	if err != nil {
		return errors.Wrap(err, "failed to load seller account id")
	}

	buyerAccountId, err := q.GetCreateAccountID(buyer)
	if err != nil {
		return errors.Wrap(err, "failed to load buyer account id")
	}

	soldAssetId, err := q.GetCreateAssetID(trade.AssetSold)
	if err != nil {
		return errors.Wrap(err, "failed to get sold asset id")
	}

	boughtAssetId, err := q.GetCreateAssetID(trade.AssetBought)
	if err != nil {
		return errors.Wrap(err, "failed to get bought asset id")
	}

	orderPreserved, baseAssetId, counterAssetId := getCanonicalAssetOrder(soldAssetId, boughtAssetId)

	var baseAccountId, counterAccountId int64
	var baseAmount, counterAmount xdr.Int64
	if orderPreserved {
		baseAccountId, baseAmount, counterAccountId, counterAmount =
			sellerAccountId, trade.AmountSold, buyerAccountId, trade.AmountBought
	} else {
		baseAccountId, baseAmount, counterAccountId, counterAmount =
			buyerAccountId, trade.AmountBought, sellerAccountId, trade.AmountSold
	}

	sql := tradesInsert.Values(
		opid,
		order,
		ledgerClosedAt.ToTime(),
		trade.OfferId,
		baseAccountId,
		baseAssetId,
		baseAmount,
		counterAccountId,
		counterAssetId,
		counterAmount,
		orderPreserved,
	)
	_, err = q.Exec(sql)
	if err != nil {
		return errors.Wrap(err, "failed to exec sql")
	}

	return nil
}

func getCanonicalAssetOrder(assetId1 int64, assetId2 int64) (orderPreserved bool, baseAssetId int64, counterAssetId int64) {
	if assetId1 < assetId2 {
		return true, assetId1, assetId2
	} else {
		return false, assetId2, assetId1
	}
}
