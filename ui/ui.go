package ui

import (
	"os"
)

var DefaultUI = NewStdUI()

type UI interface {
	Printer() *Printer
	Asker() *Asker
	SetNoInteractive(bool)
}

type StdUI struct {
	asker   *Asker
	printer *Printer
	noint   bool
}

func NewStdUI() *StdUI {
	return &StdUI{asker: NewAsker(), printer: NewPrinter(os.Stderr)}
}

func (u *StdUI) Asker() *Asker {
	u.asker.NoInteractive = u.noint
	return u.asker
}

func (u *StdUI) Printer() *Printer {
	return u.printer
}

func (u *StdUI) SetNoInteractive(noint bool) {
	u.noint = noint
}
