package handlers

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"

	"github.com/stellar/gateway/bridge/config"
	"github.com/stellar/gateway/horizon"
	"github.com/stellar/gateway/net"
	"github.com/stellar/gateway/protocols/federation"
	"github.com/stellar/gateway/protocols/stellartoml"
	"github.com/stellar/gateway/submitter"
)

type RequestHandler struct {
	Config               *config.Config                          `inject:""`
	Client               net.HttpClientInterface                 `inject:""`
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

// Used in tests
func getResponse(testServer *httptest.Server, values url.Values) (int, []byte) {
	res, err := http.PostForm(testServer.URL, values)
	if err != nil {
		panic(err)
	}
	response, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		panic(err)
	}
	return res.StatusCode, response
}
