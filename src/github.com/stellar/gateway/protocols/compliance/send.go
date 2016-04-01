package compliance

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/stellar/gateway/protocols"
)

type SendRequest struct {
	// Source account ID
	Source string `name:"source"`
	// Sender address (like alice*stellar.org)
	Sender string `name:"sender"`
	// Destination address (like bob*stellar.org)
	Destination string `name:"destination"`
	// Amount destination should receive
	Amount string `name:"amount"`
	// Code of the asset destination should receive
	AssetCode string `name:"asset_code"`
	// Issuer of the asset destination should receive
	AssetIssuer string `name:"asset_issuer"`
	// Only for path_payment
	SendMax string `name:"send_max"`
	// Only for path_payment
	SendAssetCode string `name:"send_asset_code"`
	// Only for path_payment
	SendAssetIssuer string `name:"send_asset_issuer"`
	// path[n][asset_code] path[n][asset_issuer]
	Path []protocols.Asset `name:"path"`
	// Extra memo should contain sender KYC data
	ExtraMemo string `name:"extra_memo"`

	formRequest protocols.FormRequest
}

// Will populate SendRequest fields using http.Request.
func (request *SendRequest) FromRequest(r *http.Request) {
	request.formRequest.FromRequest(r, request)
}

// Will create url.Values from SendRequest.
func (request *SendRequest) ToValues() url.Values {
	return request.formRequest.ToValues(request)
}

// Validates if SendRequest fields are valid. Useful when checking if a request is correct.
func (request *SendRequest) Validate() {
	// TODO
}

type SendResponse struct {
	protocols.SuccessResponse
	// xdr.Transaction base64-encoded. Sequence number of this transaction will be equal 0.
	TransactionXdr string `json:"transaction_xdr"`
}

func (response *SendResponse) Marshal() []byte {
	json, _ := json.MarshalIndent(response, "", "  ")
	return json
}
