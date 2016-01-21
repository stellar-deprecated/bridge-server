package gateway

import (
	"encoding/json"
	log "github.com/Sirupsen/logrus"
	"net/http"
	"strconv"

	"github.com/stellar/gateway/horizon"
	b "github.com/stellar/go-stellar-base/build"
	"github.com/stellar/go-stellar-base/keypair"
)

type RequestHandler struct {
	config               *Config
	horizon              horizon.HorizonInterface
	transactionSubmitter TransactionSubmitterInterface
}

func (rh *RequestHandler) Payment(w http.ResponseWriter, r *http.Request) {
	source := r.URL.Query().Get("source")
	sourceKeypair, err := keypair.Parse(source)
	if err != nil {
		log.WithFields(log.Fields{"source": source}).Print("Invalid source parameter")
		errorBadRequest(w, errorResponseString("invalid_source", "source parameter is invalid"))
		return
	}

	destination := r.URL.Query().Get("destination")
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

	amount := r.URL.Query().Get("amount")
	assetCode := r.URL.Query().Get("asset_code")
	assetIssuer := r.URL.Query().Get("asset_issuer")

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

	memoType := r.URL.Query().Get("memo_type")
	memo := r.URL.Query().Get("memo")

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

	accountResponse, err := rh.horizon.LoadAccount(sourceKeypair.Address())
	if err != nil {
		log.Error("Cannot load source account ", err)
		errorServerError(w)
		return
	}

	transactionMutators := []b.TransactionMutator{
		b.SourceAccount{source},
		b.Sequence{accountResponse.SequenceNumber + 1},
		paymentOperation,
	}

	if memoMutator != nil {
		transactionMutators = append(transactionMutators, memoMutator.(b.TransactionMutator))
	}

	tx := b.Transaction(transactionMutators...)

	txe := tx.Sign(source)
	txeB64, err := txe.Base64()

	if err != nil {
		log.WithFields(log.Fields{"error": err}).Error("Cannot encode transaction envelope")
		errorServerError(w)
		return
	}

	submitResponse, err := rh.horizon.SubmitTransaction(txeB64)
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
	w.Write(response)
}

func (rh *RequestHandler) Authorize(w http.ResponseWriter, r *http.Request) {
	accountId := r.URL.Query().Get("accountId")
	assetCode := r.URL.Query().Get("assetCode")

	_, err := keypair.Parse(accountId)
	if err != nil {
		log.Print("Invalid accountId parameter: ", accountId)
		errorBadRequest(w, errorResponseString("invalid_account_id", "accountId parameter is invalid"))
		return
	}

	if !rh.isAssetAllowed(assetCode) {
		log.Print("Asset code not allowed: ", assetCode)
		errorBadRequest(w, errorResponseString("invalid_asset_code", "Given assetCode not allowed"))
		return
	}

	operation := b.AllowTrust(
		b.Trustor{accountId},
		b.Authorize{true},
		b.AllowTrustAsset{assetCode},
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

func (rh *RequestHandler) Send(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	destination := q.Get("destination")
	assetCode := q.Get("assetCode")
	amount := q.Get("amount")

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

	issuingKeypair, err := keypair.Parse(rh.config.Accounts.IssuingSeed)
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

	submitResponse, err := rh.transactionSubmitter.SubmitTransaction(
		rh.config.Accounts.IssuingSeed,
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

// TODO this is duplicated in PaymentListener
func (rh *RequestHandler) isAssetAllowed(code string) bool {
	for _, b := range rh.config.Assets {
		if b == code {
			return true
		}
	}
	return false
}
