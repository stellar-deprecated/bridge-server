package federation

type Response struct {
	AccountId string `json:"account_id"`
	MemoType  string `json:"memo_type"`
	Memo      string `json:"memo"`
}
