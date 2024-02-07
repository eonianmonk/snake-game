// snake game logic package
package game

// cell types
type CellType uint8

const (
	BlankCell CellType = iota
	SnakeBody
	SnakeHead
	Food
)

// direction
type Direction uint8

const (
	Up Direction = iota
	Down
	Left
	Right
	None // no snake in cell
)
