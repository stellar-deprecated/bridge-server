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
	"strings"

	ch "github.com/stellar/gateway/compliance/handlers"
	"github.com/stellar/gateway/horizon"
	b "github.com/stellar/go-stellar-base/build"
	"github.com/stellar/go-stellar-base/keypair"
	"github.com/stellar/go-stellar-base/xdr"
)

func (rh *RequestHandler) Payment(w http.ResponseWriter, r *http.Request) {

	source := r.PostFormValue("source")
	sourceKeypair, err := keypair.Parse(source)
	if err != nil {
		log.WithFields(log.Fields{"source": source}).Print("Invalid source parameter")
		writeError(w, horizon.PaymentInvalidSource)
		return
	}

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
		operationBuilder, errorResponse = rh.createPaymentOperation(r, destinationObject)
	case "path_payment":
		log.Println("path_payment")
		operationBuilder, errorResponse = rh.createPathPaymentOperation(r, destinationObject)
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

	accountResponse, err := rh.Horizon.LoadAccount(sourceKeypair.Address())
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Error("Cannot load source account")
		writeError(w, horizon.PaymentSourceNotExist)
		return
	}

	sequenceNumber, err := strconv.ParseUint(accountResponse.SequenceNumber, 10, 64)
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Error("Cannot convert SequenceNumber")
		writeError(w, horizon.ServerError)
		return
	}

	transactionMutators := []b.TransactionMutator{
		b.SourceAccount{source},
		b.Sequence{sequenceNumber + 1},
		b.Network{rh.Config.NetworkPassphrase},
		operationBuilder.(b.TransactionMutator),
	}

	if memoMutator != nil {
		transactionMutators = append(transactionMutators, memoMutator.(b.TransactionMutator))
	}

	tx := b.Transaction(transactionMutators...)

	if tx.Err != nil {
		log.WithFields(log.Fields{"err": tx.Err}).Print("Transaction builder error")
		// TODO when build.OperationBuilder interface is ready check for
		// create_account and payment errors separately
		switch {
		case tx.Err.Error() == "Asset code length is invalid":
			writeError(w, horizon.PaymentMalformedAssetCode)
		case strings.Contains(tx.Err.Error(), "cannot parse amount"):
			writeError(w, horizon.PaymentInvalidAmount)
		default:
			log.WithFields(log.Fields{"err": tx.Err}).Print("Transaction builder error")
			writeError(w, horizon.ServerError)
		}
		return
	}

	txe := tx.Sign(source)
	txeB64, err := txe.Base64()

	if err != nil {
		log.WithFields(log.Fields{"error": err}).Error("Cannot encode transaction envelope")
		writeError(w, horizon.ServerError)
		return
	}

	submitResponse, err := rh.Horizon.SubmitTransaction(txeB64)
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Error("Error submitting transaction")
		writeError(w, horizon.ServerError)
		return
	}

	write(w, submitResponse)
}

func (rh *RequestHandler) createPaymentOperation(r *http.Request, destinationObject StellarDestination) (operationBuilder interface{}, errorResponse *horizon.SubmitTransactionResponseError) {
	amount := r.PostFormValue("amount")
	assetCode := r.PostFormValue("asset_code")
	assetIssuer := r.PostFormValue("asset_issuer")

	if assetCode != "" && assetIssuer != "" {
		issuerKeypair, err := keypair.Parse(assetIssuer)
		if err != nil {
			log.WithFields(log.Fields{"asset_issuer": assetIssuer}).Print("Invalid asset_issuer parameter")
			errorResponse = horizon.PaymentInvalidIssuer
			return
		}

		operationBuilder = b.Payment(
			b.Destination{destinationObject.AccountId},
			b.CreditAmount{assetCode, issuerKeypair.Address(), amount},
		)

		if operationBuilder.(b.PaymentBuilder).Err != nil {
			log.WithFields(log.Fields{"err": operationBuilder.(b.PaymentBuilder).Err}).Print("Error building operation")
			errorResponse = horizon.ServerError
			return
		}
	} else if assetCode == "" && assetIssuer == "" {
		mutators := []interface{}{
			b.Destination{destinationObject.AccountId},
			b.NativeAmount{amount},
		}

		// Check if destination account exist
		_, err := rh.Horizon.LoadAccount(destinationObject.AccountId)
		if err != nil {
			log.WithFields(log.Fields{"error": err}).Error("Error loading account")
			operationBuilder = b.CreateAccount(mutators...)
			if operationBuilder.(b.CreateAccountBuilder).Err != nil {
				log.WithFields(log.Fields{"err": operationBuilder.(b.CreateAccountBuilder).Err}).Print("Error building operation")
				errorResponse = horizon.ServerError
				return
			}
		} else {
			operationBuilder = b.Payment(mutators...)
			if operationBuilder.(b.PaymentBuilder).Err != nil {
				log.WithFields(log.Fields{"err": operationBuilder.(b.PaymentBuilder).Err}).Print("Error building operation")
				errorResponse = horizon.ServerError
				return
			}
		}
	} else {
		log.Print("Missing asset param.")
		errorResponse = horizon.PaymentMissingParamAsset
		return
	}
	return
}

func (rh *RequestHandler) createPathPaymentOperation(r *http.Request, destinationObject StellarDestination) (operationBuilder interface{}, errorResponse *horizon.SubmitTransactionResponseError) {
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
