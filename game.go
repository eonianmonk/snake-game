package main

import (
	"fmt"
	"time"

	"github.com/eonianmonk/snake-game/game"
	"github.com/eonianmonk/snake-game/layout"
)

func Run() {
	gameSize := 10
	rules := &game.GameRules{
		TickInterval:     time.Second * 1,
		GameSize:         gameSize,
		TransparantWalls: false,
	}
	dirChan := make(chan game.Direction)
	signal := make(chan struct{})

	app := layout.App(gameSize, dirChan) //, dirChan)
	game := game.NewGame(rules, dirChan, signal)

	loop(app, game, signal)
}

func loop(app *layout.GameUI, gm *game.Game, signal chan struct{}) {
	errch := make(chan error)

	var err error
	runUi := func(errc chan error) {
		err = app.StartUI()
		errc <- fmt.Errorf("ui error: %s", err.Error())
	}
	runGame := func(errc chan error) {
		err = gm.Start()
		errc <- fmt.Errorf("game error: %s", err.Error())
	}

	go runUi(errch)
	go runGame(errch)
	//signal <- struct{}{} // initial render
	for {
		select {
		case err := <-errch:
			panic(err)
		case <-signal: // game tick
			// update ui
			app.Update(gm.Grid)
			app.UpdateScore(gm.Score())
			app.Redraw()
		}
	}
}
