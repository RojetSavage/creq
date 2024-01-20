package args

import (
	"req"
)

func ExecuteFlags(flags []UserFlag) {
	for _, f := range flags {
		switch f.F {
		case "P", "port":
			req.R.SetPort(f.Parameter)
		case "p", "path":
			req.R.SetPath(f.Parameter)
		case "j", "json":
			req.R.SetBodyJson(f.Parameter)
		case "d", "data":
			req.R.SetBodyData(f.Parameter)
		case "c", "contentType":
			req.R.SetHeaderContentType(f.Parameter)
		case "get":
			req.R.SetMethod("GET")
		case "put":
			req.R.SetMethod("PUT")
		case "post":
			req.R.SetMethod("POST")
		case "head":
			req.R.SetMethod("HEAD")
		case "delete":
			req.R.SetMethod("DELETE")
		}

	}
}
