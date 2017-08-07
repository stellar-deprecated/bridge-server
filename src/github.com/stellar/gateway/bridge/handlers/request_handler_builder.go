package handlers

import (
	"encoding/json"
	log "github.com/Sirupsen/logrus"
	"net/http"
	"strconv"

	"github.com/stellar/gateway/protocols"
	"github.com/stellar/gateway/protocols/bridge"
	"github.com/stellar/gateway/server"
	b "github.com/stellar/go/build"
)

// Builder implements /builder endpoint
func (rh *RequestHandler) Builder(w http.ResponseWriter, r *http.Request) {
	var request bridge.BuilderRequest
	var sequenceNumber uint64

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&request)
	if err != nil {
		log.WithFields(log.Fields{"err": err}).Error("Error decoding request")
		server.Write(w, protocols.NewInvalidParameterError("", "", "Request body is not a valid JSON"))
		return
	}

	err = request.Process()
	if err != nil {
		errorResponse := err.(*protocols.ErrorResponse)
		log.WithFields(errorResponse.LogData).Error(errorResponse.Error())
		server.Write(w, errorResponse)
		return
	}

	err = request.Validate()
	if err != nil {
		errorResponse := err.(*protocols.ErrorResponse)
		log.WithFields(errorResponse.LogData).Error(errorResponse.Error())
		server.Write(w, errorResponse)
		return
	}

	if request.SequenceNumber == "" {
		accountResponse, err := rh.Horizon.LoadAccount(request.Source)
		if err != nil {
			log.WithFields(log.Fields{"err": err}).Error("Error when loading account")
			server.Write(w, protocols.InternalServerError)
			return
		}
		sequenceNumber, err = strconv.ParseUint(accountResponse.SequenceNumber, 10, 64)
	}else{
		sequenceNumber, err = strconv.ParseUint(request.SequenceNumber, 10, 64)
	}

	if err != nil {
		errorResponse := protocols.NewInvalidParameterError("sequence_number", request.SequenceNumber, "Sequence number must be a number")
		log.WithFields(errorResponse.LogData).Error(errorResponse.Error())
		server.Write(w, errorResponse)
		return
	}

	mutators := []b.TransactionMutator{
		b.SourceAccount{request.Source},
		b.Sequence{sequenceNumber},
		b.Network{rh.Config.NetworkPassphrase},
	}

	for _, operation := range request.Operations {
		mutators = append(mutators, operation.Body.ToTransactionMutator())
	}

	tx := b.Transaction(mutators...)

	if tx.Err != nil {
		log.WithFields(log.Fields{"err": err, "request": request}).Error("TransactionBuilder returned error")
		server.Write(w, protocols.InternalServerError)
		return
	}

	txe := tx.Sign(request.Signers...)
	txeB64, err := txe.Base64()
	if err != nil {
		log.WithFields(log.Fields{"err": err, "request": request}).Error("Error encoding transaction envelope")
		server.Write(w, protocols.InternalServerError)
		return
	}

	server.Write(w, &bridge.BuilderResponse{TransactionEnvelope: txeB64})
}
