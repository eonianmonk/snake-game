package layout

import (
	"github.com/eonianmonk/snake-game/game"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type GameGrid struct {
	*tview.Grid
	cells [][]*tview.Box
}

func NewGameGrid(size int) *GameGrid {

	ph := make([]int, size)
	for i := range ph {
		ph[i] = -1
	}

	grid := tview.NewGrid().
		SetRows(ph...).
		SetColumns(ph...)

	cells := make([][]*tview.Box, size)
	for i := range cells {
		cells[i] = make([]*tview.Box, size)
		for j := range cells[i] {

			cell := tview.NewBox()

			cell.SetBackgroundColor(BlankColor)

			cells[i][j] = cell
			grid.AddItem(cell, i, j, 1, 1, 0, 0, false)
		}
	}

	return &GameGrid{
		Grid:  grid,
		cells: cells,
	}
}

func (ui *GameGrid) Update(grid *game.Grid) {
	gridSize := len(ui.cells)

	for i := 0; i < gridSize; i++ {
		for j := 0; j < gridSize; j++ {
			var color tcell.Color
			switch grid.CheckCell(i, j) {
			case game.Food:
				color = FoodColor
			case game.SnakeBody:
				color = BodyColor
			case game.SnakeHead:
				color = HeadColor
			case game.BlankCell:
				color = BlankColor
			default:
				panic("failed to id cell type while updating ui")
			}
			ui.cells[i][j].SetBackgroundColor(color)
		}
	}
}
