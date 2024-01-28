package repl

import (
	"args"
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"req"
	"strings"
)

func RunRepl() {
	fmt.Printf(">>> ")
	input := []string{","}

	c := req.NewClient()
	r := req.NewRequest()

	for input[0] != "exit" {
		printCurrentRequestInfo(r)

		cmdLineArgs := strings.Split(getUserInput(), " ")
		flags := args.ParseArgs(cmdLineArgs)
		args.ValidateUserFlags(flags, true)

		req.ApplyFlagsToClient(c, flags)
		req.ApplyFlagsToRequest(r, flags)

		if send := isSendable(flags); send {
			res, err := req.SendRequest(c, r)
			defer res.Body.Close()

			if err != nil {
				log.Fatalln(err)
			}

			prettyPrint(res.Body)
		}

	}

	os.Exit(0)
}

func printCurrentRequestInfo(r *req.Request) {
	fmt.Printf("Current Request: %v -> %v, Body: %v\n", r.Method, r.URL.String(), r.Body)
}

func getUserInput() string {
	input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	input = strings.Trim(input, "\n")
	input = strings.Trim(input, " ")
	return input
}

func isSendable(flags []args.UserFlag) bool {
	if len(flags) == 0 {
		return true
	}

	//if check for 'send' flag

	return false
}

func prettyPrint(body io.ReadCloser) {
	io.Copy(os.Stdout, body)
}
