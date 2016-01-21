package horizon

type SubmitTransactionResponse struct {
	Ledger *uint64 `json:"ledger"`
	Errors *SubmitTransactionResponseError `json:"errors"`
	Extras *struct {
		EnvelopeXdr string `json:"envelope_xdr"`
		ResultXdr   string `json:"result_xdr"`
	} `json:"extras"`
}

type SubmitTransactionResponseError struct {
	TransactionErrorCode string
	OperationErrorCode   string
}
