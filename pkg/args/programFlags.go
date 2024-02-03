package args

func setProgramFlags() {
	var f = []flag{
		//repl flags

		{
			Flag:          "send",
			Short:         "s",
			ParamRequired: false,
			DefaultValue:  "",
			Description:   "sends request",
			ReplOnly:      true,
		},
		{
			Flag:          "reset",
			Short:         "x",
			ParamRequired: false,
			DefaultValue:  "",
			Description:   "returns state of http request to default",
			ReplOnly:      true,
		},

		//body
		{
			Flag:          "json",
			Short:         "j",
			ParamRequired: true,
			DefaultValue:  "",
			Description:   "Add JSON to the req body",
			ReplOnly:      false,
		},
		{
			Flag:          "data",
			Short:         "d",
			ParamRequired: true,
			DefaultValue:  "",
			Description:   "Add key value pairs which are marshalled into JSON and added to the req body",
			ReplOnly:      false,
		},

		//url
		{
			Flag:          "url",
			Short:         "",
			ParamRequired: true,
			DefaultValue:  "http://localhost.com:8080",
			Description:   "Sets the URL",
			ReplOnly:      false,
		},
		{
			Flag:          "port",
			Short:         "P",
			ParamRequired: true,
			DefaultValue:  "8080",
			Description:   "Sets the req port value",
			ReplOnly:      false,
		},
		{
			Flag:          "path",
			Short:         "p",
			ParamRequired: true,
			DefaultValue:  "",
			Description:   "Appends path to req URL. Default URL is localhost.com. If URL is given as first cmd line arg, path is appended.",
			ReplOnly:      false,
		},
		{
			Flag:          "scheme",
			Short:         "s",
			ParamRequired: true,
			DefaultValue:  "http",
			Description:   "Sets the req protocol",
			ReplOnly:      false,
		},

		//headers
		{
			Flag:          "c",
			Short:         "c",
			ParamRequired: true,
			DefaultValue:  "application/json",
			Description:   "Sets the header 'content-type'",
			ReplOnly:      false,
		},
		{
			Flag:          "cookie",
			Short:         "C",
			ParamRequired: true,
			DefaultValue:  "",
			Description:   "Sets cookie",
			ReplOnly:      false,
		},
		//methods
		{
			Flag:          "get",
			Short:         "g",
			ParamRequired: false,
			DefaultValue:  "",
			Description:   "Sets the http method",
			ReplOnly:      false,
		},
		{
			Flag:          "put",
			Short:         "",
			ParamRequired: false,
			DefaultValue:  "",
			Description:   "Sets the http method",
			ReplOnly:      false,
		},
		{
			Flag:          "post",
			Short:         "",
			ParamRequired: false,
			DefaultValue:  "",
			Description:   "Sets the http method",
			ReplOnly:      false,
		},
		{
			Flag:          "delete",
			Short:         "",
			ParamRequired: false,
			DefaultValue:  "",
			Description:   "Sets the http method",
			ReplOnly:      false,
		},
		{
			Flag:          "head",
			Short:         "",
			ParamRequired: false,
			DefaultValue:  "",
			Description:   "Sets the http method",
			ReplOnly:      false,
		},
		{
			Flag:          "patch",
			Short:         "",
			ParamRequired: false,
			DefaultValue:  "",
			Description:   "Sets the http method",
			ReplOnly:      false,
		},
		{
			Flag:          "trace",
			Short:         "",
			ParamRequired: false,
			DefaultValue:  "",
			Description:   "Sets the http method",
			ReplOnly:      false,
		},
		{
			Flag:          "connect-timeout",
			Short:         "",
			ParamRequired: true,
			DefaultValue:  "",
			Description:   "Sets a limit in seconds for a connection request",
			ReplOnly:      false,
		},
	}

	ProgramFlags = f
}
