package ui

import (
	"fmt"
	"io"
	"os"

	"github.com/gosuri/racer/ui/speakeasy"
)

// DefaultPrompter is the default prompter for the package
var DefaultPrompter = NewPrompter()

// PromptString uses the default prompter to prompt the user for input when
// the string is missing in interactive mode
func PromptString(str *string, prompt string) {
	DefaultPrompter.PromptString(str, prompt)
}

// PromptString uses the default prompter to prompt the user for input and
// hides the input when the string is missing.  It is used for capturing sensitive
// data (passwords). Will not prompt when no interactive is true
func PromptHiddenString(str *string, prompt string) error {
	return DefaultPrompter.PromptHiddenString(str, prompt)
}

// Prompter represent an interactive prompter that captures inputs from the user
type Prompter struct {
	// Reader is the input reader to read user input from
	Reader io.Reader

	// Writer is the output writer to present the prompt
	Writer io.Writer

	// NoInteractive when enabled disables prompting user
	NoInteractive bool
}

// NewPrompter returns a new instance of a prompter
func NewPrompter() *Prompter {
	return &Prompter{Reader: os.Stdin, Writer: os.Stderr}
}

// PromptString prompts the user for input when the string is missing when in interactive mode
func (a *Prompter) PromptString(str *string, prompt string) {
	if a.NoInteractive {
		return
	}
	if len(*str) == 0 {
		fmt.Fprint(a.Writer, prompt)
		fmt.Fscanln(a.Reader, str)
	}
}

// PromptHiddenString prompts the user for input and hides the input when the string is missing.
// It used for capturing sensitive data (passwords). Will not prompt when no interactive is true
func (a *Prompter) PromptHiddenString(str *string, prompt string) error {
	if a.NoInteractive || len(*str) == 0 {
		return nil
	}
	input, err := speakeasy.Ask(prompt)
	if err != nil {
		return err
	}
	*str = input
	return nil
}
