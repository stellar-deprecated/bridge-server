package stellartoml

import (
	"fmt"
	"net/http"

	"github.com/BurntSushi/toml"
	"github.com/stellar/go/address"
	"github.com/stellar/go/support/errors"
)

// GetStellarToml returns stellar.toml file for a given domain
func (c *Client) GetStellarToml(domain string) (resp *Response, err error) {
	var hresp *http.Response
	hresp, err = c.HTTP.Get(c.url(domain))
	if err != nil {
		err = errors.Wrap(err, "http request errored")
		return
	}
	defer hresp.Body.Close()

	if !(hresp.StatusCode >= 200 && hresp.StatusCode < 300) {
		err = errors.New("http request failed with non-200 status code")
		return
	}

	_, err = toml.DecodeReader(hresp.Body, &resp)
	if err != nil {
		err = errors.Wrap(err, "toml decode failed")
		return
	}

	return
}

// GetStellarTomlByAddress returns stellar.toml file of a domain fetched from a
// given address
func (c *Client) GetStellarTomlByAddress(addy string) (*Response, error) {
	_, domain, err := address.Split(addy)
	if err != nil {
		return nil, errors.Wrap(err, "parse address failed")
	}

	return c.GetStellarToml(domain)
}

// url returns the appropriate url to load for resolving domain's stellar.toml
// file
func (c *Client) url(domain string) string {
	var scheme string

	if c.UseHTTP {
		scheme = "http"
	} else {
		scheme = "https"
	}

	return fmt.Sprintf("%s://www.%s%s", scheme, domain, WellKnownPath)
}
