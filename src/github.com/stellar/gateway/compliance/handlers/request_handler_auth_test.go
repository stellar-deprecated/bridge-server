package handlers

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"time"

	"github.com/facebookgo/inject"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/stellar/gateway/compliance/config"
	"github.com/stellar/gateway/db/entities"
	"github.com/stellar/gateway/mocks"
	"github.com/stellar/gateway/net"
	"github.com/stellar/gateway/protocols/compliance"
	"github.com/stellar/gateway/protocols/memo"
	"github.com/stellar/gateway/protocols/stellartoml"
	"github.com/stellar/gateway/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/zenazn/goji/web"
)

func TestRequestHandlerAuth(t *testing.T) {
	c := &config.Config{
		NetworkPassphrase: "Test SDF Network ; September 2015",
		Keys: config.Keys{
			// GBYJZW5XFAI6XV73H5SAIUYK6XZI4CGGVBUBO3ANA2SV7KKDAXTV6AEB
			SigningSeed: "SDWTLFPALQSP225BSMX7HPZ7ZEAYSUYNDLJ5QI3YGVBNRUIIELWH3XUV",
		},
	}

	mockHTTPClient := new(mocks.MockHTTPClient)
	mockEntityManager := new(mocks.MockEntityManager)
	mockRepository := new(mocks.MockRepository)
	mockFederationResolver := new(mocks.MockFederationResolver)
	mockSignerVerifier := new(mocks.MockSignerVerifier)
	mockStellartomlResolver := new(mocks.MockStellartomlResolver)
	requestHandler := RequestHandler{}

	// Inject mocks
	var g inject.Graph

	err := g.Provide(
		&inject.Object{Value: &requestHandler},
		&inject.Object{Value: c},
		&inject.Object{Value: mockHTTPClient},
		&inject.Object{Value: mockEntityManager},
		&inject.Object{Value: mockRepository},
		&inject.Object{Value: mockFederationResolver},
		&inject.Object{Value: mockSignerVerifier},
		&inject.Object{Value: mockStellartomlResolver},
	)
	if err != nil {
		panic(err)
	}

	if err := g.Populate(); err != nil {
		panic(err)
	}

	httpHandle := func(w http.ResponseWriter, r *http.Request) {
		requestHandler.HandlerAuth(web.C{}, w, r)
	}

	testServer := httptest.NewServer(http.HandlerFunc(httpHandle))
	defer testServer.Close()

	Convey("Given auth request (no sanctions check)", t, func() {
		Convey("When data param is missing", func() {
			statusCode, response := net.GetResponse(testServer, url.Values{})
			responseString := strings.TrimSpace(string(response))
			assert.Equal(t, 400, statusCode)
			expected := test.StringToJSONMap(`{
  "code": "missing_parameter",
  "message": "Required parameter is missing.",
  "data": {
    "name": "data"
  }
}`)
			assert.Equal(t, expected, test.StringToJSONMap(responseString))
		})

		Convey("When data is invalid", func() {
			params := url.Values{
				"data": {"hello world"},
				"sig":  {"bad sig"},
			}

			statusCode, response := net.GetResponse(testServer, params)
			responseString := strings.TrimSpace(string(response))
			assert.Equal(t, 400, statusCode)
			expected := test.StringToJSONMap(`{
  "code": "invalid_parameter",
  "message": "Invalid parameter.",
  "data": {
    "name": "data"
  }
}`)
			assert.Equal(t, expected, test.StringToJSONMap(responseString))
		})

		Convey("When sender's stellar.toml does not contain signing key", func() {
			mockStellartomlResolver.On(
				"GetStellarTomlByAddress",
				"alice*stellar.org",
			).Return(stellartoml.StellarToml{}, nil).Once()

			params := url.Values{
				"data": {"{\"sender\":\"alice*stellar.org\",\"need_info\":false,\"tx\":\"AAAAAC3/58Z9rycNLmF6voWX9VmDETFVGhFoWf66mcMuir/DAAAAZAAAAAAAAAAAAAAAAAAAAAO5TSe5k00+CKUuUtfafav6xITv43pTgO6QiPes4u/N6QAAAAEAAAAAAAAAAgAAAAFVU0QAAAAAAEbpO2riZmlZMkHuBxUBYAAas3hWyo9VL1IOdnfXAVFBAAAAADuaygAAAAAAGVL83DJFwH0sKmy6AIgJYD7GexiD0YuzSMioBCAUOJwAAAABVVNEAAAAAAAZUvzcMkXAfSwqbLoAiAlgPsZ7GIPRi7NIyKgEIBQ4nAAAAAAL68IAAAAAAgAAAAAAAAABRVVSAAAAAAALt4SwWfv1PIJvDRMenW0zu91YxZbphRFLA4O+gbAaigAAAAA=\",\"memo\":\"hello world\"}"},
				"sig":  {"bad sig"},
			}

			statusCode, response := net.GetResponse(testServer, params)
			responseString := strings.TrimSpace(string(response))
			assert.Equal(t, 400, statusCode)
			expected := test.StringToJSONMap(`{
  "code": "invalid_parameter",
  "message": "Invalid parameter.",
  "data": {
    "name": "data.sender"
  }
}`)
			assert.Equal(t, expected, test.StringToJSONMap(responseString))
		})

		Convey("When signature is invalid", func() {
			mockStellartomlResolver.On(
				"GetStellarTomlByAddress",
				"alice*stellar.org",
			).Return(stellartoml.StellarToml{
				SigningKey: "GBYJZW5XFAI6XV73H5SAIUYK6XZI4CGGVBUBO3ANA2SV7KKDAXTV6AEB",
			}, nil).Once()

			params := url.Values{
				"data": {"{\"sender\":\"alice*stellar.org\",\"need_info\":false,\"tx\":\"AAAAAC3/58Z9rycNLmF6voWX9VmDETFVGhFoWf66mcMuir/DAAAAZAAAAAAAAAAAAAAAAAAAAAO5TSe5k00+CKUuUtfafav6xITv43pTgO6QiPes4u/N6QAAAAEAAAAAAAAAAgAAAAFVU0QAAAAAAEbpO2riZmlZMkHuBxUBYAAas3hWyo9VL1IOdnfXAVFBAAAAADuaygAAAAAAGVL83DJFwH0sKmy6AIgJYD7GexiD0YuzSMioBCAUOJwAAAABVVNEAAAAAAAZUvzcMkXAfSwqbLoAiAlgPsZ7GIPRi7NIyKgEIBQ4nAAAAAAL68IAAAAAAgAAAAAAAAABRVVSAAAAAAALt4SwWfv1PIJvDRMenW0zu91YxZbphRFLA4O+gbAaigAAAAA=\",\"memo\":\"hello world\"}"},
				"sig":  {"ACamNqa0dF8gf97URhFVKWSD7fmvZKc5At+8dCLM5ySR0HsHySF3G2WuwYP2nKjeqjKmu3U9Z3+u1P10w1KBCA=="},
			}

			mockSignerVerifier.On(
				"Verify",
				"GBYJZW5XFAI6XV73H5SAIUYK6XZI4CGGVBUBO3ANA2SV7KKDAXTV6AEB",
				mock.AnythingOfType("[]uint8"),
				mock.AnythingOfType("[]uint8"),
			).Return(errors.New("Verify error")).Once()

			statusCode, response := net.GetResponse(testServer, params)
			responseString := strings.TrimSpace(string(response))
			assert.Equal(t, 400, statusCode)
			expected := test.StringToJSONMap(`{
  "code": "invalid_parameter",
  "message": "Invalid parameter.",
  "data": {
    "name": "sig"
  }
}`)
			assert.Equal(t, expected, test.StringToJSONMap(responseString))
		})

		Convey("When all params are valid", func() {
			params := url.Values{
				"data": {"{\"sender\":\"alice*stellar.org\",\"need_info\":false,\"tx\":\"AAAAAC3/58Z9rycNLmF6voWX9VmDETFVGhFoWf66mcMuir/DAAAAZAAAAAAAAAAAAAAAAAAAAANEE2+jVbNnihFGrRb36GSelPtPwh/nfoMQwGD2HKr/igAAAAEAAAAAAAAAAgAAAAFVU0QAAAAAAEbpO2riZmlZMkHuBxUBYAAas3hWyo9VL1IOdnfXAVFBAAAAADuaygAAAAAAGVL83DJFwH0sKmy6AIgJYD7GexiD0YuzSMioBCAUOJwAAAABVVNEAAAAAAAZUvzcMkXAfSwqbLoAiAlgPsZ7GIPRi7NIyKgEIBQ4nAAAAAAL68IAAAAAAgAAAAAAAAABRVVSAAAAAAALt4SwWfv1PIJvDRMenW0zu91YxZbphRFLA4O+gbAaigAAAAA=\",\"memo\":\"{}\"}"},
				"sig":  {"ACamNqa0dF8gf97URhFVKWSD7fmvZKc5At+8dCLM5ySR0HsHySF3G2WuwYP2nKjeqjKmu3U9Z3+u1P10w1KBCA=="},
			}

			mockStellartomlResolver.On(
				"GetStellarTomlByAddress",
				"alice*stellar.org",
			).Return(stellartoml.StellarToml{
				SigningKey: "GBYJZW5XFAI6XV73H5SAIUYK6XZI4CGGVBUBO3ANA2SV7KKDAXTV6AEB",
			}, nil).Once()

			mockSignerVerifier.On(
				"Verify",
				"GBYJZW5XFAI6XV73H5SAIUYK6XZI4CGGVBUBO3ANA2SV7KKDAXTV6AEB",
				mock.AnythingOfType("[]uint8"),
				mock.AnythingOfType("[]uint8"),
			).Return(nil).Once()

			Convey("it returns AuthResponse", func() {
				authorizedTransaction := &entities.AuthorizedTransaction{
					TransactionID:  "29ec92f95b00dd8e8bbb0d2a2fc90db8ed5b26c396c44ac978bb13ccd8d25524",
					Memo:           "RBNvo1WzZ4oRRq0W9+hknpT7T8If536DEMBg9hyq/4o=",
					TransactionXdr: "AAAAAC3/58Z9rycNLmF6voWX9VmDETFVGhFoWf66mcMuir/DAAAAZAAAAAAAAAAAAAAAAAAAAANEE2+jVbNnihFGrRb36GSelPtPwh/nfoMQwGD2HKr/igAAAAEAAAAAAAAAAgAAAAFVU0QAAAAAAEbpO2riZmlZMkHuBxUBYAAas3hWyo9VL1IOdnfXAVFBAAAAADuaygAAAAAAGVL83DJFwH0sKmy6AIgJYD7GexiD0YuzSMioBCAUOJwAAAABVVNEAAAAAAAZUvzcMkXAfSwqbLoAiAlgPsZ7GIPRi7NIyKgEIBQ4nAAAAAAL68IAAAAAAgAAAAAAAAABRVVSAAAAAAALt4SwWfv1PIJvDRMenW0zu91YxZbphRFLA4O+gbAaigAAAAA=",
					Data:           params["data"][0],
				}

				mockEntityManager.On(
					"Persist",
					mock.AnythingOfType("*entities.AuthorizedTransaction"),
				).Run(func(args mock.Arguments) {
					value := args.Get(0).(*entities.AuthorizedTransaction)
					assert.Equal(t, authorizedTransaction.TransactionID, value.TransactionID)
					assert.Equal(t, authorizedTransaction.Memo, value.Memo)
					assert.Equal(t, authorizedTransaction.TransactionXdr, value.TransactionXdr)
					assert.WithinDuration(t, time.Now(), value.AuthorizedAt, 2*time.Second)
					assert.Equal(t, authorizedTransaction.Data, value.Data)
				}).Return(nil).Once()

				statusCode, response := net.GetResponse(testServer, params)
				responseString := strings.TrimSpace(string(response))
				assert.Equal(t, 200, statusCode)
				expected := test.StringToJSONMap(`{
  "info_status": "ok",
  "tx_status": "ok"
}`)
				assert.Equal(t, expected, test.StringToJSONMap(responseString))
			})
		})
	})

	Convey("Given auth request (sanctions check)", t, func() {
		c.Callbacks = config.Callbacks{
			Sanctions: "http://sanctions",
			AskUser:   "http://ask_user",
			FetchInfo: "http://fetch_info",
		}

		memoPreimage := memo.Memo{
			Transaction: memo.Transaction{
				Route:      "bob*acme.com",
				Note:       "Happy birthday",
				SenderInfo: "senderInfoJson",
				Extra:      "extra",
			},
		}

		Convey("When all params are valid (NeedInfo = `false`)", func() {
			authData := compliance.AuthData{
				Sender:   "alice*stellar.org",
				NeedInfo: false,
				Tx:       "AAAAAC3/58Z9rycNLmF6voWX9VmDETFVGhFoWf66mcMuir/DAAAAZAAAAAAAAAAAAAAAAAAAAANeCnOi6ZSpMNIMFUUVIfbBc5OA5kpzDbg+AJ6X8/WynAAAAAEAAAAAAAAAAgAAAAFVU0QAAAAAAEbpO2riZmlZMkHuBxUBYAAas3hWyo9VL1IOdnfXAVFBAAAAADuaygAAAAAAGVL83DJFwH0sKmy6AIgJYD7GexiD0YuzSMioBCAUOJwAAAABVVNEAAAAAAAZUvzcMkXAfSwqbLoAiAlgPsZ7GIPRi7NIyKgEIBQ4nAAAAAAL68IAAAAAAgAAAAAAAAABRVVSAAAAAAALt4SwWfv1PIJvDRMenW0zu91YxZbphRFLA4O+gbAaigAAAAA=",
				Memo:     string(memoPreimage.Marshal()),
			}

			params := url.Values{
				"data": {string(authData.Marshal())},
				"sig":  {"Q2cQVOn/A+aOxrLLeUPwHmBm3LMvlfXN8tDHo4Oi6SxWWueMTDfRkC4XvRX4emLij+Npo7/GfrZ82CnT5yB5Dg=="},
			}

			mockStellartomlResolver.On(
				"GetStellarTomlByAddress",
				"alice*stellar.org",
			).Return(stellartoml.StellarToml{
				SigningKey: "GBYJZW5XFAI6XV73H5SAIUYK6XZI4CGGVBUBO3ANA2SV7KKDAXTV6AEB",
			}, nil).Once()

			mockSignerVerifier.On(
				"Verify",
				"GBYJZW5XFAI6XV73H5SAIUYK6XZI4CGGVBUBO3ANA2SV7KKDAXTV6AEB",
				mock.AnythingOfType("[]uint8"),
				mock.AnythingOfType("[]uint8"),
			).Return(nil).Once()

			Convey("when sanctions server returns forbidden it returns tx_status `denied`", func() {
				mockHTTPClient.On(
					"PostForm",
					"http://sanctions",
					url.Values{"sender": {memoPreimage.Transaction.SenderInfo}},
				).Return(
					net.BuildHTTPResponse(403, "forbidden"),
					nil,
				).Once()

				statusCode, response := net.GetResponse(testServer, params)
				responseString := strings.TrimSpace(string(response))
				assert.Equal(t, 200, statusCode)
				expected := test.StringToJSONMap(`{
		  "info_status": "ok",
		  "tx_status": "denied"
		}`)
				assert.Equal(t, expected, test.StringToJSONMap(responseString))
			})

			Convey("when sanctions server returns accepted it returns tx_status `pending`", func() {
				mockHTTPClient.On(
					"PostForm",
					"http://sanctions",
					url.Values{"sender": {memoPreimage.Transaction.SenderInfo}},
				).Return(
					net.BuildHTTPResponse(202, "pending"),
					nil,
				).Once()

				statusCode, response := net.GetResponse(testServer, params)
				responseString := strings.TrimSpace(string(response))
				assert.Equal(t, 200, statusCode)
				expected := test.StringToJSONMap(`{
  "info_status": "ok",
  "tx_status": "pending",
  "pending": 600
}`)
				assert.Equal(t, expected, test.StringToJSONMap(responseString))
			})

			Convey("when sanctions server returns ok it returns tx_status `ok` and persists transaction", func() {
				mockHTTPClient.On(
					"PostForm",
					"http://sanctions",
					url.Values{"sender": {memoPreimage.Transaction.SenderInfo}},
				).Return(
					net.BuildHTTPResponse(200, "ok"),
					nil,
				).Once()

				authorizedTransaction := &entities.AuthorizedTransaction{
					TransactionID:  "f62589932eb9fcf0bf28fe95510bf614caf3169c67a85e75475a390a79b5ecc9",
					Memo:           "XgpzoumUqTDSDBVFFSH2wXOTgOZKcw24PgCel/P1spw=",
					TransactionXdr: "AAAAAC3/58Z9rycNLmF6voWX9VmDETFVGhFoWf66mcMuir/DAAAAZAAAAAAAAAAAAAAAAAAAAANeCnOi6ZSpMNIMFUUVIfbBc5OA5kpzDbg+AJ6X8/WynAAAAAEAAAAAAAAAAgAAAAFVU0QAAAAAAEbpO2riZmlZMkHuBxUBYAAas3hWyo9VL1IOdnfXAVFBAAAAADuaygAAAAAAGVL83DJFwH0sKmy6AIgJYD7GexiD0YuzSMioBCAUOJwAAAABVVNEAAAAAAAZUvzcMkXAfSwqbLoAiAlgPsZ7GIPRi7NIyKgEIBQ4nAAAAAAL68IAAAAAAgAAAAAAAAABRVVSAAAAAAALt4SwWfv1PIJvDRMenW0zu91YxZbphRFLA4O+gbAaigAAAAA=",
					Data:           params["data"][0],
				}

				mockEntityManager.On(
					"Persist",
					mock.AnythingOfType("*entities.AuthorizedTransaction"),
				).Run(func(args mock.Arguments) {
					value := args.Get(0).(*entities.AuthorizedTransaction)
					assert.Equal(t, authorizedTransaction.TransactionID, value.TransactionID)
					assert.Equal(t, authorizedTransaction.Memo, value.Memo)
					assert.Equal(t, authorizedTransaction.TransactionXdr, value.TransactionXdr)
					assert.WithinDuration(t, time.Now(), value.AuthorizedAt, 2*time.Second)
					assert.Equal(t, authorizedTransaction.Data, value.Data)
				}).Return(nil).Once()

				statusCode, response := net.GetResponse(testServer, params)
				responseString := strings.TrimSpace(string(response))
				assert.Equal(t, 200, statusCode)
				expected := test.StringToJSONMap(`{
  "info_status": "ok",
  "tx_status": "ok"
}`)
				assert.Equal(t, expected, test.StringToJSONMap(responseString))
			})
		})

		Convey("When all params are valid (NeedInfo = `true`)", func() {
			authData := compliance.AuthData{
				Sender:   "alice*stellar.org",
				NeedInfo: true,
				Tx:       "AAAAAC3/58Z9rycNLmF6voWX9VmDETFVGhFoWf66mcMuir/DAAAAZAAAAAAAAAAAAAAAAAAAAANeCnOi6ZSpMNIMFUUVIfbBc5OA5kpzDbg+AJ6X8/WynAAAAAEAAAAAAAAAAgAAAAFVU0QAAAAAAEbpO2riZmlZMkHuBxUBYAAas3hWyo9VL1IOdnfXAVFBAAAAADuaygAAAAAAGVL83DJFwH0sKmy6AIgJYD7GexiD0YuzSMioBCAUOJwAAAABVVNEAAAAAAAZUvzcMkXAfSwqbLoAiAlgPsZ7GIPRi7NIyKgEIBQ4nAAAAAAL68IAAAAAAgAAAAAAAAABRVVSAAAAAAALt4SwWfv1PIJvDRMenW0zu91YxZbphRFLA4O+gbAaigAAAAA=",
				Memo:     string(memoPreimage.Marshal()),
			}

			params := url.Values{
				"data": {string(authData.Marshal())},
				"sig":  {"Q2cQVOn/A+aOxrLLeUPwHmBm3LMvlfXN8tDHo4Oi6SxWWueMTDfRkC4XvRX4emLij+Npo7/GfrZ82CnT5yB5Dg=="},
			}

			mockStellartomlResolver.On(
				"GetStellarTomlByAddress",
				"alice*stellar.org",
			).Return(stellartoml.StellarToml{
				SigningKey: "GBYJZW5XFAI6XV73H5SAIUYK6XZI4CGGVBUBO3ANA2SV7KKDAXTV6AEB",
			}, nil).Once()

			mockSignerVerifier.On(
				"Verify",
				"GBYJZW5XFAI6XV73H5SAIUYK6XZI4CGGVBUBO3ANA2SV7KKDAXTV6AEB",
				mock.AnythingOfType("[]uint8"),
				mock.AnythingOfType("[]uint8"),
			).Return(nil).Once()

			// Make sanctions checks successful (tested in the previous test case)
			mockHTTPClient.On(
				"PostForm",
				"http://sanctions",
				url.Values{
					"sender": {memoPreimage.Transaction.SenderInfo},
				},
			).Return(
				net.BuildHTTPResponse(200, "ok"),
				nil,
			).Once()

			Convey("when ask_user server returns forbidden it returns info_status `denied`", func() {
				mockHTTPClient.On(
					"PostForm",
					"http://ask_user",
					url.Values{
						"sender":       {memoPreimage.Transaction.SenderInfo},
						"note":         {memoPreimage.Transaction.Note},
						"amount":       {"20.0000000"},
						"asset_code":   {"USD"},
						"asset_issuer": {"GAMVF7G4GJC4A7JMFJWLUAEIBFQD5RT3DCB5DC5TJDEKQBBACQ4JZVEE"},
					},
				).Return(
					net.BuildHTTPResponse(403, "forbidden"),
					nil,
				).Once()

				statusCode, response := net.GetResponse(testServer, params)
				responseString := strings.TrimSpace(string(response))
				assert.Equal(t, 200, statusCode)
				expected := test.StringToJSONMap(`{
  "info_status": "denied",
  "tx_status": "ok"
}`)
				assert.Equal(t, expected, test.StringToJSONMap(responseString))
			})

			Convey("when ask_user server returns pending it returns info_status `pending`", func() {
				mockHTTPClient.On(
					"PostForm",
					"http://ask_user",
					url.Values{
						"sender":       {memoPreimage.Transaction.SenderInfo},
						"note":         {memoPreimage.Transaction.Note},
						"amount":       {"20.0000000"},
						"asset_code":   {"USD"},
						"asset_issuer": {"GAMVF7G4GJC4A7JMFJWLUAEIBFQD5RT3DCB5DC5TJDEKQBBACQ4JZVEE"},
					},
				).Return(
					net.BuildHTTPResponse(202, "{\"pending\": 300}"),
					nil,
				).Once()

				statusCode, response := net.GetResponse(testServer, params)
				responseString := strings.TrimSpace(string(response))
				assert.Equal(t, 200, statusCode)
				expected := test.StringToJSONMap(`{
  "info_status": "pending",
  "tx_status": "ok",
  "pending": 300
}`)
				assert.Equal(t, expected, test.StringToJSONMap(responseString))
			})

			Convey("when ask_user server returns pending but invalid response body it returns info_status `pending` (600 seconds)", func() {
				mockHTTPClient.On(
					"PostForm",
					"http://ask_user",
					url.Values{
						"sender":       {memoPreimage.Transaction.SenderInfo},
						"note":         {memoPreimage.Transaction.Note},
						"amount":       {"20.0000000"},
						"asset_code":   {"USD"},
						"asset_issuer": {"GAMVF7G4GJC4A7JMFJWLUAEIBFQD5RT3DCB5DC5TJDEKQBBACQ4JZVEE"},
					},
				).Return(
					net.BuildHTTPResponse(202, "pending"),
					nil,
				).Once()

				statusCode, response := net.GetResponse(testServer, params)
				responseString := strings.TrimSpace(string(response))
				assert.Equal(t, 200, statusCode)
				expected := test.StringToJSONMap(`{
  "info_status": "pending",
  "tx_status": "ok",
  "pending": 600
}`)
				assert.Equal(t, expected, test.StringToJSONMap(responseString))
			})

			Convey("when ask_user server returns ok it returns info_status `ok` and DestInfo and persists transaction", func() {
				mockHTTPClient.On(
					"PostForm",
					"http://ask_user",
					url.Values{
						"sender":       {memoPreimage.Transaction.SenderInfo},
						"note":         {memoPreimage.Transaction.Note},
						"amount":       {"20.0000000"},
						"asset_code":   {"USD"},
						"asset_issuer": {"GAMVF7G4GJC4A7JMFJWLUAEIBFQD5RT3DCB5DC5TJDEKQBBACQ4JZVEE"},
					},
				).Return(
					net.BuildHTTPResponse(200, "ok"),
					nil,
				).Once()

				mockHTTPClient.On(
					"PostForm",
					"http://fetch_info",
					url.Values{"address": {"bob*acme.com"}},
				).Return(
					net.BuildHTTPResponse(200, "user data"),
					nil,
				).Once()

				authorizedTransaction := &entities.AuthorizedTransaction{
					TransactionID:  "f62589932eb9fcf0bf28fe95510bf614caf3169c67a85e75475a390a79b5ecc9",
					Memo:           "XgpzoumUqTDSDBVFFSH2wXOTgOZKcw24PgCel/P1spw=",
					TransactionXdr: "AAAAAC3/58Z9rycNLmF6voWX9VmDETFVGhFoWf66mcMuir/DAAAAZAAAAAAAAAAAAAAAAAAAAANeCnOi6ZSpMNIMFUUVIfbBc5OA5kpzDbg+AJ6X8/WynAAAAAEAAAAAAAAAAgAAAAFVU0QAAAAAAEbpO2riZmlZMkHuBxUBYAAas3hWyo9VL1IOdnfXAVFBAAAAADuaygAAAAAAGVL83DJFwH0sKmy6AIgJYD7GexiD0YuzSMioBCAUOJwAAAABVVNEAAAAAAAZUvzcMkXAfSwqbLoAiAlgPsZ7GIPRi7NIyKgEIBQ4nAAAAAAL68IAAAAAAgAAAAAAAAABRVVSAAAAAAALt4SwWfv1PIJvDRMenW0zu91YxZbphRFLA4O+gbAaigAAAAA=",
					Data:           params["data"][0],
				}

				mockEntityManager.On(
					"Persist",
					mock.AnythingOfType("*entities.AuthorizedTransaction"),
				).Run(func(args mock.Arguments) {
					value := args.Get(0).(*entities.AuthorizedTransaction)
					assert.Equal(t, authorizedTransaction.TransactionID, value.TransactionID)
					assert.Equal(t, authorizedTransaction.Memo, value.Memo)
					assert.Equal(t, authorizedTransaction.TransactionXdr, value.TransactionXdr)
					assert.WithinDuration(t, time.Now(), value.AuthorizedAt, 2*time.Second)
					assert.Equal(t, authorizedTransaction.Data, value.Data)
				}).Return(nil).Once()

				statusCode, response := net.GetResponse(testServer, params)
				responseString := strings.TrimSpace(string(response))
				assert.Equal(t, 200, statusCode)
				expected := test.StringToJSONMap(`{
  "info_status": "ok",
  "tx_status": "ok",
  "dest_info": "user data"
}`)
				assert.Equal(t, expected, test.StringToJSONMap(responseString))
			})

			Convey("When no callbacks.ask_user server", func() {
				c.Callbacks.AskUser = ""

				Convey("when FI allowed it returns info_status = `ok` and DestInfo and persists transaction", func() {
					mockRepository.On(
						"GetAllowedFiByDomain",
						"stellar.org", // sender = `alice*stellar.org`
					).Return(
						&entities.AllowedFi{}, // It just returns existing record
						nil,
					).Once()

					mockHTTPClient.On(
						"PostForm",
						"http://fetch_info",
						url.Values{"address": {"bob*acme.com"}},
					).Return(
						net.BuildHTTPResponse(200, "user data"),
						nil,
					).Once()

					authorizedTransaction := &entities.AuthorizedTransaction{
						TransactionID:  "f62589932eb9fcf0bf28fe95510bf614caf3169c67a85e75475a390a79b5ecc9",
						Memo:           "XgpzoumUqTDSDBVFFSH2wXOTgOZKcw24PgCel/P1spw=",
						TransactionXdr: "AAAAAC3/58Z9rycNLmF6voWX9VmDETFVGhFoWf66mcMuir/DAAAAZAAAAAAAAAAAAAAAAAAAAANeCnOi6ZSpMNIMFUUVIfbBc5OA5kpzDbg+AJ6X8/WynAAAAAEAAAAAAAAAAgAAAAFVU0QAAAAAAEbpO2riZmlZMkHuBxUBYAAas3hWyo9VL1IOdnfXAVFBAAAAADuaygAAAAAAGVL83DJFwH0sKmy6AIgJYD7GexiD0YuzSMioBCAUOJwAAAABVVNEAAAAAAAZUvzcMkXAfSwqbLoAiAlgPsZ7GIPRi7NIyKgEIBQ4nAAAAAAL68IAAAAAAgAAAAAAAAABRVVSAAAAAAALt4SwWfv1PIJvDRMenW0zu91YxZbphRFLA4O+gbAaigAAAAA=",
						Data:           params["data"][0],
					}

					mockEntityManager.On(
						"Persist",
						mock.AnythingOfType("*entities.AuthorizedTransaction"),
					).Run(func(args mock.Arguments) {
						value := args.Get(0).(*entities.AuthorizedTransaction)
						assert.Equal(t, authorizedTransaction.TransactionID, value.TransactionID)
						assert.Equal(t, authorizedTransaction.Memo, value.Memo)
						assert.Equal(t, authorizedTransaction.TransactionXdr, value.TransactionXdr)
						assert.WithinDuration(t, time.Now(), value.AuthorizedAt, 2*time.Second)
						assert.Equal(t, authorizedTransaction.Data, value.Data)
					}).Return(nil).Once()

					statusCode, response := net.GetResponse(testServer, params)
					responseString := strings.TrimSpace(string(response))
					assert.Equal(t, 200, statusCode)
					expected := test.StringToJSONMap(`{
  "info_status": "ok",
  "tx_status": "ok",
  "dest_info": "user data"
}`)
					assert.Equal(t, expected, test.StringToJSONMap(responseString))
				})

				Convey("when FI not allowed but User is allowed it returns info_status = `ok` and DestInfo and persists transaction", func() {
					mockRepository.On(
						"GetAllowedFiByDomain",
						"stellar.org", // sender = `alice*stellar.org`
					).Return(
						nil,
						nil,
					).Once()

					mockRepository.On(
						"GetAllowedUserByDomainAndUserID",
						"stellar.org", // sender = `alice*stellar.org`
						"alice",
					).Return(
						&entities.AllowedUser{},
						nil,
					).Once()

					mockHTTPClient.On(
						"PostForm",
						"http://fetch_info",
						url.Values{"address": {"bob*acme.com"}},
					).Return(
						net.BuildHTTPResponse(200, "user data"),
						nil,
					).Once()

					authorizedTransaction := &entities.AuthorizedTransaction{
						TransactionID:  "f62589932eb9fcf0bf28fe95510bf614caf3169c67a85e75475a390a79b5ecc9",
						Memo:           "XgpzoumUqTDSDBVFFSH2wXOTgOZKcw24PgCel/P1spw=",
						TransactionXdr: "AAAAAC3/58Z9rycNLmF6voWX9VmDETFVGhFoWf66mcMuir/DAAAAZAAAAAAAAAAAAAAAAAAAAANeCnOi6ZSpMNIMFUUVIfbBc5OA5kpzDbg+AJ6X8/WynAAAAAEAAAAAAAAAAgAAAAFVU0QAAAAAAEbpO2riZmlZMkHuBxUBYAAas3hWyo9VL1IOdnfXAVFBAAAAADuaygAAAAAAGVL83DJFwH0sKmy6AIgJYD7GexiD0YuzSMioBCAUOJwAAAABVVNEAAAAAAAZUvzcMkXAfSwqbLoAiAlgPsZ7GIPRi7NIyKgEIBQ4nAAAAAAL68IAAAAAAgAAAAAAAAABRVVSAAAAAAALt4SwWfv1PIJvDRMenW0zu91YxZbphRFLA4O+gbAaigAAAAA=",
						Data:           params["data"][0],
					}

					mockEntityManager.On(
						"Persist",
						mock.AnythingOfType("*entities.AuthorizedTransaction"),
					).Run(func(args mock.Arguments) {
						value := args.Get(0).(*entities.AuthorizedTransaction)
						assert.Equal(t, authorizedTransaction.TransactionID, value.TransactionID)
						assert.Equal(t, authorizedTransaction.Memo, value.Memo)
						assert.Equal(t, authorizedTransaction.TransactionXdr, value.TransactionXdr)
						assert.WithinDuration(t, time.Now(), value.AuthorizedAt, 2*time.Second)
						assert.Equal(t, authorizedTransaction.Data, value.Data)
					}).Return(nil).Once()

					statusCode, response := net.GetResponse(testServer, params)
					responseString := strings.TrimSpace(string(response))
					assert.Equal(t, 200, statusCode)
					expected := test.StringToJSONMap(`{
  "info_status": "ok",
  "tx_status": "ok",
  "dest_info": "user data"
}`)
					assert.Equal(t, expected, test.StringToJSONMap(responseString))
				})

				Convey("when neither FI nor User is allowed it returns info_status = `denied`", func() {
					mockRepository.On(
						"GetAllowedFiByDomain",
						"stellar.org", // sender = `alice*stellar.org`
					).Return(
						nil,
						nil,
					).Once()

					mockRepository.On(
						"GetAllowedUserByDomainAndUserID",
						"stellar.org", // sender = `alice*stellar.org`
						"alice",
					).Return(
						nil,
						nil,
					).Once()

					statusCode, response := net.GetResponse(testServer, params)
					responseString := strings.TrimSpace(string(response))
					assert.Equal(t, 200, statusCode)
					expected := test.StringToJSONMap(`{
  "info_status": "denied",
  "tx_status": "ok"
}`)
					assert.Equal(t, expected, test.StringToJSONMap(responseString))
				})
			})
		})
	})
}
