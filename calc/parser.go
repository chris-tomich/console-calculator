package calc

import (
	"strconv"
)

// NewParser will create a new calculator parser.
func NewParser() *Parser {
	p := &Parser{}

	return p
}

// Parser is a parser for the basic calculator language.
type Parser struct {
}

// Evaluate will return the value of the output and an error if the command failed for the given input.
func (p *Parser) Evaluate(input string) (string, error) {
	if input == "" {
		return "noop", nil
	}

	seperatedFields := FieldsFuncWithSeparator(input, SeparatorTester("/*-+"), SeparatorTester(" "))

	ops := make([]Operation, 0, 10)

	for i, field := range seperatedFields {
		switch (field) {
		case "+":
			append(ops, &Addition{})
		case "-":
			append(ops, &Subtraction{})
		case "*":
			append(ops, &Multiplication{})
		case "/":
			append(ops, &Division{})
		default:
			append(ops, &Constant{ A: strconv.Atoi(field) })
		}
	}

	for _, f := range ops {
		println(f)
	}

	return "completed", nil
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

	if start != len(s) {
		spans = append(spans, s[start:])
	}

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
