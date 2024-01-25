package main

import (
	"fmt"
	"io"
	"log"
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
}

func main() {
	if !runRepl {
		args.ParseArgs(os.Args)
		flags := args.ValidateUserFlags()
		c, r := args.ApplyFlagsToClientAndRequest(req.NewClient(), req.NewRequestWrapper(), flags)

		res, err := req.SendRequest(c, r)
		defer res.Body.Close()

		if err != nil {
			log.Fatalln(err)
		}

		io.Copy(os.Stdout, res.Body)
	} else {
		fmt.Println("Run REPL")
	}
}
