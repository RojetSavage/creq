package req

import (
	"net/http"
	"strings"
)

func SetHttpMethod(r *http.Request, s string) {
	if s == "GET" || s == "POST" || s == "DELETE" || s == "PUT" || s == "HEAD" || s == "PATCH" || s == "TRACE" {
		r.Method = s
	}
}

func AddHeader(r *http.Request, k string, v string) {
	r.Header.Set(k, v)
}

func AddCookie(r *http.Request, s string) {
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
