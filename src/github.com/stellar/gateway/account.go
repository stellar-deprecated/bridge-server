package gateway

type Account struct {
	AccountId      string `json:"id"`
	SequenceNumber uint64 `json:"sequence"`
}
