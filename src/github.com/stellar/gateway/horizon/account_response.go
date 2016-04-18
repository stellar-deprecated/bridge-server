package horizon

// AccountResponse contains account data returned by Horizon
type AccountResponse struct {
	AccountID      string `json:"id"`
	SequenceNumber string `json:"sequence"`
}
