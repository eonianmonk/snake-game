package game

import "sync"

type Grid struct {
	sync.RWMutex
	cells [][]CellType
}

func NewGrid(size int) *Grid {
	gridCells := make([][]CellType, size)
	for i := 0; i < size; i++ {
		gridCells[i] = make([]CellType, size)
	}
	return &Grid{cells: gridCells}
}

func (g *Grid) CheckCell(row, col int) CellType {
	return g.cells[row][col]
}
