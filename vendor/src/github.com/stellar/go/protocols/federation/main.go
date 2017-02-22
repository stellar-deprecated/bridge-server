package federation

// Response represents the result of a federation request.
type Response struct {
	StellarAddress string `json:"stellar_address,omitempty"`
	AccountID      string `json:"account_id"`
	MemoType       string `json:"memo_type,omitempty"`
	Memo           string `json:"memo,omitempty"`
}
