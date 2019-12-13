package main

import (
	"fmt"
	"log"
	"time"

	"github.com/karlhepler/aoc2019/11.1/robot"
	"github.com/karlhepler/aoc2019/input"
)

func main() {
	start := time.Now()

	rob := robot.New()
	rob.Computer.Load(<-input.Lines("input/11.1"))
	numPaintedPanels, err := rob.Activate()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Number of Painted Panels: %d\n", numPaintedPanels)

	fmt.Printf("Time: %v\n", time.Since(start))
}
