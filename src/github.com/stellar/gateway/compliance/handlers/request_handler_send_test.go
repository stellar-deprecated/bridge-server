package handlers

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/facebookgo/inject"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/stellar/gateway/compliance/config"
	"github.com/stellar/gateway/db/entities"
	"github.com/stellar/gateway/mocks"
	"github.com/stellar/gateway/net"
	"github.com/stellar/gateway/test"
	"github.com/stellar/go/build"
	"github.com/stellar/go/clients/stellartoml"
	"github.com/stellar/go/protocols/compliance"
	"github.com/stellar/go/protocols/federation"
	"github.com/stellar/go/xdr"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/zenazn/goji/web"
)

func TestRequestHandlerSend(t *testing.T) {
	c := &config.Config{
		NetworkPassphrase: "Test SDF Network ; September 2015",
		Keys: config.Keys{
			// GBYJZW5XFAI6XV73H5SAIUYK6XZI4CGGVBUBO3ANA2SV7KKDAXTV6AEB
			SigningSeed: "SDWTLFPALQSP225BSMX7HPZ7ZEAYSUYNDLJ5QI3YGVBNRUIIELWH3XUV",
		},
		Callbacks: config.Callbacks{
			FetchInfo: "http://fetch_info",
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
		&inject.Object{Value: &TestNonceGenerator{}},
	)
	if err != nil {
		panic(err)
	}

	if err := g.Populate(); err != nil {
		panic(err)
	}

	httpHandle := func(w http.ResponseWriter, r *http.Request) {
		requestHandler.HandlerSend(web.C{}, w, r)
	}

	testServer := httptest.NewServer(http.HandlerFunc(httpHandle))
	defer testServer.Close()

	Convey("Given send request", t, func() {
		Convey("When id param is missing", func() {
			statusCode, response := net.GetResponse(testServer, url.Values{})
			responseString := strings.TrimSpace(string(response))
			assert.Equal(t, 400, statusCode)
			expected := test.StringToJSONMap(`{
			  "code": "missing_parameter",
			  "message": "Required parameter is missing.",
			  "data": {
			    "name": "id"
			  }
			}`)
			assert.Equal(t, expected, test.StringToJSONMap(responseString))
		})

		Convey("When source param is missing", func() {
			statusCode, response := net.GetResponse(testServer, url.Values{"id": {"id"}})
			responseString := strings.TrimSpace(string(response))
			assert.Equal(t, 400, statusCode)
			expected := test.StringToJSONMap(`{
			  "code": "missing_parameter",
			  "message": "Required parameter is missing.",
			  "data": {
			    "name": "source"
			  }
			}`)
			assert.Equal(t, expected, test.StringToJSONMap(responseString))
		})

		Convey("When source param is invalid", func() {
			params := url.Values{
				"id":           {"id"},
				"source":       {"bad"},
				"sender":       {"alice*stellar.org"}, // GAW77Z6GPWXSODJOMF5L5BMX6VMYGEJRKUNBC2CZ725JTQZORK74HQQD
				"destination":  {"bob*stellar.org"},   // GAMVF7G4GJC4A7JMFJWLUAEIBFQD5RT3DCB5DC5TJDEKQBBACQ4JZVEE
				"amount":       {"20"},
				"asset_code":   {"USD"},
				"asset_issuer": {"GAMVF7G4GJC4A7JMFJWLUAEIBFQD5RT3DCB5DC5TJDEKQBBACQ4JZVEE"},
				"extra_memo":   {"hello world"},
			}

			statusCode, response := net.GetResponse(testServer, params)
			responseString := strings.TrimSpace(string(response))
			assert.Equal(t, 400, statusCode)
			expected := test.StringToJSONMap(`{
			  "code": "invalid_parameter",
			  "message": "Invalid parameter.",
			  "data": {
			    "name": "source"
			  }
			}`)
			assert.Equal(t, expected, test.StringToJSONMap(responseString, "more_info"))
		})

		Convey("When params are valid", func() {
			params := url.Values{
				"id":           {"id"},
				"source":       {"GAW77Z6GPWXSODJOMF5L5BMX6VMYGEJRKUNBC2CZ725JTQZORK74HQQD"},
				"sender":       {"alice*stellar.org"}, // GAW77Z6GPWXSODJOMF5L5BMX6VMYGEJRKUNBC2CZ725JTQZORK74HQQD
				"destination":  {"bob*stellar.org"},   // GAMVF7G4GJC4A7JMFJWLUAEIBFQD5RT3DCB5DC5TJDEKQBBACQ4JZVEE
				"amount":       {"20"},
				"asset_code":   {"USD"},
				"asset_issuer": {"GAMVF7G4GJC4A7JMFJWLUAEIBFQD5RT3DCB5DC5TJDEKQBBACQ4JZVEE"},
				"extra_memo":   {"hello world"},
			}

			mockRepository.Mock.On("GetAuthData", "id").Return(nil, nil).Once()
			mockEntityManager.On("Persist", mock.AnythingOfType("*entities.AuthData")).Run(func(args mock.Arguments) {
				entity, ok := args.Get(0).(*entities.AuthData)
				assert.True(t, ok, "Invalid conversion")
				assert.Equal(t, "id", entity.RequestID)
				assert.Equal(t, "stellar.org", entity.Domain)
			}).Return(nil).Once()

			senderInfo := compliance.SenderInfo{FirstName: "John", LastName: "Doe"}
			senderInfoMap, err := senderInfo.Map()
			require.NoError(t, err)

			Convey("it returns SendResponse when success (payment)", func() {
				authServer := "https://acme.com/auth"

				mockFederationResolver.On(
					"LookupByAddress",
					"bob*stellar.org",
				).Return(&federation.NameResponse{
					AccountID: "GAMVF7G4GJC4A7JMFJWLUAEIBFQD5RT3DCB5DC5TJDEKQBBACQ4JZVEE",
					MemoType:  "text",
					Memo:      federation.Memo{"bob"},
				}, nil).Once()

				mockStellartomlResolver.On(
					"GetStellarToml",
					"stellar.org",
				).Return(&stellartoml.Response{AuthServer: authServer}, nil).Once()

				attachment := compliance.Attachment{
					Nonce: "nonce",
					Transaction: compliance.Transaction{
						Route:      "bob",
						Note:       "",
						SenderInfo: senderInfoMap,
						Extra:      "hello world",
					},
				}

				attachmentJSON, err := attachment.Marshal()
				require.NoError(t, err)
				attachHash, err := attachment.Hash()
				require.NoError(t, err)

				txBuilder, err := build.Transaction(
					build.SourceAccount{"GAW77Z6GPWXSODJOMF5L5BMX6VMYGEJRKUNBC2CZ725JTQZORK74HQQD"},
					build.Sequence{0},
					build.TestNetwork,
					build.MemoHash{attachHash},
					build.Payment(
						build.Destination{"GAMVF7G4GJC4A7JMFJWLUAEIBFQD5RT3DCB5DC5TJDEKQBBACQ4JZVEE"},
						build.CreditAmount{"USD", "GAMVF7G4GJC4A7JMFJWLUAEIBFQD5RT3DCB5DC5TJDEKQBBACQ4JZVEE", "20"},
					),
				)
				require.NoError(t, err)
				txB64, err := xdr.MarshalBase64(txBuilder.TX)
				require.NoError(t, err)

				authData := compliance.AuthData{
					Sender:         "alice*stellar.org",
					NeedInfo:       false,
					Tx:             txB64,
					AttachmentJSON: string(attachmentJSON),
				}

				authDataJSON, err := authData.Marshal()
				require.NoError(t, err)

				authRequest := compliance.AuthRequest{
					DataJSON:  string(authDataJSON),
					Signature: "YeMlOYWNysyGBfsAe40z9dGgpRsKSQrqFIGAEsyJQ8osnXlLPynvJ2WQDGcBq2n5AA96YZdABhQz5ymqvxfQDw==",
				}

				authResponse := compliance.AuthResponse{
					InfoStatus: compliance.AuthStatusOk,
					TxStatus:   compliance.AuthStatusOk,
				}

				authResponseJSON, err := authResponse.Marshal()
				require.NoError(t, err)

				mockHTTPClient.On(
					"PostForm",
					c.Callbacks.FetchInfo,
					url.Values{"address": {"alice*stellar.org"}},
				).Return(
					net.BuildHTTPResponse(200, "{\"first_name\": \"John\", \"last_name\": \"Doe\"}"),
					nil,
				).Once()

				mockHTTPClient.On(
					"PostForm",
					authServer,
					authRequest.ToURLValues(),
				).Return(
					net.BuildHTTPResponse(200, string(authResponseJSON)),
					nil,
				).Once()

				mockSignerVerifier.On(
					"Sign",
					c.Keys.SigningSeed,
					[]byte(authRequest.DataJSON),
				).Return(authRequest.Signature, nil).Once()

				statusCode, response := net.GetResponse(testServer, params)
				responseString := strings.TrimSpace(string(response))
				assert.Equal(t, 200, statusCode)
				expected := test.StringToJSONMap(`{
				  "auth_response": {
				    "info_status": "ok",
				    "tx_status": "ok"
				  },
				  "transaction_xdr": "` + txB64 + `"
				}`)
				assert.Equal(t, expected, test.StringToJSONMap(responseString))
			})

			Convey("it returns SendResponse when success (path payment)", func() {
				params["send_max"] = []string{"100"}
				params["send_asset_code"] = []string{"USD"}
				params["send_asset_issuer"] = []string{"GBDOSO3K4JTGSWJSIHXAOFIBMAABVM3YK3FI6VJPKIHHM56XAFIUCGD6"}

				// Native
				params["path[0][asset_code]"] = []string{""}
				params["path[0][asset_issuer]"] = []string{""}
				// Credit
				params["path[1][asset_code]"] = []string{"EUR"}
				params["path[1][asset_issuer]"] = []string{"GAF3PBFQLH57KPECN4GRGHU5NUZ3XXKYYWLOTBIRJMBYHPUBWANIUCZU"}

				authServer := "https://acme.com/auth"

				mockFederationResolver.On(
					"LookupByAddress",
					"bob*stellar.org",
				).Return(&federation.NameResponse{
					AccountID: "GAMVF7G4GJC4A7JMFJWLUAEIBFQD5RT3DCB5DC5TJDEKQBBACQ4JZVEE",
					MemoType:  "text",
					Memo:      federation.Memo{"bob"},
				}, nil).Once()

				mockStellartomlResolver.On(
					"GetStellarToml",
					"stellar.org",
				).Return(&stellartoml.Response{AuthServer: authServer}, nil).Once()

				attachment := compliance.Attachment{
					Nonce: "nonce",
					Transaction: compliance.Transaction{
						Route:      "bob",
						Note:       "",
						SenderInfo: senderInfoMap,
						Extra:      "hello world",
					},
				}

				attachmentJSON, err := attachment.Marshal()
				require.NoError(t, err)
				attachHash, err := attachment.Hash()
				require.NoError(t, err)

				txBuilder, err := build.Transaction(
					build.SourceAccount{"GAW77Z6GPWXSODJOMF5L5BMX6VMYGEJRKUNBC2CZ725JTQZORK74HQQD"},
					build.Sequence{0},
					build.TestNetwork,
					build.MemoHash{attachHash},
					build.Payment(
						build.Destination{"GAMVF7G4GJC4A7JMFJWLUAEIBFQD5RT3DCB5DC5TJDEKQBBACQ4JZVEE"},
						build.CreditAmount{"USD", "GAMVF7G4GJC4A7JMFJWLUAEIBFQD5RT3DCB5DC5TJDEKQBBACQ4JZVEE", "20"},
						build.PayWithPath{
							build.CreditAsset("USD", "GBDOSO3K4JTGSWJSIHXAOFIBMAABVM3YK3FI6VJPKIHHM56XAFIUCGD6"),
							"100",
							[]build.Asset{
								build.NativeAsset(),
								build.CreditAsset("EUR", "GAF3PBFQLH57KPECN4GRGHU5NUZ3XXKYYWLOTBIRJMBYHPUBWANIUCZU"),
							},
						}),
				)
				require.NoError(t, err)

				txB64, _ := xdr.MarshalBase64(txBuilder.TX)

				authData := compliance.AuthData{
					Sender:         "alice*stellar.org",
					NeedInfo:       false,
					Tx:             txB64,
					AttachmentJSON: string(attachmentJSON),
				}

				authDataJSON, err := authData.Marshal()
				require.NoError(t, err)

				authRequest := compliance.AuthRequest{
					DataJSON:  string(authDataJSON),
					Signature: "ACamNqa0dF8gf97URhFVKWSD7fmvZKc5At+8dCLM5ySR0HsHySF3G2WuwYP2nKjeqjKmu3U9Z3+u1P10w1KBCA==",
				}

				authResponse := compliance.AuthResponse{
					InfoStatus: compliance.AuthStatusOk,
					TxStatus:   compliance.AuthStatusOk,
				}

				authResponseJSON, err := authResponse.Marshal()
				require.NoError(t, err)

				mockHTTPClient.On(
					"PostForm",
					c.Callbacks.FetchInfo,
					url.Values{"address": {"alice*stellar.org"}},
				).Return(
					net.BuildHTTPResponse(200, "{\"first_name\": \"John\", \"last_name\": \"Doe\"}"),
					nil,
				).Once()

				mockHTTPClient.On(
					"PostForm",
					authServer,
					authRequest.ToURLValues(),
				).Return(
					net.BuildHTTPResponse(200, string(authResponseJSON)),
					nil,
				).Once()

				mockSignerVerifier.On(
					"Sign",
					c.Keys.SigningSeed,
					[]byte(authRequest.DataJSON),
				).Return(authRequest.Signature, nil).Once()

				statusCode, response := net.GetResponse(testServer, params)
				responseString := strings.TrimSpace(string(response))
				assert.Equal(t, 200, statusCode)
				expected := test.StringToJSONMap(`{
				  "auth_response": {
				    "info_status": "ok",
				    "tx_status": "ok"
				  },
				  "transaction_xdr": "` + txB64 + `"
				}`)
				assert.Equal(t, expected, test.StringToJSONMap(responseString))
			})

			Convey("it returns SendResponse when success (forward request)", func() {
				params["destination"] = []string{""}
				params["forward_destination[domain]"] = []string{"stellar.org"}
				params["forward_destination[fields][federation_type]"] = []string{"bank_account"}
				params["forward_destination[fields][swift]"] = []string{"BOPBPHMM"}
				params["forward_destination[fields][acct]"] = []string{"2382376"}

				authServer := "https://acme.com/auth"

				mockFederationResolver.On(
					"ForwardRequest",
					"stellar.org",
					url.Values{
						"federation_type": []string{"bank_account"},
						"swift":           []string{"BOPBPHMM"},
						"acct":            []string{"2382376"},
					},
				).Return(&federation.NameResponse{
					AccountID: "GAMVF7G4GJC4A7JMFJWLUAEIBFQD5RT3DCB5DC5TJDEKQBBACQ4JZVEE",
					MemoType:  "text",
					Memo:      federation.Memo{"bob"},
				}, nil).Once()

				mockStellartomlResolver.On(
					"GetStellarToml",
					"stellar.org",
				).Return(&stellartoml.Response{AuthServer: authServer}, nil).Once()

				attachment := compliance.Attachment{
					Nonce: "nonce",
					Transaction: compliance.Transaction{
						Route:      "bob",
						Note:       "",
						SenderInfo: senderInfoMap,
						Extra:      "hello world",
					},
				}

				attachmentJSON, err := attachment.Marshal()
				require.NoError(t, err)
				attachHash, err := attachment.Hash()
				require.NoError(t, err)

				txBuilder, err := build.Transaction(
					build.SourceAccount{"GAW77Z6GPWXSODJOMF5L5BMX6VMYGEJRKUNBC2CZ725JTQZORK74HQQD"},
					build.Sequence{0},
					build.TestNetwork,
					build.MemoHash{attachHash},
					build.Payment(
						build.Destination{"GAMVF7G4GJC4A7JMFJWLUAEIBFQD5RT3DCB5DC5TJDEKQBBACQ4JZVEE"},
						build.CreditAmount{"USD", "GAMVF7G4GJC4A7JMFJWLUAEIBFQD5RT3DCB5DC5TJDEKQBBACQ4JZVEE", "20"},
					),
				)
				require.NoError(t, err)
				txB64, err := xdr.MarshalBase64(txBuilder.TX)
				require.NoError(t, err)

				authData := compliance.AuthData{
					Sender:         "alice*stellar.org",
					NeedInfo:       false,
					Tx:             txB64,
					AttachmentJSON: string(attachmentJSON),
				}

				authDataJSON, err := authData.Marshal()
				require.NoError(t, err)

				authRequest := compliance.AuthRequest{
					DataJSON:  string(authDataJSON),
					Signature: "YeMlOYWNysyGBfsAe40z9dGgpRsKSQrqFIGAEsyJQ8osnXlLPynvJ2WQDGcBq2n5AA96YZdABhQz5ymqvxfQDw==",
				}

				authResponse := compliance.AuthResponse{
					InfoStatus: compliance.AuthStatusOk,
					TxStatus:   compliance.AuthStatusOk,
				}

				authResponseJSON, err := authResponse.Marshal()
				require.NoError(t, err)

				mockHTTPClient.On(
					"PostForm",
					c.Callbacks.FetchInfo,
					url.Values{"address": {"alice*stellar.org"}},
				).Return(
					net.BuildHTTPResponse(200, "{\"first_name\": \"John\", \"last_name\": \"Doe\"}"),
					nil,
				).Once()

				mockHTTPClient.On(
					"PostForm",
					authServer,
					authRequest.ToURLValues(),
				).Return(
					net.BuildHTTPResponse(200, string(authResponseJSON)),
					nil,
				).Once()

				mockSignerVerifier.On(
					"Sign",
					c.Keys.SigningSeed,
					[]byte(authRequest.DataJSON),
				).Return(authRequest.Signature, nil).Once()

				statusCode, response := net.GetResponse(testServer, params)
				responseString := strings.TrimSpace(string(response))
				assert.Equal(t, 200, statusCode)
				expected := test.StringToJSONMap(`{
				  "auth_response": {
				    "info_status": "ok",
				    "tx_status": "ok"
				  },
				  "transaction_xdr": "` + txB64 + `"
				}`)
				assert.Equal(t, expected, test.StringToJSONMap(responseString))
			})

			Convey("it returns SendResponse when pending, then success (payment)", func() {
				authServer := "https://acme.com/auth"

				mockFederationResolver.On(
					"LookupByAddress",
					"bob*stellar.org",
				).Return(&federation.NameResponse{
					AccountID: "GAMVF7G4GJC4A7JMFJWLUAEIBFQD5RT3DCB5DC5TJDEKQBBACQ4JZVEE",
					MemoType:  "text",
					Memo:      federation.Memo{"bob"},
				}, nil).Once()

				mockStellartomlResolver.On(
					"GetStellarToml",
					"stellar.org",
				).Return(&stellartoml.Response{AuthServer: authServer}, nil).Once()

				attachment := compliance.Attachment{
					Nonce: "nonce",
					Transaction: compliance.Transaction{
						Route:      "bob",
						Note:       "",
						SenderInfo: senderInfoMap,
						Extra:      "hello world",
					},
				}

				attachmentJSON, err := attachment.Marshal()
				require.NoError(t, err)
				attachHash, err := attachment.Hash()
				require.NoError(t, err)

				txBuilder, err := build.Transaction(
					build.SourceAccount{"GAW77Z6GPWXSODJOMF5L5BMX6VMYGEJRKUNBC2CZ725JTQZORK74HQQD"},
					build.Sequence{0},
					build.TestNetwork,
					build.MemoHash{attachHash},
					build.Payment(
						build.Destination{"GAMVF7G4GJC4A7JMFJWLUAEIBFQD5RT3DCB5DC5TJDEKQBBACQ4JZVEE"},
						build.CreditAmount{"USD", "GAMVF7G4GJC4A7JMFJWLUAEIBFQD5RT3DCB5DC5TJDEKQBBACQ4JZVEE", "20"},
					),
				)
				require.NoError(t, err)
				txB64, err := xdr.MarshalBase64(txBuilder.TX)
				require.NoError(t, err)

				authData := compliance.AuthData{
					Sender:         "alice*stellar.org",
					NeedInfo:       false,
					Tx:             txB64,
					AttachmentJSON: string(attachmentJSON),
				}

				authDataJSON, err := authData.Marshal()
				require.NoError(t, err)

				authRequest := compliance.AuthRequest{
					DataJSON:  string(authDataJSON),
					Signature: "YeMlOYWNysyGBfsAe40z9dGgpRsKSQrqFIGAEsyJQ8osnXlLPynvJ2WQDGcBq2n5AA96YZdABhQz5ymqvxfQDw==",
				}

				authResponse := compliance.AuthResponse{
					InfoStatus: compliance.AuthStatusPending,
					TxStatus:   compliance.AuthStatusPending,
					Pending:    60,
				}

				authResponseJSON, err := authResponse.Marshal()
				require.NoError(t, err)

				mockHTTPClient.On(
					"PostForm",
					c.Callbacks.FetchInfo,
					url.Values{"address": {"alice*stellar.org"}},
				).Return(
					net.BuildHTTPResponse(200, "{\"first_name\": \"John\", \"last_name\": \"Doe\"}"),
					nil,
				).Once()

				mockHTTPClient.On(
					"PostForm",
					authServer,
					authRequest.ToURLValues(),
				).Return(
					net.BuildHTTPResponse(200, string(authResponseJSON)),
					nil,
				).Once()

				mockSignerVerifier.On(
					"Sign",
					c.Keys.SigningSeed,
					[]byte(authRequest.DataJSON),
				).Return(authRequest.Signature, nil).Once()

				statusCode, response := net.GetResponse(testServer, params)
				responseString := strings.TrimSpace(string(response))
				assert.Equal(t, 200, statusCode)
				expected := test.StringToJSONMap(`{
				  "auth_response": {
				    "info_status": "pending",
				    "tx_status": "pending",
				    "pending": 60
				  },
				  "transaction_xdr": "` + txB64 + `"
				}`)
				assert.Equal(t, expected, test.StringToJSONMap(responseString))

				// Sending the request again
				mockRepository.Mock.On("GetAuthData", "id").Return(&entities.AuthData{
					RequestID: "id",
					Domain:    "stellar.org",
					AuthData:  string(authDataJSON),
				}, nil).Once()

				authResponse = compliance.AuthResponse{
					InfoStatus: compliance.AuthStatusOk,
					TxStatus:   compliance.AuthStatusOk,
				}

				authResponseJSON, err = authResponse.Marshal()
				require.NoError(t, err)

				mockStellartomlResolver.On(
					"GetStellarToml",
					"stellar.org",
				).Return(&stellartoml.Response{AuthServer: authServer}, nil).Once()

				mockHTTPClient.On(
					"PostForm",
					authServer,
					authRequest.ToURLValues(),
				).Return(
					net.BuildHTTPResponse(200, string(authResponseJSON)),
					nil,
				).Once()

				mockSignerVerifier.On(
					"Sign",
					c.Keys.SigningSeed,
					[]byte(authRequest.DataJSON),
				).Return(authRequest.Signature, nil).Once()

				statusCode, response = net.GetResponse(testServer, params)
				responseString = strings.TrimSpace(string(response))
				assert.Equal(t, 200, statusCode)
				expected = test.StringToJSONMap(`{
				  "auth_response": {
				    "info_status": "ok",
				    "tx_status": "ok"
				  },
				  "transaction_xdr": "` + txB64 + `"
				}`)
				assert.Equal(t, expected, test.StringToJSONMap(responseString))
			})
		})
	})
}
