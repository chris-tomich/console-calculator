package console

import (
	"bufio"
	"fmt"
	"os"

	"github.com/pkg/errors"
)

// New will create a new Console object.
func New() *Console {
	console := &Console{}
	console.reader = bufio.NewReader(os.Stdin)

	return console
}

// Console in an abstraction of console interaction.
type Console struct {
	reader *bufio.Reader
}

// Prompt will create the prompt and wait for user input
func (console *Console) Prompt() (string, error) {
	fmt.Print(">")

	rawInput, hasMore, err := console.reader.ReadLine()

	if err != nil {
		return "", errors.Wrap(err, "issue reading from STDIN")
	}

	input := string(rawInput)

	if hasMore {
		rawInput, hasMore, err = console.reader.ReadLine()

		if err != nil {
			return "", errors.Wrap(err, "issue reading additional characters in buffer")
		}

		input += string(rawInput)
	}

	return input, nil
}

// Write will write the given output to the screen
func (console *Console) Write(output string) {
	fmt.Println(output)
}
