package repl

import (
	"args"
	"bufio"
	"fmt"
	"io"
	"os"
	"req"
	"strings"
)

func PrintCurrentRequestInfo(r req.RequestHandler) {
	fmt.Printf("\nCurrent Request:\n%v %v \n", r.Method, r.URL.String())
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

func GetCommandLineArgs() []string {
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

func PrettyPrint(body io.ReadCloser) {
	//TODO: MAKE LOOK GOOD
	io.Copy(os.Stdout, body)
	fmt.Fprint(os.Stdout, "\n")
}
