// Package external contains helper types for external packages.
package external

import (
	"github.com/stellar/go/clients/stellartoml"
)

type StellarTomlClientInterface interface {
	GetStellarToml(domain string) (resp *stellartoml.Response, err error)
	GetStellarTomlByAddress(addy string) (*stellartoml.Response, error)
}
