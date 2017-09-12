package handlers

import (
	"github.com/stellar/gateway/bridge/config"
	"github.com/stellar/gateway/db"
	"github.com/stellar/gateway/external"
	"github.com/stellar/gateway/horizon"
	"github.com/stellar/gateway/listener"
	"github.com/stellar/gateway/net"
	"github.com/stellar/gateway/submitter"
)

// RequestHandler implements bridge server request handlers
type RequestHandler struct {
	Config               *config.Config                          `inject:""`
	Client               net.HTTPClientInterface                 `inject:""`
	Horizon              horizon.HorizonInterface                `inject:""`
	Driver               db.Driver                               `inject:""`
	Repository           db.RepositoryInterface                  `inject:""`
	StellarTomlResolver  external.StellarTomlClientInterface     `inject:""`
	FederationResolver   external.FederationClientInterface      `inject:""`
	TransactionSubmitter submitter.TransactionSubmitterInterface `inject:""`
	PaymentListener      *listener.PaymentListener               `inject:""`
}

func (rh *RequestHandler) isAssetAllowed(code string, issuer string) bool {
	for _, asset := range rh.Config.Assets {
		if asset.Code == code && asset.Issuer == issuer {
			return true
		}
	}
	return false
}
