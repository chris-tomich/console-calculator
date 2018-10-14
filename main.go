package main

import (
	"fmt"
	"strings"

	"github.com/chris-tomich/console-calculator/console"
)

func main() {
	c := console.NewConsole()

	var input string
	var err error

	for CalculatorFunctions(input) {
		input, err = c.Prompt()

		if err != nil {
			fmt.Println(err)

			return
		}
	}
}

// CalculatorFunctions will perform the operation that has been given as user input.
func CalculatorFunctions(input string) bool {
	switch strings.ToLower(input) {
	case "exit":
		return false
	case "quit":
		return false
	}

	seperatedFields := FieldsFuncWithSeparator(input, SeparatorTester("/*-+"), SeparatorTester(" "))

	for _, field := range seperatedFields {
		fmt.Println("'" + field + "'")
	}

	return true
}

// FieldsFuncWithSeparator will separate a string given the separators but include the separators.
func FieldsFuncWithSeparator(s string, separatorsToInclude func(rune) bool, separatorsToExclude func(rune) bool) []string {
	spans := make([]string, 0, 32)

	start := 0

	for i, r := range s {
		if separatorsToInclude(r) {
			if start != i {
				spans = append(spans, s[start:i])
			}

			spans = append(spans, s[i:i+1])
			start = i + 1
		} else if separatorsToExclude(r) {
			if start != i {
				spans = append(spans, s[start:i])
			}

			start = i + 1
		}
	}

	spans = append(spans, s[start:])

	return spans
}

// SeparatorTester will return a tester for a string of separators.
func SeparatorTester(seps string) func(rune) bool {
	return func(r rune) bool {
		for _, sep := range seps {
			if r == sep {
				return true
			}
		}

		return false
	}
}
