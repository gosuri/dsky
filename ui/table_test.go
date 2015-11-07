package ui

import (
	"testing"
)

func TestTableFormat(t *testing.T) {
	table := NewTable("foo", "bar")
	table.Seperator = " "
	table.AddRow("f1", "b1")

	got := table.Format()
	expect := "foo bar \nf1  b1  \n"
	if got != expect {
		t.Fatalf("want: %q got: %q", expect, got)
	}
}

func ExampleTable() {
	type hacker struct {
		Name string
		Bio  string
	}
	var hackers = []hacker{
		{"Ada Lovelace", "Ada was a British mathematician and writer, chiefly known for her work on Charles Babbage's early mechanical general-purpose computer, the Analytical Engine"},
		{"Alan Turing", "Alan was a British pioneering computer scientist, mathematician, logician, cryptanalyst and theoretical biologist"},
	}
	table := NewTable("NAME", "BIO")
	table.MaxCellWidth = 20
	for _, hacker := range hackers {
		table.AddRow(hacker.Name, hacker.Bio)
	}
	NewStdUI().Printer().Add(table).Print()
}
