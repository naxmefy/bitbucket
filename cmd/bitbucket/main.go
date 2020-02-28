package main

import (
	"fmt"
	"log"
	"os"

	"github.com/naxmefy/bitbucket/internal/ui"
	"github.com/urfave/cli/v2"
)

//VERSION of Bitbucket CLI
const VERSION = "1.0.0"

func main() {
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "lang",
				Value: "english",
				Usage: "language for the greeting",
			},
		},
		Commands: []*cli.Command{
			{
				Name: "ui",
				// Aliases: []string{"c"},
				Usage: "start the bitbucket ui tool",
				Action: func(c *cli.Context) error {
					return ui.ShowUI()
				},
			},
		},
		Version: VERSION,
		Action: func(c *cli.Context) error {
			name := "Naxmefy"
			if c.NArg() > 0 {
				name = c.Args().Get(0)
			}
			if c.String("lang") == "spanish" {
				fmt.Println("Hola", name)
			} else {
				fmt.Println("Hello", name)
			}
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
