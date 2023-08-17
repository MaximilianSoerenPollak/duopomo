package timer

import (
	"fmt"
	"math"
	"strings"
)

type Bar struct {
	percent int64  // progress percentage
	cur     int64  // current progress
	total   int64  // total value for progress
	rate    string // the actual progress bar to be printed
	graph   string // the fill value for progress bar
}

func (bar *Bar) NewOption(start int64) {
	bar.cur = start
	bar.total = 100 
	if bar.graph == "" {
		bar.graph = "#"
	}
	bar.percent = bar.getPercent()
	for i := 0; i < int(bar.percent); i += 2 {
		bar.rate += bar.graph
	}
}

// This works for ANY total not just 100.
func (bar *Bar) getPercent() int64 {
	return int64((float32(bar.cur) / float32(bar.total)) * 100)
}

func (bar *Bar) Play() {
	bar.cur = bar.cur + 1
	last := bar.percent
	bar.percent = bar.getPercent()
	if bar.percent != last && bar.percent%2 == 0 {
		repeatNumber := int(math.Round(50 * (float64(bar.cur) / float64(bar.total))))
		graph := strings.Repeat("#", repeatNumber)
		bar.rate = graph
	}
	fmt.Printf("\r[%-50s]%3d%% %8d/%d", bar.rate, bar.percent, bar.cur, bar.total)
}
func BarInit() Bar {
	var bar Bar
	bar.NewOption(0)
	bar.Play()
	return bar
}
func (bar *Bar) Finish() {
	fmt.Println()
}
