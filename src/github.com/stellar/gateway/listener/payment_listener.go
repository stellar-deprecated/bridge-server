package listener

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/stellar/gateway/bridge/config"
	"github.com/stellar/gateway/db"
	"github.com/stellar/gateway/db/entities"
	"github.com/stellar/gateway/horizon"
	"github.com/stellar/gateway/protocols/compliance"
)

// PaymentListener is listening for a new payments received by ReceivingAccount
type PaymentListener struct {
	config        *config.Config
	entityManager db.EntityManagerInterface
	horizon       horizon.HorizonInterface
	log           *logrus.Entry
	repository    db.RepositoryInterface
	now           func() time.Time
}

const hookTimeout = 10 * time.Second

// NewPaymentListener creates a new PaymentListener
func NewPaymentListener(
	config *config.Config,
	entityManager db.EntityManagerInterface,
	horizon horizon.HorizonInterface,
	repository db.RepositoryInterface,
	now func() time.Time,
) (pl PaymentListener, err error) {
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
func (pl PaymentListener) Listen() (err error) {
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
			pl.log.Info("Streaming connection closed. Restarting...")
		}
	}()

	return
}

func (pl PaymentListener) onPayment(payment horizon.PaymentResponse) (err error) {
	pl.log.WithFields(logrus.Fields{"id": payment.ID}).Info("New received payment")

	id, err := strconv.ParseInt(payment.ID, 10, 64)
	if err != nil {
		pl.log.WithFields(logrus.Fields{"err": err}).Error("Error converting ID to int64")
		return err
	}

	existingPayment, err := pl.repository.GetReceivedPaymentByID(id)
	if err != nil {
		pl.log.WithFields(logrus.Fields{"err": err}).Error("Error checking if receive payment exists")
		return err
	}

	if existingPayment != nil {
		pl.log.WithFields(logrus.Fields{"id": payment.ID}).Info("Payment already exists")
		return
	}

	dbPayment := entities.ReceivedPayment{
		OperationID: payment.ID,
		ProcessedAt: pl.now(),
		PagingToken: payment.PagingToken,
	}

	savePayment := func(payment *entities.ReceivedPayment) (err error) {
		err = pl.entityManager.Persist(payment)
		return
	}

	if payment.Type != "payment" && payment.Type != "path_payment" {
		dbPayment.Status = "Not a payment operation"
		savePayment(&dbPayment)
		return
	}

	if payment.To != pl.config.Accounts.ReceivingAccountID {
		dbPayment.Status = "Operation sent not received"
		savePayment(&dbPayment)
		return nil
	}

	if !pl.isAssetAllowed(payment.AssetCode, payment.AssetIssuer) {
		dbPayment.Status = "Asset not allowed"
		savePayment(&dbPayment)
		return nil
	}

	err = pl.horizon.LoadMemo(&payment)
	if err != nil {
		pl.log.Error("Unable to load transaction memo")
		return err
	}

	var receiveResponse compliance.ReceiveResponse

	// Request extra_memo from compliance server
	if pl.config.Compliance != "" && payment.Memo.Type == "hash" {
		resp, err := http.PostForm(
			pl.config.Compliance+"/receive",
			url.Values{"memo": {string(payment.Memo.Value)}},
		)
		if err != nil {
			pl.log.WithFields(logrus.Fields{"err": err}).Error("Error sending request to compliance server")
			return err
		}

		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			pl.log.Error("Error reading compliance server response")
			return err
		}

		if resp.StatusCode != 200 {
			pl.log.WithFields(logrus.Fields{
				"status": resp.StatusCode,
				"body":   string(body),
			}).Error("Error response from compliance server")
			return err
		}

		err = json.Unmarshal([]byte(body), &receiveResponse)
		if err != nil {
			pl.log.WithFields(logrus.Fields{"err": err}).Error("Cannot unmarshal receiveResponse")
			return err
		}
	}

	client := http.Client{
		Timeout: hookTimeout,
	}
	resp, err := client.PostForm(
		pl.config.Hooks.Receive,
		url.Values{
			"id":         {payment.ID},
			"from":       {payment.From},
			"amount":     {payment.Amount},
			"asset_code": {payment.AssetCode},
			"memo_type":  {payment.Memo.Type},
			"memo":       {payment.Memo.Value},
			"extra_memo": {receiveResponse.Memo},
		},
	)
	if err != nil {
		pl.log.Error("Error sending request to receive hook")
		return err
	}

	if resp.StatusCode != 200 {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			pl.log.Error("Error reading receive hook response")
			return err
		}

		pl.log.WithFields(logrus.Fields{
			"status": resp.StatusCode,
			"body":   string(body),
		}).Error("Error response from receive hook")
		return errors.New("Error response from receive hook")
	}

	dbPayment.Status = "Success"
	err = savePayment(&dbPayment)
	if err != nil {
		pl.log.Error("Error saving payment to the DB")
		return err
	}

	return nil
}

func (pl PaymentListener) isAssetAllowed(code string, issuer string) bool {
	for _, asset := range pl.config.Assets {
		if asset.Code == code && asset.Issuer == issuer {
			return true
		}
	}
	return false
}
