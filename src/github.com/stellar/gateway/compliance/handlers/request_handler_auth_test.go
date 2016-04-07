package handlers

import (
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
	"github.com/stellar/gateway/protocols/stellartoml"
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

	mockHttpClient := new(mocks.MockHttpClient)
	mockEntityManager := new(mocks.MockEntityManager)
	mockRepository := new(mocks.MockRepository)
	mockFederationResolver := new(mocks.MockFederationResolver)
	mockStellartomlResolver := new(mocks.MockStellartomlResolver)
	requestHandler := RequestHandler{}

	// Inject mocks
	var g inject.Graph

	err := g.Provide(
		&inject.Object{Value: &requestHandler},
		&inject.Object{Value: c},
		&inject.Object{Value: mockHttpClient},
		&inject.Object{Value: mockEntityManager},
		&inject.Object{Value: mockRepository},
		&inject.Object{Value: mockFederationResolver},
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
			assert.Equal(t, "{\n  \"code\": \"missing_parameter\",\n  \"message\": \"Required parameter is missing.\",\n  \"data\": {\n    \"name\": \"data\"\n  }\n}", responseString)
		})

		Convey("When data is invalid", func() {
			params := url.Values{
				"data": {"hello world"},
				"sig":  {"bad sig"},
			}

			statusCode, response := net.GetResponse(testServer, params)
			responseString := strings.TrimSpace(string(response))
			assert.Equal(t, 400, statusCode)
			assert.Equal(t, "{\n  \"code\": \"invalid_parameter\",\n  \"message\": \"Invalid parameter.\",\n  \"data\": {\n    \"name\": \"data\"\n  }\n}", responseString)
		})

		Convey("When sender's stellar.toml does not contain signing key", func() {
			mockStellartomlResolver.On(
				"GetStellarTomlByAddress",
				"alice*stellar.org",
			).Return(stellartoml.StellarToml{}, nil).Once()

			params := url.Values{
				"data": {"{\"Sender\":\"alice*stellar.org\",\"NeedInfo\":false,\"Tx\":\"AAAAAC3/58Z9rycNLmF6voWX9VmDETFVGhFoWf66mcMuir/DAAAAZAAAAAAAAAAAAAAAAAAAAAO5TSe5k00+CKUuUtfafav6xITv43pTgO6QiPes4u/N6QAAAAEAAAAAAAAAAgAAAAFVU0QAAAAAAEbpO2riZmlZMkHuBxUBYAAas3hWyo9VL1IOdnfXAVFBAAAAADuaygAAAAAAGVL83DJFwH0sKmy6AIgJYD7GexiD0YuzSMioBCAUOJwAAAABVVNEAAAAAAAZUvzcMkXAfSwqbLoAiAlgPsZ7GIPRi7NIyKgEIBQ4nAAAAAAL68IAAAAAAgAAAAAAAAABRVVSAAAAAAALt4SwWfv1PIJvDRMenW0zu91YxZbphRFLA4O+gbAaigAAAAA=\",\"Memo\":\"hello world\"}"},
				"sig":  {"bad sig"},
			}

			statusCode, response := net.GetResponse(testServer, params)
			responseString := strings.TrimSpace(string(response))
			assert.Equal(t, 400, statusCode)
			assert.Equal(t, "{\n  \"code\": \"invalid_parameter\",\n  \"message\": \"Invalid parameter.\",\n  \"data\": {\n    \"name\": \"data.sender\"\n  }\n}", responseString)
		})

		Convey("When signature is invalid", func() {
			mockStellartomlResolver.On(
				"GetStellarTomlByAddress",
				"alice*stellar.org",
			).Return(stellartoml.StellarToml{
				SigningKey: "GBYJZW5XFAI6XV73H5SAIUYK6XZI4CGGVBUBO3ANA2SV7KKDAXTV6AEB",
			}, nil).Once()

			params := url.Values{
				"data": {"{\"Sender\":\"alice*stellar.org\",\"NeedInfo\":false,\"Tx\":\"AAAAAC3/58Z9rycNLmF6voWX9VmDETFVGhFoWf66mcMuir/DAAAAZAAAAAAAAAAAAAAAAAAAAAO5TSe5k00+CKUuUtfafav6xITv43pTgO6QiPes4u/N6QAAAAEAAAAAAAAAAgAAAAFVU0QAAAAAAEbpO2riZmlZMkHuBxUBYAAas3hWyo9VL1IOdnfXAVFBAAAAADuaygAAAAAAGVL83DJFwH0sKmy6AIgJYD7GexiD0YuzSMioBCAUOJwAAAABVVNEAAAAAAAZUvzcMkXAfSwqbLoAiAlgPsZ7GIPRi7NIyKgEIBQ4nAAAAAAL68IAAAAAAgAAAAAAAAABRVVSAAAAAAALt4SwWfv1PIJvDRMenW0zu91YxZbphRFLA4O+gbAaigAAAAA=\",\"Memo\":\"hello world\"}"},
				"sig":  {"XIh4u5TcdqUmpy/JLcsTIlD8c7fvJiRC+AwxekjBeOCbtRgE2kzN/8VRQjtKm+zNTt/nuvbM2cfYrs7uu4hnBg=="},
			}

			statusCode, response := net.GetResponse(testServer, params)
			responseString := strings.TrimSpace(string(response))
			assert.Equal(t, 400, statusCode)
			assert.Equal(t, "{\n  \"code\": \"invalid_parameter\",\n  \"message\": \"Invalid parameter.\",\n  \"data\": {\n    \"name\": \"sig\"\n  }\n}", responseString)
		})

		Convey("When all params are valid", func() {
			params := url.Values{
				"data": {"{\"Sender\":\"alice*stellar.org\",\"NeedInfo\":false,\"Tx\":\"AAAAAC3/58Z9rycNLmF6voWX9VmDETFVGhFoWf66mcMuir/DAAAAZAAAAAAAAAAAAAAAAAAAAAO5TSe5k00+CKUuUtfafav6xITv43pTgO6QiPes4u/N6QAAAAEAAAAAAAAAAgAAAAFVU0QAAAAAAEbpO2riZmlZMkHuBxUBYAAas3hWyo9VL1IOdnfXAVFBAAAAADuaygAAAAAAGVL83DJFwH0sKmy6AIgJYD7GexiD0YuzSMioBCAUOJwAAAABVVNEAAAAAAAZUvzcMkXAfSwqbLoAiAlgPsZ7GIPRi7NIyKgEIBQ4nAAAAAAL68IAAAAAAgAAAAAAAAABRVVSAAAAAAALt4SwWfv1PIJvDRMenW0zu91YxZbphRFLA4O+gbAaigAAAAA=\",\"Memo\":\"hello world\"}"},
				"sig":  {"oxfLEvW3dasnYYbVIi98nlXNJZ50yoiS4RAY15g06UhHOENcS2MUzKwLyeKfUtN6r8AFMuX5qXz8LuXgPsZwAg=="},
			}

			mockStellartomlResolver.On(
				"GetStellarTomlByAddress",
				"alice*stellar.org",
			).Return(stellartoml.StellarToml{
				SigningKey: "GBYJZW5XFAI6XV73H5SAIUYK6XZI4CGGVBUBO3ANA2SV7KKDAXTV6AEB",
			}, nil).Once()

			Convey("it returns AuthResponse", func() {
				memo := "uU0nuZNNPgilLlLX2n2r+sSE7+N6U4DukIj3rOLvzek="

				authorizedTransaction := &entities.AuthorizedTransaction{
					TransactionId:  "c18abd4c69315ee7ebbbfd5ee118fb9bd785f9e5f90f7bc06b4583c862333bcd",
					Memo:           &memo,
					TransactionXdr: "AAAAAC3/58Z9rycNLmF6voWX9VmDETFVGhFoWf66mcMuir/DAAAAZAAAAAAAAAAAAAAAAAAAAAO5TSe5k00+CKUuUtfafav6xITv43pTgO6QiPes4u/N6QAAAAEAAAAAAAAAAgAAAAFVU0QAAAAAAEbpO2riZmlZMkHuBxUBYAAas3hWyo9VL1IOdnfXAVFBAAAAADuaygAAAAAAGVL83DJFwH0sKmy6AIgJYD7GexiD0YuzSMioBCAUOJwAAAABVVNEAAAAAAAZUvzcMkXAfSwqbLoAiAlgPsZ7GIPRi7NIyKgEIBQ4nAAAAAAL68IAAAAAAgAAAAAAAAABRVVSAAAAAAALt4SwWfv1PIJvDRMenW0zu91YxZbphRFLA4O+gbAaigAAAAA=",
					Data:           params["data"][0],
				}

				mockEntityManager.On(
					"Persist",
					mock.AnythingOfType("*entities.AuthorizedTransaction"),
				).Run(func(args mock.Arguments) {
					value := args.Get(0).(*entities.AuthorizedTransaction)
					assert.Equal(t, authorizedTransaction.TransactionId, value.TransactionId)
					assert.Equal(t, authorizedTransaction.Memo, value.Memo)
					assert.Equal(t, authorizedTransaction.TransactionXdr, value.TransactionXdr)
					assert.WithinDuration(t, time.Now(), value.AuthorizedAt, 2*time.Second)
					assert.Equal(t, authorizedTransaction.Data, value.Data)
				}).Return(nil).Once()

				statusCode, response := net.GetResponse(testServer, params)
				responseString := strings.TrimSpace(string(response))
				assert.Equal(t, 200, statusCode)
				assert.Equal(t, "{\n  \"info_status\": \"ok\",\n  \"tx_status\": \"ok\"\n}", responseString)
			})
		})
	})

	Convey("Given auth request (sanctions check)", t, func() {
		c.Callbacks = config.Callbacks{
			Sanctions: "http://sanctions",
			AskUser:   "http://ask_user",
			FetchInfo: "http://fetch_info",
		}

		Convey("When all params are valid (NeedInfo = `false`)", func() {
			params := url.Values{
				"data": {"{\"Sender\":\"alice*stellar.org\",\"NeedInfo\":false,\"Tx\":\"AAAAAC3/58Z9rycNLmF6voWX9VmDETFVGhFoWf66mcMuir/DAAAAZAAAAAAAAAAAAAAAAAAAAAO5TSe5k00+CKUuUtfafav6xITv43pTgO6QiPes4u/N6QAAAAEAAAAAAAAAAgAAAAFVU0QAAAAAAEbpO2riZmlZMkHuBxUBYAAas3hWyo9VL1IOdnfXAVFBAAAAADuaygAAAAAAGVL83DJFwH0sKmy6AIgJYD7GexiD0YuzSMioBCAUOJwAAAABVVNEAAAAAAAZUvzcMkXAfSwqbLoAiAlgPsZ7GIPRi7NIyKgEIBQ4nAAAAAAL68IAAAAAAgAAAAAAAAABRVVSAAAAAAALt4SwWfv1PIJvDRMenW0zu91YxZbphRFLA4O+gbAaigAAAAA=\",\"Memo\":\"hello world\"}"},
				"sig":  {"oxfLEvW3dasnYYbVIi98nlXNJZ50yoiS4RAY15g06UhHOENcS2MUzKwLyeKfUtN6r8AFMuX5qXz8LuXgPsZwAg=="},
			}

			mockStellartomlResolver.On(
				"GetStellarTomlByAddress",
				"alice*stellar.org",
			).Return(stellartoml.StellarToml{
				SigningKey: "GBYJZW5XFAI6XV73H5SAIUYK6XZI4CGGVBUBO3ANA2SV7KKDAXTV6AEB",
			}, nil).Once()

			Convey("when sanctions server returns forbidden it returns tx_status `denied`", func() {
				mockHttpClient.On(
					"PostForm",
					"http://sanctions",
					url.Values{"data": params["data"]},
				).Return(
					net.BuildHttpResponse(403, "forbidden"),
					nil,
				).Once()

				statusCode, response := net.GetResponse(testServer, params)
				responseString := strings.TrimSpace(string(response))
				assert.Equal(t, 200, statusCode)
				assert.Equal(t, "{\n  \"info_status\": \"ok\",\n  \"tx_status\": \"denied\"\n}", responseString)
			})

			Convey("when sanctions server returns accepted it returns tx_status `pending`", func() {
				mockHttpClient.On(
					"PostForm",
					"http://sanctions",
					url.Values{"data": params["data"]},
				).Return(
					net.BuildHttpResponse(202, "pending"),
					nil,
				).Once()

				statusCode, response := net.GetResponse(testServer, params)
				responseString := strings.TrimSpace(string(response))
				assert.Equal(t, 200, statusCode)
				assert.Equal(t, "{\n  \"info_status\": \"ok\",\n  \"tx_status\": \"pending\",\n  \"pending\": 600\n}", responseString)
			})

			Convey("when sanctions server returns ok it returns tx_status `ok` and persists transaction", func() {
				mockHttpClient.On(
					"PostForm",
					"http://sanctions",
					url.Values{"data": params["data"]},
				).Return(
					net.BuildHttpResponse(200, "ok"),
					nil,
				).Once()

				memo := "uU0nuZNNPgilLlLX2n2r+sSE7+N6U4DukIj3rOLvzek="

				authorizedTransaction := &entities.AuthorizedTransaction{
					TransactionId:  "c18abd4c69315ee7ebbbfd5ee118fb9bd785f9e5f90f7bc06b4583c862333bcd",
					Memo:           &memo,
					TransactionXdr: "AAAAAC3/58Z9rycNLmF6voWX9VmDETFVGhFoWf66mcMuir/DAAAAZAAAAAAAAAAAAAAAAAAAAAO5TSe5k00+CKUuUtfafav6xITv43pTgO6QiPes4u/N6QAAAAEAAAAAAAAAAgAAAAFVU0QAAAAAAEbpO2riZmlZMkHuBxUBYAAas3hWyo9VL1IOdnfXAVFBAAAAADuaygAAAAAAGVL83DJFwH0sKmy6AIgJYD7GexiD0YuzSMioBCAUOJwAAAABVVNEAAAAAAAZUvzcMkXAfSwqbLoAiAlgPsZ7GIPRi7NIyKgEIBQ4nAAAAAAL68IAAAAAAgAAAAAAAAABRVVSAAAAAAALt4SwWfv1PIJvDRMenW0zu91YxZbphRFLA4O+gbAaigAAAAA=",
					Data:           params["data"][0],
				}

				mockEntityManager.On(
					"Persist",
					mock.AnythingOfType("*entities.AuthorizedTransaction"),
				).Run(func(args mock.Arguments) {
					value := args.Get(0).(*entities.AuthorizedTransaction)
					assert.Equal(t, authorizedTransaction.TransactionId, value.TransactionId)
					assert.Equal(t, authorizedTransaction.Memo, value.Memo)
					assert.Equal(t, authorizedTransaction.TransactionXdr, value.TransactionXdr)
					assert.WithinDuration(t, time.Now(), value.AuthorizedAt, 2*time.Second)
					assert.Equal(t, authorizedTransaction.Data, value.Data)
				}).Return(nil).Once()

				statusCode, response := net.GetResponse(testServer, params)
				responseString := strings.TrimSpace(string(response))
				assert.Equal(t, 200, statusCode)
				assert.Equal(t, "{\n  \"info_status\": \"ok\",\n  \"tx_status\": \"ok\"\n}", responseString)
			})
		})

		Convey("When all params are valid (NeedInfo = `true`)", func() {
			params := url.Values{
				"data": {"{\"Sender\":\"alice*stellar.org\",\"NeedInfo\":true,\"Tx\":\"AAAAAC3/58Z9rycNLmF6voWX9VmDETFVGhFoWf66mcMuir/DAAAAZAAAAAAAAAAAAAAAAAAAAAO5TSe5k00+CKUuUtfafav6xITv43pTgO6QiPes4u/N6QAAAAEAAAAAAAAAAgAAAAFVU0QAAAAAAEbpO2riZmlZMkHuBxUBYAAas3hWyo9VL1IOdnfXAVFBAAAAADuaygAAAAAAGVL83DJFwH0sKmy6AIgJYD7GexiD0YuzSMioBCAUOJwAAAABVVNEAAAAAAAZUvzcMkXAfSwqbLoAiAlgPsZ7GIPRi7NIyKgEIBQ4nAAAAAAL68IAAAAAAgAAAAAAAAABRVVSAAAAAAALt4SwWfv1PIJvDRMenW0zu91YxZbphRFLA4O+gbAaigAAAAA=\",\"Memo\":\"hello world\"}"},
				"sig":  {"XIh4u5TcdqUmpy/JLcsAIlD8c8fvJiRC+AwxekjBeOCbtRgE2kzN/8VRQjtKm+zNTt/nuvbM2cfYrs7uu4hnBg=="},
			}

			mockStellartomlResolver.On(
				"GetStellarTomlByAddress",
				"alice*stellar.org",
			).Return(stellartoml.StellarToml{
				SigningKey: "GBYJZW5XFAI6XV73H5SAIUYK6XZI4CGGVBUBO3ANA2SV7KKDAXTV6AEB",
			}, nil).Once()

			// Make sanctions checks successful (tested previously)
			mockHttpClient.On(
				"PostForm",
				"http://sanctions",
				url.Values{"data": params["data"]},
			).Return(
				net.BuildHttpResponse(200, "ok"),
				nil,
			).Once()

			Convey("when ask_user server returns forbidden it returns info_status `denied`", func() {
				mockHttpClient.On(
					"PostForm",
					"http://ask_user",
					url.Values{"data": params["data"]},
				).Return(
					net.BuildHttpResponse(403, "forbidden"),
					nil,
				).Once()

				statusCode, response := net.GetResponse(testServer, params)
				responseString := strings.TrimSpace(string(response))
				assert.Equal(t, 200, statusCode)
				assert.Equal(t, "{\n  \"info_status\": \"denied\",\n  \"tx_status\": \"ok\"\n}", responseString)
			})

			Convey("when ask_user server returns pending it returns info_status `pending`", func() {
				mockHttpClient.On(
					"PostForm",
					"http://ask_user",
					url.Values{"data": params["data"]},
				).Return(
					net.BuildHttpResponse(202, "pending"),
					nil,
				).Once()

				statusCode, response := net.GetResponse(testServer, params)
				responseString := strings.TrimSpace(string(response))
				assert.Equal(t, 200, statusCode)
				assert.Equal(t, "{\n  \"info_status\": \"pending\",\n  \"tx_status\": \"ok\",\n  \"pending\": 600\n}", responseString)
			})

			Convey("when ask_user server returns ok it returns info_status `ok` and DestInfo and persists transaction", func() {
				mockHttpClient.On(
					"PostForm",
					"http://ask_user",
					url.Values{"data": params["data"]},
				).Return(
					net.BuildHttpResponse(200, "ok"),
					nil,
				).Once()

				mockHttpClient.On(
					"PostForm",
					"http://fetch_info",
					url.Values{"data": params["data"]},
				).Return(
					net.BuildHttpResponse(200, "user data"),
					nil,
				).Once()

				memo := "uU0nuZNNPgilLlLX2n2r+sSE7+N6U4DukIj3rOLvzek="

				authorizedTransaction := &entities.AuthorizedTransaction{
					TransactionId:  "c18abd4c69315ee7ebbbfd5ee118fb9bd785f9e5f90f7bc06b4583c862333bcd",
					Memo:           &memo,
					TransactionXdr: "AAAAAC3/58Z9rycNLmF6voWX9VmDETFVGhFoWf66mcMuir/DAAAAZAAAAAAAAAAAAAAAAAAAAAO5TSe5k00+CKUuUtfafav6xITv43pTgO6QiPes4u/N6QAAAAEAAAAAAAAAAgAAAAFVU0QAAAAAAEbpO2riZmlZMkHuBxUBYAAas3hWyo9VL1IOdnfXAVFBAAAAADuaygAAAAAAGVL83DJFwH0sKmy6AIgJYD7GexiD0YuzSMioBCAUOJwAAAABVVNEAAAAAAAZUvzcMkXAfSwqbLoAiAlgPsZ7GIPRi7NIyKgEIBQ4nAAAAAAL68IAAAAAAgAAAAAAAAABRVVSAAAAAAALt4SwWfv1PIJvDRMenW0zu91YxZbphRFLA4O+gbAaigAAAAA=",
					Data:           params["data"][0],
				}

				mockEntityManager.On(
					"Persist",
					mock.AnythingOfType("*entities.AuthorizedTransaction"),
				).Run(func(args mock.Arguments) {
					value := args.Get(0).(*entities.AuthorizedTransaction)
					assert.Equal(t, authorizedTransaction.TransactionId, value.TransactionId)
					assert.Equal(t, authorizedTransaction.Memo, value.Memo)
					assert.Equal(t, authorizedTransaction.TransactionXdr, value.TransactionXdr)
					assert.WithinDuration(t, time.Now(), value.AuthorizedAt, 2*time.Second)
					assert.Equal(t, authorizedTransaction.Data, value.Data)
				}).Return(nil).Once()

				statusCode, response := net.GetResponse(testServer, params)
				responseString := strings.TrimSpace(string(response))
				assert.Equal(t, 200, statusCode)
				assert.Equal(t, "{\n  \"info_status\": \"ok\",\n  \"tx_status\": \"ok\",\n  \"dest_info\": \"user data\"\n}", responseString)
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

					mockHttpClient.On(
						"PostForm",
						"http://fetch_info",
						url.Values{"data": params["data"]},
					).Return(
						net.BuildHttpResponse(200, "user data"),
						nil,
					).Once()

					memo := "uU0nuZNNPgilLlLX2n2r+sSE7+N6U4DukIj3rOLvzek="

					authorizedTransaction := &entities.AuthorizedTransaction{
						TransactionId:  "c18abd4c69315ee7ebbbfd5ee118fb9bd785f9e5f90f7bc06b4583c862333bcd",
						Memo:           &memo,
						TransactionXdr: "AAAAAC3/58Z9rycNLmF6voWX9VmDETFVGhFoWf66mcMuir/DAAAAZAAAAAAAAAAAAAAAAAAAAAO5TSe5k00+CKUuUtfafav6xITv43pTgO6QiPes4u/N6QAAAAEAAAAAAAAAAgAAAAFVU0QAAAAAAEbpO2riZmlZMkHuBxUBYAAas3hWyo9VL1IOdnfXAVFBAAAAADuaygAAAAAAGVL83DJFwH0sKmy6AIgJYD7GexiD0YuzSMioBCAUOJwAAAABVVNEAAAAAAAZUvzcMkXAfSwqbLoAiAlgPsZ7GIPRi7NIyKgEIBQ4nAAAAAAL68IAAAAAAgAAAAAAAAABRVVSAAAAAAALt4SwWfv1PIJvDRMenW0zu91YxZbphRFLA4O+gbAaigAAAAA=",
						Data:           params["data"][0],
					}

					mockEntityManager.On(
						"Persist",
						mock.AnythingOfType("*entities.AuthorizedTransaction"),
					).Run(func(args mock.Arguments) {
						value := args.Get(0).(*entities.AuthorizedTransaction)
						assert.Equal(t, authorizedTransaction.TransactionId, value.TransactionId)
						assert.Equal(t, authorizedTransaction.Memo, value.Memo)
						assert.Equal(t, authorizedTransaction.TransactionXdr, value.TransactionXdr)
						assert.WithinDuration(t, time.Now(), value.AuthorizedAt, 2*time.Second)
						assert.Equal(t, authorizedTransaction.Data, value.Data)
					}).Return(nil).Once()

					statusCode, response := net.GetResponse(testServer, params)
					responseString := strings.TrimSpace(string(response))
					assert.Equal(t, 200, statusCode)
					assert.Equal(t, "{\n  \"info_status\": \"ok\",\n  \"tx_status\": \"ok\",\n  \"dest_info\": \"user data\"\n}", responseString)
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
						"GetAllowedUserByDomainAndUserId",
						"stellar.org", // sender = `alice*stellar.org`
						"alice",
					).Return(
						&entities.AllowedUser{},
						nil,
					).Once()

					mockHttpClient.On(
						"PostForm",
						"http://fetch_info",
						url.Values{"data": params["data"]},
					).Return(
						net.BuildHttpResponse(200, "user data"),
						nil,
					).Once()

					memo := "uU0nuZNNPgilLlLX2n2r+sSE7+N6U4DukIj3rOLvzek="

					authorizedTransaction := &entities.AuthorizedTransaction{
						TransactionId:  "c18abd4c69315ee7ebbbfd5ee118fb9bd785f9e5f90f7bc06b4583c862333bcd",
						Memo:           &memo,
						TransactionXdr: "AAAAAC3/58Z9rycNLmF6voWX9VmDETFVGhFoWf66mcMuir/DAAAAZAAAAAAAAAAAAAAAAAAAAAO5TSe5k00+CKUuUtfafav6xITv43pTgO6QiPes4u/N6QAAAAEAAAAAAAAAAgAAAAFVU0QAAAAAAEbpO2riZmlZMkHuBxUBYAAas3hWyo9VL1IOdnfXAVFBAAAAADuaygAAAAAAGVL83DJFwH0sKmy6AIgJYD7GexiD0YuzSMioBCAUOJwAAAABVVNEAAAAAAAZUvzcMkXAfSwqbLoAiAlgPsZ7GIPRi7NIyKgEIBQ4nAAAAAAL68IAAAAAAgAAAAAAAAABRVVSAAAAAAALt4SwWfv1PIJvDRMenW0zu91YxZbphRFLA4O+gbAaigAAAAA=",
						Data:           params["data"][0],
					}

					mockEntityManager.On(
						"Persist",
						mock.AnythingOfType("*entities.AuthorizedTransaction"),
					).Run(func(args mock.Arguments) {
						value := args.Get(0).(*entities.AuthorizedTransaction)
						assert.Equal(t, authorizedTransaction.TransactionId, value.TransactionId)
						assert.Equal(t, authorizedTransaction.Memo, value.Memo)
						assert.Equal(t, authorizedTransaction.TransactionXdr, value.TransactionXdr)
						assert.WithinDuration(t, time.Now(), value.AuthorizedAt, 2*time.Second)
						assert.Equal(t, authorizedTransaction.Data, value.Data)
					}).Return(nil).Once()

					statusCode, response := net.GetResponse(testServer, params)
					responseString := strings.TrimSpace(string(response))
					assert.Equal(t, 200, statusCode)
					assert.Equal(t, "{\n  \"info_status\": \"ok\",\n  \"tx_status\": \"ok\",\n  \"dest_info\": \"user data\"\n}", responseString)
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
						"GetAllowedUserByDomainAndUserId",
						"stellar.org", // sender = `alice*stellar.org`
						"alice",
					).Return(
						nil,
						nil,
					).Once()

					statusCode, response := net.GetResponse(testServer, params)
					responseString := strings.TrimSpace(string(response))
					assert.Equal(t, 200, statusCode)
					assert.Equal(t, "{\n  \"info_status\": \"denied\",\n  \"tx_status\": \"ok\"\n}", responseString)
				})
			})
		})
	})
}
