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
	rob.PaintedPanels[robot.Coord{0, 0}] = robot.White
	_, err := rob.Activate()
	if err != nil {
		log.Fatal(err)
	}

	if err := rob.Render(log.Writer()); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\nTime: %v\n", time.Since(start))
}
