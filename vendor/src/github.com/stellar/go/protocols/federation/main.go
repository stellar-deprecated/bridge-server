package federation

// NameResponse represents the result of a federation request
// for `name` and `forward` requests.
type NameResponse struct {
	AccountID string `json:"account_id"`
	MemoType  string `json:"memo_type,omitempty"`
	Memo      string `json:"memo,omitempty"`
}

// IDResponse represents the result of a federation request
// for `id` request.
type IDResponse struct {
	Address string `json:"stellar_address"`
}
