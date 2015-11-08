package main

import (
	"github.com/gosuri/dsky/ui"
)

type hacker struct {
	Name     string
	Birthday string
	Bio      string
}

var hackers = []hacker{
	{"Ada Lovelace", "December 10, 1815", "Ada was a British mathematician and writer, chiefly known for her work on Charles Babbage's early mechanical general-purpose computer, the Analytical Engine"},
	{"Alan Turing", "23 June, 1912", "Alan was a British pioneering computer scientist, mathematician, logician, cryptanalyst and theoretical biologist"},
}

func main() {
	table := ui.NewTable("NAME", "BIRTHDATE", "BIO")
	table.MaxCellWidth = 20
	for _, hacker := range hackers {
		table.AddRow(hacker.Name, hacker.Birthday, hacker.Bio)
	}
	ui.Printer().Add(table).Flush()
}
