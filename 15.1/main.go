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

	// TODO(kjh): How to move it around? (https://adventofcode.com/2019/day/15)
	// I need a droid.Droid!
	for i := 0; ; i++ {
		code := move(ctrl, i%4+1)
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
func move(ctrl *droid.Controller, direction int) (code int) {
	var res droid.MoveResponse
	for res = range ctrl.Move(direction) {
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
	case droid.StatusMovedStop:
		fmt.Print("X")
	}
}
