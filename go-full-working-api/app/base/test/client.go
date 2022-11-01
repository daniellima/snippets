package test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
)

type Client struct {
	serveMux *http.ServeMux
}

func NewClient(serveMux *http.ServeMux) *Client {
	return &Client{serveMux}
}

func (this *Client) requestForTest(method string, path string, body io.Reader) *httptest.ResponseRecorder {
	request := httptest.NewRequest(method, path, body)
	recorder := httptest.NewRecorder()

	this.serveMux.ServeHTTP(recorder, request)

	return recorder
}

func (this *Client) RequestPOST(path string, body string) *httptest.ResponseRecorder {
	return this.requestForTest(http.MethodPost, path, strings.NewReader(body))
}

func (this *Client) RequestGET(path string) *httptest.ResponseRecorder {
	return this.requestForTest(http.MethodGet, path, nil)
}
