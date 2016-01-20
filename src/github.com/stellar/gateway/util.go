package gateway

import (
	"encoding/json"
	"net/http"
)

func errorServerError(w http.ResponseWriter) {
	http.Error(w, errorResponseString("server_error", "Server error"), http.StatusInternalServerError)
}

func errorForbidden(w http.ResponseWriter, responseString string) {
	http.Error(w, responseString, http.StatusForbidden)
}

func errorBadRequest(w http.ResponseWriter, responseString string) {
	http.Error(w, responseString, http.StatusBadRequest)
}

func errorResponseString(code string, message string) string {
	error := ErrorResponse{Code: code, Message: message}
	json, _ := json.Marshal(error)
	return string(json)
}
