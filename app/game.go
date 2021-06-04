package app

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

const (
	STEP_NORTH = "N"
	STEP_SOUTH = "S"
	STEP_EAST  = "L"
	STEP_WEST  = "O"
)

var (
	InvalidStepsError       = errors.New("Invalid steps got")
	GameAlreadyStartedError = errors.New("Game already started")
	GameNotFinishedError    = errors.New("Game not finished")
)

type Game interface {
	SetSteps(steps string) error
	Start() error
	GetPokemonCount() (uint, error)
}

type PokemonGame struct {
	storage  Storage
	position Position
	steps    string
	finished bool
}

func NewPokemonGame(storage Storage, position Position, steps string) (Game, error) {
	p := &PokemonGame{
		storage:  storage,
		position: position,
		finished: false,
	}

	err := p.SetSteps(steps)

	if err != nil {
		return nil, err
	}

	return p, nil
}

func (p *PokemonGame) SetSteps(steps string) error {
	ok := p.isValidSteps(steps)
	if !ok {
		return InvalidStepsError
	}

	p.steps = steps
	return nil
}

func (p *PokemonGame) Start() error {
	if p.finished {
		return GameAlreadyStartedError
	}

	p.storage.AddOne()
	p.position.RegisterPosition()

	for _, step := range strings.Split(p.steps, "") {
		if STEP_NORTH == step {
			p.position.MoveToNorth()
		}

		if STEP_SOUTH == step {
			p.position.MoveToSouth()
		}

		if STEP_EAST == step {
			p.position.MoveToEast()
		}

		if STEP_WEST == step {
			p.position.MoveToWest()
		}

		if !p.position.IsAlreadyVisited() {
			p.storage.AddOne()
			p.position.RegisterPosition()
		}
	}

	p.finished = true
	return nil
}

func (p *PokemonGame) GetPokemonCount() (uint, error) {
	if !p.finished {
		return 0, GameNotFinishedError
	}

	return p.storage.Count(), nil
}

func (p *PokemonGame) isValidSteps(steps string) bool {
	matched, _ := regexp.MatchString(p.getStepsRegex(), steps)

	return !matched
}

func (p *PokemonGame) getStepsRegex() string {
	return fmt.Sprintf(`[^%s%s%s%s]{1,}`, STEP_NORTH, STEP_SOUTH, STEP_EAST, STEP_WEST)
}
