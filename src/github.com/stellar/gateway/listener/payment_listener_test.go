package listener

import (
	"errors"
	"net/url"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/stellar/gateway/bridge/config"
	"github.com/stellar/gateway/db/entities"
	"github.com/stellar/gateway/horizon"
	"github.com/stellar/gateway/mocks"
	"github.com/stellar/gateway/net"
	"github.com/stretchr/testify/assert"
)

func TestPaymentListener(t *testing.T) {
	mockEntityManager := new(mocks.MockEntityManager)
	mockHorizon := new(mocks.MockHorizon)
	mockRepository := new(mocks.MockRepository)
	mockHTTPClient := new(mocks.MockHTTPClient)

	config := &config.Config{
		Assets: []config.Asset{
			{Code: "USD", Issuer: "GD4I7AFSLZGTDL34TQLWJOM2NHLIIOEKD5RHHZUW54HERBLSIRKUOXRR"},
			{Code: "EUR", Issuer: "GD4I7AFSLZGTDL34TQLWJOM2NHLIIOEKD5RHHZUW54HERBLSIRKUOXRR"},
		},
		Accounts: config.Accounts{
			IssuingAccountID:   "GATKP6ZQM5CSLECPMTAC5226PE367QALCPM6AFHTSULPPZMT62OOPMQB",
			ReceivingAccountID: "GATKP6ZQM5CSLECPMTAC5226PE367QALCPM6AFHTSULPPZMT62OOPMQB",
		},
		Callbacks: config.Callbacks{
			Receive: "http://receive_callback",
		},
	}

	paymentListener, _ := NewPaymentListener(
		config,
		mockEntityManager,
		mockHorizon,
		mockRepository,
		mocks.Now,
	)

	paymentListener.client = mockHTTPClient

	Convey("PaymentListener", t, func() {
		operation := horizon.PaymentResponse{
			ID:          "1",
			From:        "GBIHSMPXC2KJ3NJVHEYTG3KCHYEUQRT45X6AWYWXMAXZOAX4F5LFZYYQ",
			PagingToken: "2",
			Amount:      "200",
		}

		mocks.PredefinedTime = time.Now()

		dbPayment := entities.ReceivedPayment{
			OperationID: operation.ID,
			ProcessedAt: mocks.PredefinedTime,
			PagingToken: operation.PagingToken,
		}

		Convey("When operation exists", func() {
			operation.Type = "payment"
			mockRepository.On("GetReceivedPaymentByID", int64(1)).Return(&entities.ReceivedPayment{}, nil).Once()

			Convey("it should save the status", func() {
				err := paymentListener.onPayment(operation)
				assert.Nil(t, err)
				mockEntityManager.AssertExpectations(t)
			})
		})

		Convey("When operation is not a payment", func() {
			operation.Type = "create_account"
			dbPayment.Status = "Not a payment operation"
			mockEntityManager.On("Persist", &dbPayment).Return(nil).Once()
			mockRepository.On("GetReceivedPaymentByID", int64(1)).Return(nil, nil).Once()

			Convey("it should save the status", func() {
				err := paymentListener.onPayment(operation)
				assert.Nil(t, err)
				mockEntityManager.AssertExpectations(t)
			})
		})

		Convey("When payment is sent not received", func() {
			operation.Type = "payment"
			operation.To = "GDNXBMIJLLLXZYKZBHXJ45WQ4AJQBRVT776YKGQTDBHTSPMNAFO3OZOS"
			dbPayment.Status = "Operation sent not received"
			mockEntityManager.On("Persist", &dbPayment).Return(nil).Once()
			mockRepository.On("GetReceivedPaymentByID", int64(1)).Return(nil, nil).Once()

			Convey("it should save the status", func() {
				err := paymentListener.onPayment(operation)
				assert.Nil(t, err)
				mockEntityManager.AssertExpectations(t)
			})
		})

		Convey("When asset is not allowed (issuer)", func() {
			operation.Type = "payment"
			operation.To = "GATKP6ZQM5CSLECPMTAC5226PE367QALCPM6AFHTSULPPZMT62OOPMQB"
			operation.AssetCode = "USD"
			operation.AssetIssuer = "GC4WWLMUGZJMRVJM7JUVVZBY3LJ5HL4RKIPADEGKEMLAAJEDRONUGYG7"
			dbPayment.Status = "Asset not allowed"
			mockEntityManager.On("Persist", &dbPayment).Return(nil).Once()
			mockRepository.On("GetReceivedPaymentByID", int64(1)).Return(nil, nil).Once()

			Convey("it should save the status", func() {
				err := paymentListener.onPayment(operation)
				assert.Nil(t, err)
				mockEntityManager.AssertExpectations(t)
			})
		})

		Convey("When asset is not allowed (code)", func() {
			operation.Type = "payment"
			operation.To = "GATKP6ZQM5CSLECPMTAC5226PE367QALCPM6AFHTSULPPZMT62OOPMQB"
			operation.AssetCode = "GBP"
			operation.AssetIssuer = "GD4I7AFSLZGTDL34TQLWJOM2NHLIIOEKD5RHHZUW54HERBLSIRKUOXRR"
			dbPayment.Status = "Asset not allowed"
			mockEntityManager.On("Persist", &dbPayment).Return(nil).Once()
			mockRepository.On("GetReceivedPaymentByID", int64(1)).Return(nil, nil).Once()

			Convey("it should save the status", func() {
				err := paymentListener.onPayment(operation)
				assert.Nil(t, err)
				mockEntityManager.AssertExpectations(t)
			})
		})

		Convey("When unable to load transaction memo", func() {
			operation.Type = "payment"
			operation.To = "GATKP6ZQM5CSLECPMTAC5226PE367QALCPM6AFHTSULPPZMT62OOPMQB"
			operation.AssetCode = "USD"
			operation.AssetIssuer = "GD4I7AFSLZGTDL34TQLWJOM2NHLIIOEKD5RHHZUW54HERBLSIRKUOXRR"

			mockRepository.On("GetReceivedPaymentByID", int64(1)).Return(nil, nil).Once()
			mockHorizon.On("LoadMemo", &operation).Return(errors.New("Connection error")).Once()

			Convey("it should return error", func() {
				err := paymentListener.onPayment(operation)
				assert.Error(t, err)
				mockHorizon.AssertExpectations(t)
				mockEntityManager.AssertNotCalled(t, "Persist")
			})
		})

		Convey("When receive callback returns error", func() {
			operation.Type = "payment"
			operation.To = "GATKP6ZQM5CSLECPMTAC5226PE367QALCPM6AFHTSULPPZMT62OOPMQB"
			operation.AssetCode = "USD"
			operation.AssetIssuer = "GD4I7AFSLZGTDL34TQLWJOM2NHLIIOEKD5RHHZUW54HERBLSIRKUOXRR"
			operation.Memo.Type = "text"
			operation.Memo.Value = "testing"

			mockRepository.On("GetReceivedPaymentByID", int64(1)).Return(nil, nil).Once()
			mockHorizon.On("LoadMemo", &operation).Return(nil).Once()

			mockHTTPClient.On(
				"PostForm",
				"http://receive_callback",
				url.Values{
					"id":         {"1"},
					"from":       {"GBIHSMPXC2KJ3NJVHEYTG3KCHYEUQRT45X6AWYWXMAXZOAX4F5LFZYYQ"},
					"amount":     {"200"},
					"asset_code": {"USD"},
					"memo_type":  {"text"},
					"memo":       {"testing"},
					"data":       {""},
				},
			).Return(
				net.BuildHTTPResponse(503, "ok"),
				nil,
			).Once()

			Convey("it should save the status", func() {
				err := paymentListener.onPayment(operation)
				assert.Error(t, err)
				mockHorizon.AssertExpectations(t)
				mockEntityManager.AssertNotCalled(t, "Persist")
			})
		})

		Convey("When receive callback returns success", func() {
			operation.Type = "payment"
			operation.To = "GATKP6ZQM5CSLECPMTAC5226PE367QALCPM6AFHTSULPPZMT62OOPMQB"
			operation.AssetCode = "USD"
			operation.AssetIssuer = "GD4I7AFSLZGTDL34TQLWJOM2NHLIIOEKD5RHHZUW54HERBLSIRKUOXRR"
			operation.Memo.Type = "text"
			operation.Memo.Value = "testing"

			dbPayment.Status = "Success"

			mockRepository.On("GetReceivedPaymentByID", int64(1)).Return(nil, nil).Once()
			mockHorizon.On("LoadMemo", &operation).Return(nil).Once()
			mockEntityManager.On("Persist", &dbPayment).Return(nil).Once()

			mockHTTPClient.On(
				"PostForm",
				"http://receive_callback",
				url.Values{
					"id":         {"1"},
					"from":       {"GBIHSMPXC2KJ3NJVHEYTG3KCHYEUQRT45X6AWYWXMAXZOAX4F5LFZYYQ"},
					"amount":     {"200"},
					"asset_code": {"USD"},
					"memo_type":  {"text"},
					"memo":       {"testing"},
					"data":       {""},
				},
			).Return(
				net.BuildHTTPResponse(200, "ok"),
				nil,
			).Once()

			Convey("it should save the status", func() {
				err := paymentListener.onPayment(operation)
				assert.Nil(t, err)
				mockHorizon.AssertExpectations(t)
				mockEntityManager.AssertExpectations(t)
			})
		})

		Convey("When receive callback returns success (no memo)", func() {
			operation.Type = "payment"
			operation.To = "GATKP6ZQM5CSLECPMTAC5226PE367QALCPM6AFHTSULPPZMT62OOPMQB"
			operation.AssetCode = "USD"
			operation.AssetIssuer = "GD4I7AFSLZGTDL34TQLWJOM2NHLIIOEKD5RHHZUW54HERBLSIRKUOXRR"

			dbPayment.Status = "Success"

			mockRepository.On("GetReceivedPaymentByID", int64(1)).Return(nil, nil).Once()
			mockHorizon.On("LoadMemo", &operation).Return(nil).Once()
			mockEntityManager.On("Persist", &dbPayment).Return(nil).Once()

			mockHTTPClient.On(
				"PostForm",
				"http://receive_callback",
				url.Values{
					"id":         {"1"},
					"from":       {"GBIHSMPXC2KJ3NJVHEYTG3KCHYEUQRT45X6AWYWXMAXZOAX4F5LFZYYQ"},
					"amount":     {"200"},
					"asset_code": {"USD"},
					"memo_type":  {""},
					"memo":       {""},
					"data":       {""},
				},
			).Return(
				net.BuildHTTPResponse(200, "ok"),
				nil,
			).Once()

			Convey("it should save the status", func() {
				err := paymentListener.onPayment(operation)
				assert.Nil(t, err)
				mockHorizon.AssertExpectations(t)
				mockEntityManager.AssertExpectations(t)
			})
		})

		Convey("When receive callback returns success and compliance server is connected", func() {
			paymentListener.config.Compliance = "http://compliance"

			operation.Type = "payment"
			operation.To = "GATKP6ZQM5CSLECPMTAC5226PE367QALCPM6AFHTSULPPZMT62OOPMQB"
			operation.AssetCode = "USD"
			operation.AssetIssuer = "GD4I7AFSLZGTDL34TQLWJOM2NHLIIOEKD5RHHZUW54HERBLSIRKUOXRR"
			operation.Memo.Type = "hash"
			operation.Memo.Value = "b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9"

			dbPayment.Status = "Success"

			mockRepository.On("GetReceivedPaymentByID", int64(1)).Return(nil, nil).Once()
			mockHorizon.On("LoadMemo", &operation).Return(nil).Once()
			mockEntityManager.On("Persist", &dbPayment).Return(nil).Once()

			mockHTTPClient.On(
				"PostForm",
				"http://compliance/receive",
				url.Values{"memo": {"b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9"}},
			).Return(
				net.BuildHTTPResponse(200, "{\"data\": \"hello world\"}"),
				nil,
			).Once()

			mockHTTPClient.On(
				"PostForm",
				"http://receive_callback",
				url.Values{
					"id":         {"1"},
					"from":       {"GBIHSMPXC2KJ3NJVHEYTG3KCHYEUQRT45X6AWYWXMAXZOAX4F5LFZYYQ"},
					"amount":     {"200"},
					"asset_code": {"USD"},
					"memo_type":  {"hash"},
					"memo":       {"b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9"},
					"data":       {"hello world"},
				},
			).Return(
				net.BuildHTTPResponse(200, "ok"),
				nil,
			).Once()

			Convey("it should save the status", func() {
				err := paymentListener.onPayment(operation)
				assert.Nil(t, err)
				mockHorizon.AssertExpectations(t)
				mockEntityManager.AssertExpectations(t)
			})
		})
	})
}
