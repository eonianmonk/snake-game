package main

import (
	"time"

	"github.com/rivo/tview"
)

func main() {
	app := tgrid(25)
	go changeScore(app)
	if err := app.Run(); err != nil {
		panic(err)
	}
}

func changeScore(app *tview.Application) {
	for {
		time.Sleep(time.Second * 1)
		score += 1
		app.Draw()
	}
}
