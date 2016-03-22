package compliance

import (
	"encoding/json"
	"github.com/stellar/gateway/protocols"
)

type SendResponse struct {
	protocols.SuccessResponse
	// xdr.Transaction base64-encoded. Sequence number of this transaction will be equal 0.
	TransactionXdr string `json:"transaction_xdr"`
}

func (response *SendResponse) Marshal() []byte {
	json, _ := json.MarshalIndent(response, "", "  ")
	return json
}
