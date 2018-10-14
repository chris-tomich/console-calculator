package main

import (
	"fmt"

	"github.com/chris-tomich/console-calculator/console"
)

func main() {
	c := console.NewConsole()

	var input string
	var err error

	for input != "exit" && input != "quit" {
		input, err = c.Prompt()

		if err != nil {
			fmt.Println(err)

			return
		}

		c.Write(input)
	}
}
