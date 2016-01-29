package handlers

import (
	"encoding/json"
	"net/http"
)

func errorServerError(w http.ResponseWriter) {
	http.Error(w, getServerErrorResponseString(), http.StatusInternalServerError)
}

func errorForbidden(w http.ResponseWriter, responseString string) {
	http.Error(w, responseString, http.StatusForbidden)
}

func errorBadRequest(w http.ResponseWriter, responseString string) {
	http.Error(w, responseString, http.StatusBadRequest)
}

func errorResponseString(code string, message string) string {
	error := ErrorResponse{Code: code, Message: message}
	json, _ := json.MarshalIndent(error, "", "  ")
	return string(json)
}

func getServerErrorResponseString() string {
	return errorResponseString("server_error", "Server error")
}
