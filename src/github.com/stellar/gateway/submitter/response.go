package submitter

import (
	"encoding/json"

	"github.com/stellar/go/clients/horizon"
)

// Response contains result of submitting transaction to Stellar network
type Response struct {
	Hash        string `json:"hash,omitempty"`
	SendAmount  string `json:"send_amount,omitempty"` // Path payment only.
	Ledger      int32  `json:"ledger"`
	EnvelopeXdr string `json:"envelope_xdr"`
	ResultXdr   string `json:"result_xdr"`
}

// HTTPStatus implements protocols.SuccessResponse interface
func (response *Response) HTTPStatus() int {
	return 200
}

// Marshal marshals response
func (response *Response) Marshal() []byte {
	json, _ := json.MarshalIndent(response, "", "  ")
	return json
}

// Populate sets the fields on response to the appropriate values to reflect the
// result of `hresp`.
func (response *Response) Populate(hresp *horizon.TransactionSuccess) error {
	response.Hash = hresp.Hash
	response.Ledger = hresp.Ledger
	response.EnvelopeXdr = hresp.Env
	response.ResultXdr = hresp.Result
	// TODO: response.SendAmount = hresp.SendAmount
	return nil
}
