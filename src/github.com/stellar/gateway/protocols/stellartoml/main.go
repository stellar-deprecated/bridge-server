package stellartoml

import (
	"errors"
	"net/http"
	"strings"

	"github.com/BurntSushi/toml"
)

type ResolverInterface interface {
	GetStellarToml(domain string) (stellarToml StellarToml, err error)
	GetStellarTomlByAddress(address string) (stellarToml StellarToml, err error)
}

type Resolver struct{}

func (r *Resolver) GetStellarToml(domain string) (stellarToml StellarToml, err error) {
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

func (r *Resolver) GetStellarTomlByAddress(address string) (stellarToml StellarToml, err error) {
	// TESTING
	// authServer := "http://localhost:8001"
	// signingKey := "GCXR2UP4RIOADMJAVYXAFCFFLISC65CKY4HZBVTGD4TSGUHMCTFSXW5T"
	// return StellarToml{
	// 	AuthServer: &authServer,
	// 	SigningKey: &signingKey,
	// }, nil

	tokens := strings.Split(address, "*")
	if len(tokens) == 2 {
		stellarToml, err = r.GetStellarToml(tokens[1])
	} else {
		err = errors.New("Malformed Stellar address")
	}
	return
}
