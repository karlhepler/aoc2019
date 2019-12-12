package main

import (
	"log"

	"github.com/karlhepler/aoc2019/10.2/asteroid"
	"github.com/karlhepler/aoc2019/input"
)

func main() {
	m := asteroid.BuildMap(input.Lines("10.1"))
	station, visible := asteroid.MonitoringStation(m)

	log.Println(station, len(visible))
}
