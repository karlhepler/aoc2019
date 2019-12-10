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

	for perms := range quickPerm.GeneratePermutationsInt([]int{5, 6, 7, 8, 9}) {
		phaseSettings := [5]int{}
		copy(phaseSettings[:], perms)

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