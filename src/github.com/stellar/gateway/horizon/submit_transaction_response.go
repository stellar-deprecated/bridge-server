package horizon

import (
	"encoding/json"
	"net/http"
)

var (
	ServerError = &SubmitTransactionResponseError{Code: "server_error", Message: "Server error", Status: http.StatusInternalServerError}

	// global errors
	TransactionBadSequence         = &SubmitTransactionResponseError{Code: "transaction_bad_seq", Message: "Bad Sequence. Please, try again.", Status: http.StatusBadRequest}
	TransactionBadAuth             = &SubmitTransactionResponseError{Code: "transaction_bad_auth", Message: "Invalid network or too few signatures.", Status: http.StatusBadRequest}
	TransactionInsufficientBalance = &SubmitTransactionResponseError{Code: "transaction_insufficient_balance", Message: "Transaction fee would bring account below reserve.", Status: http.StatusBadRequest}
	TransactionNoAccount           = &SubmitTransactionResponseError{Code: "transaction_no_account", Message: "Source account not found.", Status: http.StatusBadRequest}
	TransactionInsufficientFee     = &SubmitTransactionResponseError{Code: "transaction_insufficient_fee", Message: "Transaction fee is too small.", Status: http.StatusBadRequest}
	TransactionBadAuthExtra        = &SubmitTransactionResponseError{Code: "transaction_bad_auth_extra", Message: "Unused signatures attached to transaction.", Status: http.StatusBadRequest}

	// /authorize:

	// input errors
	AllowTrustInvalidAccountId    = &SubmitTransactionResponseError{Code: "invalid_account_id", Message: "accountId parameter is invalid.", Status: http.StatusBadRequest}
	AllowTrustAssetCodeNotAllowed = &SubmitTransactionResponseError{Code: "asset_code_not_allowed", Message: "Given asset_code not allowed.", Status: http.StatusBadRequest}

	// allow_trust op errors
	AllowTrustMalformed        = &SubmitTransactionResponseError{Code: "allow_trust_malformed", Message: "Asset name is malformed.", Status: http.StatusBadRequest}
	AllowTrustNoTrustline      = &SubmitTransactionResponseError{Code: "allow_trust_no_trustline", Message: "Trustor does not have a trustline yet.", Status: http.StatusBadRequest}
	AllowTrustTrustNotRequired = &SubmitTransactionResponseError{Code: "allow_trust_trust_not_required", Message: "Authorizing account does not require allowing trust. Set AUTH_REQUIRED_FLAG on your account to use this feature.", Status: http.StatusBadRequest}
	AllowTrustCantRevoke       = &SubmitTransactionResponseError{Code: "allow_trust_cant_revoke", Message: "Authorizing account has AUTH_REVOCABLE_FLAG set. Can't revoke the trustline.", Status: http.StatusBadRequest}

	// /send & /payment

	// input errors
	PaymentInvalidType              = &SubmitTransactionResponseError{Code: "invalid_type", Message: "Invalid operation type.", Status: http.StatusBadRequest}
	PaymentInvalidSource            = &SubmitTransactionResponseError{Code: "invalid_source", Message: "source parameter is invalid.", Status: http.StatusBadRequest}
	PaymentCannotResolveDestination = &SubmitTransactionResponseError{Code: "cannot_resolve_destination", Message: "Cannot resolve federated Stellar address.", Status: http.StatusBadRequest}
	PaymentInvalidDestination       = &SubmitTransactionResponseError{Code: "invalid_destination", Message: "destination parameter is invalid.", Status: http.StatusBadRequest}
	PaymentInvalidIssuer            = &SubmitTransactionResponseError{Code: "invalid_issuer", Message: "asset_issuer parameter is invalid.", Status: http.StatusBadRequest}
	PaymentMissingParamAsset        = &SubmitTransactionResponseError{Code: "missing_param_asset", Message: "When passing asset both params: `asset_code`, `asset_issuer` are required.", Status: http.StatusBadRequest}
	PaymentMissingParamMemo         = &SubmitTransactionResponseError{Code: "missing_param_memo", Message: "When passing memo both params: `memo_type`, `memo` are required.", Status: http.StatusBadRequest}
	PaymentCannotUseMemo            = &SubmitTransactionResponseError{Code: "cannot_use_memo", Message: "Memo given in request but federation returned memo fields.", Status: http.StatusBadRequest}
	PaymentInvalidMemo              = &SubmitTransactionResponseError{Code: "invalid_memo", Message: "Invalid or unsupported memo.", Status: http.StatusBadRequest}
	PaymentSourceNotExist           = &SubmitTransactionResponseError{Code: "source_not_exist", Message: "Source account does not exist.", Status: http.StatusBadRequest}
	PaymentInvalidAmount            = &SubmitTransactionResponseError{Code: "invalid_amount", Message: "amount parameter is invalid.", Status: http.StatusBadRequest}
	PaymentMalformedAssetCode       = &SubmitTransactionResponseError{Code: "malformed_asset_code", Message: "asset_code parameter is malformed.", Status: http.StatusBadRequest}
	PaymentAssetCodeNotAllowed      = &SubmitTransactionResponseError{Code: "asset_code_not_allowed", Message: "Given asset_code not allowed.", Status: http.StatusBadRequest}

	// payment op errors
	PaymentMalformed        = &SubmitTransactionResponseError{Code: "payment_malformed", Message: "Operation is malformed.", Status: http.StatusBadRequest}
	PaymentUnderfunded      = &SubmitTransactionResponseError{Code: "payment_underfunded", Message: "Not enough funds to send this transaction.", Status: http.StatusBadRequest}
	PaymentSrcNoTrust       = &SubmitTransactionResponseError{Code: "payment_src_no_trust", Message: "No trustline on source account.", Status: http.StatusBadRequest}
	PaymentSrcNotAuthorized = &SubmitTransactionResponseError{Code: "payment_src_not_authorized", Message: "Source not authorized to transfer.", Status: http.StatusBadRequest}
	PaymentNoDestination    = &SubmitTransactionResponseError{Code: "payment_no_destination", Message: "Destination account does not exist.", Status: http.StatusBadRequest}
	PaymentNoTrust          = &SubmitTransactionResponseError{Code: "payment_no_trust", Message: "Destination missing a trust line for asset.", Status: http.StatusBadRequest}
	PaymentNotAuthorized    = &SubmitTransactionResponseError{Code: "payment_not_authorized", Message: "Destination not authorized to trust asset. It needs to be allowed first by using /authorize endpoint.", Status: http.StatusBadRequest}
	PaymentLineFull         = &SubmitTransactionResponseError{Code: "payment_line_full", Message: "Sending this payment would make a destination go above their limit.", Status: http.StatusBadRequest}
	PaymentNoIssuer         = &SubmitTransactionResponseError{Code: "payment_no_issuer", Message: "Missing issuer on asset.", Status: http.StatusBadRequest}
)

type SubmitTransactionResponse struct {
	Ledger *uint64                          `json:"ledger"`
	Error  *SubmitTransactionResponseError  `json:"error"`
	Extras *SubmitTransactionResponseExtras `json:"extras"`
}

type SubmitTransactionResponseError struct {
	Status  int    `json:"-"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

type SubmitTransactionResponseExtras struct {
	EnvelopeXdr string `json:"envelope_xdr"`
	ResultXdr   string `json:"result_xdr"`
}

func NewErrorResponse(code, message string) (resp SubmitTransactionResponse) {
	error := SubmitTransactionResponseError{
		Code:    code,
		Message: message,
	}
	resp.Error = &error
	return
}

func (response *SubmitTransactionResponse) Marshal() []byte {
	json, _ := json.MarshalIndent(response, "", "  ")
	return json
}

func (error *SubmitTransactionResponseError) Equals(otherError *SubmitTransactionResponseError) bool {
	return error.Code == otherError.Code
}
