package game

// cell types
type CellType uint8

const (
	SnakeHead CellType = iota
	SnakeBody
	Food
	BlankCell
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
