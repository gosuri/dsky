package ui

import (
	"io"

	"github.com/gosuri/dsky/pkg/strutil/color"
	"github.com/gosuri/dsky/ui/pb"
	"github.com/gosuri/dsky/ui/prompt"
)

// DefaultUI is the default UI for the package
var DefaultUI = New()

// Printer returns the object that prints from the Default UI
func Printer() *BufferedPrinter {
	return DefaultUI.Printer()
}

// NewProgressBar returns a new progress bar object from the Default UI
func NewProgressBar(total int) *pb.ProgressBar {
	return DefaultUI.NewProgressBar(total)
}

// Color returns an instance of color from the Default UI
func Color() *color.Color {
	return DefaultUI.Color()
}

// UserInterface is the interface that defines UI interation functionality to the CLI
type UserInterface interface {
	// Printer must return the object that prints to the UI
	Printer() *BufferedPrinter

	// Prompter must return the object that prompts the user for input
	Prompt() *prompt.Prompter

	// Color returns an instance of color
	Color() *color.Color

	// NewProgressBar must return a new progress bar object
	NewProgressBar(total int) *pb.ProgressBar

	// SetNoInteractive must set the interactive mode for the UI
	SetNoInteractive(bool)

	// SetNoColor must set the color mode for output
	SetNoColor(nocolor bool)
}

// UI implements UserInterface and uses standard i/o
type UI struct {
	Writer io.Writer
	Reader io.Reader

	nocolor  bool
	color    *color.Color
	prompter *prompt.Prompter
	printer  *BufferedPrinter
	noint    bool
}

// NewStdUI returns an instance of the StdUI
func New() *UI {
	return &UI{prompter: prompt.New(), printer: NewPrinter()}
}

// Prompter returns the user prompter for the UI
func (u *UI) Prompt() *prompt.Prompter {
	u.prompter.NoInteractive = u.noint
	return u.prompter
}

// Printer returns the printer for the UI
func (u *UI) Printer() *BufferedPrinter {
	return u.printer
}

// SetNoInteractive sets the interactive mode for the UI
func (u *UI) SetNoInteractive(noint bool) {
	u.noint = noint
}

// SetNoColor sets if the output should have color
func (u *UI) SetNoColor(nocolor bool) {
	u.nocolor = nocolor
}

// NewProgressBar returns a new progress bar object
func (u *UI) NewProgressBar(total int) *pb.ProgressBar {
	return pb.New64(int64(total))
}

// Color returns an instance of color
func (u *UI) Color() *color.Color {
	if u.color == nil {
		u.color = &color.Color{}
	}
	if u.nocolor {
		u.color.Disable()
	}
	return u.color
}
