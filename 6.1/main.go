package main

import (
	"log"

	"github.com/karlhepler/aoc2019/6.1/orbit"
	"github.com/karlhepler/aoc2019/input"
)

func main() {
	m := orbit.BuildMap(input.Lines("6.1"))
	log.Printf("Number of Orbits: %d", m.NumOrbits())
}
