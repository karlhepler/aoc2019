package main

import (
	"fmt"
	"log"
	"math"
	"time"

	"github.com/karlhepler/aoc2019/3.1/wire"
)

func main() {
	start := time.Now()

	diagram, xings := wire.BuildDiagram()
	origin := diagram.Origin()

	// Find which crossing is closest
	var closest float64
	for _, xing := range xings {
		if closest == 0 {
			closest = xing.Manhattan(origin)
			continue
		}

		closest = math.Min(closest, xing.Manhattan(origin))
	}

	log.Println(closest)

	fmt.Printf("Time: %v\n", time.Since(start))
}
