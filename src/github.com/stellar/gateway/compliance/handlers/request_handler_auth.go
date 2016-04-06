package handlers

import (
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	log "github.com/Sirupsen/logrus"
	"net/http"
	"strings"
	"time"

	"github.com/stellar/gateway/crypto"
	"github.com/stellar/gateway/db/entities"
	"github.com/stellar/gateway/protocols"
	"github.com/stellar/gateway/protocols/compliance"
	"github.com/stellar/gateway/server"
	"github.com/stellar/gateway/submitter"
	"github.com/stellar/go-stellar-base/hash"
	"github.com/stellar/go-stellar-base/xdr"
	"github.com/zenazn/goji/web"
)

func (rh *RequestHandler) HandlerAuth(c web.C, w http.ResponseWriter, r *http.Request) {
	request := &compliance.AuthRequest{}
	request.FromRequest(r)

	err := request.Validate()
	if err != nil {
		errorResponse := err.(*protocols.ErrorResponse)
		log.WithFields(errorResponse.LogData).Error(errorResponse.Error())
		server.Write(w, errorResponse)
		return
	}

	var authData compliance.AuthData
	err = json.Unmarshal([]byte(request.Data), &authData)
	if err != nil {
		errorResponse := protocols.NewInvalidParameterError("data", request.Data)
		log.WithFields(errorResponse.LogData).Warn(errorResponse.Error())
		server.Write(w, errorResponse)
		return
	}

	senderStellarToml, err := rh.StellarTomlResolver.GetStellarTomlByAddress(authData.Sender)
	if err != nil {
		log.WithFields(log.Fields{"err": err, "sender": authData.Sender}).Warn("Cannot get stellar.toml of sender")
		server.Write(w, protocols.InvalidParameterError)
		return
	}

	if senderStellarToml.SigningKey == nil {
		errorResponse := protocols.NewInvalidParameterError("data.sender", authData.Sender)
		log.WithFields(errorResponse.LogData).Warn("No SIGNING_KEY in stellar.toml of sender")
		server.Write(w, errorResponse)
		return
	}

	// Verify signature
	signatureBytes, err := base64.StdEncoding.DecodeString(request.Signature)
	if err != nil {
		errorResponse := protocols.NewInvalidParameterError("sig", request.Signature)
		log.WithFields(errorResponse.LogData).Warn("Error decoding signature")
		server.Write(w, errorResponse)
		return
	}
	err = crypto.Verify(*senderStellarToml.SigningKey, []byte(request.Data), signatureBytes)
	if err != nil {
		log.WithFields(log.Fields{
			"signing_key": *senderStellarToml.SigningKey,
			"data":        request.Data,
			"sig":         request.Signature,
		}).Warn("Invalid signature")
		errorResponse := protocols.NewInvalidParameterError("sig", request.Signature)
		server.Write(w, errorResponse)
		return
	}

	b64r := base64.NewDecoder(base64.StdEncoding, strings.NewReader(authData.Tx))
	var tx xdr.Transaction
	_, err = xdr.Unmarshal(b64r, &tx)

	if err != nil {
		log.WithFields(log.Fields{"err": err}).Warn("Error decoding Transaction XDR")
		server.Write(w, protocols.InvalidParameterError)
		return
	}

	var memo *string

	if tx.Memo.Hash != nil {
		memoBytes := [32]byte(*tx.Memo.Hash)
		memoBase64 := base64.StdEncoding.EncodeToString(memoBytes[:])
		memo = &memoBase64
	}

	transactionHashBytes, err := submitter.TransactionHash(&tx, rh.Config.NetworkPassphrase)
	if err != nil {
		log.WithFields(log.Fields{"err": err}).Warn("Error calculating tx hash")
		server.Write(w, protocols.InternalServerError)
		return
	}

	transactionHash := hash.Hash(transactionHashBytes[:])

	authorizedTransaction := &entities.AuthorizedTransaction{
		TransactionId:  hex.EncodeToString(transactionHash[:]),
		Memo:           memo,
		TransactionXdr: authData.Tx,
		AuthorizedAt:   time.Now(),
		Data:           request.Data,
	}
	err = rh.EntityManager.Persist(authorizedTransaction)
	if err != nil {
		log.WithFields(log.Fields{"err": err}).Warn("Error persisting AuthorizedTransaction")
		server.Write(w, protocols.InternalServerError)
		return
	}

	response := compliance.AuthResponse{
		InfoStatus: compliance.AuthStatusDenied,
		TxStatus:   compliance.AuthStatusOk,
	}
	server.Write(w, &response)
}
