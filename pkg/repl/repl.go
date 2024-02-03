package repl

import (
	"args"
	"bufio"
	"fmt"
	"io"
	"os"
	"req"
	"strings"
	"time"
)

func RunRepl() {
	fmt.Printf(">>> ")
	readyToQuit := false

	c := req.NewClient()
	r := req.NewRequest()

	for !readyToQuit {
		readyToSend := false

		for !readyToSend {
			printCurrentRequestInfo(r)
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

func printCurrentRequestInfo(r *req.Request) {
	fmt.Printf("%v\nCurrent Request:\n%v %v \n", time.Now(), r.Method, r.URL.String())
	if r.Header != nil {
		for k, v := range r.Header {
			fmt.Println(k, ":", v)
		}
	}
	if r.Cookies() != nil {
		for _, cookie := range r.Cookies() {
			fmt.Printf(cookie.Name, ":", cookie.Value)
		}
	}

	if r.Body != nil {
		fmt.Println(r.Body)
	}
}

func getCommandLineArgs() []string {
	input := getUserInput()

	if len(input) == 0 {
		return nil
	}

	args := strings.Split(input, " ")
	return args
}

func getUserInput() string {
	input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	input = strings.Trim(input, "\n")
	input = strings.Trim(input, " ")
	return input
}

func isReadyToSend(flags []args.UserFlag) bool {
	if len(flags) == 0 {
		return true
	}

	for _, flag := range flags {
		if flag.F == "send" || flag.F == "s" {
			return true
		}
	}

	return false
}

func prettyPrint(body io.ReadCloser) {
	io.Copy(os.Stdout, body)
	fmt.Fprint(os.Stdout, "\n\n")
}
