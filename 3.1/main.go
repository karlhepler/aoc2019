package main

import (
	"bufio"
	"log"
	"strconv"
	"strings"

	"github.com/karlhepler/aoc2019/input"
)

type Vector [2]int

func main() {
	var diagrams []map[Vector]bool

	for path := range input.Lines("3.1") {
		origin := Vector{0, 0}
		diagram := make(map[Vector]bool)

		for distance := range MoveAlong(path) {
			// don't draw the origin
		}

		diagrams = append(diagrams, diagram)
	}
}

func MoveAlong(path string) <-chan Vector {
	v := make(chan Vector)

	go func() {
		defer close(v)

		scanner := bufio.NewScanner(strings.NewReader(path))
		scanner.Split(ScanCSVs)

		for scanner.Scan() {
			switch mov := scanner.Bytes(); mov[0] {
			case 'U':
				v <- Vector{0, Btoi(mov[1:])}
			case 'R':
				v <- Vector{Btoi(mov[1:]), 0}
			case 'D':
				v <- Vector{0, -Btoi(mov[1:])}
			case 'L':
				v <- Vector{-Btoi(mov[1:]), 0}
			default:
				log.Fatalf("%s is an invalid direction", string(mov[0]))
			}
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}()

	return v
}

func Btoi(b []byte) int {
	i, err := strconv.Atoi(string(b))
	if err != nil {
		log.Fatal(err)
	}
	return i
}

func ScanCSVs(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) > 0 {
		return len(data), data, nil
	}

	for i, d := range data {
		if d == ',' {
			return i + 1, data[:i], nil
		}
	}

	return 0, nil, nil
}
