package main

import (
	"github.com/gosuri/racer/ui"

	"time"
)

func main() {
	count := 5000
	bar := ui.NewProgressBar(count)

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
