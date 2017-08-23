package listener

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"encoding/base64"

	"github.com/Sirupsen/logrus"
	"github.com/stellar/gateway/bridge/config"
	"github.com/stellar/gateway/db"
	"github.com/stellar/gateway/db/entities"
	"github.com/stellar/gateway/horizon"
	callback "github.com/stellar/gateway/protocols/compliance"
	"github.com/stellar/go/protocols/compliance"
	"github.com/stellar/go/strkey"
	"github.com/stellar/go/support/errors"
)

// PaymentListener is listening for a new payments received by ReceivingAccount
type PaymentListener struct {
	client        HTTP
	config        *config.Config
	entityManager db.EntityManagerInterface
	horizon       horizon.HorizonInterface
	log           *logrus.Entry
	repository    db.RepositoryInterface
	now           func() time.Time
}

// HTTP represents an http client that a payment listener can use to make HTTP
// requests.
type HTTP interface {
	Do(req *http.Request) (resp *http.Response, err error)
}

const callbackTimeout = 60 * time.Second

// NewPaymentListener creates a new PaymentListener
func NewPaymentListener(
	config *config.Config,
	entityManager db.EntityManagerInterface,
	horizon horizon.HorizonInterface,
	repository db.RepositoryInterface,
	now func() time.Time,
) (pl PaymentListener, err error) {
	pl.client = &http.Client{
		Timeout: callbackTimeout,
	}
	pl.config = config
	pl.entityManager = entityManager
	pl.horizon = horizon
	pl.repository = repository
	pl.now = now
	pl.log = logrus.WithFields(logrus.Fields{
		"service": "PaymentListener",
	})
	return
}

// Listen starts listening for new payments
func (pl *PaymentListener) Listen() (err error) {
	accountID := pl.config.Accounts.ReceivingAccountID

	_, err = pl.horizon.LoadAccount(accountID)
	if err != nil {
		return
	}

	go func() {
		for {
			cursor, err := pl.repository.GetLastCursorValue()
			if err != nil {
				pl.log.WithFields(logrus.Fields{"error": err}).Error("Could not load last cursor from the DB")
				return
			}

			var cursorValue string
			if cursor != nil {
				cursorValue = *cursor
			} else {
				// If no last cursor saved set it to: `now`
				cursorValue = "now"
				cursor = &cursorValue
			}

			pl.log.WithFields(logrus.Fields{
				"accountId": accountID,
				"cursor":    cursorValue,
			}).Info("Started listening for new payments")

			err = pl.horizon.StreamPayments(
				accountID,
				cursor,
				pl.onPayment,
			)
			if err != nil {
				pl.log.Error("Error while streaming: ", err)
				pl.log.Info("Sleeping...")
				time.Sleep(10 * time.Second)
			}
		}
	}()

	return
}

func (pl *PaymentListener) ReprocessPayment(payment horizon.PaymentResponse, force bool) error {
	pl.log.WithFields(logrus.Fields{"id": payment.ID}).Info("Reprocessing a payment")

	id, err := strconv.ParseInt(payment.ID, 10, 64)
	if err != nil {
		pl.log.WithFields(logrus.Fields{"err": err}).Error("Error converting ID to int64")
		return err
	}

	existingPayment, err := pl.repository.GetReceivedPaymentByOperationID(id)
	if err != nil {
		pl.log.WithFields(logrus.Fields{"err": err}).Error("Error checking if receive payment exists")
		return err
	}

	if existingPayment == nil {
		pl.log.WithFields(logrus.Fields{"id": payment.ID}).Info("Payment has not been processed yet")
		return errors.New("Payment has not been processed yet")
	}

	if existingPayment.Status == "Success" && !force {
		pl.log.WithFields(logrus.Fields{"id": payment.ID}).Info("Trying to reprocess successful transaction without force")
		return errors.New("Trying to reprocess successful transaction without force")
	}

	existingPayment.Status = "Reprocessing..."
	existingPayment.ProcessedAt = pl.now()

	err = pl.entityManager.Persist(existingPayment)
	if err != nil {
		return err
	}

	err = pl.process(payment)

	if err != nil {
		pl.log.WithFields(logrus.Fields{"err": err}).Error("Payment reprocessed with errors")
		existingPayment.Status = err.Error()
	} else {
		pl.log.Info("Payment successfully reprocessed")
		existingPayment.Status = "Success"
	}

	return pl.entityManager.Persist(existingPayment)
}

func (pl *PaymentListener) onPayment(payment horizon.PaymentResponse) (err error) {
	pl.log.WithFields(logrus.Fields{"id": payment.ID}).Info("New received payment")

	id, err := strconv.ParseInt(payment.ID, 10, 64)
	if err != nil {
		pl.log.WithFields(logrus.Fields{"err": err}).Error("Error converting ID to int64")
		return err
	}

	existingPayment, err := pl.repository.GetReceivedPaymentByOperationID(id)
	if err != nil {
		pl.log.WithFields(logrus.Fields{"err": err}).Error("Error checking if receive payment exists")
		return err
	}

	if existingPayment != nil {
		pl.log.WithFields(logrus.Fields{"id": payment.ID}).Info("Payment already exists")
		return
	}

	dbPayment := &entities.ReceivedPayment{
		OperationID: payment.ID,
		ProcessedAt: pl.now(),
		PagingToken: payment.PagingToken,
		Status:      "Processing...",
	}

	err = pl.entityManager.Persist(dbPayment)
	if err != nil {
		return
	}

	err = pl.process(payment)

	if err != nil {
		pl.log.WithFields(logrus.Fields{"err": err}).Error("Payment processed with errors")
		dbPayment.Status = err.Error()
	} else {
		pl.log.Info("Payment successfully processed")
		dbPayment.Status = "Success"
	}

	return pl.entityManager.Persist(dbPayment)
}

func (pl *PaymentListener) process(payment horizon.PaymentResponse) error {
	if payment.Type != "payment" && payment.Type != "path_payment" {
		return errors.New("Not a payment operation")
	}

	if payment.To != pl.config.Accounts.ReceivingAccountID {
		return errors.New("Operation sent not received")
	}

	if !pl.isAssetAllowed(payment.AssetType, payment.AssetCode, payment.AssetIssuer) {
		return errors.New("Asset not allowed")
	}

	err := pl.horizon.LoadMemo(&payment)
	if err != nil {
		return errors.Wrap(err, "Unable to load transaction memo")
	}

	pl.log.WithFields(logrus.Fields{"memo": payment.Memo.Value, "type": payment.Memo.Type}).Info("Loaded memo")

	var receiveResponse callback.ReceiveResponse
	var route string

	// Request extra_memo from compliance server
	if pl.config.Compliance != "" && payment.Memo.Type == "hash" {
		complianceRequestURL := pl.config.Compliance + "/receive"
		complianceRequestBody := url.Values{"memo": {string(payment.Memo.Value)}}

		pl.log.WithFields(logrus.Fields{"url": complianceRequestURL, "body": complianceRequestBody}).Info("Sending request to compliance server")
		resp, err := pl.postForm(complianceRequestURL, complianceRequestBody)
		if err != nil {
			return errors.Wrap(err, "Error sending request to compliance server")
		}

		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return errors.Wrap(err, "Error reading compliance server response")
		}

		if resp.StatusCode != 200 {
			pl.log.WithFields(logrus.Fields{
				"status": resp.StatusCode,
				"body":   string(body),
			}).Error("Error response from compliance server")
			return errors.New("Error response from compliance server")
		}

		err = json.Unmarshal([]byte(body), &receiveResponse)
		if err != nil {
			return errors.Wrap(err, "Cannot unmarshal receiveResponse")
		}

		var authData compliance.AuthData
		err = json.Unmarshal([]byte(receiveResponse.Data), &authData)
		if err != nil {
			return errors.Wrap(err, "Cannot unmarshal authData")
		}

		var attachment compliance.Attachment
		err = json.Unmarshal([]byte(authData.AttachmentJSON), &attachment)
		if err != nil {
			return errors.Wrap(err, "Cannot unmarshal memo")
		}

		route = attachment.Transaction.Route
	} else if payment.Memo.Type != "hash" {
		route = payment.Memo.Value
	}

	resp, err := pl.postForm(
		pl.config.Callbacks.Receive,
		url.Values{
			"id":           {payment.ID},
			"from":         {payment.From},
			"route":        {route},
			"amount":       {payment.Amount},
			"asset_code":   {payment.AssetCode},
			"asset_issuer": {payment.AssetIssuer},
			"memo_type":    {payment.Memo.Type},
			"memo":         {payment.Memo.Value},
			"data":         {receiveResponse.Data},
		},
	)
	if err != nil {
		return errors.Wrap(err, "Error sending request to receive callback")
	}

	if resp.StatusCode != 200 {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return errors.Wrap(err, "Error reading receive callback response")
		}

		pl.log.WithFields(logrus.Fields{
			"status": resp.StatusCode,
			"body":   string(body),
		}).Error("Error response from receive callback")
		return errors.New("Error response from receive callback")
	}

	return nil
}

func (pl *PaymentListener) isAssetAllowed(asset_type string, code string, issuer string) bool {
	for _, asset := range pl.config.Assets {
		if asset.Code == code && asset.Issuer == issuer {
			return true
		}

		if asset.Code == "XLM" && asset.Issuer == "" && asset_type == "native" {
			return true
		}

	}
	return false
}

func (pl *PaymentListener) postForm(
	url string,
	form url.Values,
) (*http.Response, error) {

	strbody := form.Encode()

	req, err := http.NewRequest("POST", url, strings.NewReader(strbody))
	if err != nil {
		return nil, errors.Wrap(err, "configure http request failed")
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	if pl.config.MACKey != "" {
		rawMAC, err := pl.getMAC(pl.config.MACKey, []byte(strbody))
		if err != nil {
			return nil, errors.Wrap(err, "getMAC failed")
		}

		encMAC := base64.StdEncoding.EncodeToString(rawMAC)
		req.Header.Set("X_PAYLOAD_MAC", encMAC)
	}

	resp, err := pl.client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "http request errored")
	}

	return resp, nil
}

func (pl *PaymentListener) getMAC(key string, raw []byte) ([]byte, error) {

	rawkey, err := strkey.Decode(strkey.VersionByteSeed, pl.config.MACKey)
	if err != nil {
		return nil, errors.Wrap(err, "invalid MAC key")
	}

	macer := hmac.New(sha256.New, rawkey)
	macer.Write(raw)
	return macer.Sum(nil), nil
}
