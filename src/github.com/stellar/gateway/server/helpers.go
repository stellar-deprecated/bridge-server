package server

import (
	"net/http"
)

// Response represents response that can be returned by a server
type Response interface {
	HTTPStatus() int
	Marshal() []byte
}

// Write writes a response to the given http.ResponseWriter
func Write(w http.ResponseWriter, response Response) {
	if response.HTTPStatus() != 200 {
		w.WriteHeader(response.HTTPStatus())
	}
	w.Write(response.Marshal())
}
