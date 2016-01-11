package gateway

import (
	"encoding/base64"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/stellar/go-stellar-base/xdr"
)

func unmarshalTransactionResult(transactionResult string) (txResult xdr.TransactionResult, err error) {
	reader := strings.NewReader(transactionResult)
	b64r := base64.NewDecoder(base64.StdEncoding, reader)
	_, err = xdr.Unmarshal(b64r, &txResult)
	return
}

func errorServerError(w http.ResponseWriter) {
	http.Error(w, errorResponseString("server_error", "Server error"), http.StatusInternalServerError)
}

func errorForbidden(w http.ResponseWriter, responseString string) {
	http.Error(w, responseString, http.StatusForbidden)
}

func errorBadRequest(w http.ResponseWriter, responseString string) {
	http.Error(w, responseString, http.StatusInternalServerError)
}

func errorResponseString(code string, message string) string {
	error := ErrorResponse{Code: code, Message: message}
	json, _ := json.Marshal(error)
	return string(json)
}
