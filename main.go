package main

import (
	"github.com/Bibob7/reqCon/pkg/engine"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	app := &cli.App{
		Name:  "reqcon",
		Usage: "make an explosive entrance",
		Action: func(c *cli.Context) error {
			url := ""
			if c.NArg() > 0 {
				url = c.Args().Get(0)
			}
			con := c.Int("concurrent")
			num := c.Int("number")
			engine.Run(url, con, num)
			return nil
		},
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:    "concurrent",
				Aliases: []string{"c"},
				Usage:   "number of concurrent requests",
				Value:   1,
			},
			&cli.IntFlag{
				Name:    "number",
				Aliases: []string{"n"},
				Usage:   "overall number of requests",
				Value:   1,
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
