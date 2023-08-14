package timer

import (
	"fmt"
	"strings"
	"time"
)
// CURRENT ISSUE:
// Bar does not print the bar.graph. If it does it does not print it correctly if time period is anything else than 1 second.
//



type Bar struct {
	percent int64  // progress percentage
	cur     int64  // current progress
	total   int64  // total value for progress
	rate    string // the actual progress bar to be printed
	graph   string // the fill value for progress bar
}

func (bar *Bar) NewOption(start, total int64) {
	bar.cur = start
	bar.total = total
	if bar.graph == "" {
		bar.graph = "#"
	}
	//if full bar is always 100 symbols
	// percent * 10 -> current bar rate
	// TODO: This does not work right now. IDK why
	bar.percent = bar.getPercent()
	for i := 0; i < int(bar.percent); i += 2 {	
		bar.rate += bar.graph
	}
}

// This works for ANY total not just 100.
func (bar *Bar) getPercent() int64 {
	return int64((float32(bar.cur) / float32(bar.total)) * 100)
}

func (bar *Bar) Play(cur int64) {
	bar.cur = cur
	last := bar.percent
	bar.percent = bar.getPercent()
	if bar.percent != last && bar.percent%2 == 0 {
		graph := strings.Repeat("#", int(bar.cur / 50))
		fmt.Println(bar.cur)
		fmt.Println(graph)
		bar.rate += graph
	}
	fmt.Printf("\r[%-50s]%3d%% %8d/%d", bar.rate, bar.percent, bar.cur, bar.total)
}

func (bar *Bar) Finish() {
	fmt.Println()
}

// If we do the timer.Minute thing it wont update until the first minute or time period has passed.
func ProgressBar(total int64, ch <-chan time.Time) {
	var bar Bar
	bar.NewOption(0, total)
	bar.Play(int64(0))
	// bar.Play(int64(total))
	// msg := <-ch
	for i := 0; i < int(total); i++ {
		time.Sleep(1 * time.Second)
		bar.Play(int64(i))
	}
	bar.Finish()

}
