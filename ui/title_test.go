package ui

import (
	"testing"
)

func TestTitle_String(t *testing.T) {
	got := NewPrinter().AddTitle("foo").AddTitle("bar").String()
	expect := "foo\n===\n\nbar\n===\n\n"
	if got != expect {
		t.Fatal("== expected\n", expect, "== got\n", got)
	}
}
