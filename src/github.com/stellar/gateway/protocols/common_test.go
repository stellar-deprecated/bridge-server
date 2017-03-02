package protocols_test

import (
	"net/http"
	"reflect"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/stellar/gateway/protocols"
	callback "github.com/stellar/gateway/protocols/bridge"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestProtocols(t *testing.T) {
	Convey("FormRequest", t, func() {
		Convey(".ToValues", func() {
			request := &callback.PaymentRequest{
				Source:          "Source",
				Sender:          "Sender",
				Destination:     "Destination",
				Amount:          "Amount",
				AssetCode:       "AssetCode",
				AssetIssuer:     "AssetIssuer",
				SendMax:         "SendMax",
				SendAssetCode:   "SendAssetCode",
				SendAssetIssuer: "SendAssetIssuer",
				UseCompliance:   true,
				ExtraMemo:       "ExtraMemo",
				Path: []protocols.Asset{
					{Code: "USD", Issuer: "BLAH"},
					{},
					{Code: "EUR", Issuer: "BLAH2"},
				},
			}

			values := request.ToValues()
			httpRequest := &http.Request{PostForm: values}
			request.FormRequest.HTTPRequest = httpRequest

			request2 := &callback.PaymentRequest{}
			err := request2.FromRequest(httpRequest)
			require.NoError(t, err)
			assert.True(t, reflect.DeepEqual(request, request2))
		})
	})
}
