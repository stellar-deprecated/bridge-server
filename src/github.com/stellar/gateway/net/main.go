package net

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
)

// HTTPClientInterface helps mocking http.Client in tests
type HTTPClientInterface interface {
	PostForm(url string, data url.Values) (resp *http.Response, err error)
}

// BuildHTTPResponse is used in tests
func BuildHTTPResponse(statusCode int, body string) *http.Response {
	return &http.Response{
		StatusCode: statusCode,
		Body:       ioutil.NopCloser(bytes.NewBufferString(body)),
	}
}

// GetResponse is used in tests
func GetResponse(testServer *httptest.Server, values url.Values) (int, []byte) {
	res, err := http.PostForm(testServer.URL, values)
	if err != nil {
		panic(err)
	}
	response, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		panic(err)
	}
	return res.StatusCode, response
}
