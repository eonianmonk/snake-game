package game

import (
	"math/rand"
	"sync"
	"time"

	linkedlist "github.com/eonianmonk/snake-game/pkg/linked_list"
)

type Pos struct {
	row, col int
}

type Grid struct {
	sync.RWMutex
	cells   [][]CellType
	size    int
	foodPos *Pos
	rand    *rand.Rand
}

func NewGrid(size int) *Grid {
	gridCells := make([][]CellType, size)
	for i := 0; i < size; i++ {
		gridCells[i] = make([]CellType, size)
	}
	g := &Grid{
		cells: gridCells,
		size:  size,
		rand:  rand.New(rand.NewSource(time.Now().Unix())),
	}
	g.genFood()
	return g
}

func (g *Grid) CheckCell(row, col int) CellType {
	g.RLock()
	defer g.RUnlock()
	return g.cells[row][col]
}

// updates grid according to new snake state
func (g *Grid) Update(snake *Snake, newFood bool) {
	g.Lock()
	defer g.Unlock()
	// draw new grid
	for i := 0; i < g.size; i++ {
		g.cells[i] = make([]CellType, g.size)
	}
	// draw snake on grid

	iter := linkedlist.NewLLIter(snake.snake)

	v := iter.Next()
	g.cells[v.row][v.col] = SnakeHead

	for !iter.Done() {
		v = iter.Next()
		g.cells[v.row][v.col] = SnakeBody
	}

	if newFood {
		g.genFood()
	} else {
		g.cells[g.foodPos.row][g.foodPos.col] = Food
	}
}

func (g *Grid) genFood() {
	// generate new food & put it on grid
	for {
		col := g.rand.Intn(g.size)
		row := g.rand.Intn(g.size)
		cell := g.cells[row][col]

		if cell == BlankCell {
			g.foodPos = &Pos{col: col, row: row}
			g.cells[row][col] = Food
			break
		}
	}
}
