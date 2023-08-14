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


func startTimer(t int) <-chan time.Time {
	ticker := time.NewTicker(time.Second)
	// not sure if this works?
	go func() <-chan time.Time {
				return ticker.C
	}()
	time.Sleep(time.Duration(t) * time.Second)
	ticker.Stop()
	fmt.Println("Wow timer is done")
	return nil
}

func SoloTimer(c *cli.Context) error {
	timerInt, err := strconv.Atoi(c.String("timer"))
	if err != nil {
		log.Println("Could not convert timer to integer")
		log.Fatal(err)
	}
	ch := startTimer(timerInt)
	ProgressBar(int64(timerInt), ch)
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
	if c.String("type") == "solo" || c.String("type") == "s" {
		SoloTimer(c)
	}
	return nil
}
