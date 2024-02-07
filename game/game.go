package game

import (
	"time"
)

type GameRules struct {
	// rows/columns size
	TickInterval time.Duration
	GameSize     int
	// allow going through walls
	TransparantWalls bool
}

type Game struct {
	*Grid
	*Snake
	rules  *GameRules
	ticker *time.Ticker
	signal chan struct{}
	input  <-chan Direction
}

func NewGame(rules *GameRules, dirChan <-chan Direction, signal chan struct{}) *Game {
	grid := NewGrid(rules.GameSize)
	snake := NewSnake(rules.GameSize)
	return &Game{
		signal: signal,
		rules:  rules,
		Grid:   grid,
		Snake:  snake,
		ticker: time.NewTicker(rules.TickInterval),
	}
}

func (g *Game) Start() error {
	return g.eventLoop()
}

func (g *Game) eventLoop() error {
	for {
		select {
		case <-g.ticker.C:
			// game loop

			ok := g.Step()
			if !ok { // lost
				return nil
			}
			g.signal <- struct{}{}
		case dir := <-g.input:
			// change direction
			g.Snake.ChangeDirection(dir)
		}
	}
}

// calculates one step
// returns true if step is valid, false if hit itself/wall
func (g *Game) Step() bool {

	// cmpFn := func(a SnakePart, b SnakePart) bool {
	// 	return a.col == b.col && a.row == b.row
	// }
	//check if valid position
	row, col := g.Snake.CalculateNextStep()

	if row >= g.rules.GameSize || row <= 0 ||
		col >= g.rules.GameSize || col <= 0 {
		if g.rules.TransparantWalls {

			col = (col + g.rules.GameSize) % g.rules.GameSize
			row = (row + g.rules.GameSize) % g.rules.GameSize
		} else {
			return false // struck the wall
		}
	}

	ct := g.Grid.CheckCell(row, col)
	switch ct {
	case Food:
		g.Snake.StepOnce(true, true)
	case BlankCell:
		g.Snake.StepOnce(false, true)
	case SnakeBody:
		return false
	case SnakeHead:
		// how did you get here???
		return false
	default:
		panic("unknown cell type!")
	}
	// update Grid
	g.Grid.Update(g.Snake, ct == Food)

	return true
}
