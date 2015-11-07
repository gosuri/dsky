package ui

import (
	"os"
	"testing"
)

func TestTitle_String(t *testing.T) {
	got := NewPrinter(os.Stdout).AddTitle("foo").AddTitle("bar").String()
	expect := "foo\n===\n\nbar\n===\n\n"
	if got != expect {
		t.Fatal("== expected\n", expect, "== got\n", got)
	}
}
