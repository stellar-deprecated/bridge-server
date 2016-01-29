package handlers

import (
	"encoding/json"
	log "github.com/Sirupsen/logrus"
	"net/http"
	"strconv"

	b "github.com/stellar/go-stellar-base/build"
	"github.com/stellar/go-stellar-base/keypair"
)

func (rh *RequestHandler) Send(w http.ResponseWriter, r *http.Request) {
	destination := r.PostFormValue("destination")
	assetCode := r.PostFormValue("asset_code")
	amount := r.PostFormValue("amount")

	destinationObject, err := rh.AddressResolver.Resolve(destination)
	if err != nil {
		log.WithFields(log.Fields{"destination": destination}).Print("Cannot resolve address")
		errorBadRequest(w, errorResponseString("invalid_destination", "Cannot resolve destination"))
		return
	}

	_, err = keypair.Parse(destinationObject.AccountId)
	if err != nil {
		log.WithFields(log.Fields{"AccountId": destinationObject.AccountId}).Print("Invalid AccountId in destination")
		errorBadRequest(w, errorResponseString("invalid_destination", "destination parameter is invalid"))
		return
	}

	if !rh.isAssetAllowed(assetCode) {
		log.Print("Asset code not allowed: ", assetCode)
		errorBadRequest(w, errorResponseString("invalid_asset_code", "Given assetCode not allowed"))
		return
	}

	issuingKeypair, err := keypair.Parse(*rh.Config.Accounts.IssuingSeed)
	if err != nil {
		log.Print("Invalid issuingSeed")
		errorServerError(w)
		return
	}

	operationMutator := b.Payment(
		b.Destination{destinationObject.AccountId},
		b.CreditAmount{assetCode, issuingKeypair.Address(), amount},
	)
	if operationMutator.Err != nil {
		log.Print("Error creating operationMutator ", operationMutator.Err)
		errorServerError(w)
		return
	}

	memoType := r.PostFormValue("memo_type")
	memo := r.PostFormValue("memo")

	if !(((memoType == "") && (memo == "")) || ((memoType != "") && (memo != ""))) {
		log.Print("Missing one of memo params.")
		errorBadRequest(w, errorResponseString("memo_missing_param", "When passing memo both params: `memo_type`, `memo` are required"))
		return
	}

	if destinationObject.MemoType != nil {
		if memoType != "" {
			log.Print("Memo given in request but federation returned memo fields.")
			errorBadRequest(w, errorResponseString("cannot_use_memo", "Memo given in request but federation returned memo fields"))
			return
		}

		memoType = *destinationObject.MemoType
		memo = *destinationObject.Memo
	}

	var memoMutator interface{}
	switch {
	case memoType == "":
		break
	case memoType == "id":
		id, err := strconv.ParseUint(memo, 10, 64)
		if err != nil {
			log.WithFields(log.Fields{"memo": memo}).Print("Cannot convert memo_id value to uint64")
			errorBadRequest(w, errorResponseString("cannot_convert_memo_id", "Cannot convert memo_id value"))
			return
		}
		memoMutator = b.MemoID{id}
	case memoType == "text":
		memoMutator = b.MemoText{memo}
	default:
		log.Print("Not supported memo type: ", memoType)
		errorBadRequest(w, errorResponseString("memo_not_supported", "Not supported memo type"))
		return
	}

	submitResponse, err := rh.TransactionSubmitter.SubmitTransaction(
		*rh.Config.Accounts.IssuingSeed,
		operationMutator,
		memoMutator,
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

	json, err := json.MarshalIndent(submitResponse, "", "  ")

	if err != nil {
		errorServerError(w)
		return
	}

	w.Write(json)
}
