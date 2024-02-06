package game

import "fmt"

type GameError struct {
	cause string
}

func (ge *GameError) Error() string {
	return fmt.Sprintf("you lost: %s", ge.cause)
}
