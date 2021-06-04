package app_test

import (
	"testing"

	"github.com/joubertredrat/pokelang/app"
)

func TestGaming(t *testing.T) {
	tests := []struct {
		name          string
		steps         string
		countExpected int
	}{
		{
			name:          "Test input1",
			steps:         "N",
			countExpected: 2,
		},
		{
			name:          "Test input2",
			steps:         "NLSO",
			countExpected: 4,
		},
		{
			name:          "Test input3",
			steps:         "NSNSNSNSNS",
			countExpected: 2,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			game, _ := app.NewPokemonGame(
				app.NewMemoryStorage(),
				app.NewGamePosition(),
				test.steps,
			)

			game.Start()
			countGot, _ := game.GetPokemonCount()

			if uint(test.countExpected) != countGot {
				t.Errorf("game.GetPokemonCount() expected %d, got %d", test.countExpected, countGot)
			}
		})
	}
}

func TestInvalidSteps(t *testing.T) {
	_, err := app.NewPokemonGame(
		app.NewMemoryStorage(),
		app.NewGamePosition(),
		"NESAO",
	)

	if err == nil {
		t.Errorf("pkg.NewPokemonGame() expected error, got nil")
	}
}

func TestGameAlreadyStarted(t *testing.T) {
	game, _ := app.NewPokemonGame(
		app.NewMemoryStorage(),
		app.NewGamePosition(),
		"NLSO",
	)

	game.Start()
	err := game.Start()

	if err == nil {
		t.Errorf("game.Start() expected error, got nil")
	}
}

func TestGameNotFinished(t *testing.T) {
	game, _ := app.NewPokemonGame(
		app.NewMemoryStorage(),
		app.NewGamePosition(),
		"NLSO",
	)

	_, err := game.GetPokemonCount()

	if err == nil {
		t.Errorf("game.GetPokemonCount() expected error, got nil")
	}
}
