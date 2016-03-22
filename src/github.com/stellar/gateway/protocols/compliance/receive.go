package compliance

import (
	"encoding/json"
	"github.com/stellar/gateway/protocols"
)

type ReceiveResponse struct {
	protocols.SuccessResponse
	// The full text of the memo the hash of this memo is included in the transaction.
	Memo string `json:"memo"`
}

func (response *ReceiveResponse) Marshal() []byte {
	json, _ := json.MarshalIndent(response, "", "  ")
	return json
}
