package resource

import (
	"errors"
	"fmt"

	"github.com/stellar/go/services/horizon/internal/db2/history"
	"github.com/stellar/go/services/horizon/internal/httpx"
	"github.com/stellar/go/services/horizon/internal/render/hal"
	"golang.org/x/net/context"
)

// PopulateFromEffect fills out the details of a trade resource from a
// history.Effect row.
func (res *TradeEffect) PopulateFromEffect(
	ctx context.Context,
	row history.Effect,
	ledger history.Ledger,
) (err error) {
	if row.Type != history.EffectTrade {
		err = errors.New("invalid effect; not a trade")
		return
	}

	if row.LedgerSequence() != ledger.Sequence {
		err = errors.New("invalid ledger; different sequence than trade")
		return
	}

	row.UnmarshalDetails(res)
	res.ID = row.PagingToken()
	res.PT = row.PagingToken()
	res.Buyer = row.Account
	res.LedgerCloseTime = ledger.ClosedAt
	res.populateLinks(ctx, res.Seller, res.Buyer, row.HistoryOperationID)

	return
}

// PagingToken implementation for hal.Pageable
func (res TradeEffect) PagingToken() string {
	return res.PT
}

func (res *TradeEffect) populateLinks(
	ctx context.Context,
	seller string,
	buyer string,
	opid int64,
) {
	lb := hal.LinkBuilder{Base: httpx.BaseURL(ctx)}
	res.Links.Seller = lb.Link("/accounts", res.Seller)
	res.Links.Buyer = lb.Link("/accounts", res.Buyer)
	res.Links.Operation = lb.Link(
		"/operations",
		fmt.Sprintf("%d", opid),
	)
}
