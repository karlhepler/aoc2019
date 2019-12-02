package main

import (
	"log"
	"math"
	"strconv"

	"github.com/karlhepler/aoc2019/input"
)

func main() {
	total := 0

	for line := range input.Lines("1.1") {
		val, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}
		total += fuelRequired(val)
	}

	log.Printf("Total: %v", total)
}

func fuelRequired(mass int) int {
	return int(math.Trunc(float64(mass/3))) - 2
}