package main

import (
	"fmt"
	"strings"

	"github.com/chris-tomich/console-calculator/calc"
	"github.com/chris-tomich/console-calculator/console"
)

func main() {
	c := console.New()

	var output string
	var input string
	var err error

	p := calc.NewParser()
	isClosed := false

	for !isClosed {
		input, err = c.Prompt()

		if err != nil {
			fmt.Println(err)

			return
		}

		if isClosed = !evaluateMgmt(input); !isClosed {
			output, err = p.Evaluate(input)

			if err != nil {
				fmt.Println(err)

				return
			}

			fmt.Println(output)
		}
	}
}

// evaluateMgmt will evaluate all the management commands (rather than feature commands) to control the utility.
func evaluateMgmt(input string) bool {
	switch strings.ToLower(input) {
	case "exit":
		return false
	case "quit":
		return false
	}

	return true
}
