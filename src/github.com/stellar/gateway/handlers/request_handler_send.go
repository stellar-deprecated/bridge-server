package handlers

import (
	"encoding/json"
	log "github.com/Sirupsen/logrus"
	"net/http"

	b "github.com/stellar/go-stellar-base/build"
	"github.com/stellar/go-stellar-base/keypair"
)

func (rh *RequestHandler) Send(w http.ResponseWriter, r *http.Request) {
	destination := r.PostFormValue("destination")
	assetCode := r.PostFormValue("asset_code")
	amount := r.PostFormValue("amount")

	_, err := keypair.Parse(destination)
	if err != nil {
		log.Print("Invalid destination parameter: ", destination)
		errorBadRequest(w, errorResponseString("invalid_destination", "destination parameter is invalid"))
		return
	}

	if !rh.isAssetAllowed(assetCode) {
		log.Print("Asset code not allowed: ", assetCode)
		errorBadRequest(w, errorResponseString("invalid_asset_code", "Given assetCode not allowed"))
		return
	}

	issuingKeypair, err := keypair.Parse(rh.Config.Accounts.IssuingSeed)
	if err != nil {
		log.Print("Invalid issuingSeed")
		errorServerError(w)
		return
	}

	operation := b.Payment(
		b.Destination{destination},
		b.CreditAmount{assetCode, issuingKeypair.Address(), amount},
	)
	if operation.Err != nil {
		log.Print("Error creating operation ", operation.Err)
		errorServerError(w)
		return
	}

	submitResponse, err := rh.TransactionSubmitter.SubmitTransaction(
		rh.Config.Accounts.IssuingSeed,
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
			case "payment_malformed":
				errorString = errorResponseString(
					"payment_malformed",
					"Operation is malformed.",
				)
			case "payment_underfunded":
				errorString = errorResponseString(
					"payment_underfunded",
					"Not enough funds to send this transaction.",
				)
			case "payment_src_no_trust":
				errorString = errorResponseString(
					"payment_src_no_trust",
					"No trustline on source account.",
				)
			case "payment_src_not_authorized":
				errorString = errorResponseString(
					"payment_src_not_authorized",
					"Source not authorized to transfer.",
				)
			case "payment_no_destination":
				errorString = errorResponseString(
					"payment_no_destination",
					"Destination account does not exist.",
				)
			case "payment_no_trust":
				errorString = errorResponseString(
					"payment_no_trust",
					"Destination missing a trust line for asset.",
				)
			case "payment_not_authorized":
				errorString = errorResponseString(
					"payment_not_authorized",
					"Destination not authorized to trust asset. It needs to be allowed first by using /authorize endpoint.",
				)
			case "payment_line_full":
				errorString = errorResponseString(
					"payment_line_full",
					"Sending this payment would make a destination go above their limit.",
				)
			case "payment_no_issuer":
				errorString = errorResponseString(
					"payment_no_issuer",
					"Missing issuer on asset.",
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
