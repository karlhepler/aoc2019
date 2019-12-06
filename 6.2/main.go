package main

import (
	"log"

	"github.com/karlhepler/aoc2019/6.2/orbit"
	"github.com/karlhepler/aoc2019/input"
)

func main() {
	m := orbit.NewMap()
	orbit.BuildMap(m, input.Lines("6.1"))
	num, err := m.NumOrbitalTransfers("YOU", "SAN")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Number of Orbital Transfers: %d", num)
}
