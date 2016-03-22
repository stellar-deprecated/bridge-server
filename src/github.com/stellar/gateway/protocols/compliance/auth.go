package compliance

import (
	"encoding/json"
	"github.com/stellar/gateway/protocols"
)

const (
	AUTH_STATUS_OK      = "ok"
	AUTH_STATUS_PENDING = "pending"
	AUTH_STATUS_DENIED  = "denied"
)

type AuthData struct {
	// The stellar address of the customer that is initiating the send.
	Sender string
	// If the caller needs the recipient's AML info in order to send the payment.
	NeedInfo bool
	// The transaction that the sender would like to send in XDR format. This transaction is unsigned.
	Tx string
	// The full text of the memo the hash of this memo is included in the transaction.
	Memo string
}

type AuthResponse struct {
	protocols.SuccessResponse
	// If this FI is willing to share AML information or not. {ok, denied, pending}
	InfoStatus string `json:"info_status"`
	// If this FI is willing to accept this transaction. {ok, denied, pending}
	TxStatus string `json:"tx_status"`
	// (only present if info_status is ok) JSON of the recipient's AML information. in the Stellar memo convention
	DestInfo string `json:"dest_info,omitempty"`
	// (only present if info_status or tx_status is pending) Estimated number of seconds till the sender can check back for a change in status. The sender should just resubmit this request after the given number of seconds.
	Pending int `json:"pending,omitempty"`
}

func (response *AuthResponse) Marshal() []byte {
	json, _ := json.MarshalIndent(response, "", "  ")
	return json
}
