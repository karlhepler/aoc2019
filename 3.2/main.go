package main

import (
	"fmt"
	"log"
	"math"
	"time"

	"github.com/karlhepler/aoc2019/3.1/wire"
	"github.com/karlhepler/aoc2019/input"
)

func main() {
	start := time.Now()

	diagram, xings := wire.BuildDiagram()

	var steps float64
	for _, xing := range xings {
		if steps == 0 {
			steps = CountSteps(diagram, xing)
			continue
		}

		steps = math.Min(steps, CountSteps(diagram, xing))
	}

	log.Println(steps)

	fmt.Printf("Time: %v\n", time.Since(start))
}

func CountSteps(diagram *wire.Diagram, stop wire.Vector) (steps float64) {
	for path := range input.Lines("3.1") {
		start := diagram.Origin()

		for move := range wire.MoveAlong(path) {
			if start == stop {
				break
			}

			delta := move.Unit()

			for {
				if move.Empty() || start == stop {
					break
				}

				start = start.Add(delta)
				move = move.Sub(delta)

				steps++
			}
		}
	}

	return steps
}
