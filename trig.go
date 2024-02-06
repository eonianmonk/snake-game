package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func tgrid() *tview.Application {
	scoreBox := tview.NewBox().
		SetBorder(true).SetBorderAttributes(tcell.AttrBold).
		SetTitle("Score").SetTitleAlign(tview.AlignLeft)

	gapBox := tview.NewBox().SetBorder(false) // to save color
	userBox := tview.NewBox().
		SetBorder(true).SetBorderAttributes(tcell.AttrBold).
		SetTitle("Username").SetTitleAlign(tview.AlignLeft)

	gameBox := tview.NewBox().
		SetBorder(true).SetBorderAttributes(tcell.AttrBold).
		SetTitle("SNAKE").SetTitleAlign(tview.AlignCenter)

	grid := tview.NewGrid().
		SetColumns(-1, -1, -1).
		SetRows(3, -1).
		AddItem(scoreBox, 0, 0, 1, 1, 2, 1, false).
		AddItem(gapBox, 0, 1, 1, 1, 2, 1, false).
		AddItem(userBox, 0, 2, 1, 1, 2, 1, false).
		AddItem(gameBox, 1, 0, 1, 3, 25, 25, false)
		//AddItem()

	return tview.NewApplication().SetRoot(grid, true).EnableMouse(true)
}
