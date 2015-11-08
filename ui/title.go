package ui

import (
	"bytes"
)

// TitleUnderliner is the underline character for the title
var TitleUnderliner = "="

// Title is a UI component that renders a title
type Title struct {
	text string
}

func NewTitle(title string) *Title {
	return &Title{text: title}
}

// Format returns the formated string of the title
func (t *Title) String() string {
	var buf bytes.Buffer
	buf.WriteString(t.text + "\n")
	for i := 0; i < len(t.text); i++ {
		buf.WriteString(TitleUnderliner)
	}
	buf.WriteString("\n")
	return buf.String()
}
