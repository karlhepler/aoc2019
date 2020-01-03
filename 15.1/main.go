package main

import (
	"fmt"
	"log"
	"time"

	"github.com/karlhepler/aoc2019/15.1/droid"
	"github.com/karlhepler/aoc2019/input"
	"github.com/karlhepler/aoc2019/intcode"
)

func main() {
	start := time.Now()

	ctrl := &droid.Controller{Droid: newDroid()}

	directions := [4]droid.MovementCommand{
		droid.MoveNorth,
		droid.MoveEast,
		droid.MoveSouth,
		droid.MoveWest,
	}
	dir := 0

	// Next, I need to create a map of visited locations

search:
	for {
		switch code := move(ctrl, directions[dir]); code {
		case droid.StatusHitWall:
			dir = (dir + 1) % 4
		case droid.StatusMoved:
		case droid.StatusFound:
			break search
		}
	}

	fmt.Printf("\nTime: %v\n", time.Since(start))
}

func newDroid() *intcode.Computer {
	comp := intcode.NewComputer()
	if err := comp.Load(<-input.Lines("../input/15.1")); err != nil {
		log.Fatal(err)
	}
	return comp
}

// move moves the droid controller in the given direction, renders all
// move response status codes, and outputs the last move response
// status code.
func move(ctrl *droid.Controller, cmd droid.MovementCommand) (code int) {
	var res droid.MoveResponse
	for res = range ctrl.Move(cmd) {
		if res.Error != nil {
			log.Fatalf("[ERROR] %s\n", res.Error)
		}
		render(res.StatusCode)
	}
	return res.StatusCode
}

func render(code int) {
	switch code {
	case droid.StatusHitWall:
		fmt.Print("#")
	case droid.StatusMoved:
		fmt.Print(".")
	case droid.StatusFound:
		fmt.Print("X")
	}
}
