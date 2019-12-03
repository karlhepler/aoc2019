package main

import (
	"bufio"
	"log"
	"strconv"
	"strings"

	"github.com/karlhepler/aoc2019/input"
)

type Diagram map[Vector]byte

func (d *Diagram) RunWire(a, b Vector) {
	// implemented straight from WP pseudocode
	dx := b[0] - a[0]
	if dx < 0 {
		dx = -dx
	}
	dy := b[1] - a[1]
	if dy < 0 {
		dy = -dy
	}
	var sx, sy int
	if a[0] < b[0] {
		sx = 1
	} else {
		sx = -1
	}
	if a[1] < b[1] {
		sy = 1
	} else {
		sy = -1
	}
	err := dx - dy

	for {
		(*d)[Vector{a[0], a[1]}] = (*d)[Vector{a[0], a[1]}] + 1
		if a[0] == b[0] && a[1] == b[1] {
			break
		}
		e2 := 2 * err
		if e2 > -dy {
			err -= dy
			a[0] += sx
		}
		if e2 < dx {
			err += dx
			a[1] += sy
		}
	}
}

type Vector [2]int

func (v Vector) Add(a Vector) Vector {
	return Vector{
		v[0] + a[0],
		v[1] + a[1],
	}
}

func (v Vector) IsOrigin() bool {
	return v[0] == 0 && v[1] == 0
}

func main() {
	diagram := &Diagram{}

	// Build the diagrams
	for path := range input.Lines("3.1") {
		orig := Vector{0, 0}

		for dest := range MoveAlong(path) {
			diagram.RunWire(orig, dest)
			orig = orig.Add(dest)
		}
	}

	// Print out where the wires cross
	for pos, val := range *diagram {
		if val > 1 && !pos.IsOrigin() {
			log.Println(pos)
		}
	}
}

// MoveAlong scans the comma-separated input path string, sending vectors for
// each step (next destination) along the path.
func MoveAlong(path string) (destination <-chan Vector) {
	dest := make(chan Vector)

	go func() {
		defer close(dest)

		scanner := bufio.NewScanner(strings.NewReader(path))
		scanner.Split(ScanCSVs)

		for scanner.Scan() {
			switch mov := scanner.Bytes(); mov[0] {
			case 'U':
				dest <- Vector{0, Btoi(mov[1:])}
			case 'R':
				dest <- Vector{Btoi(mov[1:]), 0}
			case 'D':
				dest <- Vector{0, -Btoi(mov[1:])}
			case 'L':
				dest <- Vector{-Btoi(mov[1:]), 0}
			default:
				log.Fatalf("%s is an invalid direction", string(mov[0]))
			}
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}()

	return dest
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
