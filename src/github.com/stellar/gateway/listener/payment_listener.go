package listener

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/stellar/gateway/config"
	"github.com/stellar/gateway/db"
	"github.com/stellar/gateway/horizon"
	"github.com/stellar/go-stellar-base/keypair"
)

type PaymentListener struct {
	config         *config.Config
	entityManager  db.EntityManagerInterface
	horizon        horizon.HorizonInterface
	log            *logrus.Entry
	repository     db.RepositoryInterface
	issuingAccount keypair.KP
	now            func() time.Time
}

const HOOK_TIMEOUT = 10 * time.Second

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
	pl.issuingAccount, err = keypair.Parse(*config.Accounts.IssuingSeed)
	pl.now = now
	pl.log = logrus.WithFields(logrus.Fields{
		"service": "PaymentListener",
	})
	return
}

func (pl PaymentListener) Listen() (err error) {
	accountId := *pl.config.Accounts.ReceivingAccountId

	_, err = pl.horizon.LoadAccount(accountId)
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
			}

			pl.log.WithFields(logrus.Fields{
				"accountId": accountId,
				"cursor":    cursorValue,
			}).Info("Started listening for new payments")

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
		ProcessedAt: pl.now(),
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

	if payment.To != *pl.config.Accounts.ReceivingAccountId {
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

	if payment.Memo.Type == "" || payment.Memo.Value == "" {
		dbPayment.Status = "Transaction does not have memo"
		savePayment(&dbPayment)
		return nil
	}

	client := http.Client{
		Timeout: HOOK_TIMEOUT,
	}
	resp, err := client.PostForm(
		*pl.config.Hooks.Receive,
		url.Values{
			"id":         {payment.Id},
			"from":       {payment.From},
			"amount":     {payment.Amount},
			"asset_code": {payment.AssetCode},
			"memo_type":  {payment.Memo.Type},
			"memo":       {payment.Memo.Value},
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
