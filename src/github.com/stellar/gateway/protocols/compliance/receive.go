package compliance

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/stellar/gateway/protocols"
)

type ReceiveRequest struct {
	Memo        string `name:"memo" required:""`
	formRequest protocols.FormRequest
}

// Will populate request fields using http.Request.
func (request *ReceiveRequest) FromRequest(r *http.Request) {
	request.formRequest.FromRequest(r, request)
}

// Will create url.Values from request.
func (request *ReceiveRequest) ToValues() url.Values {
	return request.formRequest.ToValues(request)
}

// Validates if request fields are valid. Useful when checking if a request is correct.
func (request *ReceiveRequest) Validate() error {
	err := request.formRequest.CheckRequired(request)
	if err != nil {
		return err
	}
	return nil
}

type ReceiveResponse struct {
	protocols.SuccessResponse
	// The full text of the memo the hash of this memo is included in the transaction.
	Memo string `json:"memo"`
}

func (response *ReceiveResponse) Marshal() []byte {
	json, _ := json.MarshalIndent(response, "", "  ")
	return json
}
