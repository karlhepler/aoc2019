package main

import (
	"log"

	"github.com/karlhepler/aoc2019/2.1/computer"
	"github.com/karlhepler/aoc2019/2.1/input"
)

func main() {
	prgm, err := computer.Exec(input.Program(12, 2))
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("What value is left at position 0: %v", prgm[0])
}
