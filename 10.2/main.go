package main

import (
	"fmt"
	"log"
	"time"

	"github.com/karlhepler/aoc2019/10.2/asteroid"
	"github.com/karlhepler/aoc2019/input"
)

func main() {
	start := time.Now()

	m := asteroid.BuildMap(input.Lines("10.1"))
	station, visible := asteroid.MonitoringStation(m)

	log.Println(station, len(visible))

	fmt.Printf("Time: %v\n", time.Since(start))
}
