package req

import (
	"net/http"
)

type Request struct {
	http.Request
}

func NewRequest() *Request {
	req, _ := http.NewRequest(http.MethodGet, "http://localhost:3001", nil)
	r := Request{*req}
	return &r
}

func SendRequest(c *http.Client, r *Request) (*http.Response, error) {
	clone := r.Clone(r.Context())
	resp, err := c.Do(clone)

	if err != nil {
		return nil, err
	}

	return resp, nil
}
