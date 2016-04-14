package bridge

import (
	"encoding/base64"
	"net/http"
	"strings"

	"github.com/stellar/gateway/horizon"
	"github.com/stellar/gateway/protocols"
	"github.com/stellar/go-stellar-base/xdr"
)

var (
	TransactionBadSequence         = &protocols.ErrorResponse{Code: "transaction_bad_seq", Message: "Bad Sequence. Please, try again.", Status: http.StatusBadRequest}
	TransactionBadAuth             = &protocols.ErrorResponse{Code: "transaction_bad_auth", Message: "Invalid network or too few signatures.", Status: http.StatusBadRequest}
	TransactionInsufficientBalance = &protocols.ErrorResponse{Code: "transaction_insufficient_balance", Message: "Transaction fee would bring account below reserve.", Status: http.StatusBadRequest}
	TransactionNoAccount           = &protocols.ErrorResponse{Code: "transaction_no_account", Message: "Source account not found.", Status: http.StatusBadRequest}
	TransactionInsufficientFee     = &protocols.ErrorResponse{Code: "transaction_insufficient_fee", Message: "Transaction fee is too small.", Status: http.StatusBadRequest}
	TransactionBadAuthExtra        = &protocols.ErrorResponse{Code: "transaction_bad_auth_extra", Message: "Unused signatures attached to transaction.", Status: http.StatusBadRequest}
)

func ErrorFromHorizonResponse(response horizon.SubmitTransactionResponse) *protocols.ErrorResponse {
	if response.Ledger == nil && response.Extras != nil {
		var txResult xdr.TransactionResult
		txResult, err := unmarshalTransactionResult(response.Extras.ResultXdr)

		if err != nil {
			return protocols.NewInternalServerError(
				"Error decoding xdr.TransactionResult",
				map[string]interface{}{"err": err},
			)
		}

		transactionResult := txResult.Result.Code
		var operationsResult *xdr.OperationResult
		if txResult.Result.Results != nil {
			operationsResultsSlice := *txResult.Result.Results
			if len(operationsResultsSlice) > 0 {
				operationsResult = &operationsResultsSlice[0]
			}
		}

		if transactionResult != xdr.TransactionResultCodeTxSuccess &&
			transactionResult != xdr.TransactionResultCodeTxFailed {
			switch transactionResult {
			case xdr.TransactionResultCodeTxBadSeq:
				return TransactionBadSequence
			case xdr.TransactionResultCodeTxBadAuth:
				return TransactionBadAuth
			case xdr.TransactionResultCodeTxInsufficientBalance:
				return TransactionInsufficientBalance
			case xdr.TransactionResultCodeTxNoAccount:
				return TransactionNoAccount
			case xdr.TransactionResultCodeTxInsufficientFee:
				return TransactionInsufficientFee
			case xdr.TransactionResultCodeTxBadAuthExtra:
				return TransactionBadAuthExtra
			default:
				return protocols.InternalServerError
			}
		} else if operationsResult != nil {
			if operationsResult.Tr.AllowTrustResult != nil {
				switch operationsResult.Tr.AllowTrustResult.Code {
				case xdr.AllowTrustResultCodeAllowTrustMalformed:
					return AllowTrustMalformed
				case xdr.AllowTrustResultCodeAllowTrustNoTrustLine:
					return AllowTrustNoTrustline
				case xdr.AllowTrustResultCodeAllowTrustTrustNotRequired:
					return AllowTrustTrustNotRequired
				case xdr.AllowTrustResultCodeAllowTrustCantRevoke:
					return AllowTrustCantRevoke
				default:
					return protocols.InternalServerError
				}
			} else if operationsResult.Tr.PaymentResult != nil {
				switch operationsResult.Tr.PaymentResult.Code {
				case xdr.PaymentResultCodePaymentMalformed:
					return PaymentMalformed
				case xdr.PaymentResultCodePaymentUnderfunded:
					return PaymentUnderfunded
				case xdr.PaymentResultCodePaymentSrcNoTrust:
					return PaymentSrcNoTrust
				case xdr.PaymentResultCodePaymentSrcNotAuthorized:
					return PaymentSrcNotAuthorized
				case xdr.PaymentResultCodePaymentNoDestination:
					return PaymentNoDestination
				case xdr.PaymentResultCodePaymentNoTrust:
					return PaymentNoTrust
				case xdr.PaymentResultCodePaymentNotAuthorized:
					return PaymentNotAuthorized
				case xdr.PaymentResultCodePaymentLineFull:
					return PaymentLineFull
				case xdr.PaymentResultCodePaymentNoIssuer:
					return PaymentNoIssuer
				default:
					return protocols.InternalServerError
				}
			} else if operationsResult.Tr.PathPaymentResult != nil {
				switch operationsResult.Tr.PathPaymentResult.Code {
				case xdr.PathPaymentResultCodePathPaymentMalformed:
					return PaymentMalformed
				case xdr.PathPaymentResultCodePathPaymentUnderfunded:
					return PaymentUnderfunded
				case xdr.PathPaymentResultCodePathPaymentSrcNoTrust:
					return PaymentSrcNoTrust
				case xdr.PathPaymentResultCodePathPaymentSrcNotAuthorized:
					return PaymentSrcNotAuthorized
				case xdr.PathPaymentResultCodePathPaymentNoDestination:
					return PaymentNoDestination
				case xdr.PathPaymentResultCodePathPaymentNoTrust:
					return PaymentNoTrust
				case xdr.PathPaymentResultCodePathPaymentNotAuthorized:
					return PaymentNotAuthorized
				case xdr.PathPaymentResultCodePathPaymentLineFull:
					return PaymentLineFull
				case xdr.PathPaymentResultCodePathPaymentNoIssuer:
					return PaymentNoIssuer
				case xdr.PathPaymentResultCodePathPaymentTooFewOffers:
					return PaymentTooFewOffers
				case xdr.PathPaymentResultCodePathPaymentOfferCrossSelf:
					return PaymentOfferCrossSelf
				case xdr.PathPaymentResultCodePathPaymentOverSendmax:
					return PaymentOverSendmax
				default:
					return protocols.InternalServerError
				}
			}
		} else {
			return protocols.InternalServerError
		}
	}

	return nil
}

func unmarshalTransactionResult(transactionResult string) (txResult xdr.TransactionResult, err error) {
	reader := strings.NewReader(transactionResult)
	b64r := base64.NewDecoder(base64.StdEncoding, reader)
	_, err = xdr.Unmarshal(b64r, &txResult)
	return
}
