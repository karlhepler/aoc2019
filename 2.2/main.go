package main

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/karlhepler/aoc2019/2.1/computer"
	"github.com/karlhepler/aoc2019/2.1/input"
)

func main() {
	start := time.Now()

	ans, err := answer()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Answer: %v\n", ans)

	fmt.Printf("Time: %v\n", time.Since(start))
}

func answer() (int, error) {
	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			prgm, err := computer.Exec(input.Program(noun, verb))
			if err != nil {
				log.Fatal(err)
			}
			if prgm[0] == 19690720 {
				return 100*noun + verb, nil
			}
		}
	}

	return 0, errors.New("Cannot find answer")
}
