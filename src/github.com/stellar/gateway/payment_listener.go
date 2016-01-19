package gateway

import (
	"io/ioutil"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/stellar/gateway/db"
	"github.com/stellar/gateway/horizon"
	"github.com/stellar/go-stellar-base/keypair"
)

type PaymentListener struct {
	config         *Config
	entityManager  *db.EntityManager
	horizon        *horizon.Horizon
	log            *logrus.Entry
	repository     *db.Repository
	issuingAccount keypair.KP
}

func NewPaymentListener(
	config *Config, 
	entityManager *db.EntityManager, 
	horizon *horizon.Horizon,
	repository *db.Repository,
) (pl PaymentListener, err error) {
	pl.config = config
	pl.entityManager = entityManager
	pl.horizon = horizon
	pl.repository = repository
	pl.issuingAccount, err = keypair.Parse(config.Accounts.IssuingSeed)
	pl.log = logrus.WithFields(logrus.Fields{
		"service": "PaymentListener",
	})
	return
}

func (pl PaymentListener) Listen() (err error) {
	accountId := pl.config.Accounts.ReceivingAccountId

	_, err = pl.horizon.LoadAccount(accountId)
	if err != nil {
		return
	}

	cursor, err := pl.repository.GetLastCursorValue()
	if err != nil {
		return
	}

	pl.log.WithFields(logrus.Fields{
		"accountId": accountId,
		"cursor": cursor,
	}).Info("Started listening for new payments")

	go func() {
		for {
			err = pl.horizon.StreamPayments(
				accountId,
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
	pl.log.WithFields(logrus.Fields{"id": payment.Id}).Info("New payment")

	dbPayment := db.ReceivedPayment{
		OperationId: payment.Id,
		ProcessedAt: time.Now(),
		PagingToken: payment.PagingToken,
	}

	savePayment := func(payment *db.ReceivedPayment) (err error) {
		err = pl.entityManager.Persist(payment)
		return
	}

	if payment.Type != "payment" {
		dbPayment.Status = "Not a payment operation"
		savePayment(&dbPayment)
		return
	}

	if payment.To != pl.config.Accounts.ReceivingAccountId {
		dbPayment.Status = "Operation sent not received"
		savePayment(&dbPayment)
		return nil
	}

	if pl.isAssetAllowed(payment.AssetCode, payment.AssetIssuer) {
		dbPayment.Status = "Asset not allowed"
		savePayment(&dbPayment)
		return nil
	}

	err = loadMemo(&payment)
	if err != nil {
		pl.log.Error("Cannot load transaction memo")
		return err
	}

	if payment.Memo.Type != "" && payment.Memo.Value != "" {
		dbPayment.Status = "Transaction does not have memo"
		savePayment(&dbPayment)
		return nil
	}

	resp, err := http.PostForm(
		pl.config.Hooks.Receive,
		url.Values{
			"amount": {payment.Amount},
			"asset_code": {payment.AssetCode},
			"memo_type": {payment.Memo.Type},
			"memo": {payment.Memo.Value},
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
			"body": body,
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
	if issuer != pl.issuingAccount.Address() {
		return false
	}

	for _, b := range pl.config.Assets {
		if b == code {
			return true
		}
	}
	return false
}

func loadMemo(p *horizon.PaymentResponse) error {
	res, err := http.Get(p.Links.Transaction.Href)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	return json.NewDecoder(res.Body).Decode(&p.Memo)
}
