package horizon

import (
	"github.com/stellar/go/services/horizon/internal/ledger"
	"github.com/stellar/go/services/horizon/internal/render/hal"
	"github.com/stellar/go/services/horizon/internal/resource"
)

// RootAction provides a summary of the horizon instance and links to various
// useful endpoints
type RootAction struct {
	Action
}

// JSON renders the json response for RootAction
func (action *RootAction) JSON() {
	action.App.UpdateStellarCoreInfo()

	var res resource.Root
	res.Populate(
		action.Ctx,
		ledger.CurrentState(),
		action.App.horizonVersion,
		action.App.coreVersion,
		action.App.networkPassphrase,
		action.App.protocolVersion,
	)

	hal.Render(action.W, res)
}
