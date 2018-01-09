package horizon

import (
	"bufio"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/stellar/go/xdr"
)

// PaymentHandler is a function that is called when a new payment is received
type PaymentHandler func(PaymentResponse) error

// HorizonInterface allows mocking Horizon struct object
type HorizonInterface interface {
	LoadAccount(accountID string) (response AccountResponse, err error)
	LoadMemo(p *PaymentResponse) (err error)
	LoadOperation(operationID string) (response PaymentResponse, err error)
	StreamPayments(accountID string, cursor *string, onPaymentHandler PaymentHandler) (err error)
	SubmitTransaction(txeBase64 string) (response SubmitTransactionResponse, err error)
}

// Horizon implements methods to get (or submit) data from Horizon server
type Horizon struct {
	ServerURL string
	log       *logrus.Entry
}

const submitTimeout = 30 * time.Second

// New creates a new Horizon instance
func New(serverURL string) (horizon Horizon) {
	horizon.ServerURL = serverURL
	horizon.log = logrus.WithFields(logrus.Fields{
		"service": "Horizon",
	})
	return
}

// LoadAccount loads a single account from Horizon server
func (h *Horizon) LoadAccount(accountID string) (response AccountResponse, err error) {
	h.log.WithFields(logrus.Fields{
		"accountID": accountID,
	}).Info("Loading account")
	resp, err := http.Get(h.ServerURL + "/accounts/" + accountID)
	if err != nil {
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	if resp.StatusCode != 200 {
		h.log.WithFields(logrus.Fields{
			"accountID": accountID,
		}).Info("Account does not exist")
		err = fmt.Errorf("StatusCode indicates error: %s", body)
		return
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return
	}

	h.log.WithFields(logrus.Fields{
		"accountID": accountID,
	}).Info("Account loaded")
	return
}

// LoadOperation loads a single operation from Horizon server
func (h *Horizon) LoadOperation(operationID string) (response PaymentResponse, err error) {
	h.log.WithFields(logrus.Fields{
		"operationID": operationID,
	}).Info("Loading operation")
	resp, err := http.Get(h.ServerURL + "/operations/" + operationID)
	if err != nil {
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	if resp.StatusCode != 200 {
		h.log.WithFields(logrus.Fields{
			"operationID": operationID,
		}).Error("Operation does not exist")
		err = fmt.Errorf("StatusCode indicates error: %s", body)
		return
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return
	}

	h.log.WithFields(logrus.Fields{
		"operationID": operationID,
	}).Info("Operation loaded")
	return
}

// LoadMemo loads memo for a transaction in PaymentResponse
func (h *Horizon) LoadMemo(p *PaymentResponse) (err error) {
	res, err := http.Get(p.Links.Transaction.Href)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	return json.NewDecoder(res.Body).Decode(&p.Memo)
}

// StreamPayments streams incoming payments
func (h *Horizon) StreamPayments(accountID string, cursor *string, onPaymentHandler PaymentHandler) (err error) {
	url := h.ServerURL + "/accounts/" + accountID + "/payments"
	if cursor != nil {
		url += "?cursor=" + *cursor
	}

	req, _ := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}
	req.Header.Set("Accept", "text/event-stream")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	scanner := bufio.NewScanner(resp.Body)
	scanner.Split(splitSSE)

	for scanner.Scan() {
		if len(scanner.Bytes()) == 0 {
			continue
		}

		ev, err := parseEvent(scanner.Bytes())
		if err != nil {
			return err
		}

		if ev.Event != "message" {
			continue
		}

		var payment PaymentResponse
		data := ev.Data.(string)
		err = json.Unmarshal([]byte(data), &payment)
		if err != nil {
			return err
		}

		for {
			err = onPaymentHandler(payment)
			if err != nil {
				h.log.Error("Error from onPaymentHandler: ", err)
				h.log.Info("Sleeping...")
				time.Sleep(10 * time.Second)
			} else {
				break
			}
		}
	}

	err = scanner.Err()
	if err == io.ErrUnexpectedEOF {
		h.log.Info("Streaming connection closed.")
		return nil
	}
	if err != nil {
		return err
	}

	return nil
}

// SubmitTransaction submits a transaction to Stellar network via Horizon server
func (h *Horizon) SubmitTransaction(txeBase64 string) (response SubmitTransactionResponse, err error) {
	v := url.Values{}
	v.Set("tx", txeBase64)

	client := http.Client{
		Timeout: submitTimeout,
	}
	resp, err := client.PostForm(h.ServerURL+"/transactions", v)
	if err != nil {
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		h.log.WithFields(logrus.Fields{
			"body": string(body),
		}).Info("Cannot unmarshal horizon response", string(body))
		return
	}

	if response.Ledger != nil {
		h.log.WithFields(logrus.Fields{
			"ledger": response.Ledger,
		}).Info("Success response from horizon")
	} else {
		h.log.WithFields(logrus.Fields{
			"envelope": response.Extras.EnvelopeXdr,
			"result":   response.Extras.ResultXdr,
		}).Info("Error response from horizon")
	}

	return
}

func unmarshalTransactionResult(transactionResult string) (txResult xdr.TransactionResult, err error) {
	reader := strings.NewReader(transactionResult)
	b64r := base64.NewDecoder(base64.StdEncoding, reader)
	_, err = xdr.Unmarshal(b64r, &txResult)
	return
}
