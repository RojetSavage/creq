package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"args"
	"repl"
	"req"
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
		_, flags := args.ParseArgs(os.Args, false)
		err, ok := args.ValidateUserFlags(flags, false)

		if !ok {
			fmt.Println(err)
		} else {
			r := req.NewRequest()
			c := req.NewClient()

			req.ApplyFlagsToClient(c, flags)
			req.ApplyFlagsToRequest(r, flags)

			res, err := req.SendRequest(c, r)

			if err != nil {
				log.Fatalln(err)
			}

			io.Copy(os.Stdout, res.Body)
		}
	} else {
		repl.RunRepl()
	}
}
