package gateway

import (
	"encoding/json"
	log "github.com/Sirupsen/logrus"
	"net/http"

	b "github.com/stellar/go-stellar-base/build"
	"github.com/stellar/go-stellar-base/keypair"
)

type RequestHandler struct {
	config               *Config
	database             Database
	transactionSubmitter *TransactionSubmitter
}

func (rh *RequestHandler) Authorize(w http.ResponseWriter, r *http.Request) {
	accountId := r.URL.Query().Get("accountId")

	_, err := keypair.Parse(accountId)
	if err != nil {
		log.Print("Invalid accountId parameter: ", accountId)
		errorBadRequest(w, errorResponseString("invalid_account_id", "accountId parameter is invalid"))
		return
	}

	operation := b.AllowTrust(
		b.Trustor{accountId},
		b.Authorize{true},
		b.AllowTrustAsset{"USD"},
	)

	submitResponse, err := rh.transactionSubmitter.SubmitTransaction(
		rh.config.Accounts.AuthorizingSeed,
		operation,
	)
	if err != nil {
		log.Print("Error submitting transaction ", err)
		errorServerError(w)
		return
	}

	if submitResponse.Errors != nil {		
		var errorString string
		if submitResponse.Errors.OperationErrorCode != "" {
			switch submitResponse.Errors.OperationErrorCode {
			case "allow_trust_malformed":
				errorString = errorResponseString(
					"allow_trust_malformed",
					"Asset name is malformed.",
				)
			case "allow_trust_not_trustline":
				errorString = errorResponseString(
					"allow_trust_not_trustline",
					"Trustor does not have a trustline yet.",
				)
			case "allow_trust_trust_not_required":
				errorString = errorResponseString(
					"allow_trust_trust_not_required",
					"Authorizing account does not require allowing trust. Set AUTH_REQUIRED_FLAG on your account to use this feature.",
				)
			case "allow_trust_trust_cant_revoke":
				errorString = errorResponseString(
					"allow_trust_trust_cant_revoke",
					"Authorizing account has AUTH_REVOCABLE_FLAG set. Can't revoke the trustline.",
				)
			default:
				errorServerError(w)
				return
			}
		} else if submitResponse.Errors.TransactionErrorCode != "" {
			switch submitResponse.Errors.TransactionErrorCode {
			case "transaction_bad_seq":
				errorString = errorResponseString(
					"transaction_bad_seq",
					"Bad Sequence. Please, try again.",
				)
			default:
				errorServerError(w)
				return
			}
		}

		errorBadRequest(w, errorString)
		return
	}

	json, err := json.Marshal(submitResponse)

	if err != nil {
		errorServerError(w)
		return
	}

	w.Write(json)
}

// func (rh *RequestHandler) Send(w http.ResponseWriter, r *http.Request) {
// 	q := r.URL.Query()
// 	destination := q.Get("destination")
// 	asset := q.Get("asset")
// 	amount := q.Get("amount")

// 	_, err := keypair.Parse(destination)
// 	if err != nil {
// 		log.Print("Invalid destination parameter: ", destination)
// 		errorBadRequest(w, errorResponseString("invalid_destination", "destination parameter is invalid"))
// 		return
// 	}

// 	issuingKeypair, err := keypair.Parse(rh.config.Accounts.IssuingSeed)
// 	if err != nil {
// 		log.Print("Invalid issuing seed")
// 		errorServerError(w)
// 		return
// 	}

// 	// TODO check if asset is allowed

// 	operation := b.Payment(
// 		b.Destination{destination},
// 		b.CreditAmount{asset, issuingKeypair.Address(), amount},
// 	)

// 	submitResponse, err := rh.transactionSubmitter.SubmitTransaction(
// 		rh.config.Accounts.IssuingSeed,
// 		operation,
// 	)
// 	if err != nil {
// 		log.Print("Error submitting transaction ", err)
// 		errorServerError(w)
// 		return
// 	}

// 	if submitResponse.Ledger == nil && submitResponse.Extras != nil {
// 		txResult, err := unmarshalTransactionResult(submitResponse.Extras.ResultXdr)

// 		if err != nil {
// 			log.Print("Cannot decode transaction result")
// 			errorServerError(w)
// 			return
// 		}

// 		// TODO add transaction result checks
// 		operationsResults := *txResult.Result.Results
// 		var errorString string
// 		switch operationsResults[0].Tr.PaymentResult.Code {
// 		case xdr.PaymentResultCodePaymentMalformed:
// 			errorString = errorResponseString(
// 				"payment_malformed",
// 				"Operation is malformed.",
// 			)
// 		case xdr.PaymentResultCodePaymentUnderfunded:
// 			errorString = errorResponseString(
// 				"payment_underfunded",
// 				"Not enough funds to send this transaction.",
// 			)
// 		case xdr.PaymentResultCodePaymentSrcNoTrust:
// 			errorString = errorResponseString(
// 				"payment_src_no_trust",
// 				"No trustline on source account.",
// 			)
// 		case xdr.PaymentResultCodePaymentSrcNotAuthorized:
// 			errorString = errorResponseString(
// 				"payment_src_not_authorized",
// 				"Source not authorized to transfer.",
// 			)
// 		case xdr.PaymentResultCodePaymentNoDestination:
// 			errorString = errorResponseString(
// 				"payment_no_destination",
// 				"Destination account does not exist.",
// 			)
// 		case xdr.PaymentResultCodePaymentNoTrust:
// 			errorString = errorResponseString(
// 				"payment_no_trust",
// 				"Destination missing a trust line for asset.",
// 			)
// 		case xdr.PaymentResultCodePaymentNotAuthorized:
// 			errorString = errorResponseString(
// 				"payment_not_authorized",
// 				"Destination not authorized to trust asset. It needs to be allowed first by using /authorize endpoint.",
// 			)
// 		case xdr.PaymentResultCodePaymentLineFull:
// 			errorString = errorResponseString(
// 				"payment_line_full",
// 				"Sending this payment would make a destination go above their limit.",
// 			)
// 		case xdr.PaymentResultCodePaymentNoIssuer:
// 			errorString = errorResponseString(
// 				"payment_no_issuer",
// 				"Missing issuer on asset.",
// 			)
// 		default:
// 			errorServerError(w)
// 			return
// 		}
// 		errorBadRequest(w, errorString)
// 		return
// 	}

// 	json, err := json.Marshal(submitResponse)

// 	if err != nil {
// 		errorServerError(w)
// 		return
// 	}

// 	w.Write(json)
// }
