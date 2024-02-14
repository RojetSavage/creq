package args

func setProgramFlags() {
	var f = []flag{
		//repl flags

		{
			Flag:          "send",
			Short:         "s",
			ParamAccepted: false,
			DefaultValue:  "",
			Description:   "sends request",
			ReplOnly:      true,
		},
		{
			Flag:          "reset",
			Short:         "x",
			ParamAccepted: false,
			DefaultValue:  "",
			Description:   "returns state of http request to default",
			ReplOnly:      true,
		},

		//body
		{
			Flag:          "json",
			Short:         "j",
			ParamAccepted: true,
			DefaultValue:  "",
			Description:   "Add JSON to the req body",
			ReplOnly:      false,
		},
		{
			Flag:          "data",
			Short:         "d",
			ParamAccepted: true,
			DefaultValue:  "",
			Description:   "Add key value pairs which are marshalled into JSON and added to the req body",
			ReplOnly:      false,
		},

		//url
		{
			Flag:          "url",
			Short:         "",
			ParamAccepted: true,
			DefaultValue:  "http://localhost.com:8080",
			Description:   "Sets the URL",
			ReplOnly:      false,
		},
		{
			Flag:          "user",
			Short:         "u",
			ParamAccepted: true,
			DefaultValue:  "",
			Description:   "Sets the user and password for authentication",
			ReplOnly:      false,
		}, {
			Flag:          "port",
			Short:         "p",
			ParamAccepted: true,
			DefaultValue:  "8080",
			Description:   "Sets the req port value",
			ReplOnly:      false,
		},
		{
			Flag:          "path",
			Short:         "P",
			ParamAccepted: true,
			DefaultValue:  "",
			Description:   "Appends path to req URL. Default URL is localhost.com. If URL is given as first cmd line arg, path is appended.",
			ReplOnly:      false,
		},
		{
			Flag:          "scheme",
			Short:         "s",
			ParamAccepted: true,
			DefaultValue:  "http",
			Description:   "Sets the req protocol",
			ReplOnly:      false,
		},

		//headers
		{
			Flag:          "c",
			Short:         "c",
			ParamAccepted: true,
			DefaultValue:  "application/json",
			Description:   "Sets the header 'content-type'",
			ReplOnly:      false,
		},
		{
			Flag:          "header",
			Short:         "h",
			ParamAccepted: true,
			DefaultValue:  "",
			Description:   "Applies header to request e.g. foo:bar",
			ReplOnly:      false,
		}, {
			Flag:          "cookie",
			Short:         "C",
			ParamAccepted: true,
			DefaultValue:  "",
			Description:   "Sets cookie",
			ReplOnly:      false,
		},
		//methods
		{
			Flag:          "get",
			Short:         "g",
			ParamAccepted: false,
			DefaultValue:  "",
			Description:   "Sets the http method",
			ReplOnly:      false,
		},
		{
			Flag:          "put",
			Short:         "",
			ParamAccepted: false,
			DefaultValue:  "",
			Description:   "Sets the http method",
			ReplOnly:      false,
		},
		{
			Flag:          "post",
			Short:         "",
			ParamAccepted: false,
			DefaultValue:  "",
			Description:   "Sets the http method",
			ReplOnly:      false,
		},
		{
			Flag:          "delete",
			Short:         "",
			ParamAccepted: false,
			DefaultValue:  "",
			Description:   "Sets the http method",
			ReplOnly:      false,
		},
		{
			Flag:          "head",
			Short:         "",
			ParamAccepted: false,
			DefaultValue:  "",
			Description:   "Sets the http method",
			ReplOnly:      false,
		},
		{
			Flag:          "patch",
			Short:         "",
			ParamAccepted: false,
			DefaultValue:  "",
			Description:   "Sets the http method",
			ReplOnly:      false,
		},
		{
			Flag:          "trace",
			Short:         "",
			ParamAccepted: false,
			DefaultValue:  "",
			Description:   "Sets the http method",
			ReplOnly:      false,
		},
		{
			Flag:          "connect-timeout",
			Short:         "",
			ParamAccepted: true,
			DefaultValue:  "",
			Description:   "Sets a limit in seconds for a connection request",
			ReplOnly:      false,
		},
		{
			Flag:          "dump-header",
			Short:         "D",
			ParamAccepted: false,
			DefaultValue:  "./header-dump.txt",
			Description:   "Writes the response headers to the specified file. Writes to ./header-dump.txt by default",
			ReplOnly:      false,
		},
	}

	ProgramFlags = f
}
