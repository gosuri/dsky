package ui

import (
	"github.com/gosuri/racer/ui/pb"
)

// DefaultUI is the default UI for the package
var DefaultUI = NewStdUI()

// UI is the interface that defines UI interation functionality to the CLI
type UI interface {
	// Printer must return the object that prints to the UI
	Printer() *Printer
	// Prompter must return the object that prompts the user for input
	Prompter() *Prompter

	// SetNoInteractive must set the interactive mode for the UI
	SetNoInteractive(bool)

	// NewProgressBar must return a new progress bar object
	NewProgressBar(total int) *pb.ProgressBar

	// SetNoColor must set the color mode for output
	SetNoColor(nocolor bool)
}

// New returns a new instance of the default UI
func New() UI {
	return NewStdUI()
}

// StdUI implements UI and uses standard i/o
type StdUI struct {
	nocolor  bool
	prompter *Prompter
	printer  *Printer
	noint    bool
}

// NewStdUI returns an instance of the StdUI
func NewStdUI() *StdUI {
	return &StdUI{prompter: NewPrompter(), printer: NewPrinter()}
}

// Prompter returns the user prompter for the UI
func (u *StdUI) Prompter() *Prompter {
	u.prompter.NoInteractive = u.noint
	return u.prompter
}

// Printer returns the printer for the UI
func (u *StdUI) Printer() *Printer {
	u.printer.NoColor = u.nocolor
	return u.printer
}

// SetNoInteractive sets the interactive mode for the UI
func (u *StdUI) SetNoInteractive(noint bool) {
	u.noint = noint
}

// SetNoColor sets if the output should have color
func (u *StdUI) SetNoColor(nocolor bool) {
	u.nocolor = nocolor
}

// NewProgressBar returns a new progress bar object
func (u *StdUI) NewProgressBar(total int) *pb.ProgressBar {
	return pb.New64(int64(total))
}
