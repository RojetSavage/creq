package req

import (
	"args"
	"net/http"
	"strings"
)

func ApplyFlagsToRequest(r *Request, flags []args.UserFlag) error {
	for _, f := range flags {
		applyFlagToRequest(r, f)
	}
	return nil
}

func applyFlagToRequest(r *Request, f args.UserFlag) error {
	var err error

	switch f.F {
	case "s", "scheme":
		err = r.ChangeUri("scheme", f.Parameter)
	case "host":
		err = r.ChangeUri("host", f.Parameter)
	case "P", "port":
		err = r.ChangeUri("port", f.Parameter)
	case "p", "path":
		err = r.ChangeUri("path", f.Parameter)
	case "q", "query":
		err = r.ChangeUri("query", f.Parameter)
	case "fragment":
		err = r.ChangeUri("fragment", f.Parameter)
	case "url":
		err = r.SetUrl(f.Parameter)

	case "d", "data":
		r.AddHeader(strings.Join([]string{"Content-Type", "x-www-form-urlencoded"}, "="))
		err = r.ChangeUri("query", f.Parameter)
	case "j", "json":
		r.AddHeader(strings.Join([]string{"Content-Type", "application/json"}, "="))
		r.SetHttpMethod(http.MethodPost)
		r.ChangeRequestBody(f.Parameter)

	case "c", "contentType":
		r.AddHeader(strings.Join([]string{"Content-Type", f.Parameter}, "="))
	case "C", "cookie":
		r.AddCookieToReq(f.Parameter)

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
	}

	if err != nil {
		panic(err)
	}

	return nil
}
