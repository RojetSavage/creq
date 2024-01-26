package args

func setProgramFlags() {
	var f = []flag{
		//body
		{
			Flag:          "json",
			Short:         "j",
			ParamRequired: true,
			DefaultValue:  "",
			Description:   "Add JSON to the req body",
		},
		{
			Flag:          "data",
			Short:         "d",
			ParamRequired: true,
			DefaultValue:  "",
			Description:   "Add key value pairs which are marshalled into JSON and added to the req body",
		},

		//url
		{
			Flag:          "url",
			Short:         "",
			ParamRequired: true,
			DefaultValue:  "http://localhost.com:8080",
			Description:   "Sets the URL",
		},
		{
			Flag:          "port",
			Short:         "P",
			ParamRequired: true,
			DefaultValue:  "8080",
			Description:   "Sets the req port value",
		},
		{
			Flag:          "path",
			Short:         "p",
			ParamRequired: true,
			DefaultValue:  "",
			Description:   "Appends path to req URL. Default URL is localhost.com. If URL is given as first cmd line arg, path is appended.",
		},
		{
			Flag:          "scheme",
			Short:         "s",
			ParamRequired: true,
			DefaultValue:  "http",
			Description:   "Sets the req protocol",
		},

		//headers
		{
			Flag:          "c",
			Short:         "c",
			ParamRequired: true,
			DefaultValue:  "application/json",
			Description:   "Sets the header 'content-type'",
		},
		{
			Flag:          "cookie",
			Short:         "C",
			ParamRequired: true,
			DefaultValue:  "",
			Description:   "Sets cookie",
		},
		//methods
		{
			Flag:          "get",
			Short:         "g",
			ParamRequired: false,
			DefaultValue:  "",
			Description:   "Sets the http method",
		},
		{
			Flag:          "put",
			Short:         "",
			ParamRequired: false,
			DefaultValue:  "",
			Description:   "Sets the http method",
		},
		{
			Flag:          "post",
			Short:         "",
			ParamRequired: false,
			DefaultValue:  "",
			Description:   "Sets the http method",
		},
		{
			Flag:          "delete",
			Short:         "",
			ParamRequired: false,
			DefaultValue:  "",
			Description:   "Sets the http method",
		},
		{
			Flag:          "head",
			Short:         "",
			ParamRequired: false,
			DefaultValue:  "",
			Description:   "Sets the http method",
		},
		{
			Flag:          "patch",
			Short:         "",
			ParamRequired: false,
			DefaultValue:  "",
			Description:   "Sets the http method",
		},
		{
			Flag:          "trace",
			Short:         "",
			ParamRequired: false,
			DefaultValue:  "",
			Description:   "Sets the http method",
		},
		{
			Flag:          "connect-timeout",
			Short:         "",
			ParamRequired: true,
			DefaultValue:  "",
			Description:   "Sets a limit in seconds for a connection request",
		},
	}

	ProgramFlags = f
}
