package main

import (
	"fmt"
	"log"
	"time"

	"github.com/karlhepler/aoc2019/input"
	"github.com/karlhepler/aoc2019/intcode"
)

func main() {
	start := time.Now()

	comp := intcode.NewComputer()
	if err := comp.Load(<-input.Lines("5.1")); err != nil {
		log.Fatal(err)
	}

	inputs := make(chan int)
	codes, halt := comp.Exec(inputs)

	inputs <- 5

	select {
	case code := <-codes:
		log.Printf("Diagnostic Code: %d", code)
	case err := <-halt:
		log.Fatal(err)
	}

	fmt.Printf("Time: %v\n", time.Since(start))
}
