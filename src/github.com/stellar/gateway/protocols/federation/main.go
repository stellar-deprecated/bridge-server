package federation

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/stellar/gateway/protocols/stellartoml"
)

type ResolverInterface interface {
	Resolve(address string) (response Response, stellarToml stellartoml.StellarToml, err error)
	GetDestination(federationUrl, address string) (response Response, err error)
}

type Resolver struct {
	StellarTomlResolver *stellartoml.Resolver `inject:""`
}

func (r *Resolver) Resolve(address string) (response Response, stellarToml stellartoml.StellarToml, err error) {
	tokens := strings.Split(address, "*")
	if len(tokens) == 1 {
		response.AccountId = address
	} else if len(tokens) == 2 {
		stellarToml, err = r.StellarTomlResolver.GetStellarToml(tokens[1])
		if err != nil {
			return
		}

		if stellarToml.FederationServer == nil {
			err = errors.New("stellar.toml does not contain FEDERATION_SERVER value")
			return
		}

		response, err = r.GetDestination(*stellarToml.FederationServer, address)
		return
	} else {
		err = errors.New("Malformed Stellar address")
	}

	return
}

func (r *Resolver) GetDestination(federationUrl, address string) (response Response, err error) {
	if !strings.HasPrefix(federationUrl, "https://") {
		err = errors.New("Only HTTPS federation servers allowed")
		return
	}

	resp, err := http.Get(federationUrl + "?type=name&q=" + address)
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

	if (response.MemoType != nil) && (response.Memo == nil) {
		err = errors.New("Invalid federation response (memo).")
	}
	return
}
