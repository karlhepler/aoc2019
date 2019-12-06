package main

import (
	"fmt"
	"log"
	"time"

	"github.com/karlhepler/aoc2019/6.1/orbit"
	"github.com/karlhepler/aoc2019/input"
)

func main() {
	start := time.Now()

	m := orbit.NewMap()
	orbit.BuildMap(m, input.Lines("6.1"))
	log.Printf("Number of Orbits: %d", m.NumOrbits())

	fmt.Printf("Time: %v\n", time.Since(start))
}
