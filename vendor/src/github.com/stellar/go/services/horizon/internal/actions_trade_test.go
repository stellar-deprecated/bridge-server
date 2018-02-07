package horizon

import (
	"net/url"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/stellar/go/services/horizon/internal/db2/history"
	. "github.com/stellar/go/services/horizon/internal/db2/history"
	"github.com/stellar/go/services/horizon/internal/resource"
	. "github.com/stellar/go/services/horizon/internal/test/trades"
	"github.com/stellar/go/xdr"
)

func TestTradeActions_Index(t *testing.T) {
	ht := StartHTTPTest(t, "trades")
	defer ht.Finish()

	// All trades
	w := ht.Get("/trades")
	if ht.Assert.Equal(200, w.Code) {
		ht.Assert.PageOf(1, w.Body)
	}

	// 	ensure created_at is populated correctly
	records := []resource.Trade{}
	ht.UnmarshalPage(w.Body, &records)

	l := history.Ledger{}
	hq := history.Q{Session: ht.HorizonSession()}
	ht.Require.NoError(hq.LedgerBySequence(&l, 6))

	ht.Assert.WithinDuration(l.ClosedAt, records[0].LedgerCloseTime, 1*time.Second)

	var q = make(url.Values)
	q.Add("base_asset_type", "credit_alphanum4")
	q.Add("base_asset_code", "USD")
	q.Add("base_asset_issuer", "GC23QF2HUE52AMXUFUH3AYJAXXGXXV2VHXYYR6EYXETPKDXZSAW67XO4")
	q.Add("counter_asset_type", "credit_alphanum4")
	q.Add("counter_asset_code", "EUR")
	q.Add("counter_asset_issuer", "GCQPYGH4K57XBDENKKX55KDTWOTK5WDWRQOH2LHEDX3EKVIQRLMESGBG")

	w = ht.GetWithParams("/trades", q)
	if ht.Assert.Equal(200, w.Code) {
		ht.Assert.PageOf(1, w.Body)

		records := []map[string]interface{}{}
		ht.UnmarshalPage(w.Body, &records)

		ht.Assert.Contains(records[0], "base_amount")
		ht.Assert.Contains(records[0], "counter_amount")
	}

	q = make(url.Values)
	q.Add("base_asset_type", "credit_alphanum4")
	q.Add("base_asset_code", "EUR")
	q.Add("base_asset_issuer", "GCQPYGH4K57XBDENKKX55KDTWOTK5WDWRQOH2LHEDX3EKVIQRLMESGBG")
	q.Add("counter_asset_type", "credit_alphanum4")
	q.Add("counter_asset_code", "USD")
	q.Add("counter_asset_issuer", "GC23QF2HUE52AMXUFUH3AYJAXXGXXV2VHXYYR6EYXETPKDXZSAW67XO4")

	w = ht.GetWithParams("/trades", q)
	if ht.Assert.Equal(200, w.Code) {
		ht.Assert.PageOf(1, w.Body)

		records := []map[string]interface{}{}
		ht.UnmarshalPage(w.Body, &records)

		ht.Assert.Contains(records[0], "base_amount")
		ht.Assert.Contains(records[0], "counter_amount")
	}

	// For offer
	w = ht.Get("/offers/1/trades")
	if ht.Assert.Equal(200, w.Code) {
		ht.Assert.PageOf(1, w.Body)
	}

	w = ht.Get("/offers/2/trades")
	if ht.Assert.Equal(200, w.Code) {
		ht.Assert.PageOf(0, w.Body)
	}

	// for an account
	w = ht.Get("/accounts/GA5WBPYA5Y4WAEHXWR2UKO2UO4BUGHUQ74EUPKON2QHV4WRHOIRNKKH2/trades")
	if ht.Assert.Equal(200, w.Code) {
		ht.Assert.PageOf(1, w.Body)
	}

	w = ht.Get("/accounts/GCXKG6RN4ONIEPCMNFB732A436Z5PNDSRLGWK7GBLCMQLIFO4S7EYWVU/trades")
	if ht.Assert.Equal(200, w.Code) {
		ht.Assert.PageOf(1, w.Body)
	}
}

// setAssetQuery adds an asset filter with a given prefix to a query
func setAssetQuery(q *url.Values, prefix string, asset xdr.Asset) {
	var assetType, assetCode, assetFilter string
	asset.Extract(&assetType, &assetCode, &assetFilter)
	q.Add(prefix+"asset_type", assetType)
	q.Add(prefix+"asset_code", assetCode)
	q.Add(prefix+"asset_issuer", assetFilter)
}

func TestTradeActions_Aggregation(t *testing.T) {
	ht := StartHTTPTest(t, "base")
	defer ht.Finish()

	const aggregationPath = "/trade_aggregations"
	const numOfTrades = 10
	const second = 1000
	const minute = 60 * second
	const hour = minute * 60

	//a realistic millis (since epoch) value to start the test from
	//it represents a round hour and is bigger than a max int32
	const start = 1510693200000

	dbQ := &Q{ht.HorizonSession()}
	ass1, ass2, err := PopulateTestTrades(dbQ, start, numOfTrades, minute, 0)
	ht.Require.NoError(err)

	//add other trades as noise, to ensure asset filtering is working
	_, _, err = PopulateTestTrades(dbQ, start, numOfTrades, minute, numOfTrades)
	ht.Require.NoError(err)

	var records []resource.TradeAggregation
	var record resource.TradeAggregation
	var nextLink string

	q := make(url.Values)
	setAssetQuery(&q, "base_", ass1)
	setAssetQuery(&q, "counter_", ass2)

	q.Add("start_time", strconv.FormatInt(start, 10))
	q.Add("end_time", strconv.FormatInt(start+hour, 10))
	q.Add("order", "asc")

	//test one bucket for all trades
	q.Add("resolution", strconv.FormatInt(hour, 10))
	w := ht.GetWithParams(aggregationPath, q)
	if ht.Assert.Equal(200, w.Code) {
		ht.Assert.PageOf(1, w.Body)
		ht.UnmarshalPage(w.Body, &records)
		record = records[0] //Save the single aggregation record for next test
		ht.Assert.Equal("0.0005500", records[0].BaseVolume)
	}

	//test reverse one bucket - make sure values don't change
	q.Set("order", "desc")
	w = ht.GetWithParams(aggregationPath, q)
	if ht.Assert.Equal(200, w.Code) {
		ht.Assert.PageOf(1, w.Body)
		ht.UnmarshalPage(w.Body, &records)
		ht.Assert.Equal(record, records[0])
	}

	//Test bucket per trade
	q.Set("order", "asc")
	q.Set("resolution", strconv.FormatInt(minute, 10))
	w = ht.GetWithParams(aggregationPath, q)
	if ht.Assert.Equal(200, w.Code) {
		if ht.Assert.PageOf(numOfTrades, w.Body) {
			//test that asset filters work
			ht.UnmarshalPage(w.Body, &records)
			ht.Assert.Equal(int64(1), records[0].TradeCount)
			ht.Assert.Equal("0.0000100", records[0].BaseVolume)
			ht.Assert.Equal("1.0000000", records[0].Average)
		}
	}

	//test partial range by modifying endTime to be one minute above half range.
	//half of the results are expected
	endTime := start + (numOfTrades/2)*minute
	q.Set("end_time", strconv.Itoa(endTime))
	w = ht.GetWithParams(aggregationPath, q)
	if ht.Assert.Equal(200, w.Code) {
		ht.Assert.PageOf(numOfTrades/2, w.Body)
	}

	//test that page limit works
	limit := 3
	q.Add("limit", strconv.Itoa(limit))
	w = ht.GetWithParams(aggregationPath, q)
	if ht.Assert.Equal(200, w.Code) {
		ht.Assert.PageOf(limit, w.Body)
	}

	//test that next page delivers the correct amount of records
	w = ht.GetWithParams(aggregationPath, q)
	nextLink = ht.UnmarshalNext(w.Body)
	//make sure the next link is a full url and not just a path
	ht.Assert.Equal(true, strings.HasPrefix(nextLink, "http://localhost"))
	w = ht.Get(nextLink)
	if ht.Assert.Equal(200, w.Code) {
		ht.Assert.PageOf(numOfTrades/2-limit, w.Body)
		ht.UnmarshalPage(w.Body, &records)
		//test for expected value on timestamp of first record on next page
		ht.Assert.Equal(int64(start+limit*minute), records[0].Timestamp)
	}

	//test direction (desc)
	q.Set("order", "desc")
	w = ht.GetWithParams(aggregationPath, q)
	if ht.Assert.Equal(200, w.Code) {
		if ht.Assert.PageOf(limit, w.Body) {
			ht.UnmarshalPage(w.Body, &records)
			ht.Assert.Equal(int64(start+(numOfTrades/2-1)*minute), records[0].Timestamp)
		}
	}

	//test next link desc
	w = ht.GetWithParams(aggregationPath, q)
	nextLink = ht.UnmarshalNext(w.Body)
	w = ht.Get(nextLink)
	if ht.Assert.Equal(200, w.Code) {
		ht.Assert.PageOf(numOfTrades/2-limit, w.Body)
	}

	//test next next link empty
	w = ht.GetWithParams(aggregationPath, q)
	nextLink = ht.UnmarshalNext(w.Body)
	w = ht.Get(nextLink)
	nextLink = ht.UnmarshalNext(w.Body)
	w = ht.Get(nextLink)
	if ht.Assert.Equal(200, w.Code) {
		ht.Assert.PageOf(0, w.Body)
	}
}

func TestTradeActions_IndexRegressions(t *testing.T) {
	ht := StartHTTPTest(t, "trades")
	defer ht.Finish()

	// Regression:  https://github.com/stellar/go/services/horizon/internal/issues/318
	var q = make(url.Values)
	q.Add("base_asset_type", "credit_alphanum4")
	q.Add("base_asset_code", "EUR")
	q.Add("base_asset_issuer", "GCQPYGH4K57XBDENKKX55KDTWOTK5WDWRQOH2LHEDX3EKVIQRLMESGBG")
	q.Add("counter_asset_type", "native")

	w := ht.Get("/trades?" + q.Encode())

	ht.Assert.Equal(404, w.Code) //This used to be 200 with length 0
}

// TestTradeActions_AggregationOrdering checks that open/close aggregation
// fields are correct for multiple trades that occur in the same ledger
// https://github.com/stellar/go/issues/215
func TestTradeActions_AggregationOrdering(t *testing.T) {

	ht := StartHTTPTest(t, "base")
	defer ht.Finish()

	seller := GetTestAccount()
	buyer := GetTestAccount()
	ass1 := GetTestAsset("usd")
	ass2 := GetTestAsset("euro")

	dbQ := &Q{ht.HorizonSession()}
	IngestTestTrade(dbQ, ass1, ass2, seller, buyer, 1, 3, 0, 3)
	IngestTestTrade(dbQ, ass1, ass2, seller, buyer, 1, 1, 0, 1)
	IngestTestTrade(dbQ, ass1, ass2, seller, buyer, 1, 2, 0, 2)

	q := make(url.Values)
	setAssetQuery(&q, "base_", ass1)
	setAssetQuery(&q, "counter_", ass2)

	q.Add("start_time", "0")
	q.Add("end_time", "10")
	q.Add("order", "asc")
	q.Add("resolution", "10")

	var records []resource.TradeAggregation
	w := ht.GetWithParams("/trade_aggregations", q)
	if ht.Assert.Equal(200, w.Code) {
		ht.Assert.PageOf(1, w.Body)
		ht.UnmarshalPage(w.Body, &records)
		ht.Assert.Equal("1.0000000", records[0].Open)
		ht.Assert.Equal("3.0000000", records[0].Close)
	}
}
