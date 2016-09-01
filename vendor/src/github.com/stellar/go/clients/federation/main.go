package federation

import (
	"net/http"

	"github.com/stellar/go/clients/horizon"
	"github.com/stellar/go/clients/stellartoml"
)

// DefaultTestNetClient is a default federation client for testnet
var DefaultTestNetClient = &Client{
	HTTP:        http.DefaultClient,
	Horizon:     horizon.DefaultTestNetClient,
	StellarTOML: stellartoml.DefaultClient,
}

// DefaultPublicNetClient is a default federation client for oubnet
var DefaultPublicNetClient = &Client{
	HTTP:        http.DefaultClient,
	Horizon:     horizon.DefaultPublicNetClient,
	StellarTOML: stellartoml.DefaultClient,
}

// Client represents a client that is capable of resolving a Stellar.toml file
// using the internet.
type Client struct {
	StellarTOML StellarTOML
	HTTP        HTTP
	Horizon     Horizon
	AllowHTTP   bool
}

// Horizon represents a horizon client that can be consulted for data when
// needed as part of the federation protocol
type Horizon interface {
	HomeDomainForAccount(aid string) (string, error)
}

// HTTP represents the http client that a federation client uses to make http
// requests.
type HTTP interface {
	Get(url string) (*http.Response, error)
}

// IDResponse represents the result of a federation request whose type is
// "id", i.e.  A reverse federation request.
type IDResponse struct {
	Address string `json:"stellar_address"`
}

// NameResponse represents the result of a federation request whose type is
// "name", i.e.  A forward federation request.
type NameResponse struct {
	AccountID string `json:"account_id"`
	MemoType  string `json:"memo_type"`
	Memo      string `json:"memo"`
}

// StellarTOML represents a client that can resolve a given domain name to
// stellar.toml file.  The response is used to find the federation server that a
// query should be made against.
type StellarTOML interface {
	GetStellarToml(domain string) (*stellartoml.Response, error)
}

// confirm interface conformity
var _ StellarTOML = stellartoml.DefaultClient
var _ HTTP = http.DefaultClient
var _ Horizon = horizon.DefaultTestNetClient
