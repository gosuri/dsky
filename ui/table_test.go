package ui

import (
	"testing"
)

func TestTableFormat(t *testing.T) {
	table := NewTable("foo", "bar")
	table.Seperator = " "
	table.AddRow("f1", "b1")

	got := table.String()
	expect := "foo bar \nf1  b1  \n"
	if got != expect {
		t.Fatalf("want: %q got: %q", expect, got)
	}
}
