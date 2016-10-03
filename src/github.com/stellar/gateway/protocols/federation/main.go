package federation

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/stellar/gateway/protocols/stellartoml"
)

// ResolverInterface helps mocking Resolver object
type ResolverInterface interface {
	Resolve(address string) (response Response, stellarToml stellartoml.StellarToml, err error)
	GetDestination(federationURL, address string) (response Response, err error)
}

// Resolver resolves federation query
type Resolver struct {
	StellarTomlResolver *stellartoml.Resolver `inject:""`
}

// Resolve resolves federation address or account ID.
func (r *Resolver) Resolve(address string) (response Response, stellarToml stellartoml.StellarToml, err error) {
	tokens := strings.Split(address, "*")
	if len(tokens) == 1 {
		response.AccountID = address
	} else if len(tokens) == 2 {
		stellarToml, err = r.StellarTomlResolver.GetStellarToml(tokens[1])
		if err != nil {
			return
		}

		if stellarToml.FederationServer == "" {
			err = errors.New("stellar.toml does not contain FEDERATION_SERVER value")
			return
		}

		response, err = r.GetDestination(stellarToml.FederationServer, address)
		return
	} else {
		err = errors.New("Malformed Stellar address")
	}

	return
}

// GetDestination resolves federation address using server specified federationURL
func (r *Resolver) GetDestination(federationURL, address string) (response Response, err error) {
	if !strings.HasPrefix(federationURL, "https://") {
		err = errors.New("Only HTTPS federation servers allowed")
		return
	}

	qstr := url.Values{}
	qstr.Add("type", "name")
	qstr.Add("q", address)

	resp, err := http.Get(federationURL + "?" + qstr.Encode())
	if err != nil {
		return
	}
	if !(resp.StatusCode >= 200 && resp.StatusCode < 300) {
		err = errors.New("Federation response status code (" + strconv.Itoa(resp.StatusCode) + ") indicates error")
		return
	}

	var bs []byte
	bs, err = ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(bs, &response)
	if err != nil {
		return
	}

	if (response.MemoType != "") && (response.Memo == "") {
		err = errors.New("Invalid federation response (memo).")
	}
	return
}
