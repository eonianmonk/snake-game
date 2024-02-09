package layout

import "github.com/gdamore/tcell/v2"

type DrawFunc func(screen tcell.Screen, x, y, width, height int) (int, int, int, int)

const (
	scoreSymbols = 6

	HeadColor  = tcell.ColorRed
	BodyColor  = tcell.ColorBlue
	FoodColor  = tcell.ColorGreen
	BlankColor = tcell.ColorGray
)
