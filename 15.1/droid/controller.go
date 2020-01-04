package droid

import (
	"fmt"

	"github.com/karlhepler/aoc2019/intcode"
)

const (
	MoveNorth MovementCommand = iota + 1
	MoveSouth
	MoveWest
	MoveEast
)

// MovementCommand represents the four movement commands understood by Move.
type MovementCommand byte

func (cmd MovementCommand) valid() bool {
	if cmd == MoveNorth || cmd == MoveSouth || cmd == MoveWest || cmd == MoveEast {
		return true
	}
	return false
}

const (
	// The repair droid hit a wall. Its position has not changed.
	StatusHitWall int = iota
	// The repair droid has moved one step in the requested direction.
	StatusMoved
	//The repair droid has moved one step in the requested direction; its new
	//position is the location of the oxygen system.
	StatusFound
)

// Controller is droid's public API.
type Controller struct {
	Droid *intcode.Computer
}

// MoveResponse is the data that Move sends through the returned response channel.
type MoveResponse struct {
	StatusCode int
	Error      error
}

// Move accepts a movement command via an input instruction,
// sends the movement command to the repair droid,
// waits for the repair droid to finish the movement operation,
// and reports on the status of the repair droid via an output instruction.
func (ctrl Controller) Move(cmd MovementCommand) <-chan MoveResponse {
	res := make(chan MoveResponse)
	go func() {
		defer close(res)

		if !cmd.valid() {
			res <- MoveResponse{Error: fmt.Errorf("%v is an invalid movement command", cmd)}
			return
		}

		input := make(chan int)
		output, done := ctrl.Droid.Exec(input)

		input <- int(cmd)
		close(input)

		for {
			select {
			case res <- MoveResponse{StatusCode: <-output}:
			case err := <-done:
				if err != nil {
					res <- MoveResponse{Error: err}
				}
				return
			}
		}
	}()
	return res
}
