package main

import (
	"github.com/Bibob7/reqCon/pkg/engine"
	"github.com/Bibob7/reqCon/pkg/engine/config"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	app := &cli.App{
		Name:  "reqcon",
		Usage: "make an explosive entrance",
		Action: func(c *cli.Context) error {
			configPath := ""
			if c.NArg() > 0 {
				configPath = c.Args().Get(0)
			}
			con, err := config.FromFile(configPath)
			if err != nil {
				return err
			}
			return engine.Run(*con)
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
