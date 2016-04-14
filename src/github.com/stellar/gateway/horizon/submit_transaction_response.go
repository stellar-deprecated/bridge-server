package horizon

import (
	"encoding/json"
)

type SubmitTransactionResponse struct {
	Ledger *uint64                          `json:"ledger"`
	Extras *SubmitTransactionResponseExtras `json:"extras"`
}

func (response *SubmitTransactionResponse) HTTPStatus() int {
	return 200
}

func (response *SubmitTransactionResponse) Marshal() []byte {
	json, _ := json.MarshalIndent(response, "", "  ")
	return json
}

type SubmitTransactionResponseExtras struct {
	EnvelopeXdr string `json:"envelope_xdr"`
	ResultXdr   string `json:"result_xdr"`
}
