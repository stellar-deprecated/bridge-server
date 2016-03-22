package handlers

import (
	log "github.com/Sirupsen/logrus"
	"net/http"

	"github.com/stellar/gateway/protocols/compliance"
	"github.com/stellar/gateway/server"
	"github.com/zenazn/goji/web"
)

func (rh *RequestHandler) HandlerReceive(c web.C, w http.ResponseWriter, r *http.Request) {
	requestMemo := r.PostFormValue("memo")

	authorizedTransaction, err := rh.Repository.GetAuthorizedTransactionByMemo(requestMemo)
	if err != nil {
		log.WithFields(log.Fields{"err": err}).Error("Error getting authorizedTransaction")
		server.Write(w, compliance.InternalServerError)
		return
	}

	if authorizedTransaction == nil {
		log.WithFields(log.Fields{"memo": requestMemo}).Warn("authorizedTransaction not found")
		server.Write(w, compliance.TransactionNotFoundError)
		return
	}

	response := compliance.ReceiveResponse{Memo: authorizedTransaction.Data}
	server.Write(w, &response)
}
