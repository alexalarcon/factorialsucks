package main

import (
	"log"
	"os"
	"time"

	"github.com/urfave/cli/v2"

	"github.com/alejoar/factorialsucks/factorial"
)

var today time.Time = time.Now()

func main() {
	log.SetFlags(0)
	app := &cli.App{
		Name:            "factorialsucks",
		Usage:           "FactorialHR auto clock in for the whole month from the command line",
		Version:         "2.1",
		Compiled:        time.Now(),
		UsageText:       "factorialsucks [options]",
		HideHelpCommand: true,
		HideVersion:     true,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "email",
				Aliases: []string{"e"},
				Usage:   "you factorial email address",
			},
			&cli.IntFlag{
				Name:        "year",
				Aliases:     []string{"y"},
				Usage:       "clock-in year `YYYY`",
				DefaultText: "current year",
				Value:       2025,
			},
			&cli.IntFlag{
				Name:        "month",
				Aliases:     []string{"m"},
				Usage:       "clock-in month `MM`",
				DefaultText: "current month",
				Value:       int(3),
			},
			&cli.StringFlag{
				Name:    "clock-in",
				Aliases: []string{"ci"},
				Usage:   "clock-in time `HH:MM`",
				Value:   "09:00",
			},
			&cli.StringFlag{
				Name:    "clock-out",
				Aliases: []string{"co"},
				Usage:   "clock-in time `HH:MM`",
				Value:   "18:00",
			},
			&cli.BoolFlag{
				Name:    "today",
				Aliases: []string{"t"},
				Usage:   "clock in for today only",
				Value:   false,
			},
			&cli.BoolFlag{
				Name:    "until-today",
				Aliases: []string{"ut"},
				Usage:   "clock in only until today",
				Value:   false,
			},
			&cli.BoolFlag{
				Name:    "dry-run",
				Aliases: []string{"dr"},
				Usage:   "do a dry run without actually clocking in",
			},
			&cli.BoolFlag{
				Name:    "reset-month",
				Aliases: []string{"rm"},
				Usage:   "delete all shifts for the given month",
				Value:   false,
			},
		},
		Action: factorialsucks,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func factorialsucks(c *cli.Context) error {
	var year, month int
	//email, password := readCredentials(c)
	email := "xxxx_XXXx"
	password := "XxxXX_XXXX"
	today_only := c.Bool("today")
	if today_only {
		year = today.Year()
		month = int(today.Month())
	} else {
		year = c.Int("year")
		month = c.Int("month")
	}
	clock_in := c.String("clock-in")
	clock_out := c.String("clock-out")
	dry_run := c.Bool("dry-run")
	until_today := c.Bool("until-today")
	reset_month := c.Bool("reset-month")
	//reset_month = true

	client := factorial.NewFactorialClient(email, password, year, month, clock_in, clock_out, today_only, until_today)
	if reset_month {
		client.ResetMonth()
	} else {
		client.ClockIn(dry_run)
	}
	return nil
}
