package handlers

import (
	"bytes"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"

	ch "github.com/stellar/gateway/compliance/handlers"
	"github.com/stellar/gateway/horizon"
	b "github.com/stellar/go-stellar-base/build"
	"github.com/stellar/go-stellar-base/keypair"
	"github.com/stellar/go-stellar-base/xdr"
)

func (rh *RequestHandler) Send(w http.ResponseWriter, r *http.Request) {
	destination := r.PostFormValue("destination")

	destinationObject, stellarToml, err := rh.AddressResolver.Resolve(destination)
	if err != nil {
		log.WithFields(log.Fields{
			"destination": destination,
			"err":         err,
		}).Print("Cannot resolve address")
		writeError(w, horizon.PaymentCannotResolveDestination)
		return
	}

	_, err = keypair.Parse(destinationObject.AccountId)
	if err != nil {
		log.WithFields(log.Fields{"AccountId": destinationObject.AccountId}).Print("Invalid AccountId in destination")
		writeError(w, horizon.PaymentInvalidDestination)
		return
	}

	var operationBuilder interface{}
	var errorResponse *horizon.SubmitTransactionResponseError

	paymentType := r.PostFormValue("type")
	switch paymentType {
	case "":
	case "payment":
		log.Println("payment")
		operationBuilder, errorResponse = rh.createSendPaymentOperation(r, destinationObject)
	case "path_payment":
		log.Println("path_payment")
		operationBuilder, errorResponse = rh.createSendPathPaymentOperation(r, destinationObject)
	default:
		writeError(w, horizon.PaymentInvalidType)
		return
	}

	if errorResponse != nil {
		writeError(w, errorResponse)
		return
	}

	memoType := r.PostFormValue("memo_type")
	memo := r.PostFormValue("memo")
	extraMemo := r.PostFormValue("extra_memo")
	var memoMutator interface{}

	if extraMemo != "" && rh.Config.Compliance != nil {
		if stellarToml.AuthServer == nil {
			log.Print("No AUTH_SERVER in stellar.toml")
			writeError(w, horizon.ServerError)
			return
		}

		memoBytes := sha256.Sum256([]byte(extraMemo))
		var b32 [32]byte
		copy(b32[:], memoBytes[0:32])
		hash := xdr.Hash(b32)
		memoMutator = &b.MemoHash{hash}

		transaction, err := rh.TransactionSubmitter.BuildTransaction(
			*rh.Config.Accounts.IssuingSeed,
			operationBuilder,
			memoMutator,
		)

		var txBytes bytes.Buffer
		_, err = xdr.Marshal(&txBytes, transaction)
		if err != nil {
			log.Print("Error mashaling transaction")
			writeError(w, horizon.ServerError)
			return
		}

		authData := ch.AuthData{
			Tx:   base64.StdEncoding.EncodeToString(txBytes.Bytes()),
			Memo: extraMemo,
		}

		data, err := json.Marshal(authData)
		if err != nil {
			writeError(w, horizon.ServerError)
			return
		}

		resp, err := http.PostForm(
			*stellarToml.AuthServer,
			url.Values{"data": {string(data)}},
		)
		if err != nil {
			log.WithFields(log.Fields{
				"auth_server": stellarToml.AuthServer,
				"err":         err,
			}).Error("Error sending request to auth server")
			writeError(w, horizon.ServerError)
			return
		}

		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Error("Error reading auth server response")
			writeError(w, horizon.ServerError)
			return
		}

		if resp.StatusCode != 200 {
			log.WithFields(log.Fields{
				"status": resp.StatusCode,
				"body":   string(body),
			}).Error("Error response from auth server")
			writeError(w, horizon.ServerError)
			return
		}
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
	}

	submitResponse, err := rh.TransactionSubmitter.SubmitTransaction(
		*rh.Config.Accounts.IssuingSeed,
		operationBuilder,
		memoMutator,
	)
	if err != nil {
		log.WithFields(log.Fields{"err": err}).Error("Error submitting transaction")
		writeError(w, horizon.ServerError)
		return
	}

	write(w, submitResponse)
}

func (rh *RequestHandler) createSendPaymentOperation(r *http.Request, destinationObject StellarDestination) (operationBuilder interface{}, errorResponse *horizon.SubmitTransactionResponseError) {
	assetCode := r.PostFormValue("asset_code")
	amount := r.PostFormValue("amount")

	if !rh.isAssetAllowed(assetCode) {
		log.Print("Asset code not allowed: ", assetCode)
		errorResponse = horizon.PaymentAssetCodeNotAllowed
		return
	}

	issuingKeypair, err := keypair.Parse(*rh.Config.Accounts.IssuingSeed)
	if err != nil {
		log.Print("Invalid issuingSeed")
		errorResponse = horizon.ServerError
		return
	}

	operationBuilder = b.Payment(
		b.Destination{destinationObject.AccountId},
		b.CreditAmount{assetCode, issuingKeypair.Address(), amount},
	)
	if operationBuilder.(b.PaymentBuilder).Err != nil {
		log.Print("Error creating operationBuilder ", operationBuilder.(b.PaymentBuilder).Err)
		errorResponse = horizon.ServerError
	}
	return
}

func (rh *RequestHandler) createSendPathPaymentOperation(r *http.Request, destinationObject StellarDestination) (operationBuilder interface{}, errorResponse *horizon.SubmitTransactionResponseError) {
	sendMax := r.PostFormValue("send_max")
	sendAssetCode := r.PostFormValue("send_asset_code")
	sendAssetIssuer := r.PostFormValue("send_asset_issuer")

	var sendAsset b.Asset
	if sendAssetCode != "" && sendAssetIssuer != "" {
		sendAsset = b.Asset{Code: sendAssetCode, Issuer: sendAssetIssuer}
	} else if sendAssetCode == "" && sendAssetIssuer == "" {
		sendAsset = b.Asset{Native: true}
	} else {
		log.Print("Missing send asset param.")
		errorResponse = horizon.PaymentMissingParamAsset
		return
	}

	destinationAmount := r.PostFormValue("destination_amount")
	destinationAssetCode := r.PostFormValue("destination_asset_code")
	destinationAssetIssuer := r.PostFormValue("destination_asset_issuer")

	var destinationAsset b.Asset
	if destinationAssetCode != "" && destinationAssetIssuer != "" {
		destinationAsset = b.Asset{Code: destinationAssetCode, Issuer: destinationAssetIssuer}
	} else if destinationAssetCode == "" && destinationAssetIssuer == "" {
		destinationAsset = b.Asset{Native: true}
	} else {
		log.Print("Missing destination asset param.")
		errorResponse = horizon.PaymentMissingParamAsset
		return
	}

	// TODO check the fields

	var path []b.Asset

	for i := 0; ; i++ {
		codeFieldName := fmt.Sprintf("path[%d][asset_code]", i)
		issuerFieldName := fmt.Sprintf("path[%d][asset_issuer]", i)

		// If the element does not exist in PostForm break the loop
		if _, exists := r.PostForm[codeFieldName]; !exists {
			break
		}

		code := r.PostFormValue(codeFieldName)
		issuer := r.PostFormValue(issuerFieldName)

		if code == "" && issuer == "" {
			path = append(path, b.Asset{Native: true})
		} else {
			path = append(path, b.Asset{Code: code, Issuer: issuer})
		}
	}

	log.Println(destinationObject.AccountId)
	log.Println(sendAsset)
	log.Println(destinationAsset)
	log.Println(path)

	operationBuilder = b.PathPayment(
		b.Destination{destinationObject.AccountId},
		b.PathSend{
			Asset:     sendAsset,
			MaxAmount: sendMax,
		},
		b.PathDestination{
			Asset:  destinationAsset,
			Amount: destinationAmount,
		},
		b.Path{Assets: path},
	)

	if operationBuilder.(b.PathPaymentBuilder).Err != nil {
		log.WithFields(log.Fields{"err": operationBuilder.(b.PathPaymentBuilder).Err}).Print("Error building operation")
		errorResponse = horizon.ServerError
		return
	}

	return
}
