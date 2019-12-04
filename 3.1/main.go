package main

import (
	"log"
	"math"

	"github.com/karlhepler/aoc2019/3.1/wire"
	"github.com/karlhepler/aoc2019/input"
)

func main() {
	diagram := &wire.Diagram{}
	origin := diagram.Origin()

	// Build the diagrams
	for path := range input.Lines("3.1") {
		start := origin

		for end := range wire.MoveAlong(path) {
			log.Printf("RunWire %v->%v\n", start, end)
			diagram.RunWire(start, end)
			start = start.Add(end)
		}
	}

	// Find which intersection is closest
	var closest float64
	for intersection, layers := range *diagram {
		if layers > 1 && !diagram.AtOrigin(intersection) {
			if closest == 0 {
				closest = intersection.Manhattan(origin)
				continue
			}

			closest = math.Min(closest, intersection.Manhattan(origin))
		}
	}

	log.Println(closest)
}
