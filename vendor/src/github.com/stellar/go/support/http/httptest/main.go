// Package httptest enhances the stdlib net/http/httptest package by integrating
// it with gopkg.in/gavv/httpexpect.v1, reducing the boilerplate needed for http
// tests
package httptest

import (
	"net/http"
	stdtest "net/http/httptest"
	"testing"

	"gopkg.in/gavv/httpexpect.v1"
)

type Server struct {
	*httpexpect.Expect
	*stdtest.Server
}

func NewServer(t *testing.T, handler http.Handler) *Server {
	server := stdtest.NewServer(handler)
	return &Server{
		Server: server,
		Expect: httpexpect.New(t, server.URL),
	}
}
