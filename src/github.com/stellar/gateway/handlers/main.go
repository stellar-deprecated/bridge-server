package handlers

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"

	"github.com/stellar/gateway/config"
	"github.com/stellar/gateway/horizon"
	"github.com/stellar/gateway/submitter"
)

type RequestHandler struct {
	Config               *config.Config
	Horizon              horizon.HorizonInterface
	TransactionSubmitter submitter.TransactionSubmitterInterface
	AddressResolver
}

// TODO this is duplicated in PaymentListener
func (rh *RequestHandler) isAssetAllowed(code string) bool {
	for _, b := range rh.Config.Assets {
		if b == code {
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
