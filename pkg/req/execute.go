package req

import (
	"args"
	"net/http"
)

func ApplyFlagsToRequest(r *http.Request, flags []args.UserFlag) error {
	for _, f := range flags {
		applyFlagToRequest(r, f)
	}
	return nil
}

func applyFlagToRequest(r *http.Request, f args.UserFlag) error {
	var err error

	switch f.F {
	case "s", "scheme":
		err = ChangeUri(r, "scheme", f.Parameter)
	case "host":
		err = ChangeUri(r, "host", f.Parameter)
	case "P", "port":
		err = ChangeUri(r, "port", f.Parameter)
	case "p", "path":
		err = ChangeUri(r, "path", f.Parameter)
	case "q", "query":
		err = ChangeUri(r, "query", f.Parameter)
	case "fragment":
		err = ChangeUri(r, "fragment", f.Parameter)
	case "url":
		err = SetUrl(r, f.Parameter)

	case "d", "data":
		AddHeader(r, "Content-Type", "x-www-form-urlencoded")
		err = ChangeUri(r, "query", f.Parameter)
	case "j", "json":
		AddHeader(r, "Content-Type", "application/json")
		SetHttpMethod(r, http.MethodPost)
		ChangeRequestBody(r, f.Parameter)

	case "c", "contentType":
		AddHeader(r, "Content-Type", f.Parameter)
	case "C", "cookie":
		AddCookie(r, f.Parameter)

	case "get":
		SetHttpMethod(r, http.MethodGet)
	case "post":
		SetHttpMethod(r, http.MethodPost)
	case "delete":
		SetHttpMethod(r, http.MethodDelete)
	case "put":
		SetHttpMethod(r, http.MethodPut)
	case "head":
		SetHttpMethod(r, http.MethodHead)
	case "patch":
		SetHttpMethod(r, http.MethodPatch)
	case "trace":
		SetHttpMethod(r, http.MethodTrace)

	}

	if err != nil {
		panic(err)
	}

	return nil
}
