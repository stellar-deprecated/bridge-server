package handlers

import (
	"bytes"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"net/http"
	"strings"
	"time"

	"github.com/stellar/gateway/db"
	"github.com/stellar/go-stellar-base/hash"
	"github.com/stellar/go-stellar-base/xdr"
	"github.com/zenazn/goji/web"
)

const (
	AUTH_STATUS_OK      = "ok"
	AUTH_STATUS_PENDING = "pending"
	AUTH_STATUS_DENIED  = "denied"
)

type AuthData struct {
	// The stellar address of the customer that is initiating the send.
	Sender string
	// If the caller needs the recipient's AML info in order to send the payment.
	NeedInfo bool
	// The transaction that the sender would like to send in XDR format. This transaction is unsigned.
	Tx string
	// The full text of the memo the hash of this memo is included in the transaction.
	Memo string
}

type AuthResponse struct {
	// If this FI is willing to share AML information or not. {ok, denied, pending}
	InfoStatus string `json:"info_status"`
	// If this FI is willing to accept this transaction. {ok, denied, pending}
	TxStatus string `json:"tx_status"`
	// (only present if info_status is ok) JSON of the recipient's AML information. in the Stellar memo convention
	DestInfo string `json:"dest_info,omitempty"`
	// (only present if info_status or tx_status is pending) Estimated number of seconds till the sender can check back for a change in status. The sender should just resubmit this request after the given number of seconds.
	Pending int `json:"pending,omitempty"`
}

func (rh *RequestHandler) HandlerAuth(c web.C, w http.ResponseWriter, r *http.Request) {
	data := r.PostFormValue("data")
	// TODO:
	//sig := r.PostFormValue("sig")

	var authData AuthData
	err := json.Unmarshal([]byte(data), &authData)
	if err != nil {
		log.WithFields(log.Fields{"data": data}).Warn("Invalid param")
		WriteError(w, InvalidParameter)
		return
	}

	b64r := base64.NewDecoder(base64.StdEncoding, strings.NewReader(authData.Tx))
	var tx xdr.Transaction
	_, err = xdr.Unmarshal(b64r, &tx)

	if err != nil {
		log.WithFields(log.Fields{"err": err}).Warn("Error decoding Transaction XDR")
		WriteError(w, InvalidParameter)
		return
	}

	// TODO check if hashed memo is equal to data.Memo

	// Transaction hash - move to helper function
	var txBytes bytes.Buffer

	_, err = fmt.Fprintf(&txBytes, "%s", hash.Hash([]byte(rh.Config.NetworkPassphrase)))
	if err != nil {
		log.WithFields(log.Fields{"err": err}).Warn("Error calculating tx hash")
		WriteError(w, InternalServerError)
		return
	}

	_, err = xdr.Marshal(&txBytes, xdr.EnvelopeTypeEnvelopeTypeTx)
	if err != nil {
		log.WithFields(log.Fields{"err": err}).Warn("Error calculating tx hash")
		WriteError(w, InternalServerError)
		return
	}

	_, err = xdr.Marshal(&txBytes, tx)
	if err != nil {
		log.WithFields(log.Fields{"err": err}).Warn("Error calculating tx hash")
		WriteError(w, InternalServerError)
		return
	}

	transactionHash := hash.Hash(txBytes.Bytes())

	authorizedTransaction := &db.AuthorizedTransaction{
		TransactionId:  hex.EncodeToString(transactionHash[:]),
		TransactionXdr: authData.Tx,
		AuthorizedAt:   time.Now(),
		Data:           data,
	}
	err = rh.EntityManager.Persist(authorizedTransaction)
	if err != nil {
		log.WithFields(log.Fields{"err": err}).Warn("Error persisting AuthorizedTransaction")
		WriteError(w, InternalServerError)
		return
	}

	response := AuthResponse{
		InfoStatus: AUTH_STATUS_DENIED,
		TxStatus:   AUTH_STATUS_OK,
	}

	responseJson, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		log.WithFields(log.Fields{"err": err}).Error("Error marshaling response")
		WriteError(w, InternalServerError)
		return
	}

	w.Write(responseJson)
}
