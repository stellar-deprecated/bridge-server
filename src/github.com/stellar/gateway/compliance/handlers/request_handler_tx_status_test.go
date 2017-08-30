package handlers

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/facebookgo/inject"
	"github.com/goji/httpauth"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/stellar/gateway/compliance/config"
	"github.com/stellar/gateway/mocks"
	"github.com/stellar/gateway/net"
	"github.com/stellar/go/support/http/httptest"
)

func TestRequestHandlerTxStatus(t *testing.T) {
	c := &config.Config{
		NetworkPassphrase: "Test SDF Network ; September 2015",
		Keys: config.Keys{
			// GBYJZW5XFAI6XV73H5SAIUYK6XZI4CGGVBUBO3ANA2SV7KKDAXTV6AEB
			SigningSeed: "SDWTLFPALQSP225BSMX7HPZ7ZEAYSUYNDLJ5QI3YGVBNRUIIELWH3XUV",
		},
	}
	c.TxStatusAuth.Username = "username"
	c.TxStatusAuth.Password = "password"

	txid := "abc123"
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
		requestHandler.HandlerTxStatus(w, r)
	}

	testServer := httptest.NewServer(t, httpauth.SimpleBasicAuth(c.TxStatusAuth.Username,
		c.TxStatusAuth.Password)(http.HandlerFunc(httpHandle)))
	defer testServer.Close()

	Convey("Given tx_status request", t, func() {
		Convey("it returns unathorised when no auth", func() {
			testServer.GET("/tx_status").
				Expect().
				Status(http.StatusUnauthorized)
		})
		Convey("it returns unauthorised when bad auth", func() {
			testServer.GET("/tx_status").
				WithBasicAuth("username", "wrong_password").
				Expect().
				Status(http.StatusUnauthorized)
		})
		Convey("it returns bad request when no parameter", func() {
			testServer.GET("/tx_status").
				WithBasicAuth("username", "password").
				Expect().
				Status(http.StatusBadRequest)
		})
		Convey("it returns unknown when no tx_status endpoint", func() {
			testServer.GET("/tx_status").
				WithBasicAuth("username", "password").
				WithQuery("id", "123").
				Expect().
				Status(http.StatusOK).
				Body().Equal(`{"status":"unknown"}` + "\n")
		})
		Convey("it returns unknown when valid endpoint returns bad request", func() {
			c.Callbacks = config.Callbacks{
				TxStatus: "http://tx_status",
			}

			mockHTTPClient.On(
				"PostForm",
				"http://tx_status",
				url.Values{"id": {txid}},
			).Return(
				net.BuildHTTPResponse(400, "badrequest"),
				nil,
			).Once()

			testServer.GET("/tx_status").
				WithBasicAuth("username", "password").
				WithQuery("id", txid).
				Expect().
				Status(http.StatusOK).
				Body().Equal(`{"status":"unknown"}` + "\n")
		})

		Convey("it returns unknown when valid endpoint returns empty data", func() {
			c.Callbacks = config.Callbacks{
				TxStatus: "http://tx_status",
			}

			mockHTTPClient.On(
				"PostForm",
				"http://tx_status",
				url.Values{"id": {txid}},
			).Return(
				net.BuildHTTPResponse(200, "{}"),
				nil,
			).Once()

			testServer.GET("/tx_status").
				WithBasicAuth("username", "password").
				WithQuery("id", txid).
				Expect().
				Status(http.StatusOK).
				Body().Equal(`{"status":"unknown"}` + "\n")
		})

		Convey("it returns response from valid endpoint with data", func() {
			c.Callbacks = config.Callbacks{
				TxStatus: "http://tx_status",
			}

			mockHTTPClient.On(
				"PostForm",
				"http://tx_status",
				url.Values{"id": {txid}},
			).Return(
				net.BuildHTTPResponse(200, `{"status":"delivered","msg":"cash paid"}`),
				nil,
			).Once()

			testServer.GET("/tx_status").
				WithBasicAuth("username", "password").
				WithQuery("id", txid).
				Expect().
				Status(http.StatusOK).
				Body().Equal(`{"status":"delivered","msg":"cash paid"}` + "\n")
		})

	})
}
