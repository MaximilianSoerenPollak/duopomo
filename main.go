package main 


import (
	cli	"github.com/urfave/cli/v2"
	timer "github.com/maximiliansoerenpollak/duopomo/timer"
	"os"
	"log"
)


func main() {
	app := cli.NewApp()
	app.Name = "Duopomo | Have solo or duo pomodoro sessions."
	app.Usage = "Select solo or duo sessions, your times and breaks. Then go concentrate."
	app.Commands =  []*cli.Command{
		{
			Name: "timer",
			HelpName: "timer",
			Action: timer.Timer,
			ArgsUsage: ` `,
			Usage: `Start the timer with solor/duo argument and the time you want`,
			Description:  `solo/duo -> what type of session do you want. Timer -> time in minutes it sould take.`,
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name: "type",
					Usage: "Determin if you want a solo or duo session",
			},
				&cli.UintFlag{
					Name: "timer",
					Usage: "How long (in minutes) should the timer be",
				},
			},
    },
	}

    if err := app.Run(os.Args); err != nil {
        log.Fatal(err)
    }}
