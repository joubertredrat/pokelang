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

type PositionGame struct {
	coordinateNorthSouth int
	coordinateEastWest   int
	positions            map[string]bool
}

func NewPositionGame() Position {
	return &PositionGame{
		coordinateNorthSouth: 0,
		coordinateEastWest:   0,
		positions:            map[string]bool{},
	}
}

func (p *PositionGame) MoveToNorth() {
	p.coordinateNorthSouth++
}

func (p *PositionGame) MoveToSouth() {
	p.coordinateNorthSouth--
}

func (p *PositionGame) MoveToEast() {
	p.coordinateEastWest--
}

func (p *PositionGame) MoveToWest() {
	p.coordinateEastWest++
}

func (p *PositionGame) RegisterPosition() {
	ok := p.IsAlreadyVisited()
	if !ok {
		key := p.getCoordinatesKey()
		p.positions[key] = true
	}
}

func (p *PositionGame) IsAlreadyVisited() bool {
	key := p.getCoordinatesKey()
	_, ok := p.positions[key]
	return ok
}

func (p *PositionGame) getCoordinatesKey() string {
	return fmt.Sprintf("%d,%d", p.coordinateNorthSouth, p.coordinateEastWest)
}
