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
	if err := comp.Load(<-input.Lines("9.1")); err != nil {
		log.Fatal(err)
	}

	inputs := make(chan int)
	outputs, halt := comp.Exec(inputs)

	inputs <- 1

	select {
	case output := <-outputs:
		fmt.Printf("BOOST Keycode: %d\n", output)
	case err := <-halt:
		log.Fatal(err)
	}

	fmt.Printf("Time: %v\n", time.Since(start))
}
