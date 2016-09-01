package stellartoml

import (
	"net/http"
	"testing"

	"github.com/stellar/go/support/http/httptest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestClientURL(t *testing.T) {
	//HACK:  we're testing an internal method rather than setting up a http client
	//mock.

	c := &Client{UseHTTP: false}
	assert.Equal(t, "https://www.stellar.org/.well-known/stellar.toml", c.url("stellar.org"))

	c = &Client{UseHTTP: true}
	assert.Equal(t, "http://www.stellar.org/.well-known/stellar.toml", c.url("stellar.org"))
}

func TestClient(t *testing.T) {
	h := httptest.NewClient()
	c := &Client{HTTP: h}

	// happy path
	h.
		On("GET", "https://www.stellar.org/.well-known/stellar.toml").
		ReturnString(http.StatusOK,
			`FEDERATION_SERVER="https://localhost/federation"`,
		)
	stoml, err := c.GetStellarToml("stellar.org")
	require.NoError(t, err)
	assert.Equal(t, "https://localhost/federation", stoml.FederationServer)

	// not found
	h.
		On("GET", "https://www.missing.org/.well-known/stellar.toml").
		ReturnNotFound()
	stoml, err = c.GetStellarToml("missing.org")
	assert.EqualError(t, err, "http request failed with non-200 status code")

	// invalid toml
	h.
		On("GET", "https://www.json.org/.well-known/stellar.toml").
		ReturnJSON(http.StatusOK, map[string]string{"hello": "world"})
	stoml, err = c.GetStellarToml("json.org")

	if assert.Error(t, err) {
		assert.Contains(t, err.Error(), "toml decode failed")
	}
}
