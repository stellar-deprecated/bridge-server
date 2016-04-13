package handlers

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/facebookgo/inject"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/stellar/gateway/bridge/config"
	"github.com/stellar/gateway/horizon"
	"github.com/stellar/gateway/mocks"
	"github.com/stellar/gateway/net"
	"github.com/stellar/gateway/protocols/compliance"
	"github.com/stellar/gateway/protocols/federation"
	"github.com/stellar/gateway/protocols/stellartoml"
	"github.com/stellar/go-stellar-base/xdr"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRequestHandlerPayment(t *testing.T) {
	c := &config.Config{
		NetworkPassphrase: "Test SDF Network ; September 2015",
		Compliance:        "http://compliance",
	}
	mockHorizon := new(mocks.MockHorizon)
	mockHttpClient := new(mocks.MockHttpClient)
	mockTransactionSubmitter := new(mocks.MockTransactionSubmitter)
	mockFederationResolver := new(mocks.MockFederationResolver)
	mockStellartomlResolver := new(mocks.MockStellartomlResolver)
	requestHandler := RequestHandler{}

	// Inject mocks
	var g inject.Graph

	err := g.Provide(
		&inject.Object{Value: &requestHandler},
		&inject.Object{Value: c},
		&inject.Object{Value: mockHorizon},
		&inject.Object{Value: mockHttpClient},
		&inject.Object{Value: mockTransactionSubmitter},
		&inject.Object{Value: mockFederationResolver},
		&inject.Object{Value: mockStellartomlResolver},
	)
	if err != nil {
		panic(err)
	}

	if err := g.Populate(); err != nil {
		panic(err)
	}

	testServer := httptest.NewServer(http.HandlerFunc(requestHandler.Payment))
	defer testServer.Close()

	Convey("Given payment request", t, func() {
		Convey("When source is invalid", func() {
			source := "SDRAS7XIQNX25UDCCX725R4EYGBFYGJE4HJ2A3DFCWJIHMRSMS7CXX43"

			Convey("it should return error", func() {
				statusCode, response := net.GetResponse(testServer, url.Values{"source": {source}})
				responseString := strings.TrimSpace(string(response))
				assert.Equal(t, 400, statusCode)
				expectedResponse := horizon.SubmitTransactionResponse{Error: horizon.PaymentInvalidSource}
				assert.Equal(t, expectedResponse.Marshal(), []byte(responseString))
			})
		})

		Convey("When destination is invalid", func() {
			source := "SDRAS7XIQNX25UDCCX725R4EYGBFYGJE4HJ2A3DFCWJIHMRSMS7CXX42"
			destination := "GD3YBOYIUVLU"

			mockFederationResolver.On(
				"Resolve",
				"GD3YBOYIUVLU",
			).Return(
				federation.Response{AccountId: "GD3YBOYIUVLU"},
				stellartoml.StellarToml{},
				nil,
			).Once()

			Convey("it should return error", func() {
				statusCode, response := net.GetResponse(testServer, url.Values{"destination": {destination}, "source": {source}})
				responseString := strings.TrimSpace(string(response))
				assert.Equal(t, 400, statusCode)
				expectedResponse := horizon.SubmitTransactionResponse{Error: horizon.PaymentInvalidDestination}
				assert.Equal(t, expectedResponse.Marshal(), []byte(responseString))
			})
		})

		Convey("When destination is a Stellar address", func() {
			params := url.Values{
				"source":      {"SDRAS7XIQNX25UDCCX725R4EYGBFYGJE4HJ2A3DFCWJIHMRSMS7CXX42"},
				"destination": {"bob*stellar.org"},
			}

			Convey("When FederationResolver returns error", func() {
				mockFederationResolver.On(
					"Resolve",
					"bob*stellar.org",
				).Return(
					federation.Response{},
					stellartoml.StellarToml{},
					errors.New("stellar.toml response status code indicates error"),
				).Once()

				Convey("it should return error", func() {
					statusCode, response := net.GetResponse(testServer, params)
					responseString := strings.TrimSpace(string(response))
					assert.Equal(t, 400, statusCode)
					expectedResponse := horizon.SubmitTransactionResponse{Error: horizon.PaymentCannotResolveDestination}
					assert.Equal(t, expectedResponse.Marshal(), []byte(responseString))
				})
			})

			Convey("When federation response is correct (no memo)", func() {
				validParams := url.Values{
					// GCF3WVYTHF75PEG6622G5G6KU26GOSDQPDHSCJ3DQD7VONH4EYVDOGKJ
					"source":      {"SDWLS4G3XCNIYPKXJWWGGJT6UDY63WV6PEFTWP7JZMQB4RE7EUJQN5XM"},
					"destination": {"bob*stellar.org"},
					"amount":      {"20"},
				}

				mockFederationResolver.On(
					"Resolve",
					"bob*stellar.org",
				).Return(
					federation.Response{AccountId: "GDSIKW43UA6JTOA47WVEBCZ4MYC74M3GNKNXTVDXFHXYYTNO5GGVN632"},
					stellartoml.StellarToml{},
					nil,
				).Once()

				// Checking if destination account exists
				mockHorizon.On(
					"LoadAccount",
					"GDSIKW43UA6JTOA47WVEBCZ4MYC74M3GNKNXTVDXFHXYYTNO5GGVN632",
				).Return(horizon.AccountResponse{}, nil).Once()

				// Loading sequence number
				mockHorizon.On(
					"LoadAccount",
					"GCF3WVYTHF75PEG6622G5G6KU26GOSDQPDHSCJ3DQD7VONH4EYVDOGKJ",
				).Return(
					horizon.AccountResponse{
						SequenceNumber: "100",
					},
					nil,
				).Once()

				var ledger uint64
				ledger = 1988728
				horizonResponse := horizon.SubmitTransactionResponse{&ledger, nil, nil}

				mockHorizon.On(
					"SubmitTransaction",
					"AAAAAIu7VxM5f9eQ3va0bpvKprxnSHB4zyEnY4D/VzT8Jio3AAAAZAAAAAAAAABlAAAAAAAAAAAAAAABAAAAAAAAAAEAAAAA5IVbm6A8mbgc/apAizxmBf4zZmqbedR3Ke+MTa7pjVYAAAAAAAAAAAvrwgAAAAAAAAAAAfwmKjcAAABAh3M6y9LXiWD0GB1KCkgNS5H1Lnyr1wS1BsfzoM1/v0muzobwNkJinV+RcWyC8VfeKqOjKBOANJnEusl+sHkcAg==",
				).Return(horizonResponse, nil).Once()

				Convey("it should return success", func() {
					statusCode, response := net.GetResponse(testServer, validParams)
					responseString := strings.TrimSpace(string(response))

					expectedResponse, err := json.MarshalIndent(horizonResponse, "", "  ")
					if err != nil {
						panic(err)
					}

					assert.Equal(t, 200, statusCode)
					assert.Equal(t, string(expectedResponse), responseString)
				})
			})

			Convey("When federation response is correct (with memo)", func() {
				validParams := url.Values{
					// GCF3WVYTHF75PEG6622G5G6KU26GOSDQPDHSCJ3DQD7VONH4EYVDOGKJ
					"source":      {"SDWLS4G3XCNIYPKXJWWGGJT6UDY63WV6PEFTWP7JZMQB4RE7EUJQN5XM"},
					"destination": {"bob*stellar.org"},
					"amount":      {"20"},
				}

				mockFederationResolver.On(
					"Resolve",
					"bob*stellar.org",
				).Return(
					federation.Response{
						AccountId: "GDSIKW43UA6JTOA47WVEBCZ4MYC74M3GNKNXTVDXFHXYYTNO5GGVN632",
						MemoType:  "text",
						Memo:      "125",
					},
					stellartoml.StellarToml{},
					nil,
				).Once()

				// Checking if destination account exists
				mockHorizon.On(
					"LoadAccount",
					"GDSIKW43UA6JTOA47WVEBCZ4MYC74M3GNKNXTVDXFHXYYTNO5GGVN632",
				).Return(horizon.AccountResponse{}, nil).Once()

				// Loading sequence number
				mockHorizon.On(
					"LoadAccount",
					"GCF3WVYTHF75PEG6622G5G6KU26GOSDQPDHSCJ3DQD7VONH4EYVDOGKJ",
				).Return(
					horizon.AccountResponse{
						SequenceNumber: "100",
					},
					nil,
				).Once()

				var ledger uint64
				ledger = 1988728
				horizonResponse := horizon.SubmitTransactionResponse{&ledger, nil, nil}

				mockHorizon.On(
					"SubmitTransaction",
					"AAAAAIu7VxM5f9eQ3va0bpvKprxnSHB4zyEnY4D/VzT8Jio3AAAAZAAAAAAAAABlAAAAAAAAAAEAAAADMTI1AAAAAAEAAAAAAAAAAQAAAADkhVuboDyZuBz9qkCLPGYF/jNmapt51Hcp74xNrumNVgAAAAAAAAAAC+vCAAAAAAAAAAAB/CYqNwAAAEAjnc8Wf31VxgBXXhEmZfLo6c4YJtROVy5MTLsWFSx7TCkoQzCskBVcC30DrjQq7Vzm0zwg+mBmSGI5wFbctKgB",
				).Return(horizonResponse, nil).Once()

				Convey("it should return success", func() {
					statusCode, response := net.GetResponse(testServer, validParams)
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

		Convey("When assetIssuer is invalid", func() {
			source := "SDRAS7XIQNX25UDCCX725R4EYGBFYGJE4HJ2A3DFCWJIHMRSMS7CXX42"
			destination := "GDSIKW43UA6JTOA47WVEBCZ4MYC74M3GNKNXTVDXFHXYYTNO5GGVN632"
			assetCode := "USD"
			assetIssuer := "GDSIKW43UA6JTOA47WVEBCZ4MYC74M3GNKNXTVDXFHXYYTNO5GGVN631"

			mockFederationResolver.On(
				"Resolve",
				"GDSIKW43UA6JTOA47WVEBCZ4MYC74M3GNKNXTVDXFHXYYTNO5GGVN632",
			).Return(
				federation.Response{AccountId: "GDSIKW43UA6JTOA47WVEBCZ4MYC74M3GNKNXTVDXFHXYYTNO5GGVN632"},
				stellartoml.StellarToml{},
				nil,
			).Once()

			Convey("it should return error", func() {
				statusCode, response := net.GetResponse(
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
				expectedResponse := horizon.SubmitTransactionResponse{Error: horizon.PaymentInvalidIssuer}
				assert.Equal(t, expectedResponse.Marshal(), []byte(responseString))
			})
		})

		Convey("When assetCode is invalid", func() {
			// GBKGH7QZVCZ2ZA5OUGZSTHFNXTBHL3MPCKSCBJUAQODGPMWP7OMMRKDW
			source := "SDRAS7XIQNX25UDCCX725R4EYGBFYGJE4HJ2A3DFCWJIHMRSMS7CXX42"
			destination := "GDSIKW43UA6JTOA47WVEBCZ4MYC74M3GNKNXTVDXFHXYYTNO5GGVN632"
			amount := "20"
			assetCode := "1234567890123"
			assetIssuer := "GDSIKW43UA6JTOA47WVEBCZ4MYC74M3GNKNXTVDXFHXYYTNO5GGVN632"

			mockFederationResolver.On(
				"Resolve",
				"GDSIKW43UA6JTOA47WVEBCZ4MYC74M3GNKNXTVDXFHXYYTNO5GGVN632",
			).Return(
				federation.Response{AccountId: "GDSIKW43UA6JTOA47WVEBCZ4MYC74M3GNKNXTVDXFHXYYTNO5GGVN632"},
				stellartoml.StellarToml{},
				nil,
			).Once()

			Convey("it should return error", func() {
				mockHorizon.On(
					"LoadAccount",
					"GBKGH7QZVCZ2ZA5OUGZSTHFNXTBHL3MPCKSCBJUAQODGPMWP7OMMRKDW",
				).Return(
					horizon.AccountResponse{
						SequenceNumber: "100",
					},
					nil,
				).Once()

				statusCode, response := net.GetResponse(
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
				expectedResponse := horizon.SubmitTransactionResponse{Error: horizon.PaymentMalformedAssetCode}
				assert.Equal(t, expectedResponse.Marshal(), []byte(responseString))
			})
		})

		Convey("When amount is invalid", func() {
			// GBKGH7QZVCZ2ZA5OUGZSTHFNXTBHL3MPCKSCBJUAQODGPMWP7OMMRKDW
			source := "SDRAS7XIQNX25UDCCX725R4EYGBFYGJE4HJ2A3DFCWJIHMRSMS7CXX42"
			destination := "GDSIKW43UA6JTOA47WVEBCZ4MYC74M3GNKNXTVDXFHXYYTNO5GGVN632"
			amount := "test"
			assetCode := "USD"
			assetIssuer := "GDSIKW43UA6JTOA47WVEBCZ4MYC74M3GNKNXTVDXFHXYYTNO5GGVN632"

			mockFederationResolver.On(
				"Resolve",
				"GDSIKW43UA6JTOA47WVEBCZ4MYC74M3GNKNXTVDXFHXYYTNO5GGVN632",
			).Return(
				federation.Response{AccountId: "GDSIKW43UA6JTOA47WVEBCZ4MYC74M3GNKNXTVDXFHXYYTNO5GGVN632"},
				stellartoml.StellarToml{},
				nil,
			).Once()

			mockHorizon.On(
				"LoadAccount",
				"GBKGH7QZVCZ2ZA5OUGZSTHFNXTBHL3MPCKSCBJUAQODGPMWP7OMMRKDW",
			).Return(
				horizon.AccountResponse{
					SequenceNumber: "100",
				},
				nil,
			).Once()

			Convey("it should return error", func() {
				statusCode, response := net.GetResponse(
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
				expectedResponse := horizon.SubmitTransactionResponse{Error: horizon.PaymentInvalidAmount}
				assert.Equal(t, expectedResponse.Marshal(), []byte(responseString))
			})
		})

		Convey("When params are valid (payment operation)", func() {
			validParams := url.Values{
				// GCF3WVYTHF75PEG6622G5G6KU26GOSDQPDHSCJ3DQD7VONH4EYVDOGKJ
				"source":       {"SDWLS4G3XCNIYPKXJWWGGJT6UDY63WV6PEFTWP7JZMQB4RE7EUJQN5XM"},
				"destination":  {"GDSIKW43UA6JTOA47WVEBCZ4MYC74M3GNKNXTVDXFHXYYTNO5GGVN632"},
				"amount":       {"20"},
				"asset_code":   {"USD"},
				"asset_issuer": {"GDSIKW43UA6JTOA47WVEBCZ4MYC74M3GNKNXTVDXFHXYYTNO5GGVN632"},
			}

			mockFederationResolver.On(
				"Resolve",
				"GDSIKW43UA6JTOA47WVEBCZ4MYC74M3GNKNXTVDXFHXYYTNO5GGVN632",
			).Return(
				federation.Response{AccountId: "GDSIKW43UA6JTOA47WVEBCZ4MYC74M3GNKNXTVDXFHXYYTNO5GGVN632"},
				stellartoml.StellarToml{},
				nil,
			).Once()

			Convey("When memo is set", func() {
				Convey("only `memo` param is set", func() {
					validParams.Add("memo", "test")
					statusCode, response := net.GetResponse(testServer, validParams)
					responseString := strings.TrimSpace(string(response))
					assert.Equal(t, 400, statusCode)
					expectedResponse := horizon.SubmitTransactionResponse{Error: horizon.PaymentMissingParamMemo}
					assert.Equal(t, expectedResponse.Marshal(), []byte(responseString))
				})

				Convey("only `memo_type` param is set", func() {
					validParams.Add("memo_type", "id")
					statusCode, response := net.GetResponse(testServer, validParams)
					responseString := strings.TrimSpace(string(response))
					assert.Equal(t, 400, statusCode)
					expectedResponse := horizon.SubmitTransactionResponse{Error: horizon.PaymentMissingParamMemo}
					assert.Equal(t, expectedResponse.Marshal(), []byte(responseString))
				})

				Convey("memo_type=hash to long", func() {
					validParams.Add("memo_type", "hash")
					validParams.Add("memo", "012345678901234567890123456789012345678901234567890123456789012")
					statusCode, response := net.GetResponse(testServer, validParams)
					responseString := strings.TrimSpace(string(response))
					assert.Equal(t, 400, statusCode)
					expectedResponse := horizon.SubmitTransactionResponse{Error: horizon.PaymentInvalidMemo}
					assert.Equal(t, expectedResponse.Marshal(), []byte(responseString))
				})

				Convey("unsupported memo_type", func() {
					validParams.Add("memo_type", "return_hash")
					validParams.Add("memo", "0123456789")
					statusCode, response := net.GetResponse(testServer, validParams)
					responseString := strings.TrimSpace(string(response))
					assert.Equal(t, 400, statusCode)
					expectedResponse := horizon.SubmitTransactionResponse{Error: horizon.PaymentInvalidMemo}
					assert.Equal(t, expectedResponse.Marshal(), []byte(responseString))
				})

				Convey("memo is attached to the transaction", func() {
					mockHorizon.On(
						"LoadAccount",
						"GCF3WVYTHF75PEG6622G5G6KU26GOSDQPDHSCJ3DQD7VONH4EYVDOGKJ",
					).Return(
						horizon.AccountResponse{
							SequenceNumber: "100",
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
					statusCode, response := net.GetResponse(testServer, validParams)
					responseString := strings.TrimSpace(string(response))

					expectedResponse, err := json.MarshalIndent(horizonResponse, "", "  ")
					if err != nil {
						panic(err)
					}

					assert.Equal(t, 200, statusCode)
					assert.Equal(t, string(expectedResponse), responseString)
				})

				Convey("memo hash is attached to the transaction", func() {
					mockHorizon.On(
						"LoadAccount",
						"GCF3WVYTHF75PEG6622G5G6KU26GOSDQPDHSCJ3DQD7VONH4EYVDOGKJ",
					).Return(
						horizon.AccountResponse{
							SequenceNumber: "100",
						},
						nil,
					).Once()

					var ledger uint64
					ledger = 1988727
					horizonResponse := horizon.SubmitTransactionResponse{&ledger, nil, nil}

					mockHorizon.On(
						"SubmitTransaction",
						"AAAAAIu7VxM5f9eQ3va0bpvKprxnSHB4zyEnY4D/VzT8Jio3AAAAZAAAAAAAAABlAAAAAAAAAAMCADrUIHRM3rjlJN62XzjLUJXTDQAAAAAAAAAAAAAAAAAAAAEAAAAAAAAAAQAAAADkhVuboDyZuBz9qkCLPGYF/jNmapt51Hcp74xNrumNVgAAAAFVU0QAAAAAAOSFW5ugPJm4HP2qQIs8ZgX+M2Zqm3nUdynvjE2u6Y1WAAAAAAvrwgAAAAAAAAAAAfwmKjcAAABAEV6Lzok4i4C1jJA3PVVARGx2+yfVw8odprnnnG0hqkUUwKnvVQcd59UJwbfzTG7oxR5DvxflV4aQ6RmZsIcmDQ==",
					).Return(horizonResponse, nil).Once()

					validParams.Add("memo_type", "hash")
					validParams.Add("memo", "02003AD420744CDEB8E524DEB65F38CB5095D30D000000000000000000000000")
					statusCode, response := net.GetResponse(testServer, validParams)
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
					statusCode, response := net.GetResponse(testServer, validParams)
					responseString := strings.TrimSpace(string(response))
					assert.Equal(t, 400, statusCode)
					expectedResponse := horizon.SubmitTransactionResponse{Error: horizon.PaymentSourceNotExist}
					assert.Equal(t, expectedResponse.Marshal(), []byte(responseString))
				})
			})

			Convey("transaction failed in horizon", func() {
				mockHorizon.On(
					"LoadAccount",
					"GCF3WVYTHF75PEG6622G5G6KU26GOSDQPDHSCJ3DQD7VONH4EYVDOGKJ",
				).Return(
					horizon.AccountResponse{
						SequenceNumber: "100",
					},
					nil,
				).Once()

				horizonResponse := horizon.SubmitTransactionResponse{
					nil,
					&horizon.SubmitTransactionResponseError{
						Status: 400,
						Code:   "transaction_failed",
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
					statusCode, response := net.GetResponse(testServer, validParams)
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

				// Checking if destination exists
				mockHorizon.On(
					"LoadAccount",
					"GDSIKW43UA6JTOA47WVEBCZ4MYC74M3GNKNXTVDXFHXYYTNO5GGVN632",
				).Return(horizon.AccountResponse{}, nil).Once()

				// Loading sequence number
				mockHorizon.On(
					"LoadAccount",
					"GCF3WVYTHF75PEG6622G5G6KU26GOSDQPDHSCJ3DQD7VONH4EYVDOGKJ",
				).Return(
					horizon.AccountResponse{
						SequenceNumber: "100",
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
					statusCode, response := net.GetResponse(testServer, validParams)
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
						SequenceNumber: "100",
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
					statusCode, response := net.GetResponse(testServer, validParams)
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

		Convey("When params are valid (path payment operation)", func() {
			validParams := url.Values{
				// GCF3WVYTHF75PEG6622G5G6KU26GOSDQPDHSCJ3DQD7VONH4EYVDOGKJ
				"source":       {"SDWLS4G3XCNIYPKXJWWGGJT6UDY63WV6PEFTWP7JZMQB4RE7EUJQN5XM"},
				"destination":  {"GDSIKW43UA6JTOA47WVEBCZ4MYC74M3GNKNXTVDXFHXYYTNO5GGVN632"},
				"amount":       {"20"},
				"asset_code":   {"USD"},
				"asset_issuer": {"GDSIKW43UA6JTOA47WVEBCZ4MYC74M3GNKNXTVDXFHXYYTNO5GGVN632"},
				"send_max":     {"100"},
			}

			// Source
			mockHorizon.On(
				"LoadAccount",
				"GCF3WVYTHF75PEG6622G5G6KU26GOSDQPDHSCJ3DQD7VONH4EYVDOGKJ",
			).Return(
				horizon.AccountResponse{
					SequenceNumber: "100",
				},
				nil,
			).Once()

			// Destination
			mockFederationResolver.On(
				"Resolve",
				"GDSIKW43UA6JTOA47WVEBCZ4MYC74M3GNKNXTVDXFHXYYTNO5GGVN632",
			).Return(
				federation.Response{AccountId: "GDSIKW43UA6JTOA47WVEBCZ4MYC74M3GNKNXTVDXFHXYYTNO5GGVN632"},
				stellartoml.StellarToml{},
				nil,
			).Once()

			mockHorizon.On(
				"LoadAccount",
				"GDSIKW43UA6JTOA47WVEBCZ4MYC74M3GNKNXTVDXFHXYYTNO5GGVN632",
			).Return(
				horizon.AccountResponse{
					SequenceNumber: "100",
				},
				nil,
			).Once()

			Convey("transaction success (send native)", func() {
				var ledger uint64
				ledger = 1988727
				horizonResponse := horizon.SubmitTransactionResponse{&ledger, nil, nil}

				mockHorizon.On(
					"SubmitTransaction",
					"AAAAAIu7VxM5f9eQ3va0bpvKprxnSHB4zyEnY4D/VzT8Jio3AAAAZAAAAAAAAABlAAAAAAAAAAAAAAABAAAAAAAAAAIAAAAAAAAAADuaygAAAAAA5IVbm6A8mbgc/apAizxmBf4zZmqbedR3Ke+MTa7pjVYAAAABVVNEAAAAAADkhVuboDyZuBz9qkCLPGYF/jNmapt51Hcp74xNrumNVgAAAAAL68IAAAAAAAAAAAAAAAAB/CYqNwAAAECx9I76SOIjCL3pgwZdmFj9KWFzvNL82dt3+Laokh6Zm2v0o1UNq1mQsrqrv0mXM6uNcA96NbkfbogtXauHmwML",
				).Return(horizonResponse, nil).Once()

				Convey("it should return success", func() {
					statusCode, response := net.GetResponse(testServer, validParams)
					responseString := strings.TrimSpace(string(response))

					expectedResponse, err := json.MarshalIndent(horizonResponse, "", "  ")
					if err != nil {
						panic(err)
					}

					assert.Equal(t, 200, statusCode)
					assert.Equal(t, string(expectedResponse), responseString)
				})
			})

			Convey("transaction success (send credit)", func() {
				validParams["send_asset_code"] = []string{"USD"}
				validParams["send_asset_issuer"] = []string{"GBDOSO3K4JTGSWJSIHXAOFIBMAABVM3YK3FI6VJPKIHHM56XAFIUCGD6"}

				var ledger uint64
				ledger = 1988727
				horizonResponse := horizon.SubmitTransactionResponse{&ledger, nil, nil}

				mockHorizon.On(
					"SubmitTransaction",
					"AAAAAIu7VxM5f9eQ3va0bpvKprxnSHB4zyEnY4D/VzT8Jio3AAAAZAAAAAAAAABlAAAAAAAAAAAAAAABAAAAAAAAAAIAAAABVVNEAAAAAABG6Ttq4mZpWTJB7gcVAWAAGrN4VsqPVS9SDnZ31wFRQQAAAAA7msoAAAAAAOSFW5ugPJm4HP2qQIs8ZgX+M2Zqm3nUdynvjE2u6Y1WAAAAAVVTRAAAAAAA5IVbm6A8mbgc/apAizxmBf4zZmqbedR3Ke+MTa7pjVYAAAAAC+vCAAAAAAAAAAAAAAAAAfwmKjcAAABA3rQOu+r9DvUhGDSOVaD05RWgzvzMJt49opYNfGLLOSo7/29rUkPIyw5PgV/1arrTwj90HRnzmVjHJK2xy+MfBQ==",
				).Return(horizonResponse, nil).Once()

				Convey("it should return success", func() {
					statusCode, response := net.GetResponse(testServer, validParams)
					responseString := strings.TrimSpace(string(response))

					expectedResponse, err := json.MarshalIndent(horizonResponse, "", "  ")
					if err != nil {
						panic(err)
					}

					assert.Equal(t, 200, statusCode)
					assert.Equal(t, string(expectedResponse), responseString)
				})
			})

			Convey("transaction success (path)", func() {
				validParams["send_asset_code"] = []string{"USD"}
				validParams["send_asset_issuer"] = []string{"GBDOSO3K4JTGSWJSIHXAOFIBMAABVM3YK3FI6VJPKIHHM56XAFIUCGD6"}

				// Native
				validParams["path[0][asset_code]"] = []string{""}
				validParams["path[0][asset_issuer]"] = []string{""}
				// Credit
				validParams["path[1][asset_code]"] = []string{"EUR"}
				validParams["path[1][asset_issuer]"] = []string{"GAF3PBFQLH57KPECN4GRGHU5NUZ3XXKYYWLOTBIRJMBYHPUBWANIUCZU"}

				var ledger uint64
				ledger = 1988727
				horizonResponse := horizon.SubmitTransactionResponse{&ledger, nil, nil}

				mockHorizon.On(
					"SubmitTransaction",
					"AAAAAIu7VxM5f9eQ3va0bpvKprxnSHB4zyEnY4D/VzT8Jio3AAAAZAAAAAAAAABlAAAAAAAAAAAAAAABAAAAAAAAAAIAAAABVVNEAAAAAABG6Ttq4mZpWTJB7gcVAWAAGrN4VsqPVS9SDnZ31wFRQQAAAAA7msoAAAAAAOSFW5ugPJm4HP2qQIs8ZgX+M2Zqm3nUdynvjE2u6Y1WAAAAAVVTRAAAAAAA5IVbm6A8mbgc/apAizxmBf4zZmqbedR3Ke+MTa7pjVYAAAAAC+vCAAAAAAIAAAAAAAAAAUVVUgAAAAAAC7eEsFn79TyCbw0THp1tM7vdWMWW6YURSwODvoGwGooAAAAAAAAAAfwmKjcAAABAyO0YxnfaIdY51J9BaPyZYNxsBY2AhWCZpK6FRlaE+ZbdmznZ9cio2G7+fJgl3hWZUrQknQHElmzAZdgsqNnZAQ==",
				).Return(horizonResponse, nil).Once()

				Convey("it should return success", func() {
					statusCode, response := net.GetResponse(testServer, validParams)
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

	Convey("Given payment compliance request", t, func() {
		Convey("When params are valid", func() {
			params := url.Values{
				// GAW77Z6GPWXSODJOMF5L5BMX6VMYGEJRKUNBC2CZ725JTQZORK74HQQD
				"source":       {"SARMR3N465GTEHQLR3TSHDD7FHFC2I22ECFLYCHAZDEJWBVED66RW7FQ"},
				"sender":       {"alice*stellar.org"}, // GAW77Z6GPWXSODJOMF5L5BMX6VMYGEJRKUNBC2CZ725JTQZORK74HQQD
				"destination":  {"bob*stellar.org"},   // GAMVF7G4GJC4A7JMFJWLUAEIBFQD5RT3DCB5DC5TJDEKQBBACQ4JZVEE
				"amount":       {"20"},
				"asset_code":   {"USD"},
				"asset_issuer": {"GAMVF7G4GJC4A7JMFJWLUAEIBFQD5RT3DCB5DC5TJDEKQBBACQ4JZVEE"},
				"extra_memo":   {"hello world"},
			}

			Convey("it should return error when compliance server returns error", func() {
				mockHttpClient.On(
					"PostForm",
					"http://compliance/send",
					mock.AnythingOfType("url.Values"),
				).Return(
					net.BuildHttpResponse(400, "error"),
					nil,
				).Run(func(args mock.Arguments) {
					values := args.Get(1).(url.Values)
					// bridge server does not send source seed to compliance
					assert.Equal(t, []string{"GAW77Z6GPWXSODJOMF5L5BMX6VMYGEJRKUNBC2CZ725JTQZORK74HQQD"}, values["source"])
					values.Del("source")
					params.Del("source")
					assert.Equal(t, values.Encode(), params.Encode())
				}).Once()

				statusCode, response := net.GetResponse(testServer, params)
				responseString := strings.TrimSpace(string(response))
				assert.Equal(t, 500, statusCode)
				expectedResponse := horizon.SubmitTransactionResponse{Error: horizon.ServerError}
				assert.Equal(t, expectedResponse.Marshal(), []byte(responseString))
			})

			Convey("it should return denied when compliance server returns denied", func() {
				mockHttpClient.On(
					"PostForm",
					"http://compliance/send",
					mock.AnythingOfType("url.Values"),
				).Return(
					net.BuildHttpResponse(200, "{\"auth_response\": {\"tx_status\": \"denied\"}}"),
					nil,
				).Run(func(args mock.Arguments) {
					values := args.Get(1).(url.Values)
					// bridge server does not send source seed to compliance
					assert.Equal(t, []string{"GAW77Z6GPWXSODJOMF5L5BMX6VMYGEJRKUNBC2CZ725JTQZORK74HQQD"}, values["source"])
					values.Del("source")
					params.Del("source")
					assert.Equal(t, values.Encode(), params.Encode())
				}).Once()

				statusCode, response := net.GetResponse(testServer, params)
				responseString := strings.TrimSpace(string(response))
				assert.Equal(t, 403, statusCode)
				expectedResponse := horizon.NewErrorResponse(horizon.PaymentDenied)
				assert.Equal(t, expectedResponse.Marshal(), []byte(responseString))
			})

			Convey("it should return pending when compliance server returns pending", func() {
				mockHttpClient.On(
					"PostForm",
					"http://compliance/send",
					mock.AnythingOfType("url.Values"),
				).Return(
					net.BuildHttpResponse(200, "{\"auth_response\": {\"info_status\": \"pending\", \"pending\": 3600}}"),
					nil,
				).Run(func(args mock.Arguments) {
					values := args.Get(1).(url.Values)
					// bridge server does not send source seed to compliance
					assert.Equal(t, []string{"GAW77Z6GPWXSODJOMF5L5BMX6VMYGEJRKUNBC2CZ725JTQZORK74HQQD"}, values["source"])
					values.Del("source")
					params.Del("source")
					assert.Equal(t, values.Encode(), params.Encode())
				}).Once()

				statusCode, response := net.GetResponse(testServer, params)
				responseString := strings.TrimSpace(string(response))
				assert.Equal(t, 202, statusCode)
				expectedResponse := horizon.NewErrorResponse(
					horizon.NewPaymentPendingError(3600),
				)
				assert.Equal(t, expectedResponse.Marshal(), []byte(responseString))
			})

			Convey("it should submit transaction when compliance server returns success", func() {
				memoBytes, _ := hex.DecodeString("b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9")
				var hashXdr xdr.Hash
				copy(hashXdr[:], memoBytes[:])
				memo, _ := xdr.NewMemo(xdr.MemoTypeMemoHash, hashXdr)

				sourceBytes, _ := hex.DecodeString("2dffe7c67daf270d2e617abe8597f559831131551a116859feba99c32e8abfc3")
				var sourceXdr xdr.Uint256
				copy(sourceXdr[:], sourceBytes[:])

				destinationBytes, _ := hex.DecodeString("1952fcdc3245c07d2c2a6cba008809603ec67b1883d18bb348c8a8042014389c")
				var destinationXdr xdr.Uint256
				copy(destinationXdr[:], destinationBytes[:])

				issuerBytes, _ := hex.DecodeString("1952fcdc3245c07d2c2a6cba008809603ec67b1883d18bb348c8a8042014389c")
				var issuerXdr xdr.Uint256
				copy(issuerXdr[:], issuerBytes[:])

				expectedTx := &xdr.Transaction{
					SourceAccount: xdr.AccountId{
						Type:    xdr.CryptoKeyTypeKeyTypeEd25519,
						Ed25519: &sourceXdr,
					},
					Fee:    100,
					SeqNum: 0,
					Memo:   memo,
					Operations: []xdr.Operation{
						xdr.Operation{
							Body: xdr.OperationBody{
								Type: xdr.OperationTypePayment,
								PaymentOp: &xdr.PaymentOp{
									Destination: xdr.AccountId{
										Type:    xdr.CryptoKeyTypeKeyTypeEd25519,
										Ed25519: &destinationXdr,
									},
									Amount: 200000000,
									Asset: xdr.Asset{
										Type: xdr.AssetTypeAssetTypeCreditAlphanum4,
										AlphaNum4: &xdr.AssetAlphaNum4{
											AssetCode: [4]byte{'U', 'S', 'D', 0},
											Issuer: xdr.AccountId{
												Type:    xdr.CryptoKeyTypeKeyTypeEd25519,
												Ed25519: &issuerXdr,
											},
										},
									},
								},
							},
						},
					},
				}

				complianceResponse := compliance.SendResponse{
					TransactionXdr: "AAAAAC3/58Z9rycNLmF6voWX9VmDETFVGhFoWf66mcMuir/DAAAAZAAAAAAAAAAAAAAAAAAAAAO5TSe5k00+CKUuUtfafav6xITv43pTgO6QiPes4u/N6QAAAAEAAAAAAAAAAQAAAAAZUvzcMkXAfSwqbLoAiAlgPsZ7GIPRi7NIyKgEIBQ4nAAAAAFVU0QAAAAAABlS/NwyRcB9LCpsugCICWA+xnsYg9GLs0jIqAQgFDicAAAAAAvrwgAAAAAA",
				}

				mockHttpClient.On(
					"PostForm",
					"http://compliance/send",
					mock.AnythingOfType("url.Values"),
				).Return(
					net.BuildHttpResponse(200, string(complianceResponse.Marshal())),
					nil,
				).Run(func(args mock.Arguments) {
					values := args.Get(1).(url.Values)
					assert.Equal(t, []string{"GAW77Z6GPWXSODJOMF5L5BMX6VMYGEJRKUNBC2CZ725JTQZORK74HQQD"}, values["source"])
					values.Del("source")
					params.Del("source")
					assert.Equal(t, values.Encode(), params.Encode())
				}).Once()

				var ledger uint64
				ledger = 1988727
				horizonResponse := horizon.SubmitTransactionResponse{&ledger, nil, nil}

				mockTransactionSubmitter.On(
					"SignAndSubmitRawTransaction",
					params.Get("source"),
					mock.AnythingOfType("*xdr.Transaction"),
				).Run(func(args mock.Arguments) {
					tx := args.Get(1).(*xdr.Transaction)
					assert.Equal(t, *tx, *expectedTx)
				}).Return(horizonResponse, nil).Once()

				statusCode, response := net.GetResponse(testServer, params)
				responseString := strings.TrimSpace(string(response))
				assert.Equal(t, 200, statusCode)
				assert.Equal(t, horizonResponse.Marshal(), []byte(responseString))
			})
		})
	})
}
