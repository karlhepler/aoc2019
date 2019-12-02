package main

import (
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
	ics := intcodes()

	for i, num := 0, len(ics); i < num; i += 4 {
		switch ics[i] {
		case OpcodeHalt:
			log.Println("HALT")
			log.Printf("Output: %v", ics)
			return
		case OpcodeAdd:
			log.Println("ADD")
			ics[i+3] = ics[i+1] + ics[i+2]
		case OpcodeMult:
			log.Println("MULTIPLY")
			ics[i+3] = ics[i+1] * ics[i+2]
		default:
			log.Fatalf("%v is an invalid opcode", ics[i])
		}
	}
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
