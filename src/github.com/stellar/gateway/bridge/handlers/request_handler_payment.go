package handlers

import (
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/stellar/gateway/horizon"
	"github.com/stellar/gateway/protocols"
	"github.com/stellar/gateway/protocols/bridge"
	callback "github.com/stellar/gateway/protocols/compliance"
	"github.com/stellar/gateway/server"
	"github.com/stellar/go/address"
	"github.com/stellar/go/amount"
	b "github.com/stellar/go/build"
	"github.com/stellar/go/protocols/compliance"
	"github.com/stellar/go/protocols/federation"
	"github.com/stellar/go/xdr"
)

// Payment implements /payment endpoint
func (rh *RequestHandler) Payment(w http.ResponseWriter, r *http.Request) {
	request := &bridge.PaymentRequest{}
	err := request.FromRequest(r)
	if err != nil {
		log.Error(err.Error())
		server.Write(w, protocols.InvalidParameterError)
		return
	}

	err = request.Validate()
	if err != nil {
		errorResponse := err.(*protocols.ErrorResponse)
		log.WithFields(errorResponse.LogData).Error(errorResponse.Error())
		server.Write(w, errorResponse)
		return
	}

	if request.Source == "" {
		request.Source = rh.Config.Accounts.BaseSeed
	}

	// Will use compliance if compliance server is connected and:
	// * User passed extra memo OR
	// * User explicitly wants to use compliance protocol
	if rh.Config.Compliance != "" &&
		(request.ExtraMemo != "" || (request.ExtraMemo == "" && request.UseCompliance)) {
		rh.complianceProtocolPayment(w, request)
	} else {
		rh.standardPayment(w, request)
	}
}

func (rh *RequestHandler) complianceProtocolPayment(w http.ResponseWriter, request *bridge.PaymentRequest) {
	var paymentID *string

	if request.ID != "" {
		paymentID = &request.ID
	}

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

	submitResponse, err := rh.TransactionSubmitter.SignAndSubmitRawTransaction(paymentID, request.Source, &tx)
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Error("Error submitting transaction")
		server.Write(w, protocols.InternalServerError)
		return
	}

	rh.handleSubmitterResponse(w, submitResponse)
}

func (rh *RequestHandler) standardPayment(w http.ResponseWriter, request *bridge.PaymentRequest) {
	var paymentID *string

	if request.ID != "" {
		sentTransaction, err := rh.Repository.GetSentTransactionByPaymentID(request.ID)
		if err != nil {
			log.WithFields(log.Fields{"err": err}).Error("Error getting sent transaction")
			server.Write(w, protocols.InternalServerError)
			return
		}

		if sentTransaction == nil {
			paymentID = &request.ID
		} else {
			log.WithFields(log.Fields{"paymentID": request.ID, "tx": sentTransaction.EnvelopeXdr}).Info("Transaction with given ID already exists, resubmitting...")
			submitResponse, err := rh.Horizon.SubmitTransaction(sentTransaction.EnvelopeXdr)
			if err != nil {
				log.WithFields(log.Fields{"error": err}).Error("Error submitting transaction")
				server.Write(w, protocols.InternalServerError)
				return
			}

			rh.handleSubmitterResponse(w, submitResponse)
			return
		}
	}

	destinationObject := &federation.NameResponse{}
	var err error

	if request.ForwardDestination == nil {
		_, _, err = address.Split(request.Destination)
		if err != nil {
			destinationObject.AccountID = request.Destination
		} else {
			destinationObject, err = rh.FederationResolver.LookupByAddress(request.Destination)
			if err != nil {
				log.WithFields(log.Fields{"destination": request.Destination, "err": err}).Print("Cannot resolve address")
				server.Write(w, bridge.PaymentCannotResolveDestination)
				return
			}
		}
	} else {
		destinationObject, err = rh.FederationResolver.ForwardRequest(request.ForwardDestination.Domain, request.ForwardDestination.Fields)
		if err != nil {
			log.WithFields(log.Fields{"destination": request.Destination, "err": err}).Print("Cannot resolve address")
			server.Write(w, bridge.PaymentCannotResolveDestination)
			return
		}
	}

	if !protocols.IsValidAccountID(destinationObject.AccountID) {
		log.WithFields(log.Fields{"AccountId": destinationObject.AccountID}).Print("Invalid AccountId in destination")
		server.Write(w, protocols.NewInvalidParameterError("destination", request.Destination, "Destination public key must start with `G`."))
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

		for _, asset := range request.Path {
			payWith = payWith.Through(asset.ToBaseAsset())
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
		memo = destinationObject.Memo.Value
	}

	var memoMutator interface{}
	switch {
	case memoType == "":
		break
	case memoType == "id":
		id, err := strconv.ParseUint(memo, 10, 64)
		if err != nil {
			log.WithFields(log.Fields{"memo": memo}).Print("Cannot convert memo_id value to uint64")
			server.Write(w, protocols.NewInvalidParameterError("memo", request.Memo, "Memo.id must be a number"))
			return
		}
		memoMutator = b.MemoID{id}
	case memoType == "text":
		memoMutator = b.MemoText{memo}
	case memoType == "hash":
		memoBytes, err := hex.DecodeString(memo)
		if err != nil || len(memoBytes) != 32 {
			log.WithFields(log.Fields{"memo": memo}).Print("Cannot decode hash memo value")
			server.Write(w, protocols.NewInvalidParameterError("memo", request.Memo, "Memo.hash must be 32 bytes and hex encoded."))
			return
		}
		var b32 [32]byte
		copy(b32[:], memoBytes[0:32])
		hash := xdr.Hash(b32)
		memoMutator = b.MemoHash{hash}
	default:
		log.Print("Not supported memo type: ", memoType)
		server.Write(w, protocols.NewInvalidParameterError("memo", request.Memo, "Memo type not supported"))
		return
	}

	submitResponse, err := rh.TransactionSubmitter.SubmitTransaction(paymentID, request.Source, operationBuilder, memoMutator)
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Error("Error submitting transaction")
		server.Write(w, protocols.InternalServerError)
		return
	}

	rh.handleSubmitterResponse(w, submitResponse)
}

func (rh *RequestHandler) handleSubmitterResponse(w http.ResponseWriter, response horizon.SubmitTransactionResponse) {
	errorResponse := bridge.ErrorFromHorizonResponse(response)
	if errorResponse != nil {
		log.WithFields(errorResponse.LogData).Error(errorResponse.Error())
		server.Write(w, errorResponse)
		return
	}

	// Path payment send amount
	if response.ResultXdr != nil {
		var transactionResult xdr.TransactionResult
		reader := strings.NewReader(*response.ResultXdr)
		b64r := base64.NewDecoder(base64.StdEncoding, reader)
		_, err := xdr.Unmarshal(b64r, &transactionResult)

		if err == nil && transactionResult.Result.Code == xdr.TransactionResultCodeTxSuccess {
			operationResult := (*transactionResult.Result.Results)[0]
			if operationResult.Tr.PathPaymentResult != nil {
				sendAmount := operationResult.Tr.PathPaymentResult.SendAmount()
				response.SendAmount = amount.String(sendAmount)
			}
		}
	}

	server.Write(w, &response)
}
