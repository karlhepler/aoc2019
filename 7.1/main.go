package main

import (
	"fmt"
	"log"
	"time"

	quickPerm "github.com/Ramshackle-Jamathon/go-quickPerm"
	"github.com/karlhepler/aoc2019/input"
	"github.com/karlhepler/aoc2019/intcode"
)

func main() {
	start := time.Now()

	var max int
	prgm := <-input.Lines("7.1")

	for phaseSettings := range quickPerm.GeneratePermutationsInt([]int{0, 1, 2, 3, 4}) {
		go func(phaseSettings []int) {
			circuit := intcode.NewAmplificationCircuit(prgm, phaseSettings...)

			output, err := circuit.Exec(0)
			if err != nil {
				log.Fatal(err)
			}

			max = maxint(output, max)
		}(phaseSettings)
	}

	log.Printf("Highest Output Signal: %d", max)

	fmt.Printf("Time: %v\n", time.Since(start))
}

func maxint(a, b int) int {
	if a < b {
		return b
	}
	return a
}
