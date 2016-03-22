package protocols

import (
	"net/http"
)

type SuccessResponse struct{}

func (response *SuccessResponse) HTTPStatus() int {
	return http.StatusOK
}
