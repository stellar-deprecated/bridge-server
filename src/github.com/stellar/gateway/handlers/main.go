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

func (rh *RequestHandler) isAssetAllowed(code string) bool {
	for _, b := range rh.Config.Assets {
		if b == code {
			return true
		}
	}
	return false
}

func write(w http.ResponseWriter, response horizon.SubmitTransactionResponse) {
	responseContent := response.Marshal()
	if response.Error != nil {
		http.Error(w, string(responseContent), response.Error.Status)
	} else {
		w.Write(responseContent)
	}
}

func writeError(w http.ResponseWriter, error *horizon.SubmitTransactionResponseError) {
	http.Error(w, getResponseString(error), error.Status)
}


func getResponseString(error *horizon.SubmitTransactionResponseError) string {
	response := horizon.SubmitTransactionResponse{Error: error}
	return string(response.Marshal())
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
