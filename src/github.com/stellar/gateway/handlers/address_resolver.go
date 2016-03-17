package handlers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/BurntSushi/toml"
)

type StellarToml struct {
	AuthServer       *string `toml:"AUTH_SERVER"`
	FederationServer *string `toml:"FEDERATION_SERVER"`
}

type StellarDestination struct {
	AccountId string  `json:"account_id"`
	MemoType  *string `json:"memo_type"`
	Memo      *string `json:"memo"`
}

type AddressResolverHelperInterface interface {
	GetStellarToml(domain string) (stellarToml StellarToml, err error)
	GetDestination(federationUrl, address string) (destination StellarDestination, err error)
}

type AddressResolver struct {
	helper AddressResolverHelperInterface
}

func (ar AddressResolver) Resolve(address string) (destination StellarDestination, stellarToml StellarToml, err error) {
	tokens := strings.Split(address, "*")
	if len(tokens) == 1 {
		destination.AccountId = address
	} else if len(tokens) == 2 {
		stellarToml, err = ar.helper.GetStellarToml(tokens[1])
		if err != nil {
			return
		}

		if stellarToml.FederationServer == nil {
			err = errors.New("stellar.toml does not contain FEDERATION_SERVER value")
			return
		}

		destination, err = ar.helper.GetDestination(*stellarToml.FederationServer, address)
		return
	} else {
		err = errors.New("Malformed Stellar address")
	}

	return
}

type AddressResolverHelper struct{}

func (ar AddressResolverHelper) GetStellarToml(domain string) (stellarToml StellarToml, err error) {
	var resp *http.Response
	resp, err = http.Get("https://www." + domain + "/.well-known/stellar.toml")
	if err != nil {
		return
	}
	if resp.StatusCode != 200 {
		err = errors.New("stellar.toml response status code indicates error")
		return
	}

	_, err = toml.DecodeReader(resp.Body, &stellarToml)
	return
}

func (ar AddressResolverHelper) GetDestination(federationUrl, address string) (destination StellarDestination, err error) {
	if !strings.HasPrefix(federationUrl, "https://") {
		err = errors.New("Only HTTPS federation servers allowed")
		return
	}

	resp, err := http.Get(federationUrl + "?type=name&q=" + address)
	if err != nil {
		return
	}
	if resp.StatusCode == 200 {
		err = errors.New("Federation response status code indicates error")
		return
	}

	var bs []byte
	bs, err = ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(bs, &destination)
	if err != nil {
		return
	}

	if (destination.MemoType != nil) && (destination.Memo == nil) {
		err = errors.New("Invalid federation response (memo).")
	}
	return
}
