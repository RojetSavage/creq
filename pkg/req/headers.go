package req

import (
	"net/http"
	"strings"
)

func (r *Request) SetHttpMethod(s string) {
	if s == http.MethodGet || s == http.MethodPost || s == http.MethodDelete || s == http.MethodPut || s == http.MethodHead || s == http.MethodPatch || s == http.MethodTrace {
		r.Method = s
	}
}

func (r *Request) AddHeader(s string) {
	a := strings.Split(s, "&")
	for _, header := range a {
		h := strings.Split(header, "=")
		r.Header.Set(h[0], h[1])
	}
}

func (r *Request) AddCookieToReq(s string) {
	a := strings.Split(s, "&")

	for _, s := range a {
		kv := strings.Split(s, "=")
		c := http.Cookie{
			Name:     kv[0],
			Value:    kv[1],
			HttpOnly: false,
			Path:     r.URL.Path,
		}
		r.AddCookie(&c)
	}
}
