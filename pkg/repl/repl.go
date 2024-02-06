package repl

import (
	"args"
	"fmt"
	"os"
	"req"
)

func RunRepl() {
	fmt.Printf(">>> ")
	readyToQuit := false

	c := req.NewClient()
	r := req.NewRequest()

	for !readyToQuit {
		readyToSend := false

		for !readyToSend {
			printCurrentRequestInfo(*r)
			cmdLineArgs := getCommandLineArgs()

			if cmdLineArgs == nil {
				readyToSend = true
				continue
			}

			if cmdLineArgs[0] == "quit" || cmdLineArgs[0] == "q" {
				readyToQuit = true
				break
			}

			parseError, flags := args.ParseArgs(cmdLineArgs, true)
			validationError, _ := args.ValidateUserFlags(flags, true)

			if parseError != nil || validationError != nil {
				fmt.Println(parseError, validationError)
				continue
			}

			clientError := req.ApplyFlagsToClient(c, flags)
			requestError := req.ApplyFlagsToRequest(r, flags)

			if clientError != nil || requestError != nil {
				fmt.Println(clientError, requestError)
				continue
			}

			readyToSend = isReadyToSend(flags)
		}

		if !readyToQuit {
			res, err := req.SendRequest(c, r)

			if err != nil {
				fmt.Println(err)
			}

			prettyPrint(res.Body)
		}
	}
	os.Exit(0)
}
