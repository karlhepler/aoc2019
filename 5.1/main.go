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

	var code int
	inputs := make(chan int)
	codes, halt := comp.Exec(inputs)

	go func() {
		defer close(inputs)
		inputs <- 1
		for code = range codes {
			//
		}
	}()

	if err := <-halt; err != nil {
		log.Fatal(err)
	}

	log.Printf("Diagnostic Code: %d", code)

	fmt.Printf("Time: %v\n", time.Since(start))
}
