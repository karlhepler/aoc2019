package droid

import (
	"fmt"
	"log"

	"github.com/karlhepler/aoc2019/intcode"
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
	MoveNorth MovementCommand = iota + 1
	MoveSouth
	MoveWest
	MoveEast
)

const (
	// The repair droid hit a wall. Its position has not changed.
	StatusHitWall int = iota
	// The repair droid has moved one step in the requested direction.
	StatusMoved
	//The repair droid has moved one step in the requested direction; its new
	//position is the location of the oxygen system.
	StatusMovedStop
)

// Controller is droid's public API.
type Controller struct {
	Droid *intcode.Computer
}

// MoveResponse is the data that Move sends through the returned response channel.
type MoveResponse struct {
	StatusCode
	Error error
}

// Move accepts a movement command via an input instruction,
// sends the movement command to the repair droid,
// waits for the repair droid to finish the movement operation,
// and reports on the status of the repair droid via an output instruction.
func (ctrl Controller) Move(cmd MovementCommand) <-chan MoveResponse {
	res := make(chan MoveResponse)
	go ctrl.move(res, cmd)
	return res
}

func (ctrl Controller) move(res chan<- MoveResponse, cmd MovementCommand) {
	if !cmd.valid() {
		res <- MoveResponse{Error: fmt.Errorf("%v is an invalid movement command", cmd)}
		return
	}

	input := make(chan int)
	output, done := ctrl.Droid.Exec(input)

	// There is no guarantee that the droid is awaiting input.
	// Therefore, this must be in a goroutine.
	go func() {
		// It's possible that the Move method could complete before we
		// send to the input channel. Therefore, defer close(input)
		// must be in here.
		input <- int(cmd)
		close(input)
	}()

	var closed bool
	go func() {
		// Relay/"Render" all droid outputs and close the response
		// when complete.
		for o := range output {
			res <- MoveResponse{StatusCode: StatusCode(o)}
		}
		close(res)
		closed = true
	}()

	// Wait until the droid is done and
	if err := <-done; err != nil {
		r := MoveResponse{Error: err}
		if closed == false {
			res <- r
		} else {
			log.Fatalf("[INTERNAL ERROR] cannot send on a closed channel\nResponse: %#v\n", r)
		}
	}
}
