package main

import (
	"github.com/gosuri/dsky/ui"

	"sync"

	"github.com/sethgrid/multibar"

	"time"
)

func main() {
	bars, _ := multibar.New()
	count1, count2 := 150, 200
	bar1 := bars.MakeBar(count1, "bar1")
	//bar1.ShowPercent = true
	bar2 := bars.MakeBar(count2, "bar2")
	//bar2.ShowPercent = true

	go bars.Listen()

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i <= count1; i++ {
			bar1(i)
			time.Sleep(time.Millisecond * 15)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i <= count2; i++ {
			bar2(i)
			time.Sleep(time.Millisecond * 15)
		}
	}()
	wg.Wait()
}

func main1() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		bar(5000)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		bar(5000)
	}()

	wg.Wait()
}

func bar(count int) {
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
