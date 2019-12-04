package main

import (
	"log"
	"math"

	"github.com/karlhepler/aoc2019/3.1/wire"
)

func main() {
	diagram, layers := wire.BuildDiagram()
	origin := diagram.Origin()

	// Find which intersection is closest
	var closest float64
	for intersection, layer := range *diagram {
		if layer > layers {
			if closest == 0 {
				closest = intersection.Manhattan(origin)
				continue
			}

			closest = math.Min(closest, intersection.Manhattan(origin))
		}
	}

	log.Println(closest)
}
