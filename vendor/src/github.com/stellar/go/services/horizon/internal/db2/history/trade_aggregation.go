package history

import (
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/stellar/go/services/horizon/internal/db2"
	. "github.com/stellar/go/support/time"
)

// Trade aggregation represents an aggregation of trades from the trades table
type TradeAggregation struct {
	Timestamp     int64   `db:"timestamp"`
	TradeCount    int64   `db:"count"`
	BaseVolume    int64   `db:"base_volume"`
	CounterVolume int64   `db:"counter_volume"`
	Average       float64 `db:"avg"`
	High          float64 `db:"high"`
	Low           float64 `db:"low"`
	Open          float64 `db:"open"`
	Close         float64 `db:"close"`
}

// TradeAggregationsQ is a helper struct to aid in configuring queries to
// bucket and aggregate trades
type TradeAggregationsQ struct {
	baseAssetId    int64
	counterAssetId int64
	resolution     int64
	startTime      Millis
	endTime        Millis
	pagingParams   db2.PageQuery
}

// GetTradeAggregationsQ initializes a TradeAggregationsQ query builder based on the required parameters
func (q Q) GetTradeAggregationsQ(baseAssetId int64, counterAssetId int64, resolution int64, pagingParams db2.PageQuery) *TradeAggregationsQ {
	return &TradeAggregationsQ{
		baseAssetId:    baseAssetId,
		counterAssetId: counterAssetId,
		resolution:     resolution,
		pagingParams:   pagingParams,
	}
}

// WithStartTime adds an optional lower time boundary filter to the trades being aggregated
func (q *TradeAggregationsQ) WithStartTime(startTime Millis) *TradeAggregationsQ {
	// Round lower boundary up, if start time is in the middle of a bucket
	q.startTime = startTime.RoundUp(q.resolution)
	return q
}

// WithEndTime adds an upper optional time boundary filter to the trades being aggregated
func (q *TradeAggregationsQ) WithEndTime(endTime Millis) *TradeAggregationsQ {
	// Round upper boundary down, to not deliver partial bucket
	q.endTime = endTime.RoundDown(q.resolution)
	return q
}

// Generate a sql statement to aggregate Trades based on given parameters
func (q *TradeAggregationsQ) GetSql() sq.SelectBuilder {
	var orderPreserved bool
	orderPreserved, q.baseAssetId, q.counterAssetId = getCanonicalAssetOrder(q.baseAssetId, q.counterAssetId)

	var bucketSql sq.SelectBuilder
	if orderPreserved {
		bucketSql = bucketTrades(q.resolution)
	} else {
		bucketSql = reverseBucketTrades(q.resolution)
	}

	bucketSql = bucketSql.From("history_trades").
		Where(sq.Eq{"base_asset_id": q.baseAssetId, "counter_asset_id": q.counterAssetId})

	//adjust time range and apply time filters
	bucketSql = bucketSql.Where(sq.GtOrEq{"ledger_closed_at": q.startTime.ToTime()})
	if !q.endTime.IsNil() {
		bucketSql = bucketSql.Where(sq.Lt{"ledger_closed_at": q.endTime.ToTime()})
	}

	//ensure open/close order for cases when multiple trades occur in the same ledger
	bucketSql = bucketSql.OrderBy("history_operation_id ", "\"order\"")

	return sq.Select(
		"timestamp",
		"count(*) as count",
		"sum(base_amount) as base_volume",
		"sum(counter_amount) as counter_volume",
		"avg(price) as avg",
		"max(price) as high",
		"min(price) as low",
		"first(price) as open",
		"last(price) as close").
		FromSelect(bucketSql, "htrd").
		GroupBy("timestamp").
		Limit(q.pagingParams.Limit).
		OrderBy("timestamp " + q.pagingParams.Order)
}

// formatBucketTimestampSelect formats a sql select clause for a bucketed timestamp, based on given resolution
func formatBucketTimestampSelect(resolution int64) string {
	return fmt.Sprintf("div(cast((extract(epoch from ledger_closed_at) * 1000 ) as bigint), %d)*%d as timestamp",
		resolution, resolution)
}

// bucketTrades generates a select statement to filter rows from the `history_trades` table in
// a compact form, with a timestamp rounded to resolution and reversed base/counter.
func bucketTrades(resolution int64) sq.SelectBuilder {
	return sq.Select(
		formatBucketTimestampSelect(resolution),
		"history_operation_id",
		"\"order\"",
		"base_asset_id",
		"base_amount",
		"counter_asset_id",
		"counter_amount",
		"counter_amount::float/base_amount as price",
	)
}

// reverseBucketTrades generates a select statement to filter rows from the `history_trades` table in
// a compact form, with a timestamp rounded to resolution and reversed base/counter.
func reverseBucketTrades(resolution int64) sq.SelectBuilder {
	return sq.Select(
		formatBucketTimestampSelect(resolution),
		"history_operation_id",
		"\"order\"",
		"counter_asset_id as base_asset_id",
		"counter_amount as base_amount",
		"base_asset_id as counter_asset_id",
		"base_amount as counter_amount",
		"base_amount::float/counter_amount as price",
	)
}
