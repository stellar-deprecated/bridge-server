package federation

import (
	"errors"
	"net/http"
	"testing"

	"github.com/stellar/go/clients/stellartoml"
	"github.com/stellar/go/support/http/httptest"
	"github.com/stretchr/testify/assert"
)

func TestLookupByAddress(t *testing.T) {
	hmock := httptest.NewClient()
	tomlmock := &stellartoml.MockClient{}
	c := &Client{StellarTOML: tomlmock, HTTP: hmock}

	// happy path
	tomlmock.On("GetStellarToml", "stellar.org").Return(&stellartoml.Response{
		FederationServer: "https://stellar.org/federation",
	}, nil)
	hmock.On("GET", "https://stellar.org/federation").
		ReturnJSON(http.StatusOK, map[string]string{
			"stellar_address": "scott*stellar.org",
			"account_id":      "GASTNVNLHVR3NFO3QACMHCJT3JUSIV4NBXDHDO4VTPDTNN65W3B2766C",
			"memo_type":       "id",
			"memo":            "123",
		})
	resp, err := c.LookupByAddress("scott*stellar.org")

	if assert.NoError(t, err) {
		assert.Equal(t, "GASTNVNLHVR3NFO3QACMHCJT3JUSIV4NBXDHDO4VTPDTNN65W3B2766C", resp.AccountID)
		assert.Equal(t, "id", resp.MemoType)
		assert.Equal(t, "123", resp.Memo)
	}

	// failed toml resolution
	tomlmock.On("GetStellarToml", "missing.org").Return(
		(*stellartoml.Response)(nil),
		errors.New("toml failed"),
	)
	resp, err = c.LookupByAddress("scott*missing.org")
	if assert.Error(t, err) {
		assert.Contains(t, err.Error(), "toml failed")
	}

	// 404 federation response
	tomlmock.On("GetStellarToml", "404.org").Return(&stellartoml.Response{
		FederationServer: "https://404.org/federation",
	}, nil)
	hmock.On("GET", "https://404.org/federation").ReturnNotFound()
	resp, err = c.LookupByAddress("scott*404.org")
	if assert.Error(t, err) {
		assert.Contains(t, err.Error(), "failed with (404)")
	}

	// connection error on federation response
	tomlmock.On("GetStellarToml", "error.org").Return(&stellartoml.Response{
		FederationServer: "https://error.org/federation",
	}, nil)
	hmock.On("GET", "https://error.org/federation").ReturnError("kaboom!")
	resp, err = c.LookupByAddress("scott*error.org")
	if assert.Error(t, err) {
		assert.Contains(t, err.Error(), "kaboom!")
	}
}

func TestLookupByID(t *testing.T) {
	// HACK: until we improve our mocking scenario, this is just a smoke test.
	// When/if it breaks, please write this test correctly.  That, or curse
	// scott's name aloud.

	// an account without a homedomain set fails
	_, err := DefaultPublicNetClient.LookupByAccountID("GASTNVNLHVR3NFO3QACMHCJT3JUSIV4NBXDHDO4VTPDTNN65W3B2766C")
	assert.Error(t, err)
	assert.Equal(t, "homedomain not set", err.Error())
}
