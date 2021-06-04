package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/joubertredrat/pokelang/app"
)

func showIndex() {
	filename := filepath.Base(os.Args[0])

	fmt.Printf("Pokelang: Bora caçar pokemons com golang? \n\n")
	fmt.Printf("Como jogar: \n")
	fmt.Printf("  %s capturar --casas [posicao{1,n}] \n\n", filename)
	fmt.Printf("Nas casas deve ser informado \"N\" para mover para norte, \"S\" para mover para sul, \"L\" para mover para leste ou \"O\" para mover para oeste. \n")
	fmt.Printf("Para cada posição informada, será capturado o pokemon nesta casa, não sendo possível capturar duas vezes em uma casa já visitada. \n")
	fmt.Printf("Pelo menos uma posição deve ser informada. \n\n ")
	fmt.Printf("Exemplos: \n")
	fmt.Printf("  %s capturar --casas N \n", filename)
	fmt.Printf("  %s capturar --casas NLSO \n", filename)
	fmt.Printf("  %s capturar --casas NSNSNSNSNS \n", filename)
	fmt.Printf("\n")
	os.Exit(0)
}

func showErrorSteps() {
	fmt.Printf("\n\n")
	fmt.Printf(" [ALERTA] Você deve ter informado alguma posição errada, tente novamente \n")
	fmt.Printf("\n\n")
	os.Exit(255)
}

func showInternalError() {
	fmt.Printf("\n\n")
	fmt.Printf(" [ERRO] Algo de errado não deu certo, tente novamente \n")
	fmt.Printf("\n\n")
	os.Exit(255)
}

func showResult(count uint) {
	fmt.Printf(" [OK] Legal! Você capturou %d pokemons \n", count)
	os.Exit(0)
}

func issetArg(arr []string, index int) bool {
	return (len(arr) > index)
}

func Execute() {
	if !issetArg(os.Args, 1) || !issetArg(os.Args, 2) || !issetArg(os.Args, 3) {
		showIndex()
	}

	if os.Args[1] != "capturar" || os.Args[2] != "--casas" {
		showIndex()
	}

	position := app.NewGamePosition()
	storage := app.NewMemoryStorage()
	steps := os.Args[3]

	game, errNewGame := app.NewPokemonGame(storage, position, steps)

	if errNewGame != nil {
		showErrorSteps()
	}

	errStart := game.Start()

	if errStart != nil {
		showInternalError()
	}

	count, errCount := game.GetPokemonCount()

	if errCount != nil {
		showInternalError()
	}

	showResult(count)
}
