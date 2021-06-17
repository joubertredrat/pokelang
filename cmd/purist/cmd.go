package purist

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/joubertredrat/pokelang/app"
	"github.com/joubertredrat/pokelang/cmd"
)

func showIndex() {
	filename := filepath.Base(os.Args[0])

	fmt.Printf("Pokelang: Bora caçar pokemons com golang? \n\n")
	fmt.Printf("Como usar: \n")
	fmt.Printf("  %s [command] \n\n", filename)
	fmt.Printf("Comandos disponíveis: \n")
	fmt.Printf("  ajuda		Exibe a ajuda de como jogar \n")
	fmt.Printf("  capturar	Capturar os pokemons \n")
	fmt.Printf("\n")
	os.Exit(0)
}

func showErrorCommand(cmd string) {
	filename := filepath.Base(os.Args[0])

	fmt.Printf("Erro: comando \"%s\" desconhecido para \"%s\" \n", cmd, filename)
	os.Exit(255)
}

func showFlagRequiredCommand(flag string) {
	fmt.Printf("Erro: flag \"%s\" obrigatória para o comando \"capturar\" \n", flag)
	os.Exit(255)
}

func issetArg(arr []string, index int) bool {
	return (len(arr) > index)
}

func Execute() {
	if !issetArg(os.Args, 1) {
		showIndex()
	}

	switch os.Args[1] {
	case "capturar":
		if !issetArg(os.Args, 2) || !issetArg(os.Args, 3) || os.Args[2] != "--casas" {
			showFlagRequiredCommand("--casas")
		}

		position := app.NewGamePosition()
		storage := app.NewMemoryStorage()
		steps := os.Args[3]

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
	case "ajuda":
		cmd.ShowHelp()
	default:
		showErrorCommand(os.Args[1])
	}
}
