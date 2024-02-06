package game

import (
	"time"

	linkedlist "github.com/eonianmonk/snake-game/pkg/linked_list"
	"github.com/rivo/tview"
)

type GameRules struct {
	// rows/columns size
	tickInterval time.Duration
	gameSize     int
	// allow going through walls
	//transparantWalls bool
}

type Game struct {
	rules  *GameRules
	app    *tview.Application
	grid   *Grid
	snake  *Snake
	ticker *time.Ticker
	input  <-chan Direction
}

func NewGame(rules *GameRules, app *tview.Application, dirChan <-chan Direction) *Game {
	grid := NewGrid(rules.gameSize)
	snake := &Snake{
		length:  1,
		bodyDir: None,
		headDir: Right,
		snake: linkedlist.NewLinkedList[SnakePart](SnakePart{
			row:    rules.gameSize / 2,
			col:    rules.gameSize / 2,
			facing: Right,
		}),
	}
	return &Game{
		rules:  rules,
		app:    app,
		grid:   grid,
		snake:  snake,
		ticker: time.NewTicker(rules.tickInterval),
	}
}

func (g *Game) Start() {
	go g.eventLoop()
}

func (g *Game) eventLoop() error {
	for {
		select {
		case <-g.ticker.C:
			g.grid.Lock()
			// game loop
			continue
		case dir := <-g.input:
			// change direction
			g.snake.ChangeDirection(dir)
		}
	}
	return nil
}

// calculates one step
// returns true if step is valid, false if hit itself/wall
func (g *Game) StepOnce() bool {

	// cmpFn := func(a SnakePart, b SnakePart) bool {
	// 	return a.col == b.col && a.row == b.row
	// }
	g.grid.Lock()
	defer g.grid.Unlock()
	//check if valid position

	//if cell is with food - add cell to the head
	//else remove tail value and add value ahead
	return true
}
