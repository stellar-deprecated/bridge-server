package handlers

import (
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/stellar/gateway/horizon"
	"github.com/stellar/gateway/protocols"
	"github.com/stellar/gateway/protocols/bridge"
	callback "github.com/stellar/gateway/protocols/compliance"
	"github.com/stellar/gateway/server"
	"github.com/stellar/go/amount"
	b "github.com/stellar/go/build"
	"github.com/stellar/go/keypair"
	"github.com/stellar/go/protocols/compliance"
	"github.com/stellar/go/xdr"
)

// Payment implements /payment endpoint
func (rh *RequestHandler) Payment(w http.ResponseWriter, r *http.Request) {
	request := &bridge.PaymentRequest{}
	request.FromRequest(r)

	err := request.Validate()
	if err != nil {
		errorResponse := err.(*protocols.ErrorResponse)
		log.WithFields(errorResponse.LogData).Error(errorResponse.Error())
		server.Write(w, errorResponse)
		return
	}

	if request.Source == "" {
		request.Source = rh.Config.Accounts.BaseSeed
	}

	sourceKeypair, _ := keypair.Parse(request.Source)

	var submitResponse horizon.SubmitTransactionResponse
	var submitError error

	if request.ExtraMemo != "" && rh.Config.Compliance != "" {
		// Compliance server part
		sendRequest := request.ToComplianceSendRequest()

		resp, err := rh.Client.PostForm(
			rh.Config.Compliance+"/send",
			sendRequest.ToValues(),
		)
		if err != nil {
			log.WithFields(log.Fields{"err": err}).Error("Error sending request to compliance server")
			server.Write(w, protocols.InternalServerError)
			return
		}

		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Error("Error reading compliance server response")
			server.Write(w, protocols.InternalServerError)
			return
		}

		if resp.StatusCode != 200 {
			log.WithFields(log.Fields{
				"status": resp.StatusCode,
				"body":   string(body),
			}).Error("Error response from compliance server")
			server.Write(w, protocols.InternalServerError)
			return
		}

		var callbackSendResponse callback.SendResponse
		err = json.Unmarshal(body, &callbackSendResponse)
		if err != nil {
			log.Error("Error unmarshalling from compliance server")
			server.Write(w, protocols.InternalServerError)
			return
		}

		if callbackSendResponse.AuthResponse.InfoStatus == compliance.AuthStatusPending ||
			callbackSendResponse.AuthResponse.TxStatus == compliance.AuthStatusPending {
			log.WithFields(log.Fields{"response": callbackSendResponse}).Info("Compliance response pending")
			server.Write(w, bridge.NewPaymentPendingError(callbackSendResponse.AuthResponse.Pending))
			return
		}

		if callbackSendResponse.AuthResponse.InfoStatus == compliance.AuthStatusDenied ||
			callbackSendResponse.AuthResponse.TxStatus == compliance.AuthStatusDenied {
			log.WithFields(log.Fields{"response": callbackSendResponse}).Info("Compliance response denied")
			server.Write(w, bridge.PaymentDenied)
			return
		}

		var tx xdr.Transaction
		err = xdr.SafeUnmarshalBase64(callbackSendResponse.TransactionXdr, &tx)
		if err != nil {
			log.Error("Error unmarshalling transaction returned by compliance server")
			server.Write(w, protocols.InternalServerError)
			return
		}

		submitResponse, submitError = rh.TransactionSubmitter.SignAndSubmitRawTransaction(request.Source, &tx)
	} else {
		// Payment without compliance server
		destinationObject, err := rh.FederationResolver.LookupByAddress(request.Destination)
		if err != nil {
			log.WithFields(log.Fields{"destination": request.Destination, "err": err}).Print("Cannot resolve address")
			server.Write(w, bridge.PaymentCannotResolveDestination)
			return
		}

		_, err = keypair.Parse(destinationObject.AccountID)
		if err != nil {
			log.WithFields(log.Fields{"AccountId": destinationObject.AccountID}).Print("Invalid AccountId in destination")
			server.Write(w, protocols.NewInvalidParameterError("destination", request.Destination))
			return
		}

		var payWithMutator *b.PayWithPath

		if request.SendMax != "" {
			// Path payment
			var sendAsset b.Asset
			if request.SendAssetCode == "" && request.SendAssetIssuer == "" {
				sendAsset = b.NativeAsset()
			} else {
				sendAsset = b.CreditAsset(request.SendAssetCode, request.SendAssetIssuer)
			}

			payWith := b.PayWith(sendAsset, request.SendMax)

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
					payWith = payWith.Through(b.NativeAsset())
				} else {
					payWith = payWith.Through(b.CreditAsset(code, issuer))
				}
			}

			payWithMutator = &payWith
		}

		var operationBuilder interface{}

		if request.AssetCode != "" && request.AssetIssuer != "" {
			mutators := []interface{}{
				b.Destination{destinationObject.AccountID},
				b.CreditAmount{request.AssetCode, request.AssetIssuer, request.Amount},
			}

			if payWithMutator != nil {
				mutators = append(mutators, *payWithMutator)
			}

			operationBuilder = b.Payment(mutators...)
		} else {
			mutators := []interface{}{
				b.Destination{destinationObject.AccountID},
				b.NativeAmount{request.Amount},
			}

			if payWithMutator != nil {
				mutators = append(mutators, *payWithMutator)
			}

			// Check if destination account exist
			_, err = rh.Horizon.LoadAccount(destinationObject.AccountID)
			if err != nil {
				log.WithFields(log.Fields{"error": err}).Error("Error loading account")
				operationBuilder = b.CreateAccount(mutators...)
			} else {
				operationBuilder = b.Payment(mutators...)
			}
		}

		memoType := request.MemoType
		memo := request.Memo

		if destinationObject.MemoType != "" {
			if request.MemoType != "" {
				log.Print("Memo given in request but federation returned memo fields.")
				server.Write(w, bridge.PaymentCannotUseMemo)
				return
			}

			memoType = destinationObject.MemoType
			memo = destinationObject.Memo
		}

		var memoMutator interface{}
		switch {
		case memoType == "":
			break
		case memoType == "id":
			id, err := strconv.ParseUint(memo, 10, 64)
			if err != nil {
				log.WithFields(log.Fields{"memo": memo}).Print("Cannot convert memo_id value to uint64")
				server.Write(w, protocols.NewInvalidParameterError("memo", request.Memo))
				return
			}
			memoMutator = b.MemoID{id}
		case memoType == "text":
			memoMutator = &b.MemoText{memo}
		case memoType == "hash":
			memoBytes, err := hex.DecodeString(memo)
			if err != nil || len(memoBytes) != 32 {
				log.WithFields(log.Fields{"memo": memo}).Print("Cannot decode hash memo value")
				server.Write(w, protocols.NewInvalidParameterError("memo", request.Memo))
				return
			}
			var b32 [32]byte
			copy(b32[:], memoBytes[0:32])
			hash := xdr.Hash(b32)
			memoMutator = &b.MemoHash{hash}
		default:
			log.Print("Not supported memo type: ", memoType)
			server.Write(w, protocols.NewInvalidParameterError("memo", request.Memo))
			return
		}

		accountResponse, err := rh.Horizon.LoadAccount(sourceKeypair.Address())
		if err != nil {
			log.WithFields(log.Fields{"error": err}).Error("Cannot load source account")
			server.Write(w, bridge.PaymentSourceNotExist)
			return
		}

		sequenceNumber, err := strconv.ParseUint(accountResponse.SequenceNumber, 10, 64)
		if err != nil {
			log.WithFields(log.Fields{"error": err}).Error("Cannot convert SequenceNumber")
			server.Write(w, protocols.InternalServerError)
			return
		}

		transactionMutators := []b.TransactionMutator{
			b.SourceAccount{request.Source},
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
				server.Write(
					w,
					protocols.NewInvalidParameterError("asset_code", request.AssetCode),
				)
			case strings.Contains(tx.Err.Error(), "cannot parse amount"):
				server.Write(
					w,
					protocols.NewInvalidParameterError("amount", request.Amount),
				)
			default:
				log.WithFields(log.Fields{"err": tx.Err}).Print("Transaction builder error")
				server.Write(w, protocols.InternalServerError)
			}
			return
		}

		txe := tx.Sign(request.Source)
		txeB64, err := txe.Base64()

		if err != nil {
			log.WithFields(log.Fields{"error": err}).Error("Cannot encode transaction envelope")
			server.Write(w, protocols.InternalServerError)
			return
		}

		submitResponse, submitError = rh.Horizon.SubmitTransaction(txeB64)
	}

	if submitError != nil {
		log.WithFields(log.Fields{"error": submitError}).Error("Error submitting transaction")
		server.Write(w, protocols.InternalServerError)
		return
	}

	errorResponse := bridge.ErrorFromHorizonResponse(submitResponse)
	if errorResponse != nil {
		log.WithFields(errorResponse.LogData).Error(errorResponse.Error())
		server.Write(w, errorResponse)
		return
	}

	// Path payment send amount
	if submitResponse.ResultXdr != nil {
		var transactionResult xdr.TransactionResult
		reader := strings.NewReader(*submitResponse.ResultXdr)
		b64r := base64.NewDecoder(base64.StdEncoding, reader)
		_, err := xdr.Unmarshal(b64r, &transactionResult)

		if err == nil && transactionResult.Result.Code == xdr.TransactionResultCodeTxSuccess {
			operationResult := (*transactionResult.Result.Results)[0]
			if operationResult.Tr.PathPaymentResult != nil {
				sendAmount := operationResult.Tr.PathPaymentResult.SendAmount()
				submitResponse.SendAmount = amount.String(sendAmount)
			}
		}
	}

	server.Write(w, &submitResponse)
}
