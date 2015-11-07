package ui

import (
	"bytes"
)

var (
	// TitleUnderliner is the underline character for the title
	TitleUnderliner = "="
)

// Title is a UI component that renders a title
type Title struct {
	text string
}

// Format returns the formated string of the title
func (t *Title) Format() string {
	var buf bytes.Buffer
	buf.WriteString(t.text + "\n")
	for i := 0; i < len(t.text); i++ {
		buf.WriteString(TitleUnderliner)
	}
	buf.WriteString("\n")
	return buf.String()
}
