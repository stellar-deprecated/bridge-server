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
	"github.com/stellar/gateway/net"
	"github.com/stellar/gateway/test"
	b "github.com/stellar/go/build"
	"github.com/stretchr/testify/assert"
)

func TestRequestHandlerAuthorize(t *testing.T) {
	mockTransactionSubmitter := new(mocks.MockTransactionSubmitter)

	config := config.Config{
		Assets: []config.Asset{
			{Code: "USD", Issuer: "GD4I7AFSLZGTDL34TQLWJOM2NHLIIOEKD5RHHZUW54HERBLSIRKUOXRR"},
			{Code: "EUR", Issuer: "GD4I7AFSLZGTDL34TQLWJOM2NHLIIOEKD5RHHZUW54HERBLSIRKUOXRR"},
		},
		Accounts: config.Accounts{
			IssuingAccountID: "GD4I7AFSLZGTDL34TQLWJOM2NHLIIOEKD5RHHZUW54HERBLSIRKUOXRR",
			// GBQXA3ABGQGTCLEVZIUTDRWWJOQD5LSAEDZAG7GMOGD2HBLWONGUVO4I
			AuthorizingSeed: "SC37TBSIAYKIDQ6GTGLT2HSORLIHZQHBXVFI5P5K4Q5TSHRTRBK3UNWG",
		},
	}

	requestHandler := RequestHandler{Config: &config, TransactionSubmitter: mockTransactionSubmitter}
	testServer := httptest.NewServer(http.HandlerFunc(requestHandler.Authorize))
	defer testServer.Close()

	Convey("Given authorize request", t, func() {
		Convey("When accountId is invalid", func() {
			accountID := "GD3YBOYIUVLU"
			assetCode := "USD"

			Convey("it should return error", func() {
				statusCode, response := net.GetResponse(testServer, url.Values{"account_id": {accountID}, "asset_code": {assetCode}})
				responseString := strings.TrimSpace(string(response))
				assert.Equal(t, 400, statusCode)
				expected := test.StringToJSONMap(`{
				  "code": "invalid_parameter",
				  "message": "Invalid parameter.",
				  "data": {
				    "name": "account_id"
				  }
				}`)
				assert.Equal(t, expected, test.StringToJSONMap(responseString, "more_info"))
			})
		})

		Convey("When assetCode is invalid", func() {
			accountID := "GDSIKW43UA6JTOA47WVEBCZ4MYC74M3GNKNXTVDXFHXYYTNO5GGVN632"
			assetCode := "GBP"

			Convey("it should return error", func() {
				statusCode, response := net.GetResponse(testServer, url.Values{"account_id": {accountID}, "asset_code": {assetCode}})
				responseString := strings.TrimSpace(string(response))
				assert.Equal(t, 400, statusCode)
				expected := test.StringToJSONMap(`{
				  "code": "invalid_parameter",
				  "message": "Invalid parameter.",
				  "data": {
				    "name": "asset_code"
				  }
				}`)
				assert.Equal(t, expected, test.StringToJSONMap(responseString, "more_info"))
			})
		})

		Convey("When params are valid", func() {
			accountID := "GDSIKW43UA6JTOA47WVEBCZ4MYC74M3GNKNXTVDXFHXYYTNO5GGVN632"
			assetCode := "USD"

			operation := b.AllowTrust(
				b.Trustor{accountID},
				b.Authorize{true},
				b.AllowTrustAsset{assetCode},
			)

			Convey("transaction fails", func() {
				mockTransactionSubmitter.On(
					"SubmitTransaction",
					(*string)(nil),
					config.Accounts.AuthorizingSeed,
					operation,
					nil,
				).Return(
					horizon.SubmitTransactionResponse{},
					errors.New("Error sending transaction"),
				).Once()

				Convey("it should return server error", func() {
					statusCode, response := net.GetResponse(testServer, url.Values{"account_id": {accountID}, "asset_code": {assetCode}})
					responseString := strings.TrimSpace(string(response))
					assert.Equal(t, 500, statusCode)
					expected := test.StringToJSONMap(`{
					  "code": "internal_server_error",
					  "message": "Internal Server Error, please try again."
					}`)
					assert.Equal(t, expected, test.StringToJSONMap(responseString))

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
					(*string)(nil),
					config.Accounts.AuthorizingSeed,
					operation,
					nil,
				).Return(expectedSubmitResponse, nil).Once()

				Convey("it should succeed", func() {
					statusCode, response := net.GetResponse(testServer, url.Values{"account_id": {accountID}, "asset_code": {assetCode}})
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
