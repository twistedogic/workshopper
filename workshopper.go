package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"

	"github.com/twistedogic/workshopper/config"
)

func main() {
	app := &cli.App{
		Name: "workshopper",
		Action: func(c *cli.Context) error {
			_, err := config.FromYAMLFile(c.Args().First())
			return err
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
