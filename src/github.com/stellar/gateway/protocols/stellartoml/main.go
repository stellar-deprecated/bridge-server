package stellartoml

import (
	"errors"
	"net/http"
	"strings"

	"github.com/BurntSushi/toml"
)

// ResolverInterface helps mocking Resolver
type ResolverInterface interface {
	GetStellarToml(domain string) (stellarToml StellarToml, err error)
	GetStellarTomlByAddress(address string) (stellarToml StellarToml, err error)
}

// Resolver resolves stellar.toml file
type Resolver struct{}

// GetStellarToml returns stellar.toml file for a given domain
func (r *Resolver) GetStellarToml(domain string) (stellarToml StellarToml, err error) {
	var resp *http.Response
	resp, err = http.Get("https://" + domain + "/.well-known/stellar.toml")
	if err != nil {
		return
	}
	if resp.StatusCode != 200 {
		err = errors.Errorf(
			"stellar.toml response status code indicates error (%d)",
			resp.StatusCode,
		)
		return
	}

	_, err = toml.DecodeReader(resp.Body, &stellarToml)
	return
}

// GetStellarTomlByAddress returns stellar.toml file of a domain fetched from a given address
func (r *Resolver) GetStellarTomlByAddress(address string) (stellarToml StellarToml, err error) {
	tokens := strings.Split(address, "*")
	if len(tokens) == 2 {
		stellarToml, err = r.GetStellarToml(tokens[1])
	} else {
		err = errors.New("Malformed Stellar address")
	}
	return
}
