package horizon

import (
	"net/url"
	"testing"
	"time"

	"github.com/stellar/go/services/horizon/internal/db2/history"
	"github.com/stellar/go/services/horizon/internal/resource"
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

	//
	var q = make(url.Values)
	q.Add("base_asset_type", "credit_alphanum4")
	q.Add("base_asset_code", "USD")
	q.Add("base_asset_issuer", "GC23QF2HUE52AMXUFUH3AYJAXXGXXV2VHXYYR6EYXETPKDXZSAW67XO4")
	q.Add("counter_asset_type", "credit_alphanum4")
	q.Add("counter_asset_code", "EUR")
	q.Add("counter_asset_issuer", "GCQPYGH4K57XBDENKKX55KDTWOTK5WDWRQOH2LHEDX3EKVIQRLMESGBG")


	w = ht.Get("/trades?" + q.Encode())
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

	w = ht.Get("/trades?" + q.Encode())
	if ht.Assert.Equal(200, w.Code) {
		ht.Assert.PageOf(1, w.Body)

		records := []map[string]interface{}{}
		ht.UnmarshalPage(w.Body, &records)

		ht.Assert.Contains(records[0], "base_amount")
		ht.Assert.Contains(records[0], "counter_amount")
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

	ht.Assert.Equal(404, w.Code)

	// Old version
	//if ht.Assert.Equal(200, w.Code) {
	//	ht.Assert.PageOf(0, w.Body)
	//}
}
