package main

import (
	"fmt"
	"log"
	"time"

	"github.com/karlhepler/aoc2019/5.1/input"
	"github.com/karlhepler/aoc2019/5.2/computer"
)

func main() {
	start := time.Now()

	code, err := computer.Exec(input.Program(), 5)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Diagnostic Code: %d", code)

	fmt.Printf("Time: %v\n", time.Since(start))
}
