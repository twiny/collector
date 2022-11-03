package main

import (
	"log"
	"os"

	"github.com/twiny/collector/cmd/collect/api"
	"github.com/urfave/cli/v2"
)

// here we go!
func main() {
	app := &cli.App{
		Name:     "Collector",
		HelpName: "collector",
		Usage:    "Website scraper ",
		Version:  "dev-v0.0.0",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "config",
				Aliases:  []string{"c"},
				Usage:    "`path` to config file",
				Required: false,
				Value:    "./config/config.yaml",
			},
		},
		Action: func(c *cli.Context) error {
			col, err := api.NewAPI(c.String("config"))
			if err != nil {
				return err
			}
			go col.Close()

			return col.Collect()
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Println(err)
		return
	}
}
