package game

import linkedlist "github.com/eonianmonk/snake-game/pkg/linked_list"

type SnakePart struct {
	row, col int
	facing   Direction
}

type Snake struct {
	length  int
	snake   *linkedlist.LinkedList[SnakePart]
	headDir Direction
	bodyDir Direction
}

func (s *Snake) ChangeDirection(dir Direction) {
	if s.length == 1 {
		s.headDir = dir
		return
	}
	// there is body
	if dir == s.bodyDir {
		// can't turn where snake's body at
		return
	}
	s.headDir = dir
}

func (s *Snake) NextStep()

// we use info from the grid
// grow - true if next cell is food
// supressCellCheck - true if snake should calculate itselt whether it bites itself
func (s *Snake) StepOnce(grow, supressCellCheck bool) error {
	if supressCellCheck {
		panic("not implemented")
	}
	node
}
