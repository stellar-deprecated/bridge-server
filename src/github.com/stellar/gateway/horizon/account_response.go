package horizon

type AccountResponse struct {
	AccountId      string `json:"id"`
	SequenceNumber uint64 `json:"sequence"`
}
