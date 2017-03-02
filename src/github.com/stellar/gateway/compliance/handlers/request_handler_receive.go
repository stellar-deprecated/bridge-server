package handlers

import (
	log "github.com/Sirupsen/logrus"
	"net/http"

	"github.com/stellar/gateway/protocols"
	callback "github.com/stellar/gateway/protocols/compliance"
	"github.com/stellar/gateway/server"
	"github.com/zenazn/goji/web"
)

// HandlerReceive implements /receive endpoint
func (rh *RequestHandler) HandlerReceive(c web.C, w http.ResponseWriter, r *http.Request) {
	request := &callback.ReceiveRequest{}
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

	authorizedTransaction, err := rh.Repository.GetAuthorizedTransactionByMemo(request.Memo)
	if err != nil {
		log.WithFields(log.Fields{"err": err}).Error("Error getting authorizedTransaction")
		server.Write(w, protocols.InternalServerError)
		return
	}

	if authorizedTransaction == nil {
		log.WithFields(log.Fields{"memo": request.Memo}).Warn("authorizedTransaction not found")
		server.Write(w, callback.TransactionNotFoundError)
		return
	}

	response := callback.ReceiveResponse{Data: authorizedTransaction.Data}
	server.Write(w, &response)
}
