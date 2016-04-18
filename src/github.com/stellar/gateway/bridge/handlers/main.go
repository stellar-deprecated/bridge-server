package handlers

import (
	"github.com/stellar/gateway/bridge/config"
	"github.com/stellar/gateway/horizon"
	"github.com/stellar/gateway/net"
	"github.com/stellar/gateway/protocols/federation"
	"github.com/stellar/gateway/protocols/stellartoml"
	"github.com/stellar/gateway/submitter"
)

// RequestHandler implements bridge server request handlers
type RequestHandler struct {
	Config               *config.Config                          `inject:""`
	Client               net.HTTPClientInterface                 `inject:""`
	Horizon              horizon.HorizonInterface                `inject:""`
	StellarTomlResolver  stellartoml.ResolverInterface           `inject:""`
	FederationResolver   federation.ResolverInterface            `inject:""`
	TransactionSubmitter submitter.TransactionSubmitterInterface `inject:""`
}

func (rh *RequestHandler) isAssetAllowed(code string, issuer string) bool {
	for _, asset := range rh.Config.Assets {
		if asset.Code == code && asset.Issuer == issuer {
			return true
		}
	}
	return false
}
