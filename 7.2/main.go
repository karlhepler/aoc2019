package main

import (
	"fmt"
	"log"
	"math"
	"time"

	quickPerm "github.com/Ramshackle-Jamathon/go-quickPerm"
	"github.com/karlhepler/aoc2019/5.1/input"
	"github.com/karlhepler/aoc2019/7.2/computer"
)

func main() {
	start := time.Now()

	var max float64

	for phaseSettings := range quickPerm.GeneratePermutationsInt([]int{0, 1, 2, 3, 4}) {
		chain := computer.NewAmplifierChain(input.Program("7.1"), phaseSettings)
		output, err := chain.Exec(0)
		if err != nil {
			log.Fatal(err)
		}

		max = math.Max(float64(output), max)
	}

	log.Printf("Highest Output Signal: %v", max)

	fmt.Printf("Time: %v\n", time.Since(start))
}
