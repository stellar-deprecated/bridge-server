package handlers

import (
	"encoding/hex"
	"encoding/json"
	log "github.com/Sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"

	h "github.com/stellar/gateway/horizon"
	"github.com/stellar/gateway/protocols/compliance"
	"github.com/stellar/gateway/protocols/federation"
	"github.com/stellar/gateway/server"
	b "github.com/stellar/go-stellar-base/build"
	"github.com/stellar/go-stellar-base/keypair"
	"github.com/stellar/go-stellar-base/xdr"
)

func (rh *RequestHandler) Send(w http.ResponseWriter, r *http.Request) {
	sender := r.PostFormValue("sender")
	destination := r.PostFormValue("destination")
	assetCode := r.PostFormValue("asset_code")
	assetIssuer := r.PostFormValue("asset_issuer")
	amount := r.PostFormValue("amount")
	memoType := r.PostFormValue("memo_type")
	memo := r.PostFormValue("memo")
	extraMemo := r.PostFormValue("extra_memo")

	if extraMemo != "" && rh.Config.Compliance != nil {
		// Compliance server part
		sourceKeypair, err := keypair.Parse(*rh.Config.Accounts.IssuingSeed)
		if err != nil {
			log.Error("Invalid IssuingSeed")
			server.Write(w, h.NewErrorResponse(h.ServerError))
			return
		}

		resp, err := http.PostForm(
			*rh.Config.Compliance+"/send",
			url.Values{
				"source":       {sourceKeypair.Address()},
				"sender":       {sender},
				"destination":  {destination},
				"amount":       {amount},
				"asset_code":   {assetCode},
				"asset_issuer": {assetIssuer},
				"extra_memo":   {extraMemo},
			},
		)
		if err != nil {
			log.WithFields(log.Fields{"err": err}).Error("Error sending request to compliance server")
			server.Write(w, h.NewErrorResponse(h.ServerError))
			return
		}

		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Error("Error reading compliance server response")
			server.Write(w, h.NewErrorResponse(h.ServerError))
			return
		}

		if resp.StatusCode != 200 {
			log.WithFields(log.Fields{
				"status": resp.StatusCode,
				"body":   string(body),
			}).Error("Error response from compliance server")
			server.Write(w, h.NewErrorResponse(h.ServerError))
			return
		}

		var complianceSendResponse compliance.SendResponse
		err = json.Unmarshal(body, &complianceSendResponse)
		if err != nil {
			log.Error("Error unmarshalling from compliance server")
			server.Write(w, h.NewErrorResponse(h.ServerError))
			return
		}

		var tx xdr.Transaction
		err = xdr.SafeUnmarshalBase64(complianceSendResponse.TransactionXdr, &tx)
		if err != nil {
			log.Error("Error unmarshalling transaction returned by compliance server")
			server.Write(w, h.NewErrorResponse(h.ServerError))
			return
		}

		submitResponse, err := rh.TransactionSubmitter.SignAndSubmitRawTransaction(*rh.Config.Accounts.IssuingSeed, &tx)
		if err != nil {
			log.WithFields(log.Fields{"err": err}).Error("Error submitting transaction")
			server.Write(w, h.NewErrorResponse(h.ServerError))
			return
		}

		server.Write(w, &submitResponse)
	} else {
		// Send without compliance server
		destinationObject, _, err := federation.Resolve(destination)
		if err != nil {
			log.WithFields(log.Fields{
				"destination": destination,
				"err":         err,
			}).Print("Cannot resolve address")
			server.Write(w, h.NewErrorResponse(h.PaymentCannotResolveDestination))
			return
		}

		_, err = keypair.Parse(destinationObject.AccountId)
		if err != nil {
			log.WithFields(log.Fields{"AccountId": destinationObject.AccountId}).Print("Invalid AccountId in destination")
			server.Write(w, h.NewErrorResponse(h.PaymentInvalidDestination))
			return
		}

		if !rh.isAssetAllowed(assetCode) {
			log.Print("Asset code not allowed: ", assetCode)
			server.Write(w, h.NewErrorResponse(h.PaymentAssetCodeNotAllowed))
			return
		}

		issuingKeypair, err := keypair.Parse(*rh.Config.Accounts.IssuingSeed)
		if err != nil {
			log.Print("Invalid issuingSeed")
			server.Write(w, h.NewErrorResponse(h.ServerError))
			return
		}

		operationMutator := b.Payment(
			b.Destination{destinationObject.AccountId},
			b.CreditAmount{assetCode, issuingKeypair.Address(), amount},
		)
		if operationMutator.Err != nil {
			log.Print("Error creating operationMutator ", operationMutator.Err)
			server.Write(w, h.NewErrorResponse(h.ServerError))
			return
		}

		var memoMutator interface{}

		if !(((memoType == "") && (memo == "")) || ((memoType != "") && (memo != ""))) {
			log.Print("Missing one of memo params.")
			server.Write(w, h.NewErrorResponse(h.PaymentMissingParamMemo))
			return
		}

		if destinationObject.MemoType != nil {
			if memoType != "" {
				log.Print("Memo given in request but federation returned memo fields.")
				server.Write(w, h.NewErrorResponse(h.PaymentCannotUseMemo))
				return
			}

			memoType = *destinationObject.MemoType
			memo = *destinationObject.Memo
		}

		switch {
		case memoType == "":
			break
		case memoType == "id":
			id, err := strconv.ParseUint(memo, 10, 64)
			if err != nil {
				log.WithFields(log.Fields{"memo": memo}).Print("Cannot convert memo_id value to uint64")
				server.Write(w, h.NewErrorResponse(h.PaymentInvalidMemo))
				return
			}
			memoMutator = b.MemoID{id}
		case memoType == "text":
			memoMutator = b.MemoText{memo}
		case memoType == "hash":
			memoBytes, err := hex.DecodeString(memo)
			if err != nil || len(memoBytes) != 32 {
				log.WithFields(log.Fields{"memo": memo}).Print("Cannot decode hash memo value")
				server.Write(w, h.NewErrorResponse(h.PaymentInvalidMemo))
				return
			}
			var b32 [32]byte
			copy(b32[:], memoBytes[0:32])
			hash := xdr.Hash(b32)
			memoMutator = &b.MemoHash{hash}
		default:
			log.Print("Not supported memo type: ", memoType)
			server.Write(w, h.NewErrorResponse(h.PaymentInvalidMemo))
			return
		}

		submitResponse, err := rh.TransactionSubmitter.SubmitTransaction(
			*rh.Config.Accounts.IssuingSeed,
			operationMutator,
			memoMutator,
		)
		if err != nil {
			log.WithFields(log.Fields{"err": err}).Error("Error submitting transaction")
			server.Write(w, h.NewErrorResponse(h.ServerError))
			return
		}

		server.Write(w, &submitResponse)
	}
}
