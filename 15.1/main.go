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

	ctrl := droid.Controller{Droid: newDroid()}

	fmt.Println("-- START --")

	for res := range ctrl.Move(1) {
		if res.Error != nil {
			log.Fatalf("[ERROR] %s\n", res.Error)
		}

		switch res.StatusCode {
		case droid.StatusHitWall:
			fmt.Print("#")
		case droid.StatusMoved:
			fmt.Print("D")
		case droid.StatusMovedStop:
			fmt.Print("X")
		}

		fmt.Println("\n-- COMPLETE --")
	}

	fmt.Printf("Time: %v\n", time.Since(start))
}

func newDroid() *intcode.Computer {
	d := intcode.NewComputer()
	if err := d.Load(<-input.Lines("../input/15.1")); err != nil {
		log.Fatal(err)
	}
	return d
}
