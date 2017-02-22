package handlers

import (
	"bytes"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	log "github.com/Sirupsen/logrus"

	"github.com/stellar/gateway/db/entities"
	"github.com/stellar/gateway/protocols"
	callback "github.com/stellar/gateway/protocols/compliance"
	"github.com/stellar/gateway/server"
	"github.com/stellar/gateway/submitter"
	baseAmount "github.com/stellar/go/amount"
	"github.com/stellar/go/protocols/compliance"
	"github.com/stellar/go/xdr"
	"github.com/zenazn/goji/web"
)

// HandlerAuth implements authorize endpoint
func (rh *RequestHandler) HandlerAuth(c web.C, w http.ResponseWriter, r *http.Request) {
	authreq := &compliance.AuthRequest{
		DataJSON:  r.PostFormValue("data"),
		Signature: r.PostFormValue("sig"),
	}

	log.WithFields(log.Fields{"data": authreq.DataJSON, "sig": authreq.Signature}).Info("HandlerAuth")

	err := authreq.Validate()
	if err != nil {
		log.WithFields(log.Fields{"err": err}).Info(err.Error())
		server.Write(w, protocols.InvalidParameterError)
		return
	}

	authData, err := authreq.Data()
	if err != nil {
		log.WithFields(log.Fields{"err": err}).Error(err.Error())
		server.Write(w, protocols.InternalServerError)
		return
	}

	senderStellarToml, err := rh.StellarTomlResolver.GetStellarTomlByAddress(authData.Sender)
	if err != nil {
		log.WithFields(log.Fields{"err": err, "sender": authData.Sender}).Warn("Cannot get stellar.toml of sender")
		server.Write(w, protocols.InvalidParameterError)
		return
	}

	if senderStellarToml.SigningKey == "" {
		errorResponse := protocols.NewInvalidParameterError("data.sender", authData.Sender)
		log.WithFields(errorResponse.LogData).Warn("No SIGNING_KEY in stellar.toml of sender")
		server.Write(w, errorResponse)
		return
	}

	// Verify signature
	signatureBytes, err := base64.StdEncoding.DecodeString(authreq.Signature)
	if err != nil {
		errorResponse := protocols.NewInvalidParameterError("sig", authreq.Signature)
		log.WithFields(errorResponse.LogData).Warn("Error decoding signature")
		server.Write(w, errorResponse)
		return
	}
	err = rh.SignatureSignerVerifier.Verify(senderStellarToml.SigningKey, []byte(authreq.DataJSON), signatureBytes)
	if err != nil {
		log.WithFields(log.Fields{
			"signing_key": senderStellarToml.SigningKey,
			"data":        authreq.Data,
			"sig":         authreq.Signature,
		}).Warn("Invalid signature")
		errorResponse := protocols.NewInvalidParameterError("sig", authreq.Signature)
		server.Write(w, errorResponse)
		return
	}

	b64r := base64.NewDecoder(base64.StdEncoding, strings.NewReader(authData.Tx))
	var tx xdr.Transaction
	_, err = xdr.Unmarshal(b64r, &tx)
	if err != nil {
		errorResponse := protocols.NewInvalidParameterError("data.tx", authData.Tx)
		log.WithFields(log.Fields{
			"err": err,
			"tx":  authData.Tx,
		}).Warn("Error decoding Transaction XDR")
		server.Write(w, errorResponse)
		return
	}

	if tx.Memo.Hash == nil {
		errorResponse := protocols.NewInvalidParameterError("data.tx", authData.Tx)
		log.WithFields(log.Fields{"tx": authData.Tx}).Warn("Transaction does not contain Memo.Hash")
		server.Write(w, errorResponse)
		return
	}

	// Validate memo preimage hash
	memoPreimageHashBytes := sha256.Sum256([]byte(authData.AttachmentJSON))
	memoBytes := [32]byte(*tx.Memo.Hash)

	if memoPreimageHashBytes != memoBytes {
		errorResponse := protocols.NewInvalidParameterError("data.tx", authData.Tx)

		h := xdr.Hash(memoPreimageHashBytes)
		tx.Memo.Hash = &h

		var txBytes bytes.Buffer
		_, err = xdr.Marshal(&txBytes, tx)
		if err != nil {
			log.Error("Error mashaling transaction")
			server.Write(w, protocols.InternalServerError)
			return
		}

		expectedTx := base64.StdEncoding.EncodeToString(txBytes.Bytes())

		log.WithFields(log.Fields{"tx": authData.Tx, "expected_tx": expectedTx}).Warn("Memo preimage hash does not equal tx Memo.Hash")
		server.Write(w, errorResponse)
		return
	}

	attachment, err := authData.Attachment()
	if err != nil {
		log.WithFields(log.Fields{"err": err}).Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	transactionHash, err := submitter.TransactionHash(&tx, rh.Config.NetworkPassphrase)
	if err != nil {
		log.WithFields(log.Fields{"err": err}).Warn("Error calculating tx hash")
		server.Write(w, protocols.InternalServerError)
		return
	}

	response := compliance.AuthResponse{}

	// Sanctions check
	if rh.Config.Callbacks.Sanctions == "" {
		response.TxStatus = compliance.AuthStatusOk
	} else {
		senderInfo, err := json.Marshal(attachment.Transaction.SenderInfo)
		if err != nil {
			log.WithFields(log.Fields{"err": err}).Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}

		resp, err := rh.Client.PostForm(
			rh.Config.Callbacks.Sanctions,
			url.Values{"sender": {string(senderInfo)}},
		)
		if err != nil {
			log.WithFields(log.Fields{
				"sanctions": rh.Config.Callbacks.Sanctions,
				"err":       err,
			}).Error("Error sending request to sanctions server")
			server.Write(w, protocols.InternalServerError)
			return
		}

		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Error("Error reading sanctions server response")
			server.Write(w, protocols.InternalServerError)
			return
		}

		switch resp.StatusCode {
		case http.StatusOK: // AuthStatusOk
			response.TxStatus = compliance.AuthStatusOk
		case http.StatusAccepted: // AuthStatusPending
			response.TxStatus = compliance.AuthStatusPending

			pendingResponse := struct {
				Pending int `json:"pending"`
			}{}
			err := json.Unmarshal(body, &pendingResponse)
			if err != nil {
				// Set default value
				response.Pending = 600
			} else {
				response.Pending = pendingResponse.Pending
			}
		case http.StatusForbidden: // AuthStatusDenied
			response.TxStatus = compliance.AuthStatusDenied
		default:
			log.WithFields(log.Fields{
				"status": resp.StatusCode,
				"body":   string(body),
			}).Error("Error response from sanctions server")
			server.Write(w, protocols.InternalServerError)
			return
		}
	}

	// User info
	if authData.NeedInfo {
		if rh.Config.Callbacks.AskUser == "" {
			response.InfoStatus = compliance.AuthStatusDenied

			// Check AllowedFi
			tokens := strings.Split(authData.Sender, "*")
			if len(tokens) != 2 {
				log.WithFields(log.Fields{
					"sender": authData.Sender,
				}).Warn("Invalid stellar address")
				server.Write(w, protocols.InternalServerError)
				return
			}

			allowedFi, err := rh.Repository.GetAllowedFiByDomain(tokens[1])
			if err != nil {
				log.WithFields(log.Fields{"err": err}).Error("Error getting AllowedFi from DB")
				server.Write(w, protocols.InternalServerError)
				return
			}

			if allowedFi == nil {
				// FI not found check AllowedUser
				allowedUser, err := rh.Repository.GetAllowedUserByDomainAndUserID(tokens[1], tokens[0])
				if err != nil {
					log.WithFields(log.Fields{"err": err}).Error("Error getting AllowedUser from DB")
					server.Write(w, protocols.InternalServerError)
					return
				}

				if allowedUser != nil {
					response.InfoStatus = compliance.AuthStatusOk
				}
			} else {
				response.InfoStatus = compliance.AuthStatusOk
			}
		} else {
			// Ask user
			var amount, assetType, assetCode, assetIssuer string

			if len(tx.Operations) > 0 {
				operationBody := tx.Operations[0].Body
				if operationBody.Type == xdr.OperationTypePayment {
					amount = baseAmount.String(operationBody.PaymentOp.Amount)
					operationBody.PaymentOp.Asset.Extract(&assetType, &assetCode, &assetIssuer)
				} else if operationBody.Type == xdr.OperationTypePathPayment {
					amount = baseAmount.String(operationBody.PathPaymentOp.DestAmount)
					operationBody.PathPaymentOp.DestAsset.Extract(&assetType, &assetCode, &assetIssuer)
				}
			}

			senderInfo, err := json.Marshal(attachment.Transaction.SenderInfo)
			if err != nil {
				log.WithFields(log.Fields{"err": err}).Error(err.Error())
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
			}

			resp, err := rh.Client.PostForm(
				rh.Config.Callbacks.AskUser,
				url.Values{
					"amount":       {amount},
					"asset_code":   {assetCode},
					"asset_issuer": {assetIssuer},
					"sender":       {string(senderInfo)},
					"note":         {attachment.Transaction.Note},
				},
			)
			if err != nil {
				log.WithFields(log.Fields{
					"ask_user": rh.Config.Callbacks.AskUser,
					"err":      err,
				}).Error("Error sending request to ask_user server")
				server.Write(w, protocols.InternalServerError)
				return
			}

			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Error("Error reading ask_user server response")
				server.Write(w, protocols.InternalServerError)
				return
			}

			switch resp.StatusCode {
			case http.StatusOK: // AuthStatusOk
				response.InfoStatus = compliance.AuthStatusOk
			case http.StatusAccepted: // AuthStatusPending
				response.InfoStatus = compliance.AuthStatusPending

				pendingResponse := struct {
					Pending int `json:"pending"`
				}{}
				err := json.Unmarshal(body, &pendingResponse)
				if err != nil {
					// Set default value
					response.Pending = 600
				} else {
					response.Pending = pendingResponse.Pending
				}
			case http.StatusForbidden: // AuthStatusDenied
				response.InfoStatus = compliance.AuthStatusDenied
			default:
				log.WithFields(log.Fields{
					"status": resp.StatusCode,
					"body":   string(body),
				}).Error("Error response from ask_user server")
				server.Write(w, protocols.InternalServerError)
				return
			}
		}

		if response.InfoStatus == compliance.AuthStatusOk {
			// Fetch Info
			fetchInfoRequest := callback.FetchInfoRequest{Address: attachment.Transaction.Route}
			resp, err := rh.Client.PostForm(
				rh.Config.Callbacks.FetchInfo,
				fetchInfoRequest.ToValues(),
			)
			if err != nil {
				log.WithFields(log.Fields{
					"fetch_info": rh.Config.Callbacks.FetchInfo,
					"err":        err,
				}).Error("Error sending request to fetch_info server")
				server.Write(w, protocols.InternalServerError)
				return
			}

			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.WithFields(log.Fields{
					"fetch_info": rh.Config.Callbacks.FetchInfo,
					"err":        err,
				}).Error("Error reading fetch_info server response")
				server.Write(w, protocols.InternalServerError)
				return
			}

			if resp.StatusCode != http.StatusOK {
				log.WithFields(log.Fields{
					"fetch_info": rh.Config.Callbacks.FetchInfo,
					"status":     resp.StatusCode,
					"body":       string(body),
				}).Error("Error response from fetch_info server")
				server.Write(w, protocols.InternalServerError)
				return
			}

			response.DestInfo = string(body)
		}
	} else {
		response.InfoStatus = compliance.AuthStatusOk
	}

	if response.TxStatus == compliance.AuthStatusOk && response.InfoStatus == compliance.AuthStatusOk {
		authorizedTransaction := &entities.AuthorizedTransaction{
			TransactionID:  hex.EncodeToString(transactionHash[:]),
			Memo:           base64.StdEncoding.EncodeToString(memoBytes[:]),
			TransactionXdr: authData.Tx,
			AuthorizedAt:   time.Now(),
			Data:           authreq.DataJSON,
		}
		err = rh.EntityManager.Persist(authorizedTransaction)
		if err != nil {
			log.WithFields(log.Fields{"err": err}).Warn("Error persisting AuthorizedTransaction")
			server.Write(w, protocols.InternalServerError)
			return
		}
	}

	responseBody, err := response.Marshal()
	if err != nil {
		log.WithFields(log.Fields{"err": err}).Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write(responseBody)
}
