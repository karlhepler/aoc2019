package main

import (
	"log"

	"github.com/karlhepler/aoc2019/5.1/computer"
	"github.com/karlhepler/aoc2019/5.1/input"
)

func main() {
	code, err := computer.Exec(input.Program("5.1"), 1)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Diagnostic Code: %d", code)
}
