// Package external contains helper types for external packages.
package external

import (
	"github.com/stellar/go/clients/stellartoml"
	fproto "github.com/stellar/go/protocols/federation"
)

type StellarTomlClientInterface interface {
	GetStellarToml(domain string) (resp *stellartoml.Response, err error)
	GetStellarTomlByAddress(addy string) (*stellartoml.Response, error)
}

type FederationClientInterface interface {
	LookupByAddress(addy string) (*fproto.Response, error)
	LookupByAccountID(aid string) (*fproto.Response, error)
}
