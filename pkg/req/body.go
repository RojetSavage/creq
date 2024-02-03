package req

import (
	"bytes"
	"io"
	"strings"
)

func parseUrlQueryParams(s string) map[string]string {
	data := make(map[string]string)

	s = strings.Trim(s, "\"")
	a := strings.Split(s, "&")
	for _, s := range a {
		kv := strings.Split(s, "=")
		data[kv[0]] = kv[1]
	}

	return data
}

func mergeMaps(m ...map[string]string) map[string]string {
	data := make(map[string]string)

	for i := 0; i < len(m); i++ {
		for k, v := range m[i] {
			data[k] = v
		}
	}

	return data
}

func (r *Request) changeRequestBody(s string) {
	buf := []byte(s)
	r.Body = io.NopCloser(bytes.NewReader(buf))
}
