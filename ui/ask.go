package ui

import (
	"fmt"
	"io"
	"os"
)

var (
	DefaultAsker = NewAsker()
)

func AskMissingString(str *string, label string) {
	DefaultAsker.AskMissingString(str, label)
}

type Asker struct {
	Reader        io.Reader
	Writer        io.Writer
	NoInteractive bool
}

func NewAsker() *Asker {
	return &Asker{Reader: os.Stdin, Writer: os.Stderr}
}

// AskMissingString prompt the user for input when the string is null when in interactive mode
func (a *Asker) AskMissingString(str *string, prompt string) {
	if a.NoInteractive {
		return
	}
	if len(*str) == 0 {
		fmt.Fprint(a.Writer, prompt)
		fmt.Fscanln(a.Reader, str)
	}
}
