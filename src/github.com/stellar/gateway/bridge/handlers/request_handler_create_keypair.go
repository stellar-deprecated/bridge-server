package handlers

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"net/http"

	"github.com/stellar/gateway/protocols"
	"github.com/stellar/gateway/server"
	"github.com/stellar/go/keypair"
)

// KeyPair struct contains key pair public and private key
type KeyPair struct {
	PublicKey  string `json:"public_key"`
	PrivateKey string `json:"private_key"`
}

// CreateKeypair implements /create-keypair endpoint
func (rh *RequestHandler) CreateKeypair(w http.ResponseWriter, r *http.Request) {
	kp, err := keypair.Random()
	if err != nil {
		log.WithFields(log.Fields{"err": err}).Error("Error generating random keypair")
		server.Write(w, protocols.InternalServerError)
	}

	response, err := json.Marshal(KeyPair{kp.Address(), kp.Seed()})
	if err != nil {
		log.WithFields(log.Fields{"err": err}).Error("Error marshalling random keypair")
		server.Write(w, protocols.InternalServerError)
	}

	w.Write(response)
}
