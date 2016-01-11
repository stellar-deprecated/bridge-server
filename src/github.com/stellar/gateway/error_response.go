package gateway

type ErrorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}
