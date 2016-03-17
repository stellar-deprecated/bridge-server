package handlers

import (
	"encoding/json"
	log "github.com/Sirupsen/logrus"
	"net/http"

	//"github.com/stellar/gateway/db"
	"github.com/zenazn/goji/web"
)

type ReceiveResponse struct {
	// The full text of the memo the hash of this memo is included in the transaction.
	Memo string `json:"memo"`
}

func (rh *RequestHandler) HandlerReceive(c web.C, w http.ResponseWriter, r *http.Request) {
	requestMemo := r.PostFormValue("memo")

	authorizedTransaction, err := rh.Repository.GetAuthorizedTransactionByMemo(requestMemo)
	if err != nil {
		log.WithFields(log.Fields{"err": err}).Error("Error getting authorizedTransaction")
		WriteError(w, InternalServerError)
		return
	}

	if authorizedTransaction == nil {
		log.WithFields(log.Fields{"memo": requestMemo}).Warn("authorizedTransaction not found")
		WriteError(w, TransactionNotFound)
		return
	}

	response := ReceiveResponse{authorizedTransaction.Data}

	responseJson, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		log.WithFields(log.Fields{"err": err}).Error("Error marshaling response")
		WriteError(w, InternalServerError)
		return
	}

	w.Write(responseJson)
}
