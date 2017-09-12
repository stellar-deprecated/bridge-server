package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"

	log "github.com/Sirupsen/logrus"
	"github.com/stellar/gateway/db/entities"
	"github.com/stellar/gateway/horizon"
	"github.com/stellar/gateway/protocols"
	callback "github.com/stellar/gateway/protocols/compliance"
	"github.com/stellar/gateway/server"
	"github.com/stellar/go/protocols/compliance"
	"github.com/stellar/go/support/errors"
	"github.com/zenazn/goji/web"
)

// AdminReceivedPayment implements /admin/received-payments/{id} endpoint
func (rh *RequestHandler) AdminReceivedPayment(c web.C, w http.ResponseWriter, r *http.Request) {
	object, err := rh.Driver.GetOne(&entities.ReceivedPayment{}, "id = ?", c.URLParams["id"])
	if err != nil {
		log.WithFields(log.Fields{"err": err}).Error("Error getting ReceivedPayments")
		server.Write(w, protocols.InternalServerError)
		return
	}

	if object == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	payment := object.(*entities.ReceivedPayment)

	paymentResponse, err := rh.Horizon.LoadOperation(payment.OperationID)
	if err != nil {
		log.WithFields(log.Fields{"err": err}).Error("Error getting operation from Horizon")
		server.Write(w, protocols.InternalServerError)
		return
	}

	err = rh.Horizon.LoadMemo(&paymentResponse)
	if err != nil {
		log.WithFields(log.Fields{"err": err}).Error("Error loading memo")
		server.Write(w, protocols.InternalServerError)
		return
	}

	var authData *compliance.AuthData
	if paymentResponse.Memo.Type == "hash" && rh.Config.Compliance != "" {
		authData, err = rh.getComplianceData(paymentResponse.Memo.Value)
		if err != nil {
			log.WithFields(log.Fields{"err": err}).Error("Error loading compliance data")
			server.Write(w, protocols.InternalServerError)
			return
		}
	}

	response := struct {
		Payment   *entities.ReceivedPayment `json:"payment"`
		Operation horizon.PaymentResponse   `json:"operation"`
		AuthData  *compliance.AuthData      `json:"auth_data"`
	}{payment, paymentResponse, authData}

	encoder := json.NewEncoder(w)
	err = encoder.Encode(response)
	if err != nil {
		log.WithFields(log.Fields{"err": err, "payments": payment}).Error("Error encoding ReceivedPayment")
		server.Write(w, protocols.InternalServerError)
		return
	}
}

func (rh *RequestHandler) getComplianceData(memo string) (*compliance.AuthData, error) {
	complianceRequestURL := rh.Config.Compliance + "/receive"
	complianceRequestBody := url.Values{"memo": {string(memo)}}

	log.WithFields(log.Fields{"url": complianceRequestURL, "body": complianceRequestBody}).Info("Sending request to compliance server")
	resp, err := rh.Client.PostForm(complianceRequestURL, complianceRequestBody)
	if err != nil {
		return nil, errors.Wrap(err, "Error sending request to compliance server")
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "Error reading compliance server response")
	}

	if resp.StatusCode == http.StatusNotFound {
		return nil, nil
	}

	if resp.StatusCode != http.StatusOK {
		log.WithFields(log.Fields{"status": resp.StatusCode, "body": string(body)}).Error("Error response from compliance server")
		return nil, errors.New("Error response from compliance server")
	}

	var receiveResponse callback.ReceiveResponse
	err = json.Unmarshal([]byte(body), &receiveResponse)
	if err != nil {
		return nil, errors.Wrap(err, "Cannot unmarshal receiveResponse")
	}

	var authData compliance.AuthData
	err = json.Unmarshal([]byte(receiveResponse.Data), &authData)
	if err != nil {
		return nil, errors.Wrap(err, "Cannot unmarshal authData")
	}

	return &authData, nil
}

// AdminReceivedPayments implements /admin/received-payments endpoint
func (rh *RequestHandler) AdminReceivedPayments(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	limit := 10

	payments, err := rh.Repository.GetReceivedPayments(page, limit)
	if err != nil {
		log.WithFields(log.Fields{"err": err}).Error("Error loading ReceivedPayments")
		server.Write(w, protocols.InternalServerError)
		return
	}

	encoder := json.NewEncoder(w)
	err = encoder.Encode(payments)
	if err != nil {
		log.WithFields(log.Fields{"err": err, "payments": payments}).Error("Error encoding ReceivedPayments")
		server.Write(w, protocols.InternalServerError)
		return
	}
}

// AdminReceivedPayments implements /admin/sent-transactions endpoint
func (rh *RequestHandler) AdminSentTransactions(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	limit := 10

	transactions, err := rh.Repository.GetSentTransactions(page, limit)
	if err != nil {
		log.WithFields(log.Fields{"err": err}).Error("Error loading SentTransactions")
		server.Write(w, protocols.InternalServerError)
		return
	}

	encoder := json.NewEncoder(w)
	err = encoder.Encode(transactions)
	if err != nil {
		log.WithFields(log.Fields{"err": err, "transactions": transactions}).Error("Error encoding SentTransactions")
		server.Write(w, protocols.InternalServerError)
		return
	}
}
