package horizon

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/Sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/stellar/go-stellar-base/xdr"
)

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
			if (operationsResults[0].Tr.AllowTrustResult != nil) {
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
			} else if (operationsResults[0].Tr.PaymentResult != nil) {
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
