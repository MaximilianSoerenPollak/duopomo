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
	done := make(chan bool, 1)
	stopchan := make(chan bool, 1)
	restartchan := make(chan bool, 1)
	// Initiate Bar
	bar := BarInit()
	onePercent := ((timerInt * 60) / 100) //Total time in minutes * 60 -> make it seconds -> / 100 to make it 1%
	go func() {
		for {
			select {
			case <-done:
				return
			case <-stopchan:
				t.Ticker.Stop()
				fmt.Println("stopped ticker")
			case <-restartchan:
				t.Ticker.Reset(time.Second)
				fmt.Println("restarted timer")
			case <-t.Ticker.C:
				// Increasing the bar by 1% each time we hit 1% of the total.
				if t.PassedTime%onePercent == 0 {

					bar.Play()
					t.PassedTime++
				} else {
					t.PassedTime++
				}
			}
		}
	}()
	go func() {
		fmt.Println("\nExceuting the if foor loop thingy\n")
		select {
		case <-stopchan:
			time.Sleep(5 * time.Second)
			restartchan <- true
			fmt.Printf("\nRestarted channel\n")
		}
	}()
	// First GO ROUTINE that is executed
	go func() {
		for {
			if bar.percent >= 11 {
				fmt.Printf("\nPassed time before stop: %d\n", t.PassedTime)
				stopchan <- true
				fmt.Printf("\nPassed time after stop: %d\n", t.PassedTime)
				break
			}
		}
	}()
	select {
	case <-restartchan:
		fmt.Println("\nI found restart to be true\n")
		break
	}
	// fmt.Printf("Passed time after restart: %d\n", t.PassedTime)
	for {
		if bar.percent >= 100 {
			done <- true
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
