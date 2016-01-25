package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/stellar/gateway/config"
	"github.com/stellar/gateway/horizon"
	"github.com/stellar/gateway/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRequestHandlerPayment(t *testing.T) {
	mockHorizon := new(mocks.MockHorizon)

	requestHandler := RequestHandler{
		Config:  &config.Config{
			NetworkPassphrase: "Test SDF Network ; September 2015",
		},
		Horizon: mockHorizon,
	}

	testServer := httptest.NewServer(http.HandlerFunc(requestHandler.Payment))
	defer testServer.Close()

	Convey("Given payment request", t, func() {
		Convey("When source is invalid", func() {
			source := "SDRAS7XIQNX25UDCCX725R4EYGBFYGJE4HJ2A3DFCWJIHMRSMS7CXX43"

			Convey("it should return error", func() {
				statusCode, response := getResponse(testServer, url.Values{"source": {source}})
				responseString := strings.TrimSpace(string(response))
				assert.Equal(t, 400, statusCode)
				assert.Equal(t, errorResponseString("invalid_source", "source parameter is invalid"), responseString)
			})
		})

		Convey("When destination is invalid", func() {
			source := "SDRAS7XIQNX25UDCCX725R4EYGBFYGJE4HJ2A3DFCWJIHMRSMS7CXX42"
			destination := "GD3YBOYIUVLU"

			Convey("it should return error", func() {
				statusCode, response := getResponse(testServer, url.Values{"destination": {destination}, "source": {source}})
				responseString := strings.TrimSpace(string(response))
				assert.Equal(t, 400, statusCode)
				assert.Equal(t, errorResponseString("invalid_destination", "destination parameter is invalid"), responseString)
			})
		})

		// Convey("When destination can't be resolved", func() {
		// 	source := "SDRAS7XIQNX25UDCCX725R4EYGBFYGJE4HJ2A3DFCWJIHMRSMS7CXX42"
		// 	destination := "bob*stellar.org"

		// 	Convey("it should return error", func() {
		// 		statusCode, response := getResponse(testServer, url.Values{"destination": {destination}, "source": {source}})
		// 		responseString := strings.TrimSpace(string(response))
		// 		assert.Equal(t, 400, statusCode)
		// 		assert.Equal(t, errorResponseString("invalid_destination", "Cannot resolve destination"), responseString)
		// 	})
		// })

		Convey("When assetIssuer is invalid", func() {
			source := "SDRAS7XIQNX25UDCCX725R4EYGBFYGJE4HJ2A3DFCWJIHMRSMS7CXX42"
			destination := "GDSIKW43UA6JTOA47WVEBCZ4MYC74M3GNKNXTVDXFHXYYTNO5GGVN632"
			assetCode := "USD"
			assetIssuer := "GDSIKW43UA6JTOA47WVEBCZ4MYC74M3GNKNXTVDXFHXYYTNO5GGVN631"

			Convey("it should return error", func() {
				statusCode, response := getResponse(
					testServer,
					url.Values{
						"source":       {source},
						"destination":  {destination},
						"asset_code":   {assetCode},
						"asset_issuer": {assetIssuer},
					},
				)
				responseString := strings.TrimSpace(string(response))
				assert.Equal(t, 400, statusCode)
				assert.Equal(t, errorResponseString("invalid_issuer", "asset_issuer parameter is invalid"), responseString)
			})
		})

		Convey("When assetCode is invalid", func() {
			source := "SDRAS7XIQNX25UDCCX725R4EYGBFYGJE4HJ2A3DFCWJIHMRSMS7CXX42"
			destination := "GDSIKW43UA6JTOA47WVEBCZ4MYC74M3GNKNXTVDXFHXYYTNO5GGVN632"
			amount := "20"
			assetCode := "1234567890123"
			assetIssuer := "GDSIKW43UA6JTOA47WVEBCZ4MYC74M3GNKNXTVDXFHXYYTNO5GGVN632"

			Convey("it should return error", func() {
				statusCode, response := getResponse(
					testServer,
					url.Values{
						"source":       {source},
						"destination":  {destination},
						"amount":       {amount},
						"asset_code":   {assetCode},
						"asset_issuer": {assetIssuer},
					},
				)
				responseString := strings.TrimSpace(string(response))
				assert.Equal(t, 400, statusCode)
				assert.Equal(t, errorResponseString("asset_code_invalid", "asset_code param is invalid"), responseString)
			})
		})

		Convey("When amount is invalid", func() {
			source := "SDRAS7XIQNX25UDCCX725R4EYGBFYGJE4HJ2A3DFCWJIHMRSMS7CXX42"
			destination := "GDSIKW43UA6JTOA47WVEBCZ4MYC74M3GNKNXTVDXFHXYYTNO5GGVN632"
			amount := "test"
			assetCode := "USD"
			assetIssuer := "GDSIKW43UA6JTOA47WVEBCZ4MYC74M3GNKNXTVDXFHXYYTNO5GGVN632"

			Convey("it should return error", func() {
				statusCode, response := getResponse(
					testServer,
					url.Values{
						"source":       {source},
						"destination":  {destination},
						"amount":       {amount},
						"asset_code":   {assetCode},
						"asset_issuer": {assetIssuer},
					},
				)
				responseString := strings.TrimSpace(string(response))
				assert.Equal(t, 400, statusCode)
				assert.Equal(t, errorResponseString("invalid_amount", "amount is invalid"), responseString)
			})
		})

		Convey("When params are valid", func() {
			validParams := url.Values{
				// GCF3WVYTHF75PEG6622G5G6KU26GOSDQPDHSCJ3DQD7VONH4EYVDOGKJ
				"source":       {"SDWLS4G3XCNIYPKXJWWGGJT6UDY63WV6PEFTWP7JZMQB4RE7EUJQN5XM"},
				"destination":  {"GDSIKW43UA6JTOA47WVEBCZ4MYC74M3GNKNXTVDXFHXYYTNO5GGVN632"},
				"amount":       {"20"},
				"asset_code":   {"USD"},
				"asset_issuer": {"GDSIKW43UA6JTOA47WVEBCZ4MYC74M3GNKNXTVDXFHXYYTNO5GGVN632"},
			}

			Convey("When memo is set", func() {
				Convey("only `memo` param is set", func() {
					validParams.Add("memo", "test")
					statusCode, response := getResponse(testServer, validParams)
					responseString := strings.TrimSpace(string(response))
					assert.Equal(t, 400, statusCode)
					assert.Equal(t, errorResponseString("memo_missing_param", "When passing memo both params: `memo_type`, `memo` are required"), responseString)
				})

				Convey("only `memo_type` param is set", func() {
					validParams.Add("memo_type", "id")
					statusCode, response := getResponse(testServer, validParams)
					responseString := strings.TrimSpace(string(response))
					assert.Equal(t, 400, statusCode)
					assert.Equal(t, errorResponseString("memo_missing_param", "When passing memo both params: `memo_type`, `memo` are required"), responseString)
				})

				Convey("unsupported memo_type", func() {
					validParams.Add("memo_type", "hash")
					validParams.Add("memo", "0123456789")
					statusCode, response := getResponse(testServer, validParams)
					responseString := strings.TrimSpace(string(response))
					assert.Equal(t, 400, statusCode)
					assert.Equal(t, errorResponseString("memo_not_supported", "Not supported memo type"), responseString)
				})

				Convey("memo is attached to the transaction", func() {
					mockHorizon.On(
						"LoadAccount",
						"GCF3WVYTHF75PEG6622G5G6KU26GOSDQPDHSCJ3DQD7VONH4EYVDOGKJ",
					).Return(
						horizon.AccountResponse{
							SequenceNumber: 100,
						},
						nil,
					).Once()

					var ledger uint64
					ledger = 1988727
					horizonResponse := horizon.SubmitTransactionResponse{&ledger, nil, nil}

					mockHorizon.On(
						"SubmitTransaction",
						"AAAAAIu7VxM5f9eQ3va0bpvKprxnSHB4zyEnY4D/VzT8Jio3AAAAZAAAAAAAAABlAAAAAAAAAAIAAAAAAAAAewAAAAEAAAAAAAAAAQAAAADkhVuboDyZuBz9qkCLPGYF/jNmapt51Hcp74xNrumNVgAAAAFVU0QAAAAAAOSFW5ugPJm4HP2qQIs8ZgX+M2Zqm3nUdynvjE2u6Y1WAAAAAAvrwgAAAAAAAAAAAfwmKjcAAABADsRVwB27jfr3OthAWlRMSrxAIDPENw1dOfga5/cahnIneJQ5NPS5g96Rp8Y5xTsOU3Y9JmBDKB8g8lXFCXdwCA==",
					).Return(horizonResponse, nil).Once()

					validParams.Add("memo_type", "id")
					validParams.Add("memo", "123")
					statusCode, response := getResponse(testServer, validParams)
					responseString := strings.TrimSpace(string(response))

					expectedResponse, err := json.MarshalIndent(horizonResponse, "", "  ")
					if err != nil {
						panic(err)
					}

					assert.Equal(t, 200, statusCode)
					assert.Equal(t, string(expectedResponse), responseString)
				})
			})

			Convey("source account does not exist", func() {
				mockHorizon.On(
					"LoadAccount",
					"GCF3WVYTHF75PEG6622G5G6KU26GOSDQPDHSCJ3DQD7VONH4EYVDOGKJ",
				).Return(horizon.AccountResponse{}, errors.New("Not found")).Once()

				Convey("it should return error", func() {
					statusCode, response := getResponse(testServer, validParams)
					responseString := strings.TrimSpace(string(response))
					assert.Equal(t, 400, statusCode)
					assert.Equal(t, errorResponseString("source_not_exist", "source account does not exist"), responseString)
				})
			})

			Convey("transaction failed in horizon", func() {
				mockHorizon.On(
					"LoadAccount",
					"GCF3WVYTHF75PEG6622G5G6KU26GOSDQPDHSCJ3DQD7VONH4EYVDOGKJ",
				).Return(
					horizon.AccountResponse{
						SequenceNumber: 100,
					},
					nil,
				).Once()

				horizonResponse := horizon.SubmitTransactionResponse{
					nil,
					&horizon.SubmitTransactionResponseError{
						TransactionErrorCode: "transaction_failed",
					},
					&horizon.SubmitTransactionResponseExtras{
						EnvelopeXdr: "envelope",
						ResultXdr:   "result",
					},
				}

				mockHorizon.On(
					"SubmitTransaction",
					mock.AnythingOfType("string"),
				).Return(horizonResponse, nil).Once()

				Convey("it should return error", func() {
					statusCode, response := getResponse(testServer, validParams)
					responseString := strings.TrimSpace(string(response))

					expectedResponse, err := json.MarshalIndent(horizonResponse, "", "  ")
					if err != nil {
						panic(err)
					}

					assert.Equal(t, 400, statusCode)
					assert.Equal(t, string(expectedResponse), responseString)
				})
			})

			Convey("transaction success (native)", func() {
				validParams := url.Values{
					// GCF3WVYTHF75PEG6622G5G6KU26GOSDQPDHSCJ3DQD7VONH4EYVDOGKJ
					"source":      {"SDWLS4G3XCNIYPKXJWWGGJT6UDY63WV6PEFTWP7JZMQB4RE7EUJQN5XM"},
					"destination": {"GDSIKW43UA6JTOA47WVEBCZ4MYC74M3GNKNXTVDXFHXYYTNO5GGVN632"},
					"amount":      {"20"},
				}

				mockHorizon.On(
					"LoadAccount",
					"GCF3WVYTHF75PEG6622G5G6KU26GOSDQPDHSCJ3DQD7VONH4EYVDOGKJ",
				).Return(
					horizon.AccountResponse{
						SequenceNumber: 100,
					},
					nil,
				).Once()

				var ledger uint64
				ledger = 1988727
				horizonResponse := horizon.SubmitTransactionResponse{&ledger, nil, nil}

				mockHorizon.On(
					"SubmitTransaction",
					"AAAAAIu7VxM5f9eQ3va0bpvKprxnSHB4zyEnY4D/VzT8Jio3AAAAZAAAAAAAAABlAAAAAAAAAAAAAAABAAAAAAAAAAEAAAAA5IVbm6A8mbgc/apAizxmBf4zZmqbedR3Ke+MTa7pjVYAAAAAAAAAAAvrwgAAAAAAAAAAAfwmKjcAAABAh3M6y9LXiWD0GB1KCkgNS5H1Lnyr1wS1BsfzoM1/v0muzobwNkJinV+RcWyC8VfeKqOjKBOANJnEusl+sHkcAg==",
				).Return(horizonResponse, nil).Once()

				Convey("it should return success", func() {
					statusCode, response := getResponse(testServer, validParams)
					responseString := strings.TrimSpace(string(response))

					expectedResponse, err := json.MarshalIndent(horizonResponse, "", "  ")
					if err != nil {
						panic(err)
					}

					assert.Equal(t, 200, statusCode)
					assert.Equal(t, string(expectedResponse), responseString)
				})
			})

			Convey("transaction success (credit)", func() {
				mockHorizon.On(
					"LoadAccount",
					"GCF3WVYTHF75PEG6622G5G6KU26GOSDQPDHSCJ3DQD7VONH4EYVDOGKJ",
				).Return(
					horizon.AccountResponse{
						SequenceNumber: 100,
					},
					nil,
				).Once()

				var ledger uint64
				ledger = 1988727
				horizonResponse := horizon.SubmitTransactionResponse{&ledger, nil, nil}

				mockHorizon.On(
					"SubmitTransaction",
					"AAAAAIu7VxM5f9eQ3va0bpvKprxnSHB4zyEnY4D/VzT8Jio3AAAAZAAAAAAAAABlAAAAAAAAAAAAAAABAAAAAAAAAAEAAAAA5IVbm6A8mbgc/apAizxmBf4zZmqbedR3Ke+MTa7pjVYAAAABVVNEAAAAAADkhVuboDyZuBz9qkCLPGYF/jNmapt51Hcp74xNrumNVgAAAAAL68IAAAAAAAAAAAH8Jio3AAAAQHbXpCBe/lDG5rWhwNpdH+DnrkYKONvMyPJDFik5mC/gcIL9xHx3FfB+u1Ik7N9gJxi8AlRRqXo+/yCyOoQQ3Ac=",
				).Return(horizonResponse, nil).Once()

				Convey("it should return success", func() {
					statusCode, response := getResponse(testServer, validParams)
					responseString := strings.TrimSpace(string(response))

					expectedResponse, err := json.MarshalIndent(horizonResponse, "", "  ")
					if err != nil {
						panic(err)
					}

					assert.Equal(t, 200, statusCode)
					assert.Equal(t, string(expectedResponse), responseString)
				})
			})
		})
	})
}
