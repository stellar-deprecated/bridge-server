package net

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/url"
)

// To be able to mock http.Client in tests
type HttpClientInterface interface {
	PostForm(url string, data url.Values) (resp *http.Response, err error)
}

// Used in tests
func BuildHttpResponse(statusCode int, body string) *http.Response {
	return &http.Response{
		StatusCode: statusCode,
		Body:       ioutil.NopCloser(bytes.NewBufferString(body)),
	}
}
