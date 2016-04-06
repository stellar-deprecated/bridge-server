package protocols

import (
	"encoding/json"
	"net/http"
)

var (
	InternalServerError   = &ErrorResponse{Code: "internal_server_error", Message: "Internal Server Error, please try again.", Status: http.StatusInternalServerError}
	InvalidParameterError = &ErrorResponse{Code: "invalid_parameter", Message: "Invalid parameter.", Status: http.StatusBadRequest}
	MissingParameterError = &ErrorResponse{Code: "missing_parameter", Message: "Required parameter is missing.", Status: http.StatusBadRequest}
)

func NewInternalServerError(logMessage string, logData map[string]interface{}) *ErrorResponse {
	return &ErrorResponse{
		Status:     InternalServerError.Status,
		Code:       InternalServerError.Code,
		Message:    InternalServerError.Message,
		LogMessage: logMessage,
		LogData:    logData,
	}
}

func NewInvalidParameterError(name, value string) *ErrorResponse {
	return &ErrorResponse{
		Status:  InvalidParameterError.Status,
		Code:    InvalidParameterError.Code,
		Message: InvalidParameterError.Message,
		Data:    map[string]interface{}{"name": name},
		LogData: map[string]interface{}{"name": name, "value": value},
	}
}

func NewMissingParameter(name string) *ErrorResponse {
	data := map[string]interface{}{"name": name}
	return &ErrorResponse{
		Status:  MissingParameterError.Status,
		Code:    MissingParameterError.Code,
		Message: MissingParameterError.Message,
		Data:    data,
		LogData: data,
	}
}

type ErrorResponse struct {
	// HTTP status code
	Status int `json:"-"`
	// Error status code
	Code string `json:"code"`
	// Error message that will be returned to API consumer
	Message string `json:"message"`
	// Error data that will be returned to API consumer
	Data map[string]interface{} `json:"data,omitempty"`
	// Error message that will be logged.
	LogMessage string `json:"-"`
	// Error data that will be logged.
	LogData map[string]interface{} `json:"-"`
}

func (error *ErrorResponse) Error() string {
	if error.LogMessage != "" {
		return error.LogMessage
	}
	return error.Message
}

func (error *ErrorResponse) HTTPStatus() int {
	return error.Status
}

func (error *ErrorResponse) Marshal() []byte {
	json, _ := json.MarshalIndent(error, "", "  ")
	return json
}
