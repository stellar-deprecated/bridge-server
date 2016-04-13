package compliance

import (
	"net/http"
	"net/url"

	"github.com/stellar/gateway/protocols"
)

type FetchInfoRequest struct {
	Address     string `name:"address" required:""`
	formRequest protocols.FormRequest
}

// Will populate request fields using http.Request.
func (request *FetchInfoRequest) FromRequest(r *http.Request) {
	request.formRequest.FromRequest(r, request)
}

// Will create url.Values from request.
func (request *FetchInfoRequest) ToValues() url.Values {
	return request.formRequest.ToValues(request)
}

type FetchInfoResponse struct {
	Name        string `json:"name"`
	Address     string `json:"address"`
	DateOfBirth string `json:"date_of_birth"`
}
