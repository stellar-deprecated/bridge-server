package horizon

import (
	"bufio"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/Sirupsen/logrus"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/stellar/go-stellar-base/xdr"
)

type PaymentHandler func(PaymentResponse) error

type HorizonInterface interface {
	LoadAccount(accountId string) (response AccountResponse, err error)
	LoadMemo(p *PaymentResponse) (err error)
	StreamPayments(accountId string, cursor *string, onPaymentHandler PaymentHandler) (err error)
	SubmitTransaction(txeBase64 string) (response SubmitTransactionResponse, err error)
}

type Horizon struct {
	ServerUrl string
	log       *logrus.Entry
}

func New(serverUrl string) (horizon Horizon) {
	horizon.ServerUrl = serverUrl
	horizon.log = logrus.WithFields(logrus.Fields{
		"service": "Horizon",
	})
	return
}

func (h *Horizon) LoadAccount(accountId string) (response AccountResponse, err error) {
	h.log.WithFields(logrus.Fields{
		"accountId": accountId,
	}).Info("Loading account")
	resp, err := http.Get(h.ServerUrl + "/accounts/" + accountId)
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
			"accountId": accountId,
		}).Error("Account does not exist")
		err = fmt.Errorf("StatusCode indicates error: %s", body)
		return
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return
	}

	h.log.WithFields(logrus.Fields{
		"accountId": accountId,
	}).Info("Account loaded")
	return
}

func (h *Horizon) LoadMemo(p *PaymentResponse) (err error) {
	res, err := http.Get(p.Links.Transaction.Href)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	return json.NewDecoder(res.Body).Decode(&p.Memo)
}

func (h *Horizon) StreamPayments(accountId string, cursor *string, onPaymentHandler PaymentHandler) (err error) {
	url := h.ServerUrl + "/accounts/" + accountId + "/payments"
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

func (h *Horizon) SubmitTransaction(txeBase64 string) (response SubmitTransactionResponse, err error) {
	v := url.Values{}
	v.Set("tx", txeBase64)
	// TODO add request timeout
	resp, err := http.PostForm(h.ServerUrl+"/transactions", v)
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

	// Decode errors
	if response.Ledger == nil && response.Extras != nil {
		var txResult xdr.TransactionResult
		txResult, err = unmarshalTransactionResult(response.Extras.ResultXdr)

		if err != nil {
			h.log.Info("Cannot decode transaction result")
			return
		}

		transactionResult := txResult.Result.Code
		operationsResults := *txResult.Result.Results
		var transactionErrorCode string
		var operationErrorCode string

		if transactionResult != xdr.TransactionResultCodeTxSuccess {
			switch transactionResult {
			case xdr.TransactionResultCodeTxFailed:
				transactionErrorCode = "transaction_failed"
			case xdr.TransactionResultCodeTxBadSeq:
				transactionErrorCode = "transaction_bad_seq"
			default:
				transactionErrorCode = "unknown"
			}
		}

		if operationsResults != nil {
			if operationsResults[0].Tr.AllowTrustResult != nil {
				switch operationsResults[0].Tr.AllowTrustResult.Code {
				case xdr.AllowTrustResultCodeAllowTrustMalformed:
					operationErrorCode = "allow_trust_malformed"
				case xdr.AllowTrustResultCodeAllowTrustNoTrustLine:
					operationErrorCode = "allow_trust_not_trustline"
				case xdr.AllowTrustResultCodeAllowTrustTrustNotRequired:
					operationErrorCode = "allow_trust_trust_not_required"
				case xdr.AllowTrustResultCodeAllowTrustCantRevoke:
					operationErrorCode = "allow_trust_trust_cant_revoke"
				default:
					operationErrorCode = "unknown"
				}
			} else if operationsResults[0].Tr.PaymentResult != nil {
				switch operationsResults[0].Tr.PaymentResult.Code {
				case xdr.PaymentResultCodePaymentMalformed:
					operationErrorCode = "payment_malformed"
				case xdr.PaymentResultCodePaymentUnderfunded:
					operationErrorCode = "payment_underfunded"
				case xdr.PaymentResultCodePaymentSrcNoTrust:
					operationErrorCode = "payment_src_no_trust"
				case xdr.PaymentResultCodePaymentSrcNotAuthorized:
					operationErrorCode = "payment_src_not_authorized"
				case xdr.PaymentResultCodePaymentNoDestination:
					operationErrorCode = "payment_no_destination"
				case xdr.PaymentResultCodePaymentNoTrust:
					operationErrorCode = "payment_no_trust"
				case xdr.PaymentResultCodePaymentNotAuthorized:
					operationErrorCode = "payment_not_authorized"
				case xdr.PaymentResultCodePaymentLineFull:
					operationErrorCode = "payment_line_full"
				case xdr.PaymentResultCodePaymentNoIssuer:
					operationErrorCode = "payment_no_issuer"
				default:
					operationErrorCode = "unknown"
				}
			}
		}

		errors := &SubmitTransactionResponseError{
			TransactionErrorCode: transactionErrorCode,
			OperationErrorCode:   operationErrorCode,
		}
		response.Errors = errors
	}

	return
}

func unmarshalTransactionResult(transactionResult string) (txResult xdr.TransactionResult, err error) {
	reader := strings.NewReader(transactionResult)
	b64r := base64.NewDecoder(base64.StdEncoding, reader)
	_, err = xdr.Unmarshal(b64r, &txResult)
	return
}
