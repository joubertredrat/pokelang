package urfavecli

import (
	"log"
	"os"

	"github.com/joubertredrat/pokelang/app"
	"github.com/joubertredrat/pokelang/cmd"
	"github.com/urfave/cli/v2"
)

func Execute() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:    "ajuda",
				Aliases: []string{},
				Usage:   "Exibe a ajuda de como jogar",
				Action: func(c *cli.Context) error {
					cmd.ShowHelp()
					return nil
				},
			},
			{
				Name:    "capturar",
				Aliases: []string{},
				Usage:   "Capturar os pokemons",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "casas",
						Aliases:  []string{},
						Value:    "",
						Usage:    "Casas para capturar os pokemons",
						Required: true,
					},
				},
				Action: func(c *cli.Context) error {
					position := app.NewGamePosition()
					storage := app.NewMemoryStorage()
					steps := c.String("casas")

					game, errNewGame := app.NewPokemonGame(storage, position, steps)

					if errNewGame != nil {
						cmd.ShowErrorSteps()
					}

					errStart := game.Start()

					if errStart != nil {
						cmd.ShowInternalError()
					}

					count, errCount := game.GetPokemonCount()

					if errCount != nil {
						cmd.ShowInternalError()
					}

					cmd.ShowResult(count)
					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
