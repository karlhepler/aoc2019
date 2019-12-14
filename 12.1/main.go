package main

import (
	"fmt"
	"time"

	"github.com/karlhepler/aoc2019/12.1/moon"
	"github.com/karlhepler/aoc2019/input"
)

func main() {
	start := time.Now()

	moons := moon.Moons{}
	for pos := range input.Lines("input/12.1") {
		moons = append(moons, moon.NewMoon(moon.NewVector(pos)))
	}

	moons.Simulate(1000)
	energy := moons.Energy()

	fmt.Printf("Total System Energy: %d\n", energy)

	fmt.Printf("Time: %v\n", time.Since(start))
}
