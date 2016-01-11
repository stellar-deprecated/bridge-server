package horizon

type SubmitTransactionResponse struct {
	Ledger *int `json:"ledger"`
	Extras *struct {
		EnvelopeXdr string `json:"envelope_xdr"`
		ResultXdr   string `json:"result_xdr"`
	} `json:"extras"`
}
