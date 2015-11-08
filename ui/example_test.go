package ui_test

import (
	"fmt"
	"github.com/gosuri/dsky/ui"
	"time"
)

func ExampleTable() {
	type hacker struct {
		Name     string
		Birthday string
		Bio      string
	}

	var hackers = []hacker{
		{"Ada Lovelace", "December 10, 1815", "Ada was a British mathematician and writer, chiefly known for her work on Charles Babbage's early mechanical general-purpose computer, the Analytical Engine"},
		{"Alan Turing", "23 June, 1912", "Alan was a British pioneering computer scientist, mathematician, logician, cryptanalyst and theoretical biologist"},
	}

	table := ui.NewTable("NAME", "BIRTHDATE", "BIO")
	table.MaxCellWidth = 20
	for _, hacker := range hackers {
		table.AddRow(hacker.Name, hacker.Birthday, hacker.Bio)
	}
	ui.Printer().Add(table).Flush()
}

func ExampleUI_NewProgressBar() {
	count := 5000
	bar := ui.NewProgressBar(count).Start()

	// show percents (by default already true)
	bar.ShowPercent = true

	// show bar (by default already true)
	bar.ShowBar = true

	// no need counters
	bar.ShowCounters = true

	bar.ShowTimeLeft = true

	// and start
	bar.Start()
	for i := 0; i < count; i++ {
		bar.Increment()
		time.Sleep(time.Millisecond)
	}
	bar.FinishPrint("The End!")
}

func ExampleUI_Color() {
	color := ui.Color()

	// Foreground
	fmt.Println(color.Black("black"))
	fmt.Println(color.Red("red"))
	fmt.Println(color.Green("green"))
	fmt.Println(color.Yellow("yellow"))
	fmt.Println(color.Blue("blue"))
	fmt.Println(color.Magenta("magenta"))
	fmt.Println(color.Cyan("cyan"))
	fmt.Println(color.White("white"))
	fmt.Println(color.Grey("grey"))

	// colored background
	fmt.Println(color.BlackBg("black background", ui.ColorWhite))
	fmt.Println(color.RedBg("red background"))
	fmt.Println(color.GreenBg("green background"))
	fmt.Println(color.YellowBg("yellow background"))
	fmt.Println(color.BlueBg("blue background"))
	fmt.Println(color.MagentaBg("magenta background"))
	fmt.Println(color.CyanBg("cyan background"))
	fmt.Println(color.WhiteBg("white background"))

	// Emphasis
	fmt.Println(color.Bold("bold"))
	fmt.Println(color.Dim("dim"))
	fmt.Println(color.Italic("italic"))
	fmt.Println(color.Underline("underline"))
	fmt.Println(color.Inverse("inverse"))
	fmt.Println(color.Hidden("hidden"))
	fmt.Println(color.Strikeout("strikeout"))

	// Mix and match
	fmt.Println(color.Green("bold green with white background", ui.StyleBold, ui.ColorWhiteBg))
	fmt.Println(color.Red("underline red", ui.StyleUnderline))
	fmt.Println(color.Yellow("dim yellow", ui.StyleDim))
	fmt.Println(color.Cyan("inverse cyan", ui.StyleInverse))
	fmt.Println(color.Blue("bold underline dim blue", ui.StyleBold, ui.StyleUnderline, ui.StyleDim))
}
