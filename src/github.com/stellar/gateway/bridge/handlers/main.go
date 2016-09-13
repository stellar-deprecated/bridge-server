package handlers

import (
	"github.com/stellar/gateway/bridge/config"
	"github.com/stellar/gateway/net"
	"github.com/stellar/gateway/submitter"
	"github.com/stellar/go/clients/federation"
	"github.com/stellar/go/clients/horizon"
	"github.com/stellar/go/clients/stellartoml"
)

// RequestHandler implements bridge server request handlers
type RequestHandler struct {
	Config               *config.Config                          `inject:""`
	Client               net.HTTPClientInterface                 `inject:""`
	HorizonClient        *horizon.Client                         `inject:""`
	StellarTomlClient    *stellartoml.Client                     `inject:""`
	FederationClient     *federation.Client                      `inject:""`
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
