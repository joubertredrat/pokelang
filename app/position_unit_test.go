package app_test

import (
	"testing"

	"github.com/joubertredrat/pokelang/app"
)

func TestPositionStart(t *testing.T) {
	position := app.NewPositionGame()

	if position.IsAlreadyVisited() {
		t.Error("position.IsAlreadyVisited() expected false, got true")
	}

	position.RegisterPosition()

	if !position.IsAlreadyVisited() {
		t.Error("position.IsAlreadyVisited() expected true, got false")
	}
}

func TestPositionsMove(t *testing.T) {
	tests := []struct {
		name        string
		getPosition func() app.Position
	}{
		{
			name: "Testing moving to North",
			getPosition: func() app.Position {
				position := app.NewPositionGame()
				position.MoveToNorth()
				return position
			},
		},
		{
			name: "Testing moving to South",
			getPosition: func() app.Position {
				position := app.NewPositionGame()
				position.MoveToSouth()
				return position
			},
		},
		{
			name: "Testing moving to East",
			getPosition: func() app.Position {
				position := app.NewPositionGame()
				position.MoveToEast()
				return position
			},
		},
		{
			name: "Testing moving to West",
			getPosition: func() app.Position {
				position := app.NewPositionGame()
				position.MoveToWest()
				return position
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			position := test.getPosition()

			if position.IsAlreadyVisited() {
				t.Error("position.IsAlreadyVisited() expected false, got true")
			}

			position.RegisterPosition()

			if !position.IsAlreadyVisited() {
				t.Error("position.IsAlreadyVisited() expected true, got false")
			}
		})
	}
}
