package layout

import (
	"github.com/eonianmonk/snake-game/game"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type GameGrid struct {
	*tview.Grid
	cells [][]*tview.Box
	game  *game.Game
}

func (gg *GameGrid) Draw(screen tcell.Screen) {
	for i := range gg.cells {
		for j := range gg.cells {
			gg.cells[i][j].Draw(screen)
		}
	}
}

func NewGameGrid(size int, game *game.Game) tview.Primitive {

	gridPlaceholder := make([]int, size)
	for i := range gridPlaceholder {
		gridPlaceholder[i] = -1
	}

	grid := tview.NewGrid().
		SetRows(gridPlaceholder...).
		SetColumns(gridPlaceholder...)

	cells := make([][]*tview.Box, size)
	for i := range cells {
		cells[i] = make([]*tview.Box, size)
		for j := range cells[i] {

			cell := tview.NewBox().
				SetBackgroundColor(tcell.ColorGray)
				//SetBorder(true).SetBorderAttributes(tcell.AttrBold).
				//SetTitle(fmt.Sprintf("%d %d", i, j))
			if (i+j)%2 == 0 {
				cell.SetBackgroundColor(tcell.ColorOlive)
			}

			cells[i][j] = cell
			grid.AddItem(cell, i, j, 1, 1, 1, 1, false)
		}
	}

	return &GameGrid{
		Grid:  grid,
		cells: cells,
		game:  game,
	}
}
