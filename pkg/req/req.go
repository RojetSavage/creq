package req

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/url"
	"strconv"
	"strings"
)

type request struct {
	url      *url.URL //if url is given, remove default port
	port     int      //used to override default
	path     string
	method   string
	body     []byte
	encoding string
	headers
}

type headers struct {
	contentType string
}

var R request = request{
	method: "GET",
	port:   8080,
	headers: headers{
		contentType: "application/json",
	},
}

func init() {
	u, err := url.ParseRequestURI("http://localhost.com:8080")
	fmt.Println("first")
	if err != nil {
		log.Fatal("Failed to init program url")
	} else {
		R.url = u
	}

	p, err := strconv.Atoi(u.Port())

	if err != nil {
		fmt.Println("No assigned port")
	}

	R.port = p
}

func (r request) buildUrl() string {
	var b strings.Builder

	scheme := r.url.Scheme
	host := r.url.Hostname()
	port := r.url.Port()
	uri := r.url.RequestURI()
	userPath := r.path

	b.WriteString(scheme)
	b.WriteString("://")
	b.WriteString(host)

	if port != "" {
		b.WriteString(":")
		b.WriteString(port)
	}

	if userPath != "" {
		b.WriteString(r.path)
	} else {
		b.WriteString(uri)
	}

	return b.String()
}

func (r *request) SetUrl(u *url.URL) {
	R.url = u
	R.port = 0
}

func (r *request) SetHeaderContentType(s string) {
	r.contentType = s
}

func (r *request) SetMethod(s string) error {
	if s == "GET" || s == "POST" || s == "DELETE" || s == "PUT" {
		R.method = s
		return nil
	}

	return errors.New("Invalid HTTP method")
}

func (r *request) SetPath(s string) {
	if strings.HasPrefix(s, "/") == true {
		r.path = s
	} else {
		b := strings.Builder{}
		b.WriteString("/")
		b.WriteString(s)
		r.path = b.String()
	}
}

func (r *request) SetPort(s string) error {
	if p, err := strconv.Atoi(s); err != nil {
		return errors.New("Can't convert port number to int")
	} else if R.url.Port() == s {
		return nil //user set port is same as url port
	} else {
		R.port = p
		return nil
	}
}

func (r *request) SetBodyData(s string) {
	data := make(map[string]string)

	s = strings.Trim(s, "\"")
	a := strings.Split(s, "&")
	for _, s := range a {
		kv := strings.Split(s, "=")
		data[kv[0]] = kv[1]
	}

	j, err := json.Marshal(data)

	if err != nil {
		log.Fatalln("Failed to marshal user data")
	}

	R.body = j
}

func (r request) PerformRequest() (string, error) {
	// fmt.Println("request", R.url.Scheme, R.url.Host, R.url.Port(), R.port, R.method, string(R.body))
	fmt.Println(r.buildUrl())
	return "", nil
}

func (r *request) SetBodyJson(s string) {
	j, err := json.Marshal(s)

	if err != nil {
		log.Fatalln("Failed to marshal user provided json")
	}

	R.body = j
}
