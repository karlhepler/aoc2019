package main

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/karlhepler/aoc2019/1.1/fuel"
	"github.com/karlhepler/aoc2019/input"
)

func main() {
	start := time.Now()

	total := 0

	for line := range input.Lines("1.1") {
		val, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}
		total += fuel.Required(val)
	}

	log.Printf("Total: %v", total)

	fmt.Printf("Time: %v\n", time.Since(start))
}
