package req

import (
	"bytes"
	"io"
	"net/http"
)

func NewRequest() *RequestHandler {
	req, _ := http.NewRequest(http.MethodGet, "http://localhost:3001", nil)
	r := RequestHandler{false, *req}
	return &r
}

func SendRequest(c *http.Client, r *RequestHandler) (*http.Response, error) {
	clone := cloneReq(r)

	res, err := c.Do(clone)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func cloneReq(r *RequestHandler) *http.Request {
	clone := r.Clone(r.Context())

	if r.Body == nil {
		return clone
	}

	body, _ := io.ReadAll(r.Body)
	bodyCopy := make([]byte, len(body))
	_ = copy(bodyCopy, body)

	r.ContentLength = int64(len(body))
	clone.ContentLength = int64(len(body))

	r.Body = io.NopCloser(bytes.NewBuffer(body))
	clone.Body = io.NopCloser(bytes.NewBuffer(bodyCopy))

	r.GetBody = func() (io.ReadCloser, error) {
		reader := bytes.NewReader(body)
		return io.NopCloser(reader), nil
	}

	clone.GetBody = func() (io.ReadCloser, error) {
		reader := bytes.NewReader(bodyCopy)
		return io.NopCloser(reader), nil
	}

	if r.GetBody != nil && r.ContentLength == 0 {
		r.Body = http.NoBody
		r.GetBody = func() (io.ReadCloser, error) { return http.NoBody, nil }
	}

	if clone.GetBody != nil && clone.ContentLength == 0 {
		clone.Body = http.NoBody
		clone.GetBody = func() (io.ReadCloser, error) { return http.NoBody, nil }
	}

	return clone
}

func (r *RequestHandler) resetRequest() {
	*r = *NewRequest()
}
