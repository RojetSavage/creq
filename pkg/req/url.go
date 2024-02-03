package req

import (
	"errors"
	"net/url"
	"strconv"
	"strings"
)

type dismantledUrl struct {
	scheme   string
	user     string
	host     string
	port     string
	path     string
	query    string
	fragment string
}

func (r *Request) setUrl(s string) error {
	url, err := url.Parse(s)

	if err != nil {
		return errors.New("Bad URL formed")
	}

	r.URL = url
	return nil
}

func explodeUrl(u *url.URL) *dismantledUrl {
	d := dismantledUrl{}

	d.scheme = u.Scheme
	d.host = u.Hostname()
	d.port = u.Port()
	d.path = u.Path
	d.query = u.RawQuery
	d.fragment = u.Fragment
	return &d
}

func reassembleUrl(u *dismantledUrl) string {
	var b = strings.Builder{}

	b.WriteString(u.scheme)
	b.WriteString("://")

	if u.user != "" {
		b.WriteString(u.user)
		b.WriteString("@")
	}

	b.WriteString(u.host)

	if u.port != "" {
		b.WriteString(":")
		b.WriteString(u.port)
	}

	if u.path != "" {
		b.WriteString(u.path)
	}

	if u.query != "" {
		b.WriteString("?")
		b.WriteString(u.query)
	}

	if u.fragment != "" {
		b.WriteString("#")
		b.WriteString(u.fragment)
	}

	return b.String()
}

func (r *Request) changeUri(urlComponent string, s string) error {
	u := explodeUrl(r.URL)

	switch urlComponent {
	case "scheme":
		u.scheme = s
	case "host":
		u.host = s
	case "port":
		u.port = validatePort(s)
	case "path":
		u.path = validatePath(s)
	case "query":
		u.query = s
	case "fragment":
		u.fragment = s
	}

	err := r.setUrl(reassembleUrl(u))
	return err
}

func validatePort(s string) string {
	if _, err := strconv.Atoi(s); err != nil {
		return ""
	}
	return s
}

func validatePath(s string) string {
	if strings.HasPrefix(s, "/") == true {
		return s
	} else {
		b := strings.Builder{}
		b.WriteString("/")
		b.WriteString(s)
		return b.String()
	}
}
