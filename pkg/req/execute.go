package req

import (
	"args"
	"net/http"
)

func ApplyFlagsToClientAndRequest(flags []args.UserFlag) (*http.Client, *Request) {
	c, r := NewClient(), NewRequestWrapper()

	for _, f := range flags {
		switch f.F {

		//client
		case "connect-timeout":
			SetClientTimeout(c, f.Parameter)

		//url
		case "s", "scheme":
			r.SetScheme(f.Parameter)
		case "P", "port":
			r.SetPort(f.Parameter)
		case "p", "path":
			r.SetPath(f.Parameter)
		case "url":
			r.SetUrl(f.Parameter)
		case "d", "data":
			r.AddHeader("Content-Type", "x-www-form-urlencoded")
			r.SetUrlQueryString(f.Parameter)

		//body
		case "j", "json":
			r.AddHeader("Content-Type", "application/json")
			r.SetHttpMethod(http.MethodPost)
			r.SetBodyJson(f.Parameter)

		//headers
		case "c", "contentType":
			r.AddHeader("Content-Type", f.Parameter)
		case "C", "cookie":
			r.AddCookie(f.Parameter)

		//methods
		case "get":
			r.SetHttpMethod(http.MethodGet)
		case "post":
			r.SetHttpMethod(http.MethodPost)
		case "delete":
			r.SetHttpMethod(http.MethodDelete)
		case "put":
			r.SetHttpMethod(http.MethodPut)
		case "head":
			r.SetHttpMethod(http.MethodHead)
		case "patch":
			r.SetHttpMethod(http.MethodPatch)
		case "trace":
			r.SetHttpMethod(http.MethodTrace)
		default:
			panic("Black Flag")
		}

	}

	return c, r
}
