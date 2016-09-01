// Package horizon provides client access to a horizon server, allowing an
// application to post transactions and lookup ledger information.
//
// Create an instance of `Client` to customize the server used, or alternatively
// use `DefaultTestNetClient` or `DefaultPublicNetClient` to access the SDF run
// horizon servers.
package horizon

import (
	"net/http"
	"net/url"

	"github.com/stellar/go/build"
)

// DefaultTestNetClient is a default client to connect to test network
var DefaultTestNetClient = &Client{
	URL:  "https://horizon-testnet.stellar.org",
	HTTP: http.DefaultClient,
}

// DefaultPublicNetClient is a default client to connect to public network
var DefaultPublicNetClient = &Client{
	URL:  "https://horizon.stellar.org",
	HTTP: http.DefaultClient,
}

// Client struct contains data required to connect to Horizon instance
type Client struct {
	// URL of Horizon server to connect
	URL string

	// HTTP client to make requests with
	HTTP HTTP
}

// Error struct contains the problem returned by Horizon
type Error struct {
	Response *http.Response
	Problem  Problem
}

// HTTP represents the HTTP client that a horizon client uses to communicate
type HTTP interface {
	Do(req *http.Request) (resp *http.Response, err error)
	Get(url string) (resp *http.Response, err error)
	PostForm(url string, data url.Values) (resp *http.Response, err error)
}

// PaymentHandler is a function that is called when a new payment is received
type PaymentHandler func(PaymentResponse)

// ensure that the horizon client can be used as a SequenceProvider
var _ build.SequenceProvider = &Client{}
