package handlers

import (
	"encoding/json"
	log "github.com/Sirupsen/logrus"
	"net/http"

	b "github.com/stellar/go-stellar-base/build"
	"github.com/stellar/go-stellar-base/keypair"
)

func (rh *RequestHandler) Authorize(w http.ResponseWriter, r *http.Request) {
	accountId := r.PostFormValue("account_id")
	assetCode := r.PostFormValue("asset_code")

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

	submitResponse, err := rh.TransactionSubmitter.SubmitTransaction(
		rh.Config.Accounts.AuthorizingSeed,
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
