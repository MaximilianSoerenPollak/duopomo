package timer

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"time"

	cli "github.com/urfave/cli/v2"
)

type pomotimer struct {
	Name   string
	Length int  // Lenght in minutes
	Break  bool // Break timer or not
}

type command struct {
	Message string
}

type trackingTicker struct {
	Ticker     *time.Ticker
	PassedTime int
}

func ConvertTimerToInt(c *cli.Context) int {
	timerInt, err := strconv.Atoi(c.String("timer"))
	if err != nil {
		log.Println("Error converting Timer")
		log.Fatal(err)
	}
	return timerInt
}

func calcTimeRemaining(t *trackingTicker, totalTime int) int {
	return t.PassedTime - totalTime
}

func SoloTimer(c *cli.Context) error {
	timerInt, err := strconv.Atoi(c.String("timer"))
	if err != nil {
		log.Println("Could not convert timer to integer")
		log.Fatal(err)
	}
	t := trackingTicker{Ticker: time.NewTicker(time.Second), PassedTime: 0}
	// Define channels
	done := make(chan bool)
	stopchan := make(chan bool)
	restartchan := make(chan bool)
	// Initiate Bar
	bar := BarInit()
	onePercent := ((timerInt * 60) / 100) //Total time in minutes * 60 -> make it seconds -> / 100 to make it 1%
	go func() {
		for {
			select {
			case <-done:
				return
			case <-t.Ticker.C:
				// Increasing the bar by 1% each time we hit 1% of the total.
				if t.PassedTime%onePercent == 0 {

					bar.Play()
					t.PassedTime++
				} else {
					t.PassedTime++
				}
			case <-stopchan:
				t.Ticker.Stop()
			case <-restartchan:
				t.Ticker.Reset(time.Second)
			}
		}
	}()
	go func() {
		if t.PassedTime == timerInt {
			fmt.Println("Inside the passedtime == timerint")
			done <- true
		}
	}()
	for {
		if <-done {
			bar.Finish()
			break
		}
	}
	
	return nil
}

func Timer(c *cli.Context) error {
	fmt.Println("I'm the timer, now doing my thingy.")
	if c.Args().Len() > 0 {
		return errors.New("no arguments expected, please use flags")
	}
	if c.String("type") == "" {
		return errors.New("need to provide a type. 'solo' or 'duo'")
	}
	if !c.IsSet("timer") {
		return errors.New("need to provide a timer time. the provided number is how long the timer is in minutes")
	}
	time := ConvertTimerToInt(c)
	if time < 2 {
		return errors.New("minimum timer length is 2 minutes. please provide a time that is longer or equal at least that")
	}
	if c.String("type") == "solo" || c.String("type") == "s" {
		SoloTimer(c)
	}
	return nil
}
