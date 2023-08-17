package timer

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"time"
	"bufio"
	"os"
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

func ConvertTimerToInt(c *cli.Context) int {
	timerInt, err := strconv.Atoi(c.String("timer"))
	if err != nil {
		log.Println("Error converting Timer")
		log.Fatal(err)
	}
	return timerInt
}


func SoloTimer(c *cli.Context) error {
	timerInt, err := strconv.Atoi(c.String("timer"))
	if err != nil {
		log.Println("Could not convert timer to integer")
		log.Fatal(err)
	}
	t := time.NewTicker(time.Second)
	done := make(chan bool)
	bar := BarInit()
	onePercent := ((timerInt * 60) / 100) //Total time in minutes * 60 -> make it seconds -> / 100 to make it 1% 
	go func() {
		i := 0
		for {
			select {
			case <-done:
				return
			case <-t.C:
				// Increasing the bar by 1% each time we hit 1% of the total.
				if i % onePercent == 0 { 
					bar.Play()
					i++
				}else{
					i++
				}
			}
		}
	}()
	time.After(time.Duration(timerInt) * 60 * time.Second)
	done <- true
	bar.Finish()
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
