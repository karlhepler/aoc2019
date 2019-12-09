package main

import (
	"fmt"
	"log"
	"time"

	"github.com/karlhepler/aoc2019/5.1/computer"
	"github.com/karlhepler/aoc2019/5.1/input"
)

func main() {
	start := time.Now()

	code, err := computer.Exec(input.Program("5.1"), 1)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Diagnostic Code: %d", code)

	fmt.Printf("Time: %v\n", time.Since(start))
}
