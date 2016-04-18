package compliance

import (
	"net/http"

	"github.com/stellar/gateway/protocols"
)

var (
	// /receive

	// TransactionNotFoundError is an error response
	TransactionNotFoundError = &protocols.ErrorResponse{Code: "transaction_not_found", Message: "Transaction not found.", Status: http.StatusNotFound}

	// /send

	// CannotResolveDestination is an error response
	CannotResolveDestination = &protocols.ErrorResponse{Code: "cannot_resolve_destination", Message: "Cannot resolve federated Stellar address.", Status: http.StatusBadRequest}
	// AuthServerNotDefined is an error response
	AuthServerNotDefined = &protocols.ErrorResponse{Code: "auth_server_not_defined", Message: "No AUTH_SERVER defined in stellar.toml file.", Status: http.StatusBadRequest}
)
