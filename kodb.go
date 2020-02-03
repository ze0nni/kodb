package main

import (
	"log"
	"os"
	"sort"

	"github.com/urfave/cli"
)

func main() {
	app := &cli.App{
		Name:  "kodb",
		Usage: "Static database",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "driver, d",
				Value: "in-memory",
				Usage: "Name of database driver",
			},
			&cli.StringFlag{
				Name:  "source, s",
				Usage: "Drivers connection string",
			},
		},
		Commands: []*cli.Command{
			{
				Name:    "interact",
				Aliases: []string{"i"},
				Usage:   "Run in interact mode",
				Action: func(c *cli.Context) error {
					return nil
				},
			},
			{
				Name:  "web",
				Usage: "Run as web server",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "host",
						Value: "127.0.0.1",
						Usage: "Hostname",
					},
					&cli.StringFlag{
						Name:  "port",
						Value: "80",
						Usage: "Port",
					},
				},
				Action: func(c *cli.Context) error {
					return nil
				},
			},
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
