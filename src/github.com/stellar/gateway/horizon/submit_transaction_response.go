package horizon

import (
	"encoding/json"
)

// SubmitTransactionResponse contains result of submitting transaction to Stellar network
type SubmitTransactionResponse struct {
	Hash       string                           `json:"hash,omitempty"`
	SendAmount string                           `json:"send_amount,omitempty"` // Path payment only.
	ResultXdr  *string                          `json:"result_xdr,omitempty"`  // Only success response.
	Ledger     *uint64                          `json:"ledger"`
	Extras     *SubmitTransactionResponseExtras `json:"extras,omitempty"`
}

// HTTPStatus implements protocols.SuccessResponse interface
func (response *SubmitTransactionResponse) HTTPStatus() int {
	return 200
}

// Marshal marshals response
func (response *SubmitTransactionResponse) Marshal() []byte {
	json, _ := json.MarshalIndent(response, "", "  ")
	return json
}

// SubmitTransactionResponseExtras contains extra information returned by Horizon
type SubmitTransactionResponseExtras struct {
	EnvelopeXdr string `json:"envelope_xdr"`
	ResultXdr   string `json:"result_xdr"`
}
