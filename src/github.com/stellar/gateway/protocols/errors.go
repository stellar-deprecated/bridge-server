package protocols

import (
	"encoding/json"
	"net/http"
)

var (
	// InternalServerError is an error response
	InternalServerError = &ErrorResponse{Code: "internal_server_error", Message: "Internal Server Error, please try again.", Status: http.StatusInternalServerError}
	// InvalidParameterError  is an error response
	InvalidParameterError = &ErrorResponse{Code: "invalid_parameter", Message: "Invalid parameter.", Status: http.StatusBadRequest}
	// MissingParameterError is an error response
	MissingParameterError = &ErrorResponse{Code: "missing_parameter", Message: "Required parameter is missing.", Status: http.StatusBadRequest}
)

// NewInternalServerError creates and returns a new InternalServerError
func NewInternalServerError(logMessage string, logData map[string]interface{}) *ErrorResponse {
	return &ErrorResponse{
		Status:     InternalServerError.Status,
		Code:       InternalServerError.Code,
		Message:    InternalServerError.Message,
		LogMessage: logMessage,
		LogData:    logData,
	}
}

// NewInvalidParameterError creates and returns a new InvalidParameterError
func NewInvalidParameterError(name, value, moreInfo string, additionalLogData ...map[string]interface{}) *ErrorResponse {
	logData := map[string]interface{}{"name": name, "value": value}
	if len(additionalLogData) == 1 {
		for k, v := range additionalLogData[0] {
			logData[k] = v
		}
	}

	return &ErrorResponse{
		Status:   InvalidParameterError.Status,
		Code:     InvalidParameterError.Code,
		Message:  InvalidParameterError.Message,
		MoreInfo: moreInfo,
		Data:     map[string]interface{}{"name": name},
		LogData:  logData,
	}
}

// NewMissingParameter creates and returns a new MissingParameterError
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

// ErrorResponse represents error response and implements server.Response and error interfaces
type ErrorResponse struct {
	// HTTP status code
	Status int `json:"-"`
	// Error status code
	Code string `json:"code"`
	// Error message that will be returned to API consumer
	Message string `json:"message"`
	// Additional information returned to API consumer
	MoreInfo string `json:"more_info,omitempty"`
	// Error data that will be returned to API consumer
	Data map[string]interface{} `json:"data,omitempty"`
	// Error message that will be logged.
	LogMessage string `json:"-"`
	// Error data that will be logged.
	LogData map[string]interface{} `json:"-"`
}

// Error returns Message or LogMessage if set
func (error *ErrorResponse) Error() string {
	if error.LogMessage != "" {
		return error.LogMessage
	}
	return error.Message
}

// HTTPStatus returns ErrorResponse.Status
func (error *ErrorResponse) HTTPStatus() int {
	return error.Status
}

// Marshal marshals ErrorResponse
func (error *ErrorResponse) Marshal() []byte {
	json, _ := json.MarshalIndent(error, "", "  ")
	return json
}
