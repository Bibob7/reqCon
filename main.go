package main

import (
	"github.com/Bibob7/reqCon/pkg/config"
	"github.com/Bibob7/reqCon/pkg/engine"
	"github.com/Bibob7/reqCon/pkg/http"
	"github.com/Bibob7/reqCon/pkg/result"
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
			body, err := con.Body()
			if err != nil {
				return err
			}
			req := http.NewRequest(
				con.RequestConfig.Method,
				con.RequestConfig.Url,
				con.RequestConfig.Header,
				body,
			)
			runner := engine.NewRunner(&http.Client{}, result.NewResultHandler())
			return runner.Run(engine.Config{
				ReqNum:  con.ReqNum,
				ConcNum: con.ConcNum,
				Request: req,
			})
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
