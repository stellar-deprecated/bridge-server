package compliance

import (
	"net/http"
	"net/url"

	"github.com/stellar/gateway/protocols"
)

// FetchInfoRequest represents a request sent to fetch_info callback
type FetchInfoRequest struct {
	Address     string `name:"address" required:""`
	formRequest protocols.FormRequest
}

// FromRequest will populate request fields using http.Request.
func (request *FetchInfoRequest) FromRequest(r *http.Request) {
	request.formRequest.FromRequest(r, request)
}

// ToValues will create url.Values from request.
func (request *FetchInfoRequest) ToValues() url.Values {
	return request.formRequest.ToValues(request)
}

// FetchInfoResponse represents a response returned by fetch_info callback
type FetchInfoResponse struct {
	Name        string `json:"name"`
	Address     string `json:"address"`
	DateOfBirth string `json:"date_of_birth"`
}
