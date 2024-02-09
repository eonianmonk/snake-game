package layout

import (
	"strconv"
	"strings"

	"github.com/eonianmonk/snake-game/game"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type GameUI struct {
	app          *tview.Application
	gameGrid     *GameGrid
	dirChan      chan game.Direction
	displayScore int
}

func drawScore(score *int) DrawFunc {
	return func(screen tcell.Screen, x, y, width, height int) (int, int, int, int) {
		underscoreWidth := width / (scoreSymbols + scoreSymbols/2)
		offset := 2
		_, _, style, _ := screen.GetContent(x, y)

		uY := height / 2

		scoreStr := strconv.Itoa(*score)
		scoreStrFull := strings.Repeat("_", scoreSymbols-len(scoreStr)) + scoreStr
		scoreRunesFull := []rune(scoreStrFull)

		for i := 0; i < scoreSymbols; i++ {
			uX := x + offset + i*underscoreWidth + i + 2

			screen.SetContent(uX, uY, scoreRunesFull[i], nil, style)

		}
		return x, y, width, height
	}
}

func App(gridSize int, dc chan game.Direction) *GameUI {
	var displayScore int = 0
	scoreBox := tview.NewBox().
		SetBorder(true).SetBorderAttributes(tcell.AttrBold).
		SetTitle("Score").SetTitleAlign(tview.AlignLeft).
		SetDrawFunc(drawScore(&displayScore))

	gapBox := tview.NewBox() // empty placeholder
	userBox := tview.NewBox().
		SetBorder(true).SetBorderAttributes(tcell.AttrBold).
		SetTitle("Username").SetTitleAlign(tview.AlignLeft)

	gameGrid := NewGameGrid(gridSize)
	grid := tview.NewGrid().
		SetColumns(-1, -1, -1).
		SetRows(3, -20).
		AddItem(scoreBox, 0, 0, 1, 1, 0, 0, false).
		AddItem(gapBox, 0, 1, 1, 1, 0, 0, false).
		AddItem(userBox, 0, 2, 1, 1, 0, 0, false).
		AddItem(gameGrid, 1, 0, 1, 3, 0, 0, false)

	app := tview.NewApplication().SetRoot(grid, true)

	ui := &GameUI{
		app:          app,
		gameGrid:     gameGrid,
		dirChan:      dc,
		displayScore: displayScore,
	}

	app.SetInputCapture(ui.processInput)

	return ui

}

func (ui *GameUI) Update(grid *game.Grid) {
	ui.gameGrid.Update(grid)
}

func (ui *GameUI) UpdateScore(score int) {
	ui.displayScore = score
}

func (ui *GameUI) StartUI() error {
	return ui.app.Run()
}

func (ui *GameUI) Redraw() {
	ui.app.Draw()
}

func (ui *GameUI) processInput(event *tcell.EventKey) *tcell.EventKey {
	switch event.Rune() {
	case 'w':
		ui.dirChan <- game.Up
	case 'd':
		ui.dirChan <- game.Right
	case 's':
		ui.dirChan <- game.Down
	case 'a':
		ui.dirChan <- game.Left
	}
	return event
}
