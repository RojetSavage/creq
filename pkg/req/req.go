package req

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
)

type Request struct {
	headers  map[string]string
	cookies  map[string]string
	method   string
	body     []byte
	encoding string

	reqUrl
}

type reqUrl struct {
	scheme   string
	user     string
	host     string
	port     string
	path     string
	query    string
	fragment string
}

func NewRequestWrapper() *Request {
	r := Request{
		method:  "GET",
		headers: map[string]string{},
		cookies: map[string]string{},
		reqUrl: reqUrl{
			scheme: "http",
			host:   "localhost",
			port:   "3001",
		},
	}

	return &r
}

func SendRequest(c *http.Client, r *Request) (*http.Response, error) {
	httpReq := createHTTPRequestFromBase(r)
	resp, err := c.Do(httpReq)

	if err != nil {
		return nil, errors.New("HTTP request failed.")
	}

	return resp, err
}

func createHTTPRequestFromBase(r *Request) *http.Request {
	fmt.Println(r.method, r.buildUrl(), bytes.NewBuffer(r.body), r.headers)
	req, _ := http.NewRequest(r.method, r.buildUrl(), bytes.NewBuffer(r.body))
	AddHeadersToRequest(req, r.headers)
	AddCookiesToRequest(req, r.cookies)

	return req
}
