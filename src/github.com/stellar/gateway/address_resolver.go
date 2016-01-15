package gateway

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/stellar/go-stellar-base/xdr"
)

type StellarDestination struct {
	AccountId string
	Memo      interface{}
}

type StellarToml struct {
	FederationServer string `toml:"FEDERATION_SERVER"`
}

type FederationResponse struct {
	StellarAddress string `json:"stellar_address"`
	AccountId      string `json:"account_id"`
	MemoType       string `json:"memo_type"`
	Memo           string `json:"memo"`
}

func resolveAddress(address string) (dest StellarDestination, err error) {
	// look for the '*'
	tokens := strings.Split(address, "*")
	if len(tokens) == 1 {
		dest.AccountId = address
	} else if len(tokens) == 2 {
		// find stellar.toml
		// ask the federation server
		var resp *http.Response
		resp, err = http.Get("https://" + tokens[2] + "/.well-known/stellar.toml")
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
				var m FederationResponse
				var bs []byte
				bs, err = ioutil.ReadAll(resp.Body)
				err = json.Unmarshal(bs, &m)
				if err != nil {
					return
				}
				dest.AccountId = m.AccountId

				var memoType xdr.MemoType
				var memoValue interface{}
				switch m.MemoType {
				case "text":
					memoType = xdr.MemoTypeMemoText
					memoValue = m.Memo
				case "id":
					var v uint64
					v, err = strconv.ParseUint(m.Memo, 10, 64)
					if err != nil {
						return
					}
					memoType = xdr.MemoTypeMemoId
					memoValue = v
				case "hash":
					memoType = xdr.MemoTypeMemoHash
					memoValue = m.Memo
				}

				dest.Memo, err = xdr.NewMemo(memoType, memoValue)
				if err != nil {
					return
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
