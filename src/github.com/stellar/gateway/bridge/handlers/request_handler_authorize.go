package handlers

import (
	"net/http"

	log "github.com/Sirupsen/logrus"

	"github.com/stellar/gateway/protocols"
	"github.com/stellar/gateway/protocols/bridge"
	"github.com/stellar/gateway/server"
	"github.com/stellar/gateway/submitter"
	b "github.com/stellar/go/build"
	"github.com/stellar/go/clients/horizon"
	"github.com/stellar/go/support/errors"
)

// Authorize implements /authorize endpoint
func (rh *RequestHandler) Authorize(w http.ResponseWriter, r *http.Request) {
	request := &bridge.AuthorizeRequest{}
	request.FromRequest(r)

	err := request.Validate(rh.Config.Assets, rh.Config.Accounts.IssuingAccountID)
	if err != nil {
		errorResponse := err.(*protocols.ErrorResponse)
		log.WithFields(errorResponse.LogData).Error(errorResponse.Error())
		server.Write(w, errorResponse)
		return
	}

	operationMutator := b.AllowTrust(
		b.Trustor{request.AccountID},
		b.Authorize{true},
		b.AllowTrustAsset{request.AssetCode},
	)

	submitResponse, err := rh.TransactionSubmitter.SubmitTransaction(
		rh.Config.Accounts.AuthorizingSeed,
		operationMutator,
		nil,
	)

	// on success
	if err == nil {
		var response submitter.Response
		response.Populate(&submitResponse)
		server.Write(w, &response)

		// on horizon error
	} else if err, ok := errors.Cause(err).(*horizon.Error); ok {

		errorResponse := bridge.ErrorFromHorizonResponse(err)
		log.WithFields(errorResponse.LogData).Error(errorResponse.Error())
		server.Write(w, errorResponse)

		// on unknown error
	} else {
		log.WithFields(log.Fields{"error": err}).Error("Error submitting transaction")
		server.Write(w, protocols.InternalServerError)
	}

	return
}
