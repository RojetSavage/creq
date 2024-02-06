package req

import (
	"net/http"
	"strings"
)

func (r *RequestHandler) setHttpMethod(s string) {
	if s == http.MethodGet || s == http.MethodPost || s == http.MethodDelete || s == http.MethodPut || s == http.MethodHead || s == http.MethodPatch || s == http.MethodTrace {
		r.Method = s
	}
}

func (r *RequestHandler) addHeader(s string) {
	h := strings.Split(s, ":")
	r.Header.Set(h[0], h[1])
}

func (r *RequestHandler) addCookieToReq(s string) {
	kv := strings.Split(s, ":")
	c := http.Cookie{
		Name:     kv[0],
		Value:    kv[1],
		HttpOnly: false,
		Path:     r.URL.Path,
	}
	r.AddCookie(&c)
}
