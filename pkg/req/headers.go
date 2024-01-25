package req

import (
	"net/http"
	"strings"
)

func (r *Request) SetHttpMethod(s string) {
	if s == "GET" || s == "POST" || s == "DELETE" || s == "PUT" || s == "HEAD" || s == "PATCH" || s == "TRACE" {
		r.method = s
	} else {
		panic("Invalid HTTP method. This shouldn't happen.")
	}
}

func (r *Request) AddHeader(k string, v string) {
	r.headers[k] = v
}

func AddHeadersToRequest(req *http.Request, headers map[string]string) {
	for k, v := range headers {
		req.Header.Set(k, v)
	}
}

func (r *Request) AddCookie(s string) {
	a := strings.Split(s, "&")

	for _, s := range a {
		kv := strings.Split(s, "=")
		r.cookies[kv[0]] = kv[1]
	}
}

func AddCookiesToRequest(req *http.Request, cookies map[string]string) {
	for k, v := range cookies {
		c := &http.Cookie{
			Name:     k,
			Value:    v,
			HttpOnly: false,
			Path:     req.URL.Path,
		}
		req.AddCookie(c)
	}
}
