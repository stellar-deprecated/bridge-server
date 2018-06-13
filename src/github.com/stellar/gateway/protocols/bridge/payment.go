package bridge

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/stellar/gateway/protocols"
	callback "github.com/stellar/gateway/protocols/compliance"
	"github.com/stellar/go/keypair"
)

var (
	// input errors

	// PaymentCannotResolveDestination is an error response
	PaymentCannotResolveDestination = &protocols.ErrorResponse{Code: "cannot_resolve_destination", Message: "Cannot resolve federated Stellar address.", Status: http.StatusBadRequest}
	// PaymentCannotUseMemo is an error response
	PaymentCannotUseMemo = &protocols.ErrorResponse{Code: "cannot_use_memo", Message: "Memo given in request but federation returned memo fields.", Status: http.StatusBadRequest}
	// PaymentSourceNotExist is an error response
	PaymentSourceNotExist = &protocols.ErrorResponse{Code: "source_not_exist", Message: "Source account does not exist.", Status: http.StatusBadRequest}
	// PaymentAssetCodeNotAllowed is an error response
	PaymentAssetCodeNotAllowed = &protocols.ErrorResponse{Code: "asset_code_not_allowed", Message: "Given asset_code not allowed.", Status: http.StatusBadRequest}

	// compliance

	// PaymentPending is an error response
	PaymentPending = &protocols.ErrorResponse{Code: "pending", Message: "Transaction pending. Repeat your request after given time.", Status: http.StatusAccepted}
	// PaymentDenied is an error response
	PaymentDenied = &protocols.ErrorResponse{Code: "denied", Message: "Transaction denied by destination.", Status: http.StatusForbidden}

	// payment op errors

	// PaymentMalformed is an error response
	PaymentMalformed = &protocols.ErrorResponse{Code: "payment_malformed", Message: "Operation is malformed.", Status: http.StatusBadRequest}
	// PaymentUnderfunded is an error response
	PaymentUnderfunded = &protocols.ErrorResponse{Code: "payment_underfunded", Message: "Not enough funds to send this transaction.", Status: http.StatusBadRequest}
	// PaymentSrcNoTrust is an error response
	PaymentSrcNoTrust = &protocols.ErrorResponse{Code: "payment_src_no_trust", Message: "No trustline on source account.", Status: http.StatusBadRequest}
	// PaymentSrcNotAuthorized is an error response
	PaymentSrcNotAuthorized = &protocols.ErrorResponse{Code: "payment_src_not_authorized", Message: "Source not authorized to transfer.", Status: http.StatusBadRequest}
	// PaymentNoDestination is an error response
	PaymentNoDestination = &protocols.ErrorResponse{Code: "payment_no_destination", Message: "Destination account does not exist.", Status: http.StatusBadRequest}
	// PaymentNoTrust is an error response
	PaymentNoTrust = &protocols.ErrorResponse{Code: "payment_no_trust", Message: "Destination missing a trust line for asset.", Status: http.StatusBadRequest}
	// PaymentNotAuthorized is an error response
	PaymentNotAuthorized = &protocols.ErrorResponse{Code: "payment_not_authorized", Message: "Destination not authorized to trust asset. It needs to be allowed first by using /authorize endpoint.", Status: http.StatusBadRequest}
	// PaymentLineFull is an error response
	PaymentLineFull = &protocols.ErrorResponse{Code: "payment_line_full", Message: "Sending this payment would make a destination go above their limit.", Status: http.StatusBadRequest}
	// PaymentNoIssuer is an error response
	PaymentNoIssuer = &protocols.ErrorResponse{Code: "payment_no_issuer", Message: "Missing issuer on asset.", Status: http.StatusBadRequest}
	// PaymentTooFewOffers is an error response
	PaymentTooFewOffers = &protocols.ErrorResponse{Code: "payment_too_few_offers", Message: "Not enough offers to satisfy path.", Status: http.StatusBadRequest}
	// PaymentOfferCrossSelf is an error response
	PaymentOfferCrossSelf = &protocols.ErrorResponse{Code: "payment_offer_cross_self", Message: "would cross one of its own offers.", Status: http.StatusBadRequest}
	// PaymentOverSendmax is an error response
	PaymentOverSendmax = &protocols.ErrorResponse{Code: "payment_over_sendmax", Message: "Could not satisfy sendmax.", Status: http.StatusBadRequest}
)

// PaymentRequest represents request made to /payment endpoint of the bridge server
type PaymentRequest struct {
	// Payment ID
	ID string `name:"id"`
	// Source account secret
	Source string `name:"source"`
	// Sender address (like alice*stellar.org)
	Sender string `name:"sender"`
	// Destination address (like bob*stellar.org)
	Destination string `name:"destination"`
	// ForwardDestination
	ForwardDestination *protocols.ForwardDestination `name:"forward_destination"`
	// Memo type
	MemoType string `name:"memo_type"`
	// Memo value
	Memo string `name:"memo"`
	// Amount destination should receive
	Amount string `name:"amount" required:""`
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
	// Determined whether to use compliance protocol or to send a simple payment.
	UseCompliance bool `name:"use_compliance"`
	// Extra memo. If set, UseCompliance value will be ignored and it will use compliance.
	ExtraMemo string `name:"extra_memo"`

	protocols.FormRequest
}

// FromRequest will populate request fields using http.Request.
func (request *PaymentRequest) FromRequest(r *http.Request) error {
	return request.FormRequest.FromRequest(r, request)
}

// ToValues will create url.Values from request.
func (request *PaymentRequest) ToValues() url.Values {
	return request.FormRequest.ToValues(request)
}

// ToComplianceSendRequest transforms PaymentRequest to callback.SendRequest
func (request *PaymentRequest) ToComplianceSendRequest() callback.SendRequest {
	sourceKeypair, _ := keypair.Parse(request.Source)
	return callback.SendRequest{
		ID: request.ID,
		// Compliance does not sign the transaction, it just needs a public key
		Source:             sourceKeypair.Address(),
		Sender:             request.Sender,
		Destination:        request.Destination,
		ForwardDestination: request.ForwardDestination,
		Amount:             request.Amount,
		AssetCode:          request.AssetCode,
		AssetIssuer:        request.AssetIssuer,
		SendMax:            request.SendMax,
		SendAssetCode:      request.SendAssetCode,
		SendAssetIssuer:    request.SendAssetIssuer,
		Path:               request.Path,
		ExtraMemo:          request.ExtraMemo,
	}
}

// Validate validates if request fields are valid. Useful when checking if a request is correct.
func (request *PaymentRequest) Validate() error {
	err := request.FormRequest.CheckRequired(request)
	if err != nil {
		return err
	}

	if request.Source != "" {
		_, err = keypair.Parse(request.Source)
		if err != nil {
			return protocols.NewInvalidParameterError("source", request.Source, "Source must be a public key (starting with `G`).")
		}
	}

	if request.Destination == "" && request.ForwardDestination == nil {
		return protocols.NewMissingParameter("destination")
	}

	if !protocols.IsValidAmount(request.Amount) {
		return protocols.NewInvalidParameterError("amount", request.Amount, "Invalid amount.")
	}

	if request.SendMax != "" {
		if !protocols.IsValidAmount(request.SendMax) {
			return protocols.NewInvalidParameterError("send_max", request.SendMax, "Invalid amount.")
		}
	}

	// Memo
	if request.MemoType == "" && request.Memo != "" {
		return protocols.NewMissingParameter("memo_type")
	}

	if request.MemoType != "" && request.Memo == "" {
		return protocols.NewMissingParameter("memo")
	}

	// Destination Asset
	if request.AssetCode == "" && request.AssetIssuer != "" {
		return protocols.NewMissingParameter("asset_code")
	}

	if request.AssetCode != "" && request.AssetIssuer == "" {
		return protocols.NewMissingParameter("asset_issuer")
	}

	destinationAsset := protocols.Asset{
		Code:   request.AssetCode,
		Issuer: request.AssetIssuer,
	}

	if !destinationAsset.Validate() {
		return protocols.NewInvalidParameterError("asset", destinationAsset.String(), "Invalid asset.")
	}

	if request.AssetIssuer != "" {
		if !protocols.IsValidAccountID(request.AssetIssuer) {
			return protocols.NewInvalidParameterError("asset_issuer", request.AssetIssuer, "Asset issuer must be a public key (starting with `G`).")
		}
	}

	// Send Asset
	if request.SendAssetCode == "" && request.SendAssetIssuer != "" {
		return protocols.NewMissingParameter("send_asset_code")
	}

	if request.SendAssetCode != "" && request.SendAssetIssuer == "" {
		return protocols.NewMissingParameter("send_asset_issuer")
	}

	sendAsset := protocols.Asset{
		Code:   request.AssetCode,
		Issuer: request.AssetIssuer,
	}

	if !sendAsset.Validate() {
		return protocols.NewInvalidParameterError("asset", sendAsset.String(), "Invalid asset.")
	}

	if request.SendAssetIssuer != "" {
		if !protocols.IsValidAccountID(request.SendAssetIssuer) {
			return protocols.NewInvalidParameterError("send_asset_issuer", request.SendAssetIssuer, "Send asset issuer must be a public key (starting with `G`).")
		}
	}

	return nil
}

func validateStellarAddress(address string) bool {
	tokens := strings.Split(address, "*")
	return len(tokens) == 2
}

// NewPaymentPendingError creates a new PaymentPending error
func NewPaymentPendingError(seconds int) *protocols.ErrorResponse {
	return &protocols.ErrorResponse{
		Status:  PaymentPending.Status,
		Code:    PaymentPending.Code,
		Message: PaymentPending.Message,
		Data:    map[string]interface{}{"pending": seconds},
	}
}
