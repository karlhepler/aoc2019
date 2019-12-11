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
	comp.Load(<-input.Lines("9.1"))

	inputs := make(chan int)
	outputs := comp.Exec(inputs)

	inputs <- 1
	output := <-outputs
	if output.Error != nil {
		log.Fatal(output.Error)
	}

	fmt.Printf("BOOST Keycode: %d\n", output.Value)

	fmt.Printf("Time: %v\n", time.Since(start))
}
