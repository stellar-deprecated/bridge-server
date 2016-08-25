// Package horizon provides client access to a horizon server, allowing an
// application to post transactions and lookup ledger information.
//
// Create an instance of `Client` to customize the server used, or alternatively
// use `DefaultTestNetClient` or `DefaultPublicNetClient` to access the SDF run
// horizon servers.
package horizon

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
	"sync"

	"github.com/stellar/go/xdr"
)

// DefaultTestNetClient is a default client to connect to test network
var DefaultTestNetClient = &Client{URL: "https://horizon-testnet.stellar.org"}

// DefaultPublicNetClient is a default client to connect to public network
var DefaultPublicNetClient = &Client{URL: "https://horizon.stellar.org"}

// Error struct contains the problem returned by Horizon
type Error struct {
	Response *http.Response
	Problem  Problem
}

func (herror *Error) Error() string {
	return "Horizon error"
}

type HorizonHttpClient interface {
	Get(url string) (resp *http.Response, err error)
	PostForm(url string, data url.Values) (resp *http.Response, err error)
}

// Client struct contains data required to connect to Horizon instance
type Client struct {
	// URL of Horizon server to connect
	URL string
	// Will be populated with &http.Client when nil. If you want to configure your http.Client make sure Timeout is at least 10 seconds.
	Client HorizonHttpClient
	// clientInit initializes http client once
	clientInit sync.Once
}

// LoadAccount loads the account state from horizon. err can be either error
// object or horizon.Error object.
func (c *Client) LoadAccount(accountID string) (account Account, err error) {
	c.initHttpClient()
	resp, err := c.Client.Get(c.URL + "/accounts/" + accountID)
	if err != nil {
		return
	}

	err = decodeResponse(resp, &account)
	return
}

// SequenceForAccount implements build.SequenceProvider
func (c *Client) SequenceForAccount(
	accountID string,
) (xdr.SequenceNumber, error) {

	a, err := c.LoadAccount(accountID)
	if err != nil {
		return 0, err
	}

	seq, err := strconv.ParseUint(a.Sequence, 10, 64)
	if err != nil {
		return 0, err
	}

	return xdr.SequenceNumber(seq), nil
}

// SubmitTransaction submits a transaction to the network. err can be either error object or horizon.Error object.
func (c *Client) SubmitTransaction(transactionEnvelopeXdr string) (response TransactionSuccess, err error) {
	v := url.Values{}
	v.Set("tx", transactionEnvelopeXdr)

	c.initHttpClient()
	resp, err := c.Client.PostForm(c.URL+"/transactions", v)
	if err != nil {
		return
	}

	err = decodeResponse(resp, &response)
	return
}

func (c *Client) initHttpClient() {
	c.clientInit.Do(func() {
		if c.Client == nil {
			c.Client = &http.Client{}
		}
	})
}

func decodeResponse(resp *http.Response, object interface{}) (err error) {
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)

	if !(resp.StatusCode >= 200 && resp.StatusCode < 300) {
		horizonError := &Error{
			Response: resp,
		}
		decodeError := decoder.Decode(&horizonError.Problem)
		if decodeError != nil {
			return decodeError
		}
		return horizonError
	}

	err = decoder.Decode(&object)
	if err != nil {
		return
	}
	return
}
