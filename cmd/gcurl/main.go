package main

import (
	"fmt"
	"os"

	"args"
	"req"
)

var (
	flags   []args.UserFlag
	runRepl bool
)

func init() {
	if len(os.Args) == 1 {
		runRepl = true
	}
	fmt.Println(os.Args)
}

func main() {
	if !runRepl {
		args.ParseArgs(os.Args)
		flags = args.ValidateUserFlags()
		args.ExecuteFlags(flags)

		req.R.PerformRequest()
		//os.copy the response to stdout
	} else {
		fmt.Println("Run REPL")
	}
}
