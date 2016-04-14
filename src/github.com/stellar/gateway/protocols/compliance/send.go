package compliance

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"

	"github.com/stellar/gateway/protocols"
	"github.com/stellar/go-stellar-base/keypair"
)

type SendRequest struct {
	// Source account ID
	Source string `name:"source" required:""`
	// Sender address (like alice*stellar.org)
	Sender string `name:"sender" required:""`
	// Destination address (like bob*stellar.org)
	Destination string `name:"destination" required:""`
	// Amount destination should receive
	Amount string `name:"amount" required:""`
	// Code of the asset destination should receive
	AssetCode string `name:"asset_code" required:""`
	// Issuer of the asset destination should receive
	AssetIssuer string `name:"asset_issuer" required:""`
	// Only for path_payment
	SendMax string `name:"send_max"`
	// Only for path_payment
	SendAssetCode string `name:"send_asset_code"`
	// Only for path_payment
	SendAssetIssuer string `name:"send_asset_issuer"`
	// path[n][asset_code] path[n][asset_issuer]
	Path []protocols.Asset `name:"path"`
	// Extra memo
	ExtraMemo string `name:"extra_memo" required:""`

	protocols.FormRequest
}

// Will populate request fields using http.Request.
func (request *SendRequest) FromRequest(r *http.Request) {
	request.FormRequest.FromRequest(r, request)
}

// Will create url.Values from request.
func (request *SendRequest) ToValues() url.Values {
	return request.FormRequest.ToValues(request)
}

// Validates if request fields are valid. Useful when checking if a request is correct.
func (request *SendRequest) Validate() error {
	err := request.FormRequest.CheckRequired(request)
	if err != nil {
		return err
	}

	_, err = keypair.Parse(request.Source)
	if err != nil {
		return protocols.NewInvalidParameterError("source", request.Source)
	}

	if !validateStellarAddress(request.Sender) {
		return protocols.NewInvalidParameterError("sender", request.Sender)
	}

	if !validateStellarAddress(request.Destination) {
		return protocols.NewInvalidParameterError("destination", request.Destination)
	}

	_, err = keypair.Parse(request.AssetIssuer)
	if err != nil {
		return protocols.NewInvalidParameterError("asset_issuer", request.AssetIssuer)
	}

	return nil
}

func validateStellarAddress(address string) bool {
	tokens := strings.Split(address, "*")
	return len(tokens) == 2
}

type SendResponse struct {
	protocols.SuccessResponse
	AuthResponse `json:"auth_response"`
	// xdr.Transaction base64-encoded. Sequence number of this transaction will be equal 0.
	TransactionXdr string `json:"transaction_xdr,omitempty"`
}

func (response *SendResponse) Marshal() []byte {
	json, _ := json.MarshalIndent(response, "", "  ")
	return json
}
