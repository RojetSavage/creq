package main

import (
	"gcurl/pkg/repl"
	"io"
	"log"
	"os"
	"req"

	"args"
)

var (
	runRepl bool
)

func init() {
	if len(os.Args) == 1 {
		runRepl = true
	}
}

func main() {
	if len(os.Args) > 1 {
		flags := args.ParseArgs(os.Args)
		args.ValidateUserFlags(flags, false)

		r := req.NewRequest()
		c := req.NewClient()

		req.ApplyFlagsToClient(c, flags)
		req.ApplyFlagsToRequest(r, flags)

		res, err := req.SendRequest(c, r)
		// defer res.Body.Close()

		if err != nil {
			log.Fatalln(err)
		}

		io.Copy(os.Stdout, res.Body)
	} else {
		repl.RunRepl()
	}
}
