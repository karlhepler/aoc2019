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

	// The droid starts at (0,0)
	// Dijkstra's and A* algoriths should be used to solve the problem
	// https://www.geeksforgeeks.org/a-search-algorithm/

	fmt.Printf("\nTime: %v\n", time.Since(start))
}

func newDroid() *intcode.Computer {
	comp := intcode.NewComputer()
	if err := comp.Load(<-input.Lines("input/15.1")); err != nil {
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
