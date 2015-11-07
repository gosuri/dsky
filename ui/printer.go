package ui

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

// Component in the interface that UI components need to implement
type Component interface {
	Format() string
}

// Printer represents the output printer for the ui
type UIPrinter struct {
	// Writer is where the output should writer to
	Writer io.Writer

	comps []Component
}

// NewPrinter returns a pointer to a new printer object
func NewPrinter() *UIPrinter {
	return &UIPrinter{Writer: os.Stdout}
}

// Add adds the components to the printer
func (p *UIPrinter) Add(c Component) *UIPrinter {
	p.comps = append(p.comps, c)
	return p
}

// AddTitle Adds a Title to the printer
func (p *UIPrinter) AddTitle(title string) *UIPrinter {
	return p.Add(&Title{text: title})
}

// String returns the formmated string of the output
func (p *UIPrinter) String() string {
	var buf bytes.Buffer
	for _, c := range p.comps {
		buf.WriteString(c.Format())
		buf.WriteString("\n")
	}
	return buf.String()
}

// Print prints the output to the writer
func (p *UIPrinter) Print() {
	fmt.Fprintln(p.Writer, p.String())
}
