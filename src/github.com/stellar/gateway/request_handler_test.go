package gateway

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/stellar/gateway/horizon"
	"github.com/stellar/gateway/mocks"
	b "github.com/stellar/go-stellar-base/build"
	"github.com/stellar/go-stellar-base/keypair"
	"github.com/stretchr/testify/assert"
)

func TestRequestHandlerAuthorize(t *testing.T) {
	mockTransactionSubmitter := new(mocks.MockTransactionSubmitter)

	config := Config{
		Assets: []string{"USD", "EUR"},
		Accounts: Accounts{
			// GD4I7AFSLZGTDL34TQLWJOM2NHLIIOEKD5RHHZUW54HERBLSIRKUOXRR
			IssuingSeed: "SC34WILLHVADXMP6ACPMIRA6TRAWJMVCLPFNW7S6MUMXJVLAZUC4EWHP",
			// GBQXA3ABGQGTCLEVZIUTDRWWJOQD5LSAEDZAG7GMOGD2HBLWONGUVO4I
			AuthorizingSeed: "SC37TBSIAYKIDQ6GTGLT2HSORLIHZQHBXVFI5P5K4Q5TSHRTRBK3UNWG",
		},
	}

	requestHandler := RequestHandler{config: &config, transactionSubmitter: mockTransactionSubmitter}
	testServer := httptest.NewServer(http.HandlerFunc(requestHandler.Authorize))
	defer testServer.Close()

	Convey("Given authorize request", t, func() {
		Convey("When accountId is invalid", func() {
			accountId := "GD3YBOYIUVLU"
			assetCode := "USD"

			Convey("it should return error", func() {
				statusCode, response := GetResponse(testServer, "?accountId="+accountId+"&assetCode="+assetCode)
				responseString := strings.TrimSpace(string(response))
				assert.Equal(t, 400, statusCode)
				assert.Equal(t, errorResponseString("invalid_account_id", "accountId parameter is invalid"), responseString)
			})
		})

		Convey("When assetCode is invalid", func() {
			accountId := "GDSIKW43UA6JTOA47WVEBCZ4MYC74M3GNKNXTVDXFHXYYTNO5GGVN632"
			assetCode := "GBP"

			Convey("it should return error", func() {
				statusCode, response := GetResponse(testServer, "?accountId="+accountId+"&assetCode="+assetCode)
				responseString := strings.TrimSpace(string(response))
				assert.Equal(t, 400, statusCode)
				assert.Equal(t, errorResponseString("invalid_asset_code", "Given assetCode not allowed"), responseString)
			})
		})

		Convey("When params are valid", func() {
			accountId := "GDSIKW43UA6JTOA47WVEBCZ4MYC74M3GNKNXTVDXFHXYYTNO5GGVN632"
			assetCode := "USD"

			operation := b.AllowTrust(
				b.Trustor{accountId},
				b.Authorize{true},
				b.AllowTrustAsset{assetCode},
			)

			Convey("transaction fails", func() {
				mockTransactionSubmitter.On(
					"SubmitTransaction",
					config.Accounts.AuthorizingSeed,
					operation,
				).Return(
					horizon.SubmitTransactionResponse{},
					errors.New("Error sending transaction"),
				).Once()

				Convey("it should return server error", func() {
					statusCode, response := GetResponse(testServer, "?accountId="+accountId+"&assetCode="+assetCode)
					responseString := strings.TrimSpace(string(response))
					assert.Equal(t, 500, statusCode)
					assert.Equal(t, "{\"code\":\"server_error\",\"message\":\"Server error\"}", responseString)
					mockTransactionSubmitter.AssertExpectations(t)
				})
			})

			Convey("transaction succeeds", func() {
				var ledger uint64
				ledger = 100
				expectedSubmitResponse := horizon.SubmitTransactionResponse{
					Ledger: &ledger,
				}

				mockTransactionSubmitter.On(
					"SubmitTransaction",
					config.Accounts.AuthorizingSeed,
					operation,
				).Return(expectedSubmitResponse, nil).Once()

				Convey("it should succeed", func() {
					statusCode, response := GetResponse(testServer, "?accountId="+accountId+"&assetCode="+assetCode)
					var actualSubmitTransactionResponse horizon.SubmitTransactionResponse
					json.Unmarshal(response, &actualSubmitTransactionResponse)
					assert.Equal(t, 200, statusCode)
					assert.Equal(t, expectedSubmitResponse, actualSubmitTransactionResponse)
					mockTransactionSubmitter.AssertExpectations(t)
				})
			})
		})
	})
}

func TestRequestHandlerSend(t *testing.T) {
	mockTransactionSubmitter := new(mocks.MockTransactionSubmitter)

	config := Config{
		Assets: []string{"USD", "EUR"},
		Accounts: Accounts{
			// GD4I7AFSLZGTDL34TQLWJOM2NHLIIOEKD5RHHZUW54HERBLSIRKUOXRR
			IssuingSeed: "SC34WILLHVADXMP6ACPMIRA6TRAWJMVCLPFNW7S6MUMXJVLAZUC4EWHP",
			// GBQXA3ABGQGTCLEVZIUTDRWWJOQD5LSAEDZAG7GMOGD2HBLWONGUVO4I
			AuthorizingSeed: "SC37TBSIAYKIDQ6GTGLT2HSORLIHZQHBXVFI5P5K4Q5TSHRTRBK3UNWG",
		},
	}

	requestHandler := RequestHandler{config: &config, transactionSubmitter: mockTransactionSubmitter}
	testServer := httptest.NewServer(http.HandlerFunc(requestHandler.Send))
	defer testServer.Close()

	Convey("Given send request", t, func() {
		Convey("When destination is invalid", func() {
			destination := "GD3YBOYIUVLU"
			assetCode := "USD"

			Convey("it should return error", func() {
				statusCode, response := GetResponse(testServer, "?destination="+destination+"&assetCode="+assetCode)
				responseString := strings.TrimSpace(string(response))
				assert.Equal(t, 400, statusCode)
				assert.Equal(t, errorResponseString("invalid_destination", "destination parameter is invalid"), responseString)
			})
		})

		Convey("When assetCode is invalid", func() {
			destination := "GDSIKW43UA6JTOA47WVEBCZ4MYC74M3GNKNXTVDXFHXYYTNO5GGVN632"
			assetCode := "GBP"

			Convey("it should return error", func() {
				statusCode, response := GetResponse(testServer, "?destination="+destination+"&assetCode="+assetCode)
				responseString := strings.TrimSpace(string(response))
				assert.Equal(t, 400, statusCode)
				assert.Equal(t, errorResponseString("invalid_asset_code", "Given assetCode not allowed"), responseString)
			})
		})

		Convey("When params are valid", func() {
			destination := "GDSIKW43UA6JTOA47WVEBCZ4MYC74M3GNKNXTVDXFHXYYTNO5GGVN632"
			assetCode := "USD"
			amount := "20"

			issuingKeypair, err := keypair.Parse(config.Accounts.IssuingSeed)
			if err != nil {
				panic(err)
			}

			operation := b.Payment(
				b.Destination{destination},
				b.CreditAmount{assetCode, issuingKeypair.Address(), amount},
			)

			Convey("transaction fails", func() {
				mockTransactionSubmitter.On(
					"SubmitTransaction",
					config.Accounts.IssuingSeed,
					operation,
				).Return(
					horizon.SubmitTransactionResponse{},
					errors.New("Error sending transaction"),
				).Once()

				Convey("it should return server error", func() {
					statusCode, response := GetResponse(testServer, "?destination="+destination+"&assetCode="+assetCode+"&amount="+amount)
					responseString := strings.TrimSpace(string(response))
					assert.Equal(t, 500, statusCode)
					assert.Equal(t, "{\"code\":\"server_error\",\"message\":\"Server error\"}", responseString)
					mockTransactionSubmitter.AssertExpectations(t)
				})
			})

			Convey("transaction succeeds", func() {
				var ledger uint64
				ledger = 100
				expectedSubmitResponse := horizon.SubmitTransactionResponse{
					Ledger: &ledger,
				}

				mockTransactionSubmitter.On(
					"SubmitTransaction",
					config.Accounts.IssuingSeed,
					operation,
				).Return(expectedSubmitResponse, nil).Once()

				Convey("it should succeed", func() {
					statusCode, response := GetResponse(testServer, "?destination="+destination+"&assetCode="+assetCode+"&amount="+amount)
					var actualSubmitTransactionResponse horizon.SubmitTransactionResponse
					json.Unmarshal(response, &actualSubmitTransactionResponse)
					assert.Equal(t, 200, statusCode)
					assert.Equal(t, expectedSubmitResponse, actualSubmitTransactionResponse)
					mockTransactionSubmitter.AssertExpectations(t)
				})
			})
		})
	})
}

func GetResponse(testServer *httptest.Server, query string) (int, []byte) {
	res, err := http.Get(testServer.URL + query)
	if err != nil {
		panic(err)
	}
	response, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		panic(err)
	}
	return res.StatusCode, response
}
