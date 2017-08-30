package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"

	log "github.com/Sirupsen/logrus"

	"github.com/stellar/gateway/protocols"
	"github.com/stellar/gateway/server"
	"github.com/zenazn/goji/web"
)

// TransactionStatusResponse represents a response from the tx_status endpoint
type TransactionStatusResponse struct {
	protocols.SuccessResponse
	Status   string `json:"status"`
	RecvCode string `json:"recv_code,omitempty"`
	RefundTx string `json:"refund_tx,omitempty"`
	Msg      string `json:"msg,omitempty"`
}

// HandlerTxStatus implements /tx_status endpoint
func (rh *RequestHandler) HandlerTxStatus(c web.C, w http.ResponseWriter, r *http.Request) {

	txid := r.URL.Query().Get("id")
	if txid == "" {
		log.Info("unable to get query parameter")
		server.Write(w, protocols.InvalidParameterError)
		return
	}
	response := TransactionStatusResponse{}

	if rh.Config.Callbacks.TxStatus == "" {
		response.Status = "unknown"
	} else {
		resp, err := rh.Client.PostForm(
			rh.Config.Callbacks.TxStatus,
			url.Values{"id": {txid}},
		)
		if err != nil {
			log.WithFields(log.Fields{
				"tx_status": rh.Config.Callbacks.TxStatus,
				"err":       err,
			}).Error("Error sending request to tx_status server")
			server.Write(w, protocols.InternalServerError)
			return
		}

		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Error("Error reading tx_status server response")
			server.Write(w, protocols.InternalServerError)
			return
		}

		switch resp.StatusCode {
		case http.StatusOK:
			err := json.Unmarshal(body, &response)
			if err != nil {
				log.WithFields(log.Fields{
					"tx_status": rh.Config.Callbacks.TxStatus,
					"body":      string(body),
				}).Error("Unable to decode tx_status response")
				server.Write(w, protocols.InternalServerError)
				return
			}
			if response.Status == "" {
				response.Status = "unknown"
			}

		default:
			response.Status = "unknown"
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
