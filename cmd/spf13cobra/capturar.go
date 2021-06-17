package spf13cobra

import (
	"github.com/joubertredrat/pokelang/app"
	"github.com/joubertredrat/pokelang/cmd"
	"github.com/spf13/cobra"
)

var steps string

func NewCapturarCommand() *cobra.Command {
	capturar := &cobra.Command{
		Use:   "capturar",
		Short: "Capturar os pokemons",
		Run: func(c *cobra.Command, args []string) {
			position := app.NewGamePosition()
			storage := app.NewMemoryStorage()

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
		},
	}

	capturar.Flags().StringVarP(&steps, "casas", "", "", "Casas para capturar os pokemons")
	capturar.MarkFlagRequired("casas")

	return capturar
}
