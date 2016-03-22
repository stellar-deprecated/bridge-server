package server

import (
	"net/http"
)

type Response interface {
	HTTPStatus() int
	Marshal() []byte
}

func Write(w http.ResponseWriter, response Response) {
	if response.HTTPStatus() != 200 {
		w.WriteHeader(response.HTTPStatus())
	}
	w.Write(response.Marshal())
}
