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
	FederationServer string `toml:"FEDERATION_SERVER"`
}

type StellarDestination struct {
	AccountId string  `json:"account_id"`
	MemoType  *string `json:"memo_type"`
	Memo      *string `json:"memo"`
}

func ResolveAddress(address string) (federation StellarDestination, err error) {
	// look for the '*'
	tokens := strings.Split(address, "*")
	if len(tokens) == 1 {
		federation.AccountId = address
	} else if len(tokens) == 2 {
		// find stellar.toml
		// ask the federation server
		var resp *http.Response
		resp, err = http.Get("https://www." + tokens[2] + "/.well-known/stellar.toml")
		if err != nil {
			return
		}
		if resp.StatusCode == 200 {
			var stellarToml StellarToml
			_, err = toml.DecodeReader(resp.Body, &stellarToml)
			if err != nil {
				return
			}

			// TODO check if stellarToml.FEDERATION_SERVER is https server
			resp, err = http.Get(stellarToml.FederationServer + "?type=name&q=" + address)
			if err != nil {
				return
			}
			if resp.StatusCode == 200 {
				var bs []byte
				bs, err = ioutil.ReadAll(resp.Body)
				err = json.Unmarshal(bs, &federation)
				if err != nil {
					return
				}

				if (federation.MemoType != nil) && (federation.Memo == nil) {
					err = errors.New("Invalid federation response (memo).")
				}

			} else { // fetching the name from the federation server failed
				err = errors.New(resp.Status)
			}
		} else { // fetching the stellar.toml failed
			err = errors.New(resp.Status)
		}
	} else {
		err = errors.New("malformed address")
	}

	return
}
