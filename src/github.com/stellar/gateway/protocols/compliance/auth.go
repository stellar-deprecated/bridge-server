package compliance

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/stellar/gateway/protocols"
)

type AuthStatus string

const (
	AuthStatusOk      AuthStatus = "ok"
	AuthStatusPending AuthStatus = "pending"
	AuthStatusDenied  AuthStatus = "denied"
)

type AuthRequest struct {
	// Stringified AuthData JSON object
	Data string `name:"data" required:""`
	// Signature of sending FI
	Signature string `name:"sig" required:""`

	formRequest protocols.FormRequest
}

type AuthData struct {
	// The stellar address of the customer that is initiating the send.
	Sender string `name:"sender" required:""`
	// If the caller needs the recipient's AML info in order to send the payment.
	NeedInfo bool `name:"need_info" required:""`
	// The transaction that the sender would like to send in XDR format. This transaction is unsigned.
	Tx string `name:"tx" required:""`
	// The full text of the memo the hash of this memo is included in the transaction.
	Memo string `name:"memo" required:""`
}

// Will populate request fields using http.Request.
func (request *AuthRequest) FromRequest(r *http.Request) {
	request.formRequest.FromRequest(r, request)
}

// Will create url.Values from request.
func (request *AuthRequest) ToValues() url.Values {
	return request.formRequest.ToValues(request)
}

// Validates if request fields are valid. Useful when checking if a request is correct.
func (request *AuthRequest) Validate() error {
	err := request.formRequest.CheckRequired(request)
	if err != nil {
		return err
	}

	return nil
}

type AuthResponse struct {
	protocols.SuccessResponse
	// If this FI is willing to share AML information or not. {ok, denied, pending}
	InfoStatus AuthStatus `json:"info_status"`
	// If this FI is willing to accept this transaction. {ok, denied, pending}
	TxStatus AuthStatus `json:"tx_status"`
	// (only present if info_status is ok) JSON of the recipient's AML information. in the Stellar memo convention
	DestInfo string `json:"dest_info,omitempty"`
	// (only present if info_status or tx_status is pending) Estimated number of seconds till the sender can check back for a change in status. The sender should just resubmit this request after the given number of seconds.
	Pending int `json:"pending,omitempty"`
}

func (response *AuthResponse) Marshal() []byte {
	json, _ := json.MarshalIndent(response, "", "  ")
	return json
}
