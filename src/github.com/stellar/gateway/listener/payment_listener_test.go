package listener

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/stellar/gateway/bridge/config"
	"github.com/stellar/gateway/db/entities"
	"github.com/stellar/gateway/horizon"
	"github.com/stellar/gateway/mocks"
	"github.com/stellar/gateway/net"
	callback "github.com/stellar/gateway/protocols/compliance"
	"github.com/stellar/go/protocols/compliance"
	"github.com/stellar/go/strkey"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func ensurePaymentStatus(t *testing.T, operation horizon.PaymentResponse, status string) func(args mock.Arguments) {
	return func(args mock.Arguments) {
		payment := args.Get(0).(*entities.ReceivedPayment)
		assert.Equal(t, operation.ID, payment.OperationID)
		assert.Equal(t, mocks.PredefinedTime, payment.ProcessedAt)
		assert.Equal(t, operation.PagingToken, payment.PagingToken)
		assert.Equal(t, status, payment.Status)
		assert.Equal(t, operation.TransactionID, payment.TransactionID)
	}
}

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

	paymentListener, err := NewPaymentListener(
		config,
		mockEntityManager,
		mockHorizon,
		mockRepository,
		mocks.Now,
	)
	require.NoError(t, err)

	paymentListener.client = mockHTTPClient

	Convey("PaymentListener", t, func() {
		operation := horizon.PaymentResponse{
			ID:          "1",
			From:        "GBIHSMPXC2KJ3NJVHEYTG3KCHYEUQRT45X6AWYWXMAXZOAX4F5LFZYYQ",
			PagingToken: "2",
			Amount:      "200",
		}

		mocks.PredefinedTime = time.Now()

		config.Assets[1].Code = "EUR"
		config.Assets[1].Issuer = "GD4I7AFSLZGTDL34TQLWJOM2NHLIIOEKD5RHHZUW54HERBLSIRKUOXRR"

		Convey("When operation exists", func() {
			operation.Type = "payment"
			mockRepository.On("GetReceivedPaymentByOperationID", int64(1)).Return(&entities.ReceivedPayment{}, nil).Once()

			Convey("it should save the status", func() {
				err := paymentListener.onPayment(operation)
				assert.Nil(t, err)
				mockEntityManager.AssertExpectations(t)
			})
		})

		Convey("When operation is not a payment", func() {
			operation.Type = "create_account"

			mockEntityManager.On("Persist", mock.AnythingOfType("*entities.ReceivedPayment")).
				Run(ensurePaymentStatus(t, operation, "Processing...")).Return(nil).Once()

			mockEntityManager.On("Persist", mock.AnythingOfType("*entities.ReceivedPayment")).
				Run(ensurePaymentStatus(t, operation, "Not a payment operation")).Return(nil).Once()

			mockRepository.On("GetReceivedPaymentByOperationID", int64(1)).Return(nil, nil).Once()

			Convey("it should save the status", func() {
				err := paymentListener.onPayment(operation)
				assert.Nil(t, err)
				mockEntityManager.AssertExpectations(t)
				mockRepository.AssertExpectations(t)
			})
		})

		Convey("When payment is sent not received", func() {
			operation.Type = "payment"
			operation.To = "GDNXBMIJLLLXZYKZBHXJ45WQ4AJQBRVT776YKGQTDBHTSPMNAFO3OZOS"

			mockEntityManager.On("Persist", mock.AnythingOfType("*entities.ReceivedPayment")).
				Run(ensurePaymentStatus(t, operation, "Processing...")).Return(nil).Once()

			mockEntityManager.On("Persist", mock.AnythingOfType("*entities.ReceivedPayment")).
				Run(ensurePaymentStatus(t, operation, "Operation sent not received")).Return(nil).Once()

			mockRepository.On("GetReceivedPaymentByOperationID", int64(1)).Return(nil, nil).Once()

			Convey("it should save the status", func() {
				err := paymentListener.onPayment(operation)
				assert.Nil(t, err)
				mockEntityManager.AssertExpectations(t)
				mockRepository.AssertExpectations(t)
			})
		})

		Convey("When asset is not allowed (issuer)", func() {
			operation.Type = "payment"
			operation.To = "GATKP6ZQM5CSLECPMTAC5226PE367QALCPM6AFHTSULPPZMT62OOPMQB"
			operation.AssetCode = "USD"
			operation.AssetIssuer = "GC4WWLMUGZJMRVJM7JUVVZBY3LJ5HL4RKIPADEGKEMLAAJEDRONUGYG7"

			mockEntityManager.On("Persist", mock.AnythingOfType("*entities.ReceivedPayment")).
				Run(ensurePaymentStatus(t, operation, "Processing...")).Return(nil).Once()

			mockEntityManager.On("Persist", mock.AnythingOfType("*entities.ReceivedPayment")).
				Run(ensurePaymentStatus(t, operation, "Asset not allowed")).Return(nil).Once()

			mockRepository.On("GetReceivedPaymentByOperationID", int64(1)).Return(nil, nil).Once()

			Convey("it should save the status", func() {
				err := paymentListener.onPayment(operation)
				assert.Nil(t, err)
				mockEntityManager.AssertExpectations(t)
				mockRepository.AssertExpectations(t)
			})
		})

		Convey("When asset is not allowed (code)", func() {
			operation.Type = "payment"
			operation.To = "GATKP6ZQM5CSLECPMTAC5226PE367QALCPM6AFHTSULPPZMT62OOPMQB"
			operation.AssetCode = "GBP"
			operation.AssetIssuer = "GD4I7AFSLZGTDL34TQLWJOM2NHLIIOEKD5RHHZUW54HERBLSIRKUOXRR"

			mockEntityManager.On("Persist", mock.AnythingOfType("*entities.ReceivedPayment")).
				Run(ensurePaymentStatus(t, operation, "Processing...")).Return(nil).Once()

			mockEntityManager.On("Persist", mock.AnythingOfType("*entities.ReceivedPayment")).
				Run(ensurePaymentStatus(t, operation, "Asset not allowed")).Return(nil).Once()

			mockRepository.On("GetReceivedPaymentByOperationID", int64(1)).Return(nil, nil).Once()

			Convey("it should save the status", func() {
				err := paymentListener.onPayment(operation)
				assert.Nil(t, err)
				mockEntityManager.AssertExpectations(t)
				mockRepository.AssertExpectations(t)
			})
		})

		Convey("When payment is XLM (no XLM asset in config)", func() {
			operation.Type = "payment"
			operation.From = "GBL27BKG2JSDU6KQ5YJKCDWTVIU24VTG4PLB63SF4K2DBZS5XZMWRPVU"
			operation.To = "GATKP6ZQM5CSLECPMTAC5226PE367QALCPM6AFHTSULPPZMT62OOPMQB"
			operation.AssetCode = ""
			operation.AssetIssuer = ""
			operation.AssetType = "native"

			mockEntityManager.On("Persist", mock.AnythingOfType("*entities.ReceivedPayment")).
				Run(ensurePaymentStatus(t, operation, "Processing...")).Return(nil).Once()

			mockEntityManager.On("Persist", mock.AnythingOfType("*entities.ReceivedPayment")).
				Run(ensurePaymentStatus(t, operation, "Asset not allowed")).Return(nil).Once()

			mockRepository.On("GetReceivedPaymentByOperationID", int64(1)).Return(nil, nil).Once()

			Convey("it should save the status", func() {
				err := paymentListener.onPayment(operation)
				assert.Nil(t, err)
				mockEntityManager.AssertExpectations(t)
				mockRepository.AssertExpectations(t)
			})
		})

		Convey("When payment is XLM (XLM asset in config)", func() {
			operation.Type = "payment"
			operation.To = "GATKP6ZQM5CSLECPMTAC5226PE367QALCPM6AFHTSULPPZMT62OOPMQB"
			operation.AssetCode = ""
			operation.AssetIssuer = ""
			operation.AssetType = "native"
			operation.Memo.Type = "book"
			operation.Memo.Value = "testing"
			config.Assets[1].Code = "XLM"
			config.Assets[1].Issuer = ""

			mockRepository.On("GetReceivedPaymentByOperationID", int64(1)).Return(nil, nil).Once()
			mockHorizon.On("LoadMemo", &operation).Return(nil).Once()

			mockEntityManager.On("Persist", mock.AnythingOfType("*entities.ReceivedPayment")).
				Run(ensurePaymentStatus(t, operation, "Processing...")).Return(nil).Once()

			mockEntityManager.On("Persist", mock.AnythingOfType("*entities.ReceivedPayment")).
				Run(ensurePaymentStatus(t, operation, "Success")).Return(nil).Once()

			mockHTTPClient.On(
				"Do",
				mock.MatchedBy(func(req *http.Request) bool {
					return req.URL.String() == "http://receive_callback"
				}),
			).Return(
				net.BuildHTTPResponse(200, "ok"),
				nil,
			).Once()

			Convey("it should save the status", func() {
				err := paymentListener.onPayment(operation)
				assert.Nil(t, err)
				mockRepository.AssertExpectations(t)
				mockHorizon.AssertExpectations(t)
				mockEntityManager.AssertExpectations(t)
			})
		})

		Convey("When unable to load transaction memo", func() {
			operation.Type = "payment"
			operation.To = "GATKP6ZQM5CSLECPMTAC5226PE367QALCPM6AFHTSULPPZMT62OOPMQB"
			operation.AssetCode = "USD"
			operation.AssetIssuer = "GD4I7AFSLZGTDL34TQLWJOM2NHLIIOEKD5RHHZUW54HERBLSIRKUOXRR"

			mockRepository.On("GetReceivedPaymentByOperationID", int64(1)).Return(nil, nil).Once()

			mockEntityManager.On("Persist", mock.AnythingOfType("*entities.ReceivedPayment")).
				Run(ensurePaymentStatus(t, operation, "Processing...")).Return(nil).Once()

			mockEntityManager.On("Persist", mock.AnythingOfType("*entities.ReceivedPayment")).
				Run(ensurePaymentStatus(t, operation, "Unable to load transaction memo: Connection error")).Return(nil).Once()

			mockHorizon.On("LoadMemo", &operation).Return(errors.New("Connection error")).Once()

			Convey("it should save the status", func() {
				err := paymentListener.onPayment(operation)
				assert.NoError(t, err)
				mockHorizon.AssertExpectations(t)
				mockEntityManager.AssertExpectations(t)
				mockRepository.AssertExpectations(t)
			})
		})

		Convey("When receive callback returns error", func() {
			operation.Type = "payment"
			operation.To = "GATKP6ZQM5CSLECPMTAC5226PE367QALCPM6AFHTSULPPZMT62OOPMQB"
			operation.AssetCode = "USD"
			operation.AssetIssuer = "GD4I7AFSLZGTDL34TQLWJOM2NHLIIOEKD5RHHZUW54HERBLSIRKUOXRR"
			operation.Memo.Type = "text"
			operation.Memo.Value = "testing"

			mockRepository.On("GetReceivedPaymentByOperationID", int64(1)).Return(nil, nil).Once()
			mockHorizon.On("LoadMemo", &operation).Return(nil).Once()

			mockEntityManager.On("Persist", mock.AnythingOfType("*entities.ReceivedPayment")).
				Run(ensurePaymentStatus(t, operation, "Processing...")).Return(nil).Once()

			mockEntityManager.On("Persist", mock.AnythingOfType("*entities.ReceivedPayment")).
				Run(ensurePaymentStatus(t, operation, "Error response from receive callback")).Return(nil).Once()

			mockHTTPClient.On(
				"Do",
				mock.MatchedBy(func(req *http.Request) bool {
					return req.URL.String() == "http://receive_callback"
				}),
			).Return(
				net.BuildHTTPResponse(503, "ok"),
				nil,
			).Once()

			Convey("it should save the status", func() {
				err := paymentListener.onPayment(operation)
				assert.NoError(t, err)
				mockHorizon.AssertExpectations(t)
				mockEntityManager.AssertExpectations(t)
				mockRepository.AssertExpectations(t)
			})
		})

		Convey("When receive callback returns success", func() {
			operation.Type = "payment"
			operation.From = "GBL27BKG2JSDU6KQ5YJKCDWTVIU24VTG4PLB63SF4K2DBZS5XZMWRPVU"
			operation.To = "GATKP6ZQM5CSLECPMTAC5226PE367QALCPM6AFHTSULPPZMT62OOPMQB"
			operation.Amount = "100"
			operation.AssetCode = "USD"
			operation.AssetIssuer = "GD4I7AFSLZGTDL34TQLWJOM2NHLIIOEKD5RHHZUW54HERBLSIRKUOXRR"
			operation.Memo.Type = "text"
			operation.Memo.Value = "testing"

			mockRepository.On("GetReceivedPaymentByOperationID", int64(1)).Return(nil, nil).Once()
			mockHorizon.On("LoadMemo", &operation).Return(nil).Once()

			mockEntityManager.On("Persist", mock.AnythingOfType("*entities.ReceivedPayment")).
				Run(ensurePaymentStatus(t, operation, "Processing...")).Return(nil).Once()

			mockEntityManager.On("Persist", mock.AnythingOfType("*entities.ReceivedPayment")).
				Run(ensurePaymentStatus(t, operation, "Success")).Return(nil).Once()

			mockHTTPClient.On(
				"Do",
				mock.MatchedBy(func(req *http.Request) bool {
					return req.URL.String() == "http://receive_callback"
				}),
			).Return(
				net.BuildHTTPResponse(200, "ok"),
				nil,
			).Run(func(args mock.Arguments) {
				req := args.Get(0).(*http.Request)

				assert.Equal(t, operation.From, req.PostFormValue("from"))
				assert.Equal(t, operation.Amount, req.PostFormValue("amount"))
				assert.Equal(t, operation.AssetCode, req.PostFormValue("asset_code"))
				assert.Equal(t, operation.AssetIssuer, req.PostFormValue("asset_issuer"))
				assert.Equal(t, operation.Memo.Type, req.PostFormValue("memo_type"))
				assert.Equal(t, operation.Memo.Value, req.PostFormValue("memo"))
			}).Once()

			Convey("it should save the status", func() {
				err := paymentListener.onPayment(operation)
				assert.Nil(t, err)
				mockHorizon.AssertExpectations(t)
				mockEntityManager.AssertExpectations(t)
				mockRepository.AssertExpectations(t)
			})
		})

		Convey("When receive callback returns success (no memo)", func() {
			operation.Type = "payment"
			operation.To = "GATKP6ZQM5CSLECPMTAC5226PE367QALCPM6AFHTSULPPZMT62OOPMQB"
			operation.AssetCode = "USD"
			operation.AssetIssuer = "GD4I7AFSLZGTDL34TQLWJOM2NHLIIOEKD5RHHZUW54HERBLSIRKUOXRR"

			mockEntityManager.On("Persist", mock.AnythingOfType("*entities.ReceivedPayment")).
				Run(ensurePaymentStatus(t, operation, "Processing...")).Return(nil).Once()

			mockEntityManager.On("Persist", mock.AnythingOfType("*entities.ReceivedPayment")).
				Run(ensurePaymentStatus(t, operation, "Success")).Return(nil).Once()

			mockRepository.On("GetReceivedPaymentByOperationID", int64(1)).Return(nil, nil).Once()
			mockHorizon.On("LoadMemo", &operation).Return(nil).Once()

			mockHTTPClient.On(
				"Do",
				mock.MatchedBy(func(req *http.Request) bool {
					return req.URL.String() == "http://receive_callback"
				}),
			).Return(
				net.BuildHTTPResponse(200, "ok"),
				nil,
			).Once()

			Convey("it should save the status", func() {
				err := paymentListener.onPayment(operation)
				assert.Nil(t, err)
				mockHorizon.AssertExpectations(t)
				mockEntityManager.AssertExpectations(t)
				mockRepository.AssertExpectations(t)
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

			mockRepository.On("GetReceivedPaymentByOperationID", int64(1)).Return(nil, nil).Once()
			mockHorizon.On("LoadMemo", &operation).Return(nil).Once()

			mockEntityManager.On("Persist", mock.AnythingOfType("*entities.ReceivedPayment")).
				Run(ensurePaymentStatus(t, operation, "Processing...")).Return(nil).Once()

			mockEntityManager.On("Persist", mock.AnythingOfType("*entities.ReceivedPayment")).
				Run(ensurePaymentStatus(t, operation, "Success")).Return(nil).Once()

			attachment := compliance.Attachment{
				Transaction: compliance.Transaction{
					Route: "jed*stellar.org",
				},
			}

			attachmentString, _ := json.Marshal(attachment)

			auth := compliance.AuthData{
				AttachmentJSON: string(attachmentString),
			}

			authString, _ := json.Marshal(auth)

			response := callback.ReceiveResponse{
				Data: string(authString),
			}

			responseString, _ := json.Marshal(response)

			mockHTTPClient.On(
				"Do",
				mock.MatchedBy(func(req *http.Request) bool {
					return req.URL.String() == "http://compliance/receive"
				}),
			).Return(
				net.BuildHTTPResponse(200, string(responseString)),
				nil,
			).Once()

			mockHTTPClient.On(
				"Do",
				mock.MatchedBy(func(req *http.Request) bool {
					return req.URL.String() == "http://receive_callback"
				}),
			).Return(
				net.BuildHTTPResponse(200, "ok"),
				nil,
			).Once()

			Convey("it should save the status", func() {
				err := paymentListener.onPayment(operation)
				assert.Nil(t, err)
				mockHorizon.AssertExpectations(t)
				mockEntityManager.AssertExpectations(t)
				mockRepository.AssertExpectations(t)
			})
		})

		Convey("Reprocessing a payment", func() {
			Convey("it should reprocess a payment when a payment exists", func() {
				operation := horizon.PaymentResponse{
					ID:          "1",
					From:        "GBIHSMPXC2KJ3NJVHEYTG3KCHYEUQRT45X6AWYWXMAXZOAX4F5LFZYYQ",
					PagingToken: "2",
					Amount:      "200",
					Type:        "payment",
					To:          "GATKP6ZQM5CSLECPMTAC5226PE367QALCPM6AFHTSULPPZMT62OOPMQB",
					AssetCode:   "USD",
					AssetIssuer: "GD4I7AFSLZGTDL34TQLWJOM2NHLIIOEKD5RHHZUW54HERBLSIRKUOXRR",
				}

				var id int64 = 3
				existingPayment := entities.ReceivedPayment{
					ID:          &id,
					OperationID: operation.ID,
					ProcessedAt: mocks.PredefinedTime,
					PagingToken: operation.PagingToken,
					Status:      "Failed",
				}
				existingPayment.SetExists()

				mockHorizon.On("LoadMemo", &operation).Return(nil).Once()
				mockRepository.On("GetReceivedPaymentByOperationID", int64(1)).Return(&existingPayment, nil).Once()

				mockEntityManager.On("Persist", &existingPayment).Return(nil).
					Run(func(args mock.Arguments) {
						dbPayment := args.Get(0).(*entities.ReceivedPayment)
						assert.Equal(t, false, dbPayment.IsNew())
						assert.Equal(t, int64(3), *dbPayment.ID)
						assert.Equal(t, "Reprocessing...", dbPayment.Status)
					}).Once()

				mockEntityManager.On("Persist", &existingPayment).Return(nil).
					Run(func(args mock.Arguments) {
						dbPayment := args.Get(0).(*entities.ReceivedPayment)
						assert.Equal(t, false, dbPayment.IsNew())
						assert.Equal(t, int64(3), *dbPayment.ID)
						assert.Equal(t, "Success", dbPayment.Status)
					}).Once()

				mockHTTPClient.On(
					"Do",
					mock.MatchedBy(func(req *http.Request) bool {
						return req.URL.String() == "http://receive_callback"
					}),
				).Return(
					net.BuildHTTPResponse(200, "ok"),
					nil,
				).Once()

				err := paymentListener.ReprocessPayment(operation, false)
				assert.Nil(t, err)
				mockHorizon.AssertExpectations(t)
				mockEntityManager.AssertExpectations(t)
			})
		})

		Convey("When transaction id is not returned (current horizon release)", func() {
			operation.Type = "payment"
			operation.To = "GATKP6ZQM5CSLECPMTAC5226PE367QALCPM6AFHTSULPPZMT62OOPMQB"
			operation.AssetCode = "USD"
			operation.AssetIssuer = "GD4I7AFSLZGTDL34TQLWJOM2NHLIIOEKD5RHHZUW54HERBLSIRKUOXRR"
			operation.Memo.Type = "text"
			operation.Memo.Value = "testing"

			mockRepository.On("GetReceivedPaymentByOperationID", int64(1)).Return(nil, nil).Once()
			mockHorizon.On("LoadMemo", &operation).Return(nil).Once()

			mockEntityManager.On("Persist", mock.AnythingOfType("*entities.ReceivedPayment")).
				Run(ensurePaymentStatus(t, operation, "Processing...")).Return(nil).Once()

			mockEntityManager.On("Persist", mock.AnythingOfType("*entities.ReceivedPayment")).
				Run(ensurePaymentStatus(t, operation, "Error response from receive callback")).Return(nil).Once()

			mockHTTPClient.On(
				"Do",
				mock.MatchedBy(func(req *http.Request) bool {
					return req.URL.String() == "http://receive_callback"
				}),
			).Return(
				net.BuildHTTPResponse(503, "ok"),
				nil,
			).Once()

			Convey("it should save the status", func() {
				err := paymentListener.onPayment(operation)
				assert.NoError(t, err)
				mockHorizon.AssertExpectations(t)
				mockEntityManager.AssertExpectations(t)
				mockRepository.AssertExpectations(t)
			})
		})

		Convey("When transaction id is not returned by horizon", func() {
			operation.Type = "payment"
			operation.From = "GBL27BKG2JSDU6KQ5YJKCDWTVIU24VTG4PLB63SF4K2DBZS5XZMWRPVU"
			operation.To = "GATKP6ZQM5CSLECPMTAC5226PE367QALCPM6AFHTSULPPZMT62OOPMQB"
			operation.Amount = "100"
			operation.AssetCode = "USD"
			operation.AssetIssuer = "GD4I7AFSLZGTDL34TQLWJOM2NHLIIOEKD5RHHZUW54HERBLSIRKUOXRR"
			operation.Memo.Type = "text"
			operation.Memo.Value = "testing"
			operation.TransactionID = "18ce6c3320d35e683cb653a3e812ce43e8f5c24ab0d0e87668d5591c679c9755"

			mockRepository.On("GetReceivedPaymentByOperationID", int64(1)).Return(nil, nil).Once()
			mockHorizon.On("LoadMemo", &operation).Return(nil).Once()

			mockEntityManager.On("Persist", mock.AnythingOfType("*entities.ReceivedPayment")).
				Run(ensurePaymentStatus(t, operation, "Processing...")).Return(nil).Once()

			mockEntityManager.On("Persist", mock.AnythingOfType("*entities.ReceivedPayment")).
				Run(ensurePaymentStatus(t, operation, "Success")).Return(nil).Once()

			mockHTTPClient.On(
				"Do",
				mock.MatchedBy(func(req *http.Request) bool {
					return req.URL.String() == "http://receive_callback"
				}),
			).Return(
				net.BuildHTTPResponse(200, "ok"),
				nil,
			).Run(func(args mock.Arguments) {
				req := args.Get(0).(*http.Request)

				assert.Equal(t, operation.From, req.PostFormValue("from"))
				assert.Equal(t, operation.Amount, req.PostFormValue("amount"))
				assert.Equal(t, operation.AssetCode, req.PostFormValue("asset_code"))
				assert.Equal(t, operation.AssetIssuer, req.PostFormValue("asset_issuer"))
				assert.Equal(t, operation.Memo.Type, req.PostFormValue("memo_type"))
				assert.Equal(t, operation.Memo.Value, req.PostFormValue("memo"))
				assert.Equal(t, operation.TransactionID, req.PostFormValue("transaction_id"))
			}).Once()

			Convey("it should save the status", func() {
				err := paymentListener.onPayment(operation)
				assert.Nil(t, err)
				mockHorizon.AssertExpectations(t)
				mockEntityManager.AssertExpectations(t)
				mockRepository.AssertExpectations(t)
			})
		})

	})
}

func TestPostForm_MACKey(t *testing.T) {
	validKey := "SABLR5HOI2IUOYB27TR4TO7HWDJIGSRJTT4UUTXXZOFVVPGQKJ5ME43J"
	rawkey, err := strkey.Decode(strkey.VersionByteSeed, validKey)
	require.NoError(t, err)

	handler := http.NewServeMux()
	handler.HandleFunc("/no_mac", func(w http.ResponseWriter, req *http.Request) {
		assert.Empty(t, req.Header.Get("X_PAYLOAD_MAC"), "unexpected MAC present")
	})
	handler.HandleFunc("/mac", func(w http.ResponseWriter, req *http.Request) {
		body, err := ioutil.ReadAll(req.Body)
		require.NoError(t, err)

		macer := hmac.New(sha256.New, rawkey)
		macer.Write(body)
		rawExpected := macer.Sum(nil)
		encExpected := base64.StdEncoding.EncodeToString(rawExpected)

		assert.Equal(t, encExpected, req.Header.Get("X_PAYLOAD_MAC"), "MAC is wrong")
	})

	srv := httptest.NewServer(handler)
	defer srv.Close()

	cfg := &config.Config{}
	pl, err := NewPaymentListener(cfg, nil, nil, nil, nil)
	require.NoError(t, err)

	// no mac if the key is not set
	_, err = pl.postForm(srv.URL+"/no_mac", url.Values{"foo": []string{"base"}})
	require.NoError(t, err)

	// generates a valid mac if a key is set.
	cfg.MACKey = validKey
	_, err = pl.postForm(srv.URL+"/mac", url.Values{"foo": []string{"base"}})
	require.NoError(t, err)

	// errors is the key is invalid
	cfg.MACKey = "broken"
	_, err = pl.postForm(srv.URL+"/mac", url.Values{"foo": []string{"base"}})

	if assert.Error(t, err) {
		assert.Contains(t, err.Error(), "invalid MAC key")
	}
}
