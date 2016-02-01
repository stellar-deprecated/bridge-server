package handlers

import (
	log "github.com/Sirupsen/logrus"
	"net/http"
	"strconv"

	"github.com/stellar/gateway/horizon"
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
		writeError(w, horizon.PaymentCannotResolveDestination)
		return
	}

	_, err = keypair.Parse(destinationObject.AccountId)
	if err != nil {
		log.WithFields(log.Fields{"AccountId": destinationObject.AccountId}).Print("Invalid AccountId in destination")
		writeError(w, horizon.PaymentInvalidDestination)
		return
	}

	if !rh.isAssetAllowed(assetCode) {
		log.Print("Asset code not allowed: ", assetCode)
		writeError(w, horizon.PaymentAssetCodeNotAllowed)
		return
	}

	issuingKeypair, err := keypair.Parse(*rh.Config.Accounts.IssuingSeed)
	if err != nil {
		log.Print("Invalid issuingSeed")
		writeError(w, horizon.ServerError)
		return
	}

	operationMutator := b.Payment(
		b.Destination{destinationObject.AccountId},
		b.CreditAmount{assetCode, issuingKeypair.Address(), amount},
	)
	if operationMutator.Err != nil {
		log.Print("Error creating operationMutator ", operationMutator.Err)
		writeError(w, horizon.ServerError)
		return
	}

	memoType := r.PostFormValue("memo_type")
	memo := r.PostFormValue("memo")

	if !(((memoType == "") && (memo == "")) || ((memoType != "") && (memo != ""))) {
		log.Print("Missing one of memo params.")
		writeError(w, horizon.PaymentMissingParamMemo)
		return
	}

	if destinationObject.MemoType != nil {
		if memoType != "" {
			log.Print("Memo given in request but federation returned memo fields.")
			writeError(w, horizon.PaymentCannotUseMemo)
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
			writeError(w, horizon.PaymentInvalidMemo)
			return
		}
		memoMutator = b.MemoID{id}
	case memoType == "text":
		memoMutator = b.MemoText{memo}
	default:
		log.Print("Not supported memo type: ", memoType)
		writeError(w, horizon.PaymentInvalidMemo)
		return
	}

	submitResponse, err := rh.TransactionSubmitter.SubmitTransaction(
		*rh.Config.Accounts.IssuingSeed,
		operationMutator,
		memoMutator,
	)
	if err != nil {
		log.Print("Error submitting transaction ", err)
		writeError(w, horizon.ServerError)
		return
	}

	write(w, submitResponse)
}
