package handlers

import (
	log "github.com/Sirupsen/logrus"
	"net/http"

	h "github.com/stellar/gateway/horizon"
	"github.com/stellar/gateway/server"
	b "github.com/stellar/go-stellar-base/build"
	"github.com/stellar/go-stellar-base/keypair"
)

func (rh *RequestHandler) Authorize(w http.ResponseWriter, r *http.Request) {
	accountId := r.PostFormValue("account_id")
	assetCode := r.PostFormValue("asset_code")

	_, err := keypair.Parse(accountId)
	if err != nil {
		log.Print("Invalid accountId parameter: ", accountId)
		server.Write(w, h.NewErrorResponse(h.AllowTrustInvalidAccountId))
		return
	}

	if !rh.isAssetAllowed(assetCode, *rh.Config.Accounts.IssuingAccountId) {
		log.WithFields(log.Fields{"asset_code": assetCode}).Warn("Asset code not allowed")
		server.Write(w, h.NewErrorResponse(h.AllowTrustAssetCodeNotAllowed))
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
		server.Write(w, h.NewErrorResponse(h.ServerError))
		return
	}

	server.Write(w, &submitResponse)
}
