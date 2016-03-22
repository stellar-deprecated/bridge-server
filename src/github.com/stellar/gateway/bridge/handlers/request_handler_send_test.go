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
	"github.com/stellar/gateway/bridge/config"
	"github.com/stellar/gateway/horizon"
	"github.com/stellar/gateway/mocks"
	b "github.com/stellar/go-stellar-base/build"
	"github.com/stellar/go-stellar-base/keypair"
	"github.com/stretchr/testify/assert"
)

func TestRequestHandlerSend(t *testing.T) {
	mockAddressResolverHelper := new(MockAddressResolverHelper)
	addressResolver := AddressResolver{mockAddressResolverHelper}

	mockTransactionSubmitter := new(mocks.MockTransactionSubmitter)

	IssuingSeed := "SC34WILLHVADXMP6ACPMIRA6TRAWJMVCLPFNW7S6MUMXJVLAZUC4EWHP"
	AuthorizingSeed := "SC37TBSIAYKIDQ6GTGLT2HSORLIHZQHBXVFI5P5K4Q5TSHRTRBK3UNWG"

	config := config.Config{
		Assets: []string{"USD", "EUR"},
		Accounts: &config.Accounts{
			// GD4I7AFSLZGTDL34TQLWJOM2NHLIIOEKD5RHHZUW54HERBLSIRKUOXRR
			IssuingSeed: &IssuingSeed,
			// GBQXA3ABGQGTCLEVZIUTDRWWJOQD5LSAEDZAG7GMOGD2HBLWONGUVO4I
			AuthorizingSeed: &AuthorizingSeed,
		},
	}

	issuingKeypair, err := keypair.Parse(*config.Accounts.IssuingSeed)
	if err != nil {
		panic(err)
	}

	requestHandler := RequestHandler{
		AddressResolver:      addressResolver,
		Config:               &config,
		TransactionSubmitter: mockTransactionSubmitter,
	}
	testServer := httptest.NewServer(http.HandlerFunc(requestHandler.Send))
	defer testServer.Close()

	Convey("Given send request", t, func() {
		Convey("When destination is invalid", func() {
			destination := "GD3YBOYIUVLU"
			assetCode := "USD"

			Convey("it should return error", func() {
				statusCode, response := getResponse(testServer, url.Values{"destination": {destination}, "asset_code": {assetCode}})
				responseString := strings.TrimSpace(string(response))
				assert.Equal(t, 400, statusCode)
				assert.Equal(t, getResponseString(horizon.PaymentInvalidDestination), responseString)
			})
		})

		Convey("When assetCode is invalid", func() {
			destination := "GDSIKW43UA6JTOA47WVEBCZ4MYC74M3GNKNXTVDXFHXYYTNO5GGVN632"
			assetCode := "GBP"

			Convey("it should return error", func() {
				statusCode, response := getResponse(testServer, url.Values{"destination": {destination}, "asset_code": {assetCode}})
				responseString := strings.TrimSpace(string(response))
				assert.Equal(t, 400, statusCode)
				assert.Equal(t, getResponseString(horizon.PaymentAssetCodeNotAllowed), responseString)
			})
		})

		Convey("When destination is a Stellar address", func() {
			params := url.Values{
				"asset_code":  {"USD"},
				"amount":      {"20"},
				"destination": {"bob*stellar.org"},
			}

			Convey("When stellar.toml does not exist", func() {
				mockAddressResolverHelper.On(
					"GetStellarToml",
					"stellar.org",
				).Return(
					StellarToml{},
					errors.New("stellar.toml response status code indicates error"),
				).Once()

				Convey("it should return error", func() {
					statusCode, response := getResponse(testServer, params)
					responseString := strings.TrimSpace(string(response))
					assert.Equal(t, 400, statusCode)
					assert.Equal(t, getResponseString(horizon.PaymentCannotResolveDestination), responseString)
				})
			})

			Convey("When stellar.toml does not contain FEDERATION_SERVER", func() {
				mockAddressResolverHelper.On(
					"GetStellarToml",
					"stellar.org",
				).Return(
					StellarToml{},
					nil,
				).Once()

				Convey("it should return error", func() {
					statusCode, response := getResponse(testServer, params)
					responseString := strings.TrimSpace(string(response))
					assert.Equal(t, 400, statusCode)
					assert.Equal(t, getResponseString(horizon.PaymentCannotResolveDestination), responseString)
				})
			})

			Convey("When GetDestination() errors", func() {
				federationServer := "http://api.example.com"
				mockAddressResolverHelper.On(
					"GetStellarToml",
					"stellar.org",
				).Return(
					StellarToml{&federationServer},
					nil,
				).Once()

				mockAddressResolverHelper.On(
					"GetDestination",
					"http://api.example.com",
					"bob*stellar.org",
				).Return(
					StellarDestination{},
					errors.New("Only HTTPS federation servers allowed"),
				).Once()

				Convey("it should return error", func() {
					statusCode, response := getResponse(testServer, params)
					responseString := strings.TrimSpace(string(response))
					assert.Equal(t, 400, statusCode)
					assert.Equal(t, getResponseString(horizon.PaymentCannotResolveDestination), responseString)
				})
			})

			Convey("When federation response is correct", func() {
				federationServer := "http://api.example.com"
				mockAddressResolverHelper.On(
					"GetStellarToml",
					"stellar.org",
				).Return(
					StellarToml{&federationServer},
					nil,
				).Once()

				mockAddressResolverHelper.On(
					"GetDestination",
					"http://api.example.com",
					"bob*stellar.org",
				).Return(StellarDestination{AccountId: "GDSIKW43UA6JTOA47WVEBCZ4MYC74M3GNKNXTVDXFHXYYTNO5GGVN632"}, nil).Once()

				var ledger uint64
				ledger = 1988728
				expectedSubmitResponse := horizon.SubmitTransactionResponse{&ledger, nil, nil}

				operation := b.Payment(
					b.Destination{"GDSIKW43UA6JTOA47WVEBCZ4MYC74M3GNKNXTVDXFHXYYTNO5GGVN632"},
					b.CreditAmount{
						params.Get("asset_code"),
						issuingKeypair.Address(),
						params.Get("amount"),
					},
				)

				mockTransactionSubmitter.On(
					"SubmitTransaction",
					*config.Accounts.IssuingSeed,
					operation,
					nil,
				).Return(expectedSubmitResponse, nil).Once()

				Convey("it should return success", func() {
					statusCode, response := getResponse(testServer, params)
					responseString := strings.TrimSpace(string(response))

					expectedResponse, err := json.MarshalIndent(expectedSubmitResponse, "", "  ")
					if err != nil {
						panic(err)
					}

					assert.Equal(t, 200, statusCode)
					assert.Equal(t, string(expectedResponse), responseString)
				})
			})
		})

		Convey("When destination is an accountId", func() {
			params := url.Values{
				"asset_code":  {"USD"},
				"amount":      {"20"},
				"destination": {"GDSIKW43UA6JTOA47WVEBCZ4MYC74M3GNKNXTVDXFHXYYTNO5GGVN632"},
			}

			Convey("When params are valid", func() {
				operation := b.Payment(
					b.Destination{params.Get("destination")},
					b.CreditAmount{
						params.Get("asset_code"),
						issuingKeypair.Address(),
						params.Get("amount"),
					},
				)

				Convey("transaction fails", func() {
					mockTransactionSubmitter.On(
						"SubmitTransaction",
						*config.Accounts.IssuingSeed,
						operation,
						nil,
					).Return(
						horizon.SubmitTransactionResponse{},
						errors.New("Error sending transaction"),
					).Once()

					Convey("it should return server error", func() {
						statusCode, response := getResponse(testServer, params)
						responseString := strings.TrimSpace(string(response))
						assert.Equal(t, 500, statusCode)
						assert.Equal(t, getResponseString(horizon.ServerError), responseString)
						mockTransactionSubmitter.AssertExpectations(t)
					})
				})

				Convey("transaction succeeds (no memo)", func() {
					var ledger uint64
					ledger = 100
					expectedSubmitResponse := horizon.SubmitTransactionResponse{
						Ledger: &ledger,
					}

					mockTransactionSubmitter.On(
						"SubmitTransaction",
						*config.Accounts.IssuingSeed,
						operation,
						nil,
					).Return(expectedSubmitResponse, nil).Once()

					Convey("it should succeed", func() {
						statusCode, response := getResponse(testServer, params)
						var actualSubmitTransactionResponse horizon.SubmitTransactionResponse
						json.Unmarshal(response, &actualSubmitTransactionResponse)
						assert.Equal(t, 200, statusCode)
						assert.Equal(t, expectedSubmitResponse, actualSubmitTransactionResponse)
						mockTransactionSubmitter.AssertExpectations(t)
					})
				})

				Convey("transaction succeeds (with memo)", func() {
					params.Add("memo_type", "id")
					params.Add("memo", "123")

					var ledger uint64
					ledger = 100
					expectedSubmitResponse := horizon.SubmitTransactionResponse{
						Ledger: &ledger,
					}

					memo := b.MemoID{123}

					mockTransactionSubmitter.On(
						"SubmitTransaction",
						*config.Accounts.IssuingSeed,
						operation,
						memo,
					).Return(expectedSubmitResponse, nil).Once()

					Convey("it should succeed", func() {
						statusCode, response := getResponse(testServer, params)
						var actualSubmitTransactionResponse horizon.SubmitTransactionResponse
						json.Unmarshal(response, &actualSubmitTransactionResponse)
						assert.Equal(t, 200, statusCode)
						assert.Equal(t, expectedSubmitResponse, actualSubmitTransactionResponse)
						mockTransactionSubmitter.AssertExpectations(t)
					})
				})
			})
		})
	})
}
