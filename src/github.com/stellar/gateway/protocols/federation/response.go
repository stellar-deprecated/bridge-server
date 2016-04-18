package federation

// Response represents response returned by federation server
type Response struct {
	AccountID string `json:"account_id"`
	MemoType  string `json:"memo_type"`
	Memo      string `json:"memo"`
}
