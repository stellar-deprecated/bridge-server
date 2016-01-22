package horizon

type SubmitTransactionResponse struct {
	Ledger *uint64                          `json:"ledger"`
	Errors *SubmitTransactionResponseError  `json:"errors"`
	Extras *SubmitTransactionResponseExtras `json:"extras"`
}

type SubmitTransactionResponseError struct {
	TransactionErrorCode string `json:"transaction_error"`
	OperationErrorCode   string `json:"operation_error"`
}

type SubmitTransactionResponseExtras struct {
	EnvelopeXdr string `json:"envelope_xdr"`
	ResultXdr   string `json:"result_xdr"`
}
