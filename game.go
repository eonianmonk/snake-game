package main

import (
	"fmt"
	"time"

	"github.com/eonianmonk/snake-game/game"
	"github.com/eonianmonk/snake-game/layout"
)

func Run() {
	gameSize := 20
	rules := &game.GameRules{
		TickInterval:     time.Millisecond * 750,
		GameSize:         gameSize,
		TransparantWalls: false,
	}
	dirChan := make(chan game.Direction)
	signal := make(chan struct{})

	app := layout.App(gameSize, dirChan)
	game := game.NewGame(rules, dirChan, signal)

	loop(app, game, signal)
}

func loop(app *layout.GameUI, gm *game.Game, signal chan struct{}) {
	errch := make(chan error)

	runUi := func(errc chan error) {
		err := app.StartUI()
		errc <- fmt.Errorf("ui error: ", err.Error())
	}
	runGame := func(errc chan error) {
		err := gm.Start()
		errc <- fmt.Errorf("game error: ", err.Error())
	}

	go runUi(errch)
	go runGame(errch)

	for {
		select {
		case err := <-errch:
			panic(err)
		case <-signal: // game tick
			// update ui
			app.Update(gm.Grid)
			app.UpdateScore(gm.Score())
		}
	}
}
