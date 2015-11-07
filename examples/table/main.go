package main

import (
	"github.com/gosuri/racer/ui"
)

type hacker struct {
	Name string
	Bio  string
}

var hackers = []hacker{
	{"Ada Lovelace", "Ada was a British mathematician and writer, chiefly known for her work on Charles Babbage's early mechanical general-purpose computer, the Analytical Engine"},
	{"Alan Turing", "Alan was a British pioneering computer scientist, mathematician, logician, cryptanalyst and theoretical biologist"},
}

func main() {
	table := ui.NewTable("NAME", "BIO")
	table.MaxCellWidth = 20
	for _, hacker := range hackers {
		table.AddRow(hacker.Name, hacker.Bio)
	}
	ui.NewStdUI().Printer().Add(table).Print()
}
