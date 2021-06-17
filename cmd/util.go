package cmd

import (
	"fmt"
	"os"
	"path/filepath"
)

func ShowHelp() {
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

func ShowErrorSteps() {
	fmt.Printf(" [ALERTA] Você deve ter informado alguma posição errada, tente novamente \n")
	os.Exit(255)
}

func ShowInternalError() {
	fmt.Printf(" [ERRO] Algo de errado não deu certo, tente novamente \n")
	os.Exit(255)
}

func ShowResult(count uint) {
	fmt.Printf(" [OK] Legal! Você capturou %d pokemons \n", count)
	os.Exit(0)
}
