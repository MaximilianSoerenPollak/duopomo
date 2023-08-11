package timer 

import (
	"fmt"
	cli	"github.com/urfave/cli/v2"
	"errors"
	)


type pomotimer struct {
	Name string
	Length int // Lenght in minutes
	Break bool // Break timer or not
}


func SoloTimer(c *cli.Context) error {
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
	if c.String("type") == "solo" || c.String("type") == "s"	{
		SoloTimer(c)
	}
	return nil
}
