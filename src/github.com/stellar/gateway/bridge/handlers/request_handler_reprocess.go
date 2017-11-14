package handlers

import (
	"net/http"

	log "github.com/sirupsen/logrus"
	"github.com/stellar/gateway/protocols"
	"github.com/stellar/gateway/protocols/bridge"
	"github.com/stellar/gateway/server"
)

// Authorize implements /reprocess endpoint
func (rh *RequestHandler) Reprocess(w http.ResponseWriter, r *http.Request) {
	request := &bridge.ReprocessRequest{}
	err := request.FromRequest(r)
	if err != nil {
		log.Error(err.Error())
		server.Write(w, protocols.InvalidParameterError)
		return
	}

	err = request.Validate()
	if err != nil {
		errorResponse := err.(*protocols.ErrorResponse)
		log.WithFields(errorResponse.LogData).Error(errorResponse.Error())
		server.Write(w, errorResponse)
		return
	}

	operation, err := rh.Horizon.LoadOperation(request.OperationID)
	if err != nil {
		server.Write(w, &bridge.ReprocessResponse{Status: "error", Message: err.Error()})
		return
	}

	err = rh.PaymentListener.ReprocessPayment(operation, request.Force)

	if err != nil {
		server.Write(w, &bridge.ReprocessResponse{Status: "error", Message: err.Error()})
		return
	}

	server.Write(w, &bridge.ReprocessResponse{Status: "ok"})
}
