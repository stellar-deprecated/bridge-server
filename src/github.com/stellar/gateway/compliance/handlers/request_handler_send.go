package handlers

import (
	"bytes"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	log "github.com/Sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/stellar/gateway/crypto"
	"github.com/stellar/gateway/protocols"
	"github.com/stellar/gateway/protocols/compliance"
	"github.com/stellar/gateway/server"
	"github.com/stellar/gateway/submitter"
	b "github.com/stellar/go-stellar-base/build"
	"github.com/stellar/go-stellar-base/xdr"
	"github.com/zenazn/goji/web"
)

func (rh *RequestHandler) HandlerSend(c web.C, w http.ResponseWriter, r *http.Request) {
	request := &compliance.SendRequest{}
	request.FromRequest(r)

	err := request.Validate()
	if err != nil {
		errorResponse := err.(*protocols.ErrorResponse)
		log.WithFields(errorResponse.LogData).Error(errorResponse.Error())
		server.Write(w, errorResponse)
		return
	}

	destinationObject, stellarToml, err := rh.FederationResolver.Resolve(request.Destination)
	if err != nil {
		log.WithFields(log.Fields{
			"destination": request.Destination,
			"err":         err,
		}).Print("Cannot resolve address")
		server.Write(w, compliance.CannotResolveDestination)
		return
	}

	if stellarToml.AuthServer == nil {
		log.Print("No AUTH_SERVER in stellar.toml")
		server.Write(w, compliance.AuthServerNotDefined)
		return
	}

	var payWithMutator *b.PayWithPath

	if request.SendMax != "" {
		// Path payment
		var sendAsset b.Asset
		if request.SendAssetCode != "" && request.SendAssetIssuer != "" {
			sendAsset = b.CreditAsset(request.SendAssetCode, request.SendAssetIssuer)
		} else if request.SendAssetCode == "" && request.SendAssetIssuer == "" {
			sendAsset = b.NativeAsset()
		} else {
			log.Print("Missing send asset param.")
			server.Write(w, protocols.MissingParameterError)
			return
		}

		payWith := b.PayWith(sendAsset, request.SendMax)

		for _, asset := range request.Path {
			if asset.Code == "" && asset.Issuer == "" {
				payWith = payWith.Through(b.NativeAsset())
			} else {
				payWith = payWith.Through(b.CreditAsset(asset.Code, asset.Issuer))
			}
		}

		payWithMutator = &payWith
	}

	mutators := []interface{}{
		b.Destination{destinationObject.AccountId},
		b.CreditAmount{
			request.AssetCode,
			request.AssetIssuer,
			request.Amount,
		},
	}

	if payWithMutator != nil {
		mutators = append(mutators, *payWithMutator)
	}

	operationMutator := b.Payment(mutators...)
	if operationMutator.Err != nil {
		log.WithFields(log.Fields{
			"err": operationMutator.Err,
		}).Error("Error creating operation")
		server.Write(w, protocols.InternalServerError)
		return
	}

	memoBytes := sha256.Sum256([]byte(request.ExtraMemo))
	var b32 [32]byte
	copy(b32[:], memoBytes[0:32])
	hash := xdr.Hash(b32)
	memoMutator := &b.MemoHash{hash}

	transaction, err := submitter.BuildTransaction(
		request.Source,
		rh.Config.NetworkPassphrase,
		operationMutator,
		memoMutator,
	)

	var txBytes bytes.Buffer
	_, err = xdr.Marshal(&txBytes, transaction)
	if err != nil {
		log.Error("Error mashaling transaction")
		server.Write(w, protocols.InternalServerError)
		return
	}

	txBase64 := base64.StdEncoding.EncodeToString(txBytes.Bytes())

	authData := compliance.AuthData{
		Sender:   request.Sender,
		NeedInfo: true,
		Tx:       txBase64,
		Memo:     request.ExtraMemo,
	}

	data, err := json.Marshal(authData)
	if err != nil {
		log.Error("Error mashaling authData")
		server.Write(w, protocols.InternalServerError)
		return
	}

	sig, err := crypto.Sign(rh.Config.Keys.SigningSeed, data)
	if err != nil {
		log.Error("Error signing authData")
		server.Write(w, protocols.InternalServerError)
		return
	}

	resp, err := rh.Client.PostForm(
		*stellarToml.AuthServer,
		url.Values{
			"data": {string(data)},
			"sig":  {sig},
		},
	)
	if err != nil {
		log.WithFields(log.Fields{
			"auth_server": stellarToml.AuthServer,
			"err":         err,
		}).Error("Error sending request to auth server")
		server.Write(w, protocols.InternalServerError)
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error("Error reading auth server response")
		server.Write(w, protocols.InternalServerError)
		return
	}

	if resp.StatusCode != 200 {
		log.WithFields(log.Fields{
			"status": resp.StatusCode,
			"body":   string(body),
		}).Error("Error response from auth server")
		server.Write(w, protocols.InternalServerError)
		return
	}

	response := compliance.SendResponse{TransactionXdr: txBase64}
	server.Write(w, &response)
}
