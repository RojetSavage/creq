package main

import (
	"fmt"
	"log"
	"os"

	"args"
	"repl"
	"req"
)

func main() {
	if len(os.Args) < 1 {
		c := req.NewClient()
		r := req.NewRequest()
		op := req.NewResponseOperationsQueue()

		repl.PrintCurrentRequestInfo(*r)
		cmdLineArgs := repl.GetCommandLineArgs()

		parseError, flags := args.ParseArgs(cmdLineArgs, true)
		validationError, _ := args.ValidateUserFlags(&flags, true)

		if ok := repl.PrintError(parseError, validationError); ok == false {
			log.Fatal(parseError, validationError)
		}

		clientError := req.ApplyFlagsToClient(c, flags)
		requestError := req.ApplyFlagsToRequest(r, flags)
		responseError := req.QueueResponseOperations(op, flags)

		if ok := repl.PrintError(clientError, requestError, responseError); ok == false {
			log.Fatal(clientError, requestError, responseError)
		}

		res, err := req.SendRequest(c, r)
		modifiedRes, err := req.NewResponseHandler(res, op)

		if err != nil {
			fmt.Println(err)
		}

		repl.PrettyPrint(modifiedRes.Response.Body)
	} else {
		repl.RunRepl()
	}
}
