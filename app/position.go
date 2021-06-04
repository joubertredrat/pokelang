package app

import "fmt"

type Position interface {
	MoveToNorth()
	MoveToSouth()
	MoveToEast()
	MoveToWest()
	RegisterPosition()
	IsAlreadyVisited() bool
	getCoordinatesKey() string
}

type GamePosition struct {
	coordinateNorthSouth int
	coordinateEastWest   int
	positions            map[string]bool
}

func NewGamePosition() Position {
	return &GamePosition{
		coordinateNorthSouth: 0,
		coordinateEastWest:   0,
		positions:            map[string]bool{},
	}
}

func (p *GamePosition) MoveToNorth() {
	p.coordinateNorthSouth++
}

func (p *GamePosition) MoveToSouth() {
	p.coordinateNorthSouth--
}

func (p *GamePosition) MoveToEast() {
	p.coordinateEastWest--
}

func (p *GamePosition) MoveToWest() {
	p.coordinateEastWest++
}

func (p *GamePosition) RegisterPosition() {
	ok := p.IsAlreadyVisited()
	if !ok {
		key := p.getCoordinatesKey()
		p.positions[key] = true
	}
}

func (p *GamePosition) IsAlreadyVisited() bool {
	key := p.getCoordinatesKey()
	_, ok := p.positions[key]
	return ok
}

func (p *GamePosition) getCoordinatesKey() string {
	return fmt.Sprintf("%d,%d", p.coordinateNorthSouth, p.coordinateEastWest)
}
