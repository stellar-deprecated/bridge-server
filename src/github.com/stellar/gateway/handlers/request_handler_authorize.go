package handlers

import (
	log "github.com/Sirupsen/logrus"
	"net/http"

	"github.com/stellar/gateway/horizon"
	b "github.com/stellar/go-stellar-base/build"
	"github.com/stellar/go-stellar-base/keypair"
)

func (rh *RequestHandler) Authorize(w http.ResponseWriter, r *http.Request) {
	accountId := r.PostFormValue("account_id")
	assetCode := r.PostFormValue("asset_code")

	_, err := keypair.Parse(accountId)
	if err != nil {
		log.Print("Invalid accountId parameter: ", accountId)
		writeError(w, horizon.AllowTrustInvalidAccountId)
		return
	}

	if !rh.isAssetAllowed(assetCode) {
		log.Print("Asset code not allowed: ", assetCode)
		writeError(w, horizon.AllowTrustAssetCodeNotAllowed)
		return
	}

	operationMutator := b.AllowTrust(
		b.Trustor{accountId},
		b.Authorize{true},
		b.AllowTrustAsset{assetCode},
	)

	submitResponse, err := rh.TransactionSubmitter.SubmitTransaction(
		*rh.Config.Accounts.AuthorizingSeed,
		operationMutator,
		nil,
	)

	if err != nil {
		log.WithFields(log.Fields{"err": err}).Error("Error submitting transaction")
		writeError(w, horizon.ServerError)
		return
	}

	write(w, submitResponse)
}
