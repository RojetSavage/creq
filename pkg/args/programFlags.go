package args

func setProgramFlags() {
	var f = []Flag{
		//used to send json
		{
			Flag:          "json",
			Short:         "j",
			ParamRequired: true,
			DefaultValue:  "",
		},
		//send key/value pairs, which we parse into json
		{
			Flag:          "data",
			Short:         "d",
			ParamRequired: true,
			DefaultValue:  "",
		},
		//request only the headers
		{
			Flag:          "header",
			Short:         "h",
			ParamRequired: false,
			DefaultValue:  "",
		},
		//sets the port value
		{
			Flag:          "port",
			Short:         "P",
			ParamRequired: true,
			DefaultValue:  "8080",
		},
		//sets the path to append to url, either set url with full path or use to append path to localhost
		{
			Flag:          "path",
			Short:         "p",
			ParamRequired: true,
			DefaultValue:  "",
		},
	}

	ProgramFlags = f
}
