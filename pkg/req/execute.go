package req

import (
	"args"
	"net/http"
	"strings"
)

func ApplyFlagsToRequest(r *RequestHandler, flags []args.UserFlag) error {
	for _, f := range flags {
		applyFlagToRequest(r, f)
	}
	return nil
}

func applyFlagToRequest(r *RequestHandler, f args.UserFlag) error {
	var err error

	switch f.F {
	case "x", "reset":
		r.resetRequest()
	case "s", "scheme":
		err = r.changeUri("scheme", f.Parameter)
	case "u", "user":
		err = r.changeUri("user", f.Parameter)
	case "host":
		err = r.changeUri("host", f.Parameter)
	case "P", "port":
		err = r.changeUri("port", f.Parameter)
	case "p", "path":
		err = r.changeUri("path", f.Parameter)
	case "q", "query":
		err = r.changeUri("query", f.Parameter)
	case "fragment":
		err = r.changeUri("fragment", f.Parameter)
	case "url":
		err = r.setUrl(f.Parameter)
	case "get":
		r.setHttpMethod(http.MethodGet)
	case "post":
		r.setHttpMethod(http.MethodPost)
	case "delete":
		r.setHttpMethod(http.MethodDelete)
	case "put":
		r.setHttpMethod(http.MethodPut)
	case "head":
		r.setHttpMethod(http.MethodHead)
	case "patch":
		r.setHttpMethod(http.MethodPatch)
	case "trace":
		r.setHttpMethod(http.MethodTrace)
	case "h", "header":
		r.addHeader(f.Parameter)
	case "d", "data":
		err = r.changeUri("query", f.Parameter)
	case "c", "contentType":
		r.addHeader(strings.Join([]string{"Content-Type", f.Parameter}, "="))
	case "C", "cookie":
		r.addCookieToReq(f.Parameter)
	case "j", "json":
		r.addHeader(strings.Join([]string{"Content-Type", "application/json"}, "="))
		r.setHttpMethod(http.MethodPost)
		r.changeRequestBody(f.Parameter)
	}

	if err != nil {
		return err
	}

	return nil
}

func ApplyFlagsToResponse(r *ResponseHandler, flags []args.UserFlag) error {
	for _, f := range flags {
		applyFlagToResponse(r, f)
	}
	return nil
}

func applyFlagToResponse(r *ResponseHandler, f args.UserFlag) error {
	var err error

	switch f.F {
	case "D", "dump-header":
		r.dumpHeader(f.Parameter)
	}

	return err
}
