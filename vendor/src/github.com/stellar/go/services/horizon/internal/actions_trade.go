package horizon

import (
	"errors"

	"github.com/stellar/go/services/horizon/internal/db2"
	"github.com/stellar/go/services/horizon/internal/db2/history"
	"github.com/stellar/go/services/horizon/internal/render/hal"
	"github.com/stellar/go/services/horizon/internal/resource"
	"github.com/stellar/go/xdr"
)

type TradeIndexAction struct {
	Action
	BaseAssetFilter       xdr.Asset
	HasBaseAssetFilter    bool
	CounterAssetFilter    xdr.Asset
	HasCounterAssetFilter bool
	PagingParams          db2.PageQuery
	Records               []history.Trade
	Page                  hal.Page
}

// JSON is a method for actions.JSON
func (action *TradeIndexAction) JSON() {
	action.Do(
		action.EnsureHistoryFreshness,
		action.loadParams,
		action.loadRecords,
		action.loadPage,
		func() {
			hal.Render(action.W, action.Page)
		},
	)
}

// loadParams sets action.Query from the request params
func (action *TradeIndexAction) loadParams() {
	action.PagingParams = action.GetPageQuery()
	action.BaseAssetFilter, action.HasBaseAssetFilter = action.MaybeGetAsset("base_")
	action.CounterAssetFilter, action.HasCounterAssetFilter = action.MaybeGetAsset("counter_")
}

// loadRecords populates action.Records
func (action *TradeIndexAction) loadRecords() {
	trades := action.HistoryQ().Trades()

	if action.HasBaseAssetFilter {

		baseAssetId, err := action.HistoryQ().GetAssetID(action.BaseAssetFilter)
		if err != nil {
			action.Err = err
			return
		}

		if action.HasCounterAssetFilter {

			counterAssetId, err := action.HistoryQ().GetAssetID(action.CounterAssetFilter)
			if err != nil {
				action.Err = err
				return
			}
			trades = action.HistoryQ().TradesForAssetPair(baseAssetId, counterAssetId)
		} else {
			action.Err = errors.New("this endpoint supports asset pairs but only one asset supplied")
			return
		}
	}
	action.Err = trades.Page(action.PagingParams).Select(&action.Records)
}

// loadPage populates action.Page
func (action *TradeIndexAction) loadPage() {
	for _, record := range action.Records {
		var res resource.Trade

		action.Err = res.Populate(action.Ctx, record)
		if action.Err != nil {
			return
		}

		action.Page.Add(res)
	}

	action.Page.BaseURL = action.BaseURL()
	action.Page.BasePath = action.Path()
	action.Page.Limit = action.PagingParams.Limit
	action.Page.Cursor = action.PagingParams.Cursor
	action.Page.Order = action.PagingParams.Order
	action.Page.PopulateLinks()
}
