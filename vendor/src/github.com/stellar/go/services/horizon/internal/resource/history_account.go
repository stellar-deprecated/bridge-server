package resource

import (
	"github.com/stellar/go/services/horizon/internal/db2/history"
	"golang.org/x/net/context"
)

func (this *HistoryAccount) Populate(ctx context.Context, row history.Account) {
	this.ID = row.Address
	this.AccountID = row.Address
}
