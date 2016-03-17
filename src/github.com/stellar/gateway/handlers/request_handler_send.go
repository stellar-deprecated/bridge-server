package handlers

import (
	"encoding/hex"
	"encoding/json"
	log "github.com/Sirupsen/logrus"
	"net/http"
	"strconv"

	"github.com/stellar/gateway/horizon"
	ch "github.com/stellar/gateway/compliance/handlers"
	b "github.com/stellar/go-stellar-base/build"
	"github.com/stellar/go-stellar-base/keypair"
	"github.com/stellar/go-stellar-base/xdr"
)

func (rh *RequestHandler) Send(w http.ResponseWriter, r *http.Request) {
	destination := r.PostFormValue("destination")
	assetCode := r.PostFormValue("asset_code")
	amount := r.PostFormValue("amount")

	destinationObject, stellarToml, err := rh.AddressResolver.Resolve(destination)
	if err != nil {
		log.WithFields(log.Fields{"destination": destination}).Print("Cannot resolve address")
		writeError(w, horizon.PaymentCannotResolveDestination)
		return
	}

	_, err = keypair.Parse(destinationObject.AccountId)
	if err != nil {
		log.WithFields(log.Fields{"AccountId": destinationObject.AccountId}).Print("Invalid AccountId in destination")
		writeError(w, horizon.PaymentInvalidDestination)
		return
	}

	if !rh.isAssetAllowed(assetCode) {
		log.Print("Asset code not allowed: ", assetCode)
		writeError(w, horizon.PaymentAssetCodeNotAllowed)
		return
	}

	issuingKeypair, err := keypair.Parse(*rh.Config.Accounts.IssuingSeed)
	if err != nil {
		log.Print("Invalid issuingSeed")
		writeError(w, horizon.ServerError)
		return
	}

	operationMutator := b.Payment(
		b.Destination{destinationObject.AccountId},
		b.CreditAmount{assetCode, issuingKeypair.Address(), amount},
	)
	if operationMutator.Err != nil {
		log.Print("Error creating operationMutator ", operationMutator.Err)
		writeError(w, horizon.ServerError)
		return
	}

	memoType := r.PostFormValue("memo_type")
	memo := r.PostFormValue("memo")
	extraMemo := r.PostFormValue("extra_memo")

	if extraMemo != "" && rh.Config.Compliance != nil {
		// Create memo using extraMemo data and send it auth server
		authData := ch.AuthData{
			Sender: "",
			NeedInfo: false,
			Tx: "",
			Memo: "",
		}

		data, err := json.Marshal(data)
		if err != nil {
			writeError(w, horizon.ServerError)
			return
		}

		client := http.Client{
			Timeout: HOOK_TIMEOUT,
		}
		resp, err := client.PostForm(
			stellarToml.AuthServer,
			url.Values{"data": {data}},
		)
		if err != nil {
			pl.log.WithFields(logrus.Fields{"auth_server": stellarToml.AuthServer}).Error("Error sending request to auth server")
			writeError(w, horizon.ServerError)
			return
		}

		if resp.StatusCode != 200 {
			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				pl.log.Error("Error reading auth server response")
				writeError(w, horizon.ServerError)
				return
			}

			pl.log.WithFields(logrus.Fields{
				"status": resp.StatusCode,
				"body":   string(body),
			}).Error("Error response from auth server")
			writeError(w, horizon.ServerError)
			return
		}

		// Success
	} else {
		if !(((memoType == "") && (memo == "")) || ((memoType != "") && (memo != ""))) {
			log.Print("Missing one of memo params.")
			writeError(w, horizon.PaymentMissingParamMemo)
			return
		}

		if destinationObject.MemoType != nil {
			if memoType != "" {
				log.Print("Memo given in request but federation returned memo fields.")
				writeError(w, horizon.PaymentCannotUseMemo)
				return
			}

			memoType = *destinationObject.MemoType
			memo = *destinationObject.Memo
		}
	}

	var memoMutator interface{}
	switch {
	case memoType == "":
		break
	case memoType == "id":
		id, err := strconv.ParseUint(memo, 10, 64)
		if err != nil {
			log.WithFields(log.Fields{"memo": memo}).Print("Cannot convert memo_id value to uint64")
			writeError(w, horizon.PaymentInvalidMemo)
			return
		}
		memoMutator = b.MemoID{id}
	case memoType == "text":
		memoMutator = b.MemoText{memo}
	case memoType == "hash":
		memoBytes, err := hex.DecodeString(memo)
		if err != nil || len(memoBytes) != 32 {
			log.WithFields(log.Fields{"memo": memo}).Print("Cannot decode hash memo value")
			writeError(w, horizon.PaymentInvalidMemo)
			return
		}
		var b32 [32]byte
		copy(b32[:], memoBytes[0:32])
		hash := xdr.Hash(b32)
		memoMutator = &b.MemoHash{hash}
	default:
		log.Print("Not supported memo type: ", memoType)
		writeError(w, horizon.PaymentInvalidMemo)
		return
	}

	submitResponse, err := rh.TransactionSubmitter.SubmitTransaction(
		*rh.Config.Accounts.IssuingSeed,
		operationMutator,
		memoMutator,
	)
	if err != nil {
		log.WithFields(log.Fields{"err": err}).Error("Error submitting transaction")
		writeError(w, horizon.ServerError)
		return
	}

	write(w, submitResponse)
}
