package handlers

import (
	"encoding/json"
	"net/http"
)

var (
	InternalServerError = &ErrorResponse{Code: "internal_server_error", Message: "Internal Server Error, please try again.", Status: http.StatusInternalServerError}
	InvalidParameter    = &ErrorResponse{Code: "invalid_parameter", Message: "Invalid parameter.", Status: http.StatusBadRequest}
)

func WriteError(w http.ResponseWriter, error *ErrorResponse) {
	w.WriteHeader(error.Status)
	w.Write(error.Marshal())
}

type ErrorResponse struct {
	Status  int    `json:"-"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (response *ErrorResponse) Marshal() []byte {
	json, _ := json.MarshalIndent(response, "", "  ")
	return json
}
