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
	"github.com/stellar/gateway/protocols/compliance"
	"github.com/stellar/gateway/protocols/stellartoml"
	"github.com/stellar/gateway/server"
	"github.com/stellar/gateway/submitter"
	"github.com/stellar/go-stellar-base/hash"
	"github.com/stellar/go-stellar-base/xdr"
	"github.com/zenazn/goji/web"
)

func (rh *RequestHandler) HandlerAuth(c web.C, w http.ResponseWriter, r *http.Request) {
	data := r.PostFormValue("data")
	sig := r.PostFormValue("sig")

	var authData compliance.AuthData
	err := json.Unmarshal([]byte(data), &authData)
	if err != nil {
		log.WithFields(log.Fields{"data": data}).Warn("Invalid param")
		server.Write(w, compliance.InvalidParameterError)
		return
	}

	senderStellarToml, err := stellartoml.GetStellarTomlByAddress(authData.Sender)
	if err != nil {
		log.WithFields(log.Fields{"err": err, "sender": authData.Sender}).Warn("Cannot get stellar.toml of sender")
		server.Write(w, compliance.InvalidParameterError)
		return
	}

	if senderStellarToml.SigningKey == nil {
		log.Warn("No SIGNING_KEY in stellar.toml of sender")
		server.Write(w, compliance.InvalidParameterError)
		return
	}

	// Verify signature
	signatureBytes, err := base64.StdEncoding.DecodeString(sig)
	if err != nil {
		log.WithFields(log.Fields{"sig": sig}).Warn("Error decoding signature")
		server.Write(w, compliance.InvalidParameterError)
		return
	}
	err = crypto.Verify(*senderStellarToml.SigningKey, []byte(data), signatureBytes)
	if err != nil {
		log.WithFields(log.Fields{
			"signing_key": *senderStellarToml.SigningKey,
			"data":        data,
			"sig":         sig,
		}).Warn("Invalid signature")
		server.Write(w, compliance.InvalidParameterError)
		return
	}

	b64r := base64.NewDecoder(base64.StdEncoding, strings.NewReader(authData.Tx))
	var tx xdr.Transaction
	_, err = xdr.Unmarshal(b64r, &tx)

	if err != nil {
		log.WithFields(log.Fields{"err": err}).Warn("Error decoding Transaction XDR")
		server.Write(w, compliance.InvalidParameterError)
		return
	}

	var memo *string

	if tx.Memo.Hash != nil {
		memoBytes := [32]byte(*tx.Memo.Hash)
		memoHex := base64.StdEncoding.EncodeToString(memoBytes[:])
		memo = &memoHex
	}

	transactionHashBytes, err := submitter.TransactionHash(&tx, rh.Config.NetworkPassphrase)
	if err != nil {
		log.WithFields(log.Fields{"err": err}).Warn("Error calculating tx hash")
		server.Write(w, compliance.InternalServerError)
		return
	}

	transactionHash := hash.Hash(transactionHashBytes[:])

	authorizedTransaction := &entities.AuthorizedTransaction{
		TransactionId:  hex.EncodeToString(transactionHash[:]),
		Memo:           memo,
		TransactionXdr: authData.Tx,
		AuthorizedAt:   time.Now(),
		Data:           data,
	}
	err = rh.EntityManager.Persist(authorizedTransaction)
	if err != nil {
		log.WithFields(log.Fields{"err": err}).Warn("Error persisting AuthorizedTransaction")
		server.Write(w, compliance.InternalServerError)
		return
	}

	response := compliance.AuthResponse{
		InfoStatus: compliance.AUTH_STATUS_DENIED,
		TxStatus:   compliance.AUTH_STATUS_OK,
	}
	server.Write(w, &response)
}
