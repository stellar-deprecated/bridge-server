package compliance

// PendingResponse is a response from Sanctions and AskUser callbacks when they return 202 Accepted status
type PendingResponse struct {
	// Estimated number of seconds till the sender can check back for a change in status.
	Pending int `json:"pending"`
}
