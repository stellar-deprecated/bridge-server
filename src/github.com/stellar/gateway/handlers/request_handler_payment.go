package handlers

import (
	log "github.com/Sirupsen/logrus"
	"net/http"
	"strconv"
	"strings"

	"github.com/stellar/gateway/horizon"
	b "github.com/stellar/go-stellar-base/build"
	"github.com/stellar/go-stellar-base/keypair"
)

func (rh *RequestHandler) Payment(w http.ResponseWriter, r *http.Request) {
	source := r.PostFormValue("source")
	sourceKeypair, err := keypair.Parse(source)
	if err != nil {
		log.WithFields(log.Fields{"source": source}).Print("Invalid source parameter")
		writeError(w, horizon.PaymentInvalidSource)
		return
	}

	destination := r.PostFormValue("destination")
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

	amount := r.PostFormValue("amount")
	assetCode := r.PostFormValue("asset_code")
	assetIssuer := r.PostFormValue("asset_issuer")

	var operationBuilder interface{}

	if assetCode != "" && assetIssuer != "" {
		issuerKeypair, err := keypair.Parse(assetIssuer)
		if err != nil {
			log.WithFields(log.Fields{"asset_issuer": assetIssuer}).Print("Invalid asset_issuer parameter")
			writeError(w, horizon.PaymentInvalidIssuer)
			return
		}

		operationBuilder = b.Payment(
			b.Destination{destinationObject.AccountId},
			b.CreditAmount{assetCode, issuerKeypair.Address(), amount},
		)
	} else if assetCode == "" && assetIssuer == "" {
		mutators := []interface{}{
			b.Destination{destinationObject.AccountId},
			b.NativeAmount{amount},
		}

		// Check if destination account exist
		_, err = rh.Horizon.LoadAccount(destinationObject.AccountId)
		if err != nil {
			log.WithFields(log.Fields{"error": err}).Error("Error loading account")
			operationBuilder = b.CreateAccount(mutators...)
		} else {
			operationBuilder = b.Payment(mutators...)
		}
	} else {
		log.Print("Missing asset param.")
		writeError(w, horizon.PaymentMissingParamAsset)
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
		memoMutator = &b.MemoText{memo}
	default:
		log.Print("Not supported memo type: ", memoType)
		writeError(w, horizon.PaymentInvalidMemo)
		return
	}

	accountResponse, err := rh.Horizon.LoadAccount(sourceKeypair.Address())
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Error("Cannot load source account")
		writeError(w, horizon.PaymentSourceNotExist)
		return
	}

	sequenceNumber, err := strconv.ParseUint(accountResponse.SequenceNumber, 10, 64)
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Error("Cannot convert SequenceNumber")
		writeError(w, horizon.ServerError)
		return
	}

	transactionMutators := []b.TransactionMutator{
		b.SourceAccount{source},
		b.Sequence{sequenceNumber + 1},
		b.Network{rh.Config.NetworkPassphrase},
		operationBuilder.(b.TransactionMutator),
	}

	if memoMutator != nil {
		transactionMutators = append(transactionMutators, memoMutator.(b.TransactionMutator))
	}

	tx := b.Transaction(transactionMutators...)

	if tx.Err != nil {
		log.WithFields(log.Fields{"err": tx.Err}).Print("Transaction builder error")
		// TODO when build.OperationBuilder interface is ready check for
		// create_account and payment errors separately
		switch {
		case tx.Err.Error() == "Asset code length is invalid":
			writeError(w, horizon.PaymentMalformedAssetCode)
		case strings.Contains(tx.Err.Error(), "cannot parse amount"):
			writeError(w, horizon.PaymentInvalidAmount)
		default:
			log.WithFields(log.Fields{"err": tx.Err}).Print("Transaction builder error")
			writeError(w, horizon.ServerError)
		}
		return
	}

	txe := tx.Sign(source)
	txeB64, err := txe.Base64()

	if err != nil {
		log.WithFields(log.Fields{"error": err}).Error("Cannot encode transaction envelope")
		writeError(w, horizon.ServerError)
		return
	}

	submitResponse, err := rh.Horizon.SubmitTransaction(txeB64)
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Error("Error submitting transaction")
		writeError(w, horizon.ServerError)
		return
	}

	write(w, submitResponse)
}
