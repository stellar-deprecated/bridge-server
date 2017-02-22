package compliance

// AuthStatus represents auth status returned by Auth Server
type AuthStatus string

const (
	// AuthStatusOk is returned when authentication was successful
	AuthStatusOk AuthStatus = "ok"
	// AuthStatusPending is returned when authentication is pending
	AuthStatusPending AuthStatus = "pending"
	// AuthStatusDenied is returned when authentication was denied
	AuthStatusDenied AuthStatus = "denied"
)

// AuthRequest represents auth request sent to compliance server
type AuthRequest struct {
	// Marshalled AuthData JSON object (because of the attached signature, json can be marshalled to infinite number of valid JSON strings)
	DataJSON string `name:"data"`
	// Signature of sending FI
	Signature string `name:"sig"`
}

// AuthData represents how AuthRequest.Data field looks like.
type AuthData struct {
	// The stellar address of the customer that is initiating the send.
	Sender string `json:"sender"`
	// If the caller needs the recipient's AML info in order to send the payment.
	NeedInfo bool `json:"need_info"`
	// The transaction that the sender would like to send in XDR format. This transaction is unsigned.
	Tx string `json:"tx"`
	// The full text of the attachment the hash of this attachment is included in the transaction.
	AttachmentJSON string `json:"attachment"`
}

// AuthResponse represents response sent by auth server
type AuthResponse struct {
	// If this FI is willing to share AML information or not. {ok, denied, pending}
	InfoStatus AuthStatus `json:"info_status"`
	// If this FI is willing to accept this transaction. {ok, denied, pending}
	TxStatus AuthStatus `json:"tx_status"`
	// (only present if info_status is ok) JSON of the recipient's AML information. in the Stellar attachment convention
	DestInfo string `json:"dest_info,omitempty"`
	// (only present if info_status or tx_status is pending) Estimated number of seconds till the sender can check back for a change in status. The sender should just resubmit this request after the given number of seconds.
	Pending int `json:"pending,omitempty"`
}
