package bridge

import (
	"fmt"
	"net/http"

	"github.com/stellar/gateway/protocols"
	"github.com/stellar/go/clients/horizon"
	"github.com/stellar/go/xdr"
)

var (
	// TransactionBadSequence is an error response
	TransactionBadSequence = &protocols.ErrorResponse{Code: "transaction_bad_seq", Message: "Bad Sequence. Please, try again.", Status: http.StatusBadRequest}
	// TransactionBadAuth is an error response
	TransactionBadAuth = &protocols.ErrorResponse{Code: "transaction_bad_auth", Message: "Invalid network or too few signatures.", Status: http.StatusBadRequest}
	// TransactionInsufficientBalance is an error response
	TransactionInsufficientBalance = &protocols.ErrorResponse{Code: "transaction_insufficient_balance", Message: "Transaction fee would bring account below reserve.", Status: http.StatusBadRequest}
	// TransactionNoAccount is an error response
	TransactionNoAccount = &protocols.ErrorResponse{Code: "transaction_no_account", Message: "Source account not found.", Status: http.StatusBadRequest}
	// TransactionInsufficientFee is an error response
	TransactionInsufficientFee = &protocols.ErrorResponse{Code: "transaction_insufficient_fee", Message: "Transaction fee is too small.", Status: http.StatusBadRequest}
	// TransactionBadAuthExtra is an error response
	TransactionBadAuthExtra  = &protocols.ErrorResponse{Code: "transaction_bad_auth_extra", Message: "Unused signatures attached to transaction.", Status: http.StatusBadRequest}
	TransactionInternalError = &protocols.ErrorResponse{Code: "transaction_internal_error", Message: "Transaction triggered an internal error to stellar-core", Status: http.StatusBadRequest}
)

// ErrorFromHorizonResponse checks if horizon.SubmitTransactionResponse is an
// error response and creates ErrorResponse for it
func ErrorFromHorizonResponse(herr *horizon.Error) *protocols.ErrorResponse {

	if herr.Problem.Type == "transaction_malformed" {
		return protocols.NewInternalServerError("transaction malformed", nil)
	}

	trc, err := herr.ResultCodes()
	if err != nil {
		return protocols.NewInternalServerError("Cannot retrieve error codes", nil)
	}

	// first, we see if the error code at the transaction
	switch trc.TransactionCode {
	case "tx_bad_seq":
		return TransactionBadSequence
	case "tx_bad_auth":
		return TransactionBadAuth
	case "tx_insufficient_balance":
		return TransactionInsufficientBalance
	case "tx_no_source_account":
		return TransactionNoAccount
	case "tx_insufficient_fee":
		return TransactionInsufficientFee
	case "tx_bad_auth_extra":
		return TransactionBadAuthExtra
	case "tx_internal_error":
		return TransactionInternalError
	case "tx_failed":
		// noop; continue on to the operation inspection below
	default:
		panic(fmt.Sprintf("Unexpected transaction result code: %s", trc.TransactionCode))
	}

	if len(trc.OperationCodes) != 1 {
		return protocols.NewInternalServerError("unexpected op codes: expected exactly one", nil)
	}

	txe, err := herr.Envelope()
	if err != nil {
		return protocols.NewInternalServerError("Cannot retrieve envelope from error", nil)
	}

	opcode := trc.OperationCodes[0]
	optype := txe.Tx.Operations[0].Body.Type

	// determine the error response based upon the operation
	switch opcode {
	case "op_malformed":
		return malformedErrorResponse(optype)
	case "op_no_trustline":
		return AllowTrustNoTrustline
	case "op_not_required":
		return AllowTrustTrustNotRequired
	case "op_cant_revoke":
		return AllowTrustCantRevoke
	case "op_underfunded":
		return PaymentUnderfunded
	case "op_src_no_trust":
		return PaymentSrcNoTrust
	case "op_src_not_authorized":
		return PaymentSrcNotAuthorized
	case "op_no_destination":
		return PaymentNoDestination
	case "op_no_trust":
		return PaymentNoTrust
	case "op_not_authorized":
		return PaymentNotAuthorized
	case "op_line_full":
		return PaymentLineFull
	case "op_no_issuer":
		return PaymentNoIssuer
	case "op_too_few_offers":
		return PaymentTooFewOffers
	case "op_cross_self":
		return PaymentOfferCrossSelf
	case "op_over_source_max":
		return PaymentOverSendmax
	default:
		msg := fmt.Sprintf("unexpected operation result: %s", opcode)
		return protocols.NewInternalServerError(msg, nil)
	}
}

func malformedErrorResponse(optype xdr.OperationType) *protocols.ErrorResponse {
	switch optype {
	case xdr.OperationTypeAllowTrust:
		return AllowTrustMalformed
	case xdr.OperationTypePayment:
		return PaymentMalformed
	case xdr.OperationTypePathPayment:
		return PaymentMalformed
	default:
		msg := fmt.Sprintf("unexpected operation type: %s", optype)
		return protocols.NewInternalServerError(msg, nil)
	}
}
