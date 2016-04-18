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
	"github.com/stellar/gateway/protocols/compliance"
	"github.com/stretchr/testify/assert"
	"github.com/zenazn/goji/web"
)

func TestRequestHandlerReceive(t *testing.T) {
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
		requestHandler.HandlerReceive(web.C{}, w, r)
	}

	testServer := httptest.NewServer(http.HandlerFunc(httpHandle))
	defer testServer.Close()

	Convey("Given receive request", t, func() {
		Convey("it returns TransactionNotFoundError when memo not found", func() {
			memo := "907ba78b4545338d3539683e63ecb51cf51c10adc9dabd86e92bd52339f298b9"
			params := url.Values{"memo": {memo}}

			mockRepository.On("GetAuthorizedTransactionByMemo", memo).Return(nil, nil).Once()

			statusCode, response := net.GetResponse(testServer, params)
			responseString := strings.TrimSpace(string(response))
			assert.Equal(t, 404, statusCode)
			assert.Equal(t, compliance.TransactionNotFoundError.Marshal(), []byte(responseString))
		})

		Convey("it returns preimage when memo has been found", func() {
			memo := "bcc649cfdb8cc557053da67df7e7fcb740dcf7f721cebe1f2082597ad0d5e7d8"
			params := url.Values{"memo": {memo}}

			authorizedTransaction := entities.AuthorizedTransaction{
				Memo: memo,
				Data: "hello world",
			}

			mockRepository.On("GetAuthorizedTransactionByMemo", memo).Return(
				&authorizedTransaction,
				nil,
			).Once()

			statusCode, response := net.GetResponse(testServer, params)
			responseString := strings.TrimSpace(string(response))
			assert.Equal(t, 200, statusCode)
			assert.Equal(t, "{\n  \"memo\": \"hello world\"\n}", responseString)
		})
	})
}
