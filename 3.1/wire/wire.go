package wire

import (
	"bufio"
	"log"
	"strconv"
	"strings"
)

// MoveAlong scans the comma-separated input path string, sending vectors for
// each step (next destination) along the path.
func MoveAlong(path string) (destination <-chan Vector) {
	dest := make(chan Vector)

	go func() {
		defer close(dest)

		scanner := bufio.NewScanner(strings.NewReader(path))
		scanner.Split(scanCSVs)

		for scanner.Scan() {
			switch mov := scanner.Bytes(); mov[0] {
			case 'U':
				dest <- Vector{0, btof(mov[1:])}
			case 'R':
				dest <- Vector{btof(mov[1:]), 0}
			case 'D':
				dest <- Vector{0, -btof(mov[1:])}
			case 'L':
				dest <- Vector{-btof(mov[1:]), 0}
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

func btof(b []byte) float64 {
	i, err := strconv.Atoi(string(b))
	if err != nil {
		log.Fatal(err)
	}
	return float64(i)
}

func scanCSVs(data []byte, atEOF bool) (advance int, token []byte, err error) {
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
