package req

import (
	"errors"
	"net/url"
	"strconv"
	"strings"
)

func (r *Request) SetUrl(s string) {
	url, _ := url.ParseRequestURI(s)
	r.setReqUrl(url)
}

func (r *Request) setReqUrl(u *url.URL) {
	r.scheme = u.Scheme
	r.host = u.Hostname()
	r.port = u.Port()
	r.path = u.RequestURI()
	r.query = u.Query().Encode()
	r.fragment = u.Fragment
}

func (r Request) buildUrl() string {
	var b strings.Builder

	b.WriteString(r.scheme)
	b.WriteString("://")
	b.WriteString(r.host)

	if r.port != "" {
		b.WriteString(":")
		b.WriteString(r.port)
	}

	if r.path != "" {
		b.WriteString(r.path)
	}

	if r.query != "" {
		b.WriteString("?")
		b.WriteString(r.query)
	}

	return b.String()
}

func (r *Request) SetScheme(s string) {
	r.scheme = s
}

func (r *Request) SetPort(s string) error {
	if _, err := strconv.Atoi(s); err != nil {
		return errors.New("Can't convert port number to int")
	}
	r.port = s
	return nil
}

func (r *Request) SetPath(s string) {
	if strings.HasPrefix(s, "/") == true {
		r.path = s
	} else {
		b := strings.Builder{}
		b.WriteString("/")
		b.WriteString(s)
		r.path = b.String()
	}
}
