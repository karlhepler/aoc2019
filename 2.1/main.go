package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/karlhepler/aoc2019/input"
)

const (
	OpcodeHalt = 99
	OpcodeAdd  = 1
	OpcodeMult = 2
)

func main() {
	ics, err := exec(restore(intcodes()))
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("What value is left at position 0: %v", ics[0])
}

func exec(ics []int) ([]int, error) {
	for i, num := 0, len(ics); i < num; i += 4 {
		switch ics[i] {
		case OpcodeHalt:
			return ics, nil
		case OpcodeAdd:
			ics[ics[i+3]] = ics[ics[i+1]] + ics[ics[i+2]]
		case OpcodeMult:
			ics[ics[i+3]] = ics[ics[i+1]] * ics[ics[i+2]]
		default:
			return ics, fmt.Errorf("%v is an invalid opcode", ics[i])
		}
	}

	return ics, nil
}

func intcodes() []int {
	strs := strings.Split(<-input.Lines("2.1"), ",")
	ints := make([]int, len(strs))

	for i, s := range strs {
		var err error
		ints[i], err = strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}
	}

	return ints
}

func restore(ics []int) []int {
	ics[1] = 12
	ics[2] = 2
	return ics
}
