package handlers

import (
	"encoding/json"
	log "github.com/Sirupsen/logrus"
	"net/http"
	"strconv"
	"strings"

	b "github.com/stellar/go-stellar-base/build"
	"github.com/stellar/go-stellar-base/keypair"
)

func (rh *RequestHandler) Payment(w http.ResponseWriter, r *http.Request) {
	source := r.PostFormValue("source")
	sourceKeypair, err := keypair.Parse(source)
	if err != nil {
		log.WithFields(log.Fields{"source": source}).Print("Invalid source parameter")
		errorBadRequest(w, errorResponseString("invalid_source", "source parameter is invalid"))
		return
	}

	destination := r.PostFormValue("destination")
	destinationObject, err := ResolveAddress(destination)
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

	amount := r.PostFormValue("amount")
	assetCode := r.PostFormValue("asset_code")
	assetIssuer := r.PostFormValue("asset_issuer")

	var amountMutator interface{}

	if assetCode != "" && assetIssuer != "" {
		issuerKeypair, err := keypair.Parse(assetIssuer)
		if err != nil {
			log.WithFields(log.Fields{"asset_issuer": assetIssuer}).Print("Invalid asset_issuer parameter")
			errorBadRequest(w, errorResponseString("invalid_issuer", "asset_issuer parameter is invalid"))
			return
		}
		amountMutator = b.CreditAmount{assetCode, issuerKeypair.Address(), amount}
	} else if assetCode == "" && assetIssuer == "" {
		amountMutator = b.NativeAmount{amount}
	} else {
		log.Print("Missing asset param.")
		errorBadRequest(w, errorResponseString("asset_missing_param", "When passing asser both params: `asset_code`, `asset_issuer` are required"))
		return
	}

	paymentOperation := b.Payment(
		b.Destination{destinationObject.AccountId},
		amountMutator.(b.PaymentMutator),
	)

	if paymentOperation.Err != nil {
		log.WithFields(log.Fields{"err": paymentOperation.Err}).Print("Payment builder error")
		switch {
		case paymentOperation.Err.Error() == "Asset code length is invalid":
			errorBadRequest(w, errorResponseString("asset_code_invalid", "asset_code param is invalid"))
		case strings.Contains(paymentOperation.Err.Error(), "cannot parse amount"):
			errorBadRequest(w, errorResponseString("invalid_amount", "amount is invalid"))
		default:
			errorServerError(w)
		}
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
		memoMutator = &b.MemoText{memo}
	default:
		log.Print("Not supported memo type: ", memoType)
		errorBadRequest(w, errorResponseString("memo_not_supported", "Not supported memo type"))
		return
	}

	accountResponse, err := rh.Horizon.LoadAccount(sourceKeypair.Address())
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Error("Cannot load source account")
		errorBadRequest(w, errorResponseString("source_not_exist", "source account does not exist"))
		return
	}

	networkMutator := b.Network{rh.Config.NetworkPassphrase}

	transactionMutators := []b.TransactionMutator{
		b.SourceAccount{source},
		b.Sequence{accountResponse.SequenceNumber + 1},
		networkMutator,
		paymentOperation,
	}

	if memoMutator != nil {
		transactionMutators = append(transactionMutators, memoMutator.(b.TransactionMutator))
	}

	tx := b.Transaction(transactionMutators...)

	// TODO check for errors

	txe := tx.Sign(source)
	txeB64, err := txe.Base64()

	if err != nil {
		log.WithFields(log.Fields{"error": err}).Error("Cannot encode transaction envelope")
		errorServerError(w)
		return
	}

	submitResponse, err := rh.Horizon.SubmitTransaction(txeB64)
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Error("Error submitting transaction")
		errorServerError(w)
		return
	}

	response, err := json.MarshalIndent(submitResponse, "", "  ")
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Error("Cannot Marshal submitResponse")
		errorServerError(w)
		return
	}

	if submitResponse.Ledger != nil {
		w.Write(response)
	} else {
		errorBadRequest(w, string(response))
	}

}
