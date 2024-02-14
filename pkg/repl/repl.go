package repl

import (
	"args"
	"fmt"
	"os"
	"req"
)

var readyToQuit bool = false

func RunRepl() {
	c := req.NewClient()
	r := req.NewRequest()
	op := req.NewResponseOperationsQueue()

	for !readyToQuit {
		readyToSend := false

		for !readyToSend {
			PrintCurrentRequestInfo(*r)
			cmdLineArgs := GetCommandLineArgs()

			if noArgs := IsNoArgs(cmdLineArgs); noArgs == true {
				readyToSend = true
				continue
			}

			if quit := handleQuitRepl(cmdLineArgs[0]); quit == true {
				readyToQuit = true
				break
			}

			parseError, flags := args.ParseArgs(cmdLineArgs, true)
			validationError, _ := args.ValidateUserFlags(&flags, true)

			if ok := PrintError(parseError, validationError); ok == false {
				continue
			}

			clientError := req.ApplyFlagsToClient(c, flags)
			requestError := req.ApplyFlagsToRequest(r, flags)
			responseError := req.QueueResponseOperations(op, flags)

			if ok := PrintError(clientError, requestError, responseError); ok == false {
				continue
			}

			readyToSend = isReadyToSend(flags)
		}

		if !readyToQuit {
			res, err := req.SendRequest(c, r)
			modifiedRes, err := req.NewResponseHandler(res, op)

			if err != nil {
				fmt.Println(err)
			}

			PrettyPrint(modifiedRes.Response.Body)
		}
	}
	os.Exit(0)
}

func PrintError(e ...error) bool {
	ok := true

	for _, err := range e {
		if err != nil {
			fmt.Println(err)
			ok = false
		}
	}
	return ok
}

func IsNoArgs(a []string) bool {
	if len(a) == 0 || a == nil {
		return true
	}
	return false
}

func handleQuitRepl(s string) bool {
	if s == "quit" || s == "q" {
		return true
	}
	return false
}
