package httptest

import (
	"net/http"

	"github.com/jarcoal/httpmock"
	"github.com/stellar/go/support/errors"
)

// Return specifies the response for a ClientExpectation, which is then
// committed to the connected mock client.
func (ce *ClientExpectation) Return(r httpmock.Responder) *ClientExpectation {
	ce.Client.MockTransport.RegisterResponder(
		ce.Method,
		ce.URL,
		r,
	)
	return ce
}

// ReturnError causes this expectation to resolve to an error.
func (ce *ClientExpectation) ReturnError(msg string) *ClientExpectation {
	return ce.Return(func(*http.Request) (*http.Response, error) {
		return nil, errors.New(msg)
	})
}

// ReturnString causes this expectation to resolve to a string-based body with
// the provided status code.
func (ce *ClientExpectation) ReturnString(
	status int,
	body string,
) *ClientExpectation {
	return ce.Return(
		httpmock.NewStringResponder(status, body),
	)
}

// ReturnText causes this expectation to resolve to a json-based body with the
// provided status code.  Panics when the provided body cannot be encoded to
// JSON.
func (ce *ClientExpectation) ReturnJSON(
	status int,
	body interface{},
) *ClientExpectation {

	r, err := httpmock.NewJsonResponder(status, body)
	if err != nil {
		panic(err)
	}

	return ce.Return(r)
}

// ReturnNotFound is a simple helper that causes this expectation to resolve to
// a 404 error.  If a customized body is needed, use something like
// `ReturnString`` instead.
func (ce *ClientExpectation) ReturnNotFound() *ClientExpectation {
	return ce.ReturnString(http.StatusNotFound, "not found")
}
