package compliance

import (
	"encoding/json"
	"net/http"
)

var (
	InternalServerError   = &ErrorResponse{Code: "internal_server_error", Message: "Internal Server Error, please try again.", Status: http.StatusInternalServerError}
	InvalidParameterError = &ErrorResponse{Code: "invalid_parameter", Message: "Invalid parameter.", Status: http.StatusBadRequest}
	MissingParameter      = &ErrorResponse{Code: "missing_parameter", Message: "One of required parameters is missing.", Status: http.StatusBadRequest}

	// /receive
	TransactionNotFoundError = &ErrorResponse{Code: "transaction_not_found", Message: "Transaction not found.", Status: http.StatusNotFound}

	// /send
	CannotResolveDestination = &ErrorResponse{Code: "cannot_resolve_destination", Message: "Cannot resolve federated Stellar address.", Status: http.StatusBadRequest}
	AuthServerNotDefined     = &ErrorResponse{Code: "auth_server_not_defined", Message: "No AUTH_SERVER defined in stellar.toml file.", Status: http.StatusBadRequest}
)

type ErrorResponse struct {
	Status    int    `json:"-"`
	Code      string `json:"code"`
	Message   string `json:"message"`
	Parameter string `json:"parameter,omitempty"`
}

func (error *ErrorResponse) HTTPStatus() int {
	return error.Status
}

func (error *ErrorResponse) Marshal() []byte {
	json, _ := json.MarshalIndent(error, "", "  ")
	return json
}
