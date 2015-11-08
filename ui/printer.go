package ui

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

// Printer represents the output printer for the ui
type BufferedPrinter struct {
	// Writer is where the output should writer to
	Writer io.Writer

	comps []fmt.Stringer
}

// NewPrinter returns a pointer to a new printer object
func NewPrinter() *BufferedPrinter {
	return &BufferedPrinter{Writer: os.Stdout}
}

// Add adds the components to the printer
func (p *BufferedPrinter) Add(c fmt.Stringer) *BufferedPrinter {
	p.comps = append(p.comps, c)
	return p
}

// String returns the formmated string of the output
func (p *BufferedPrinter) String() string {
	var buf bytes.Buffer
	for _, c := range p.comps {
		buf.WriteString(c.String())
		buf.WriteString("\n")
	}
	return buf.String()
}

// Print prints the output to the writer
func (p *BufferedPrinter) Flush() {
	fmt.Fprintln(p.Writer, p)
}
