package game

import linkedlist "github.com/eonianmonk/snake-game/pkg/linked_list"

type Snake struct {
	snake    *linkedlist.LinkedList[Pos] // TODO: move to lilo queue (maybe)
	headDir  Direction
	bodyDir  Direction
	nextPart *Pos
}

func NewSnake(gameSize int) *Snake {
	return &Snake{
		bodyDir: None,
		headDir: Right,
		snake: linkedlist.NewLinkedList[Pos](Pos{
			row: gameSize / 2,
			col: gameSize / 2,
		}),
	}
}

func (s *Snake) ChangeDirection(dir Direction) {
	if s.snake.Len() == 1 {
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

// function calculates snake's next step. Required for step finalization
// should be called before StepOnce
func (s *Snake) CalculateNextStep() (row int, col int) {
	headPos := s.snake.At(0).Value()
	switch s.headDir {
	case Up:
		row = headPos.row - 1
		col = headPos.col
	case Left:
		row = headPos.row
		col = headPos.col - 1
	case Right:
		row = headPos.row
		col = headPos.col + 1
	case Down:
		row = headPos.row + 1
		col = headPos.col
	default:
		panic("unknown head direction in NextStep calculation")
	}
	s.nextPart = &Pos{row: row, col: col}
	return
}

// function to finalize snake step
// grow - true if next cell is food
// supressCellCheck - true if snake should calculate itselt whether it bites itself
func (s *Snake) StepOnce(grow, supressCellCheck bool) {
	if !supressCellCheck {
		panic("not implemented")
	}
	if s.nextPart == nil {
		panic("next step not calculated")
	}

	// nextPart calculated in CalculateNextStep
	s.snake.AppendHead(*s.nextPart)
	if !grow {
		err := s.snake.DeleteAt(-1)
		if err != nil {
			panic(err) // failed to delete snake's tail
		}
	}

	s.nextPart = nil
}

func (s *Snake) Score() int {
	return s.snake.Len() - 1
}
