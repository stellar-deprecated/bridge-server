package compliance

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/stellar/gateway/protocols"
)

// ReceiveRequest represents request sent to /receive endpoint of compliance server
type ReceiveRequest struct {
	Memo        string `name:"memo" required:""`
	formRequest protocols.FormRequest
}

// FromRequest will populate request fields using http.Request.
func (request *ReceiveRequest) FromRequest(r *http.Request) error {
	return request.formRequest.FromRequest(r, request)
}

// ToValues will create url.Values from request.
func (request *ReceiveRequest) ToValues() url.Values {
	return request.formRequest.ToValues(request)
}

// Validate validates if request fields are valid. Useful when checking if a request is correct.
func (request *ReceiveRequest) Validate() error {
	err := request.formRequest.CheckRequired(request)
	if err != nil {
		return err
	}
	return nil
}

// ReceiveResponse represents response returned by /receive endpoint
type ReceiveResponse struct {
	protocols.SuccessResponse
	// The AuthData hash of this memo.
	Data string `json:"data"`
}

// Marshal marshals ReceiveResponse
func (response *ReceiveResponse) Marshal() []byte {
	json, _ := json.MarshalIndent(response, "", "  ")
	return json
}
