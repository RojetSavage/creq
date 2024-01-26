package repl

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func RunRepl() {
	var input string
	fmt.Printf(">>>: ")

	for input != "exit" {
		input, _ = bufio.NewReader(os.Stdin).ReadString('\n')

		input = strings.Trim(input, "\n")
		input = strings.Trim(input, " ")
		// fmt.Printf("Results you Provided: \n%v\n", len(input))
		// fmt.Printf("Results you Provided: \n%v\n", s)
		// fmt.Println(s != "exit")
		fmt.Printf(">>>: ")
	}

	os.Exit(0)
}
