package req

import (
	"errors"
	"net/http"
	"net/url"
)

type Request struct {
	headers  map[string]string
	cookies  map[string]string
	method   string
	body     []byte
	encoding string
	url      *url.URL
}

func NewRequest() *http.Request {
	req, _ := http.NewRequest(http.MethodGet, "http://localhost:3001", nil)
	return req
}

func SendRequest(c *http.Client, r *http.Request) (*http.Response, error) {
	resp, err := c.Do(r)

	if err != nil {
		return nil, errors.New("HTTP request failed.")
	}

	return resp, err
}
