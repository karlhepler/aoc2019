package main

import (
	"log"
	"strconv"

	"github.com/karlhepler/aoc2019/1.2/fuel"
	"github.com/karlhepler/aoc2019/input"
)

func main() {
	total := 0

	for line := range input.Lines("1.1") {
		val, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}

		total += fuel.Required(val)
	}

	log.Printf("Total: %v", total)
}
