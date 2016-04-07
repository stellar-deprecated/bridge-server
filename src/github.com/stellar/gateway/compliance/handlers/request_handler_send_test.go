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
	"github.com/stellar/gateway/mocks"
	"github.com/stellar/gateway/net"
	"github.com/stellar/gateway/protocols/federation"
	"github.com/stellar/gateway/protocols/stellartoml"
	"github.com/stretchr/testify/assert"
	"github.com/zenazn/goji/web"
)

func TestRequestHandlerSend(t *testing.T) {
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
		requestHandler.HandlerSend(web.C{}, w, r)
	}

	testServer := httptest.NewServer(http.HandlerFunc(httpHandle))
	defer testServer.Close()

	Convey("Given send request", t, func() {
		Convey("When source param is missing", func() {
			statusCode, response := net.GetResponse(testServer, url.Values{})
			responseString := strings.TrimSpace(string(response))
			assert.Equal(t, 400, statusCode)
			assert.Equal(t, "{\n  \"code\": \"missing_parameter\",\n  \"message\": \"Required parameter is missing.\",\n  \"data\": {\n    \"name\": \"source\"\n  }\n}", responseString)
		})

		Convey("When source param is invalid", func() {
			params := url.Values{
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
			assert.Equal(t, "{\n  \"code\": \"invalid_parameter\",\n  \"message\": \"Invalid parameter.\",\n  \"data\": {\n    \"name\": \"source\"\n  }\n}", responseString)
		})

		Convey("When params are valid", func() {
			params := url.Values{
				"source":       {"GAW77Z6GPWXSODJOMF5L5BMX6VMYGEJRKUNBC2CZ725JTQZORK74HQQD"},
				"sender":       {"alice*stellar.org"}, // GAW77Z6GPWXSODJOMF5L5BMX6VMYGEJRKUNBC2CZ725JTQZORK74HQQD
				"destination":  {"bob*stellar.org"},   // GAMVF7G4GJC4A7JMFJWLUAEIBFQD5RT3DCB5DC5TJDEKQBBACQ4JZVEE
				"amount":       {"20"},
				"asset_code":   {"USD"},
				"asset_issuer": {"GAMVF7G4GJC4A7JMFJWLUAEIBFQD5RT3DCB5DC5TJDEKQBBACQ4JZVEE"},
				"extra_memo":   {"hello world"},
			}

			Convey("it returns SendResponse when success (payment)", func() {
				authServer := "https://acme.com/auth"

				mockFederationResolver.On(
					"Resolve",
					"bob*stellar.org",
				).Return(federation.Response{
					AccountId: "GAMVF7G4GJC4A7JMFJWLUAEIBFQD5RT3DCB5DC5TJDEKQBBACQ4JZVEE",
				}, stellartoml.StellarToml{
					AuthServer: authServer,
				}, nil).Once()

				transactionXdr := "AAAAAC3/58Z9rycNLmF6voWX9VmDETFVGhFoWf66mcMuir/DAAAAZAAAAAAAAAAAAAAAAAAAAAO5TSe5k00+CKUuUtfafav6xITv43pTgO6QiPes4u/N6QAAAAEAAAAAAAAAAQAAAAAZUvzcMkXAfSwqbLoAiAlgPsZ7GIPRi7NIyKgEIBQ4nAAAAAFVU0QAAAAAABlS/NwyRcB9LCpsugCICWA+xnsYg9GLs0jIqAQgFDicAAAAAAvrwgAAAAAA"
				data := "{\"Sender\":\"alice*stellar.org\",\"NeedInfo\":false,\"Tx\":\"" + transactionXdr + "\",\"Memo\":\"hello world\"}"
				sig := "XG5RkV+Cd+cdzBo/tKoylJmmwaP+YbL0DBdCH0ulL+nJot1QEy+EFm/bQ3ZGagU8sURMKCQunX3UHAak2h4ACQ=="

				mockHttpClient.On(
					"PostForm",
					authServer,
					url.Values{"data": {data}, "sig": {sig}},
				).Return(
					net.BuildHttpResponse(200, "ok"),
					nil,
				).Once()

				statusCode, response := net.GetResponse(testServer, params)
				responseString := strings.TrimSpace(string(response))
				assert.Equal(t, 200, statusCode)
				assert.Equal(t, "{\n  \"transaction_xdr\": \"AAAAAC3/58Z9rycNLmF6voWX9VmDETFVGhFoWf66mcMuir/DAAAAZAAAAAAAAAAAAAAAAAAAAAO5TSe5k00+CKUuUtfafav6xITv43pTgO6QiPes4u/N6QAAAAEAAAAAAAAAAQAAAAAZUvzcMkXAfSwqbLoAiAlgPsZ7GIPRi7NIyKgEIBQ4nAAAAAFVU0QAAAAAABlS/NwyRcB9LCpsugCICWA+xnsYg9GLs0jIqAQgFDicAAAAAAvrwgAAAAAA\"\n}", responseString)
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
					"Resolve",
					"bob*stellar.org",
				).Return(federation.Response{
					AccountId: "GAMVF7G4GJC4A7JMFJWLUAEIBFQD5RT3DCB5DC5TJDEKQBBACQ4JZVEE",
				}, stellartoml.StellarToml{
					AuthServer: authServer,
				}, nil).Once()

				transactionXdr := "AAAAAC3/58Z9rycNLmF6voWX9VmDETFVGhFoWf66mcMuir/DAAAAZAAAAAAAAAAAAAAAAAAAAAO5TSe5k00+CKUuUtfafav6xITv43pTgO6QiPes4u/N6QAAAAEAAAAAAAAAAgAAAAFVU0QAAAAAAEbpO2riZmlZMkHuBxUBYAAas3hWyo9VL1IOdnfXAVFBAAAAADuaygAAAAAAGVL83DJFwH0sKmy6AIgJYD7GexiD0YuzSMioBCAUOJwAAAABVVNEAAAAAAAZUvzcMkXAfSwqbLoAiAlgPsZ7GIPRi7NIyKgEIBQ4nAAAAAAL68IAAAAAAgAAAAAAAAABRVVSAAAAAAALt4SwWfv1PIJvDRMenW0zu91YxZbphRFLA4O+gbAaigAAAAA="
				data := "{\"Sender\":\"alice*stellar.org\",\"NeedInfo\":false,\"Tx\":\"" + transactionXdr + "\",\"Memo\":\"hello world\"}"
				sig := "oxfLEvW3dasnYYbVIi98nlXNJZ50yoiS4RAY15g06UhHOENcS2MUzKwLyeKfUtN6r8AFMuX5qXz8LuXgPsZwAg=="

				mockHttpClient.On(
					"PostForm",
					authServer,
					url.Values{"data": {data}, "sig": {sig}},
				).Return(
					net.BuildHttpResponse(200, "ok"),
					nil,
				).Once()

				statusCode, response := net.GetResponse(testServer, params)
				responseString := strings.TrimSpace(string(response))
				assert.Equal(t, 200, statusCode)
				assert.Equal(t, "{\n  \"transaction_xdr\": \"AAAAAC3/58Z9rycNLmF6voWX9VmDETFVGhFoWf66mcMuir/DAAAAZAAAAAAAAAAAAAAAAAAAAAO5TSe5k00+CKUuUtfafav6xITv43pTgO6QiPes4u/N6QAAAAEAAAAAAAAAAgAAAAFVU0QAAAAAAEbpO2riZmlZMkHuBxUBYAAas3hWyo9VL1IOdnfXAVFBAAAAADuaygAAAAAAGVL83DJFwH0sKmy6AIgJYD7GexiD0YuzSMioBCAUOJwAAAABVVNEAAAAAAAZUvzcMkXAfSwqbLoAiAlgPsZ7GIPRi7NIyKgEIBQ4nAAAAAAL68IAAAAAAgAAAAAAAAABRVVSAAAAAAALt4SwWfv1PIJvDRMenW0zu91YxZbphRFLA4O+gbAaigAAAAA=\"\n}", responseString)
			})
		})
	})
}
