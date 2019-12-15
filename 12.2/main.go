package main

import (
	"fmt"
	"log"
	"time"

	"github.com/karlhepler/aoc2019/12.2/moon"
	"github.com/karlhepler/aoc2019/input"
)

func main() {
	start := time.Now()

	moons := moon.Moons{}
	for pos := range input.Lines("input/12.1") {
		moons = append(moons, moon.NewMoon(moon.NewVector(pos)))
	}

	log.Println(moons.Period())

	fmt.Printf("Time: %v\n", time.Since(start))
}
