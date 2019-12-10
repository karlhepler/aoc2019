package computer_test

import (
	"testing"

	"github.com/karlhepler/aoc2019/7.2/computer"
)

func TestExec(t *testing.T) {
	tcs := []struct {
		prgm     []int
		input    int
		expected int
	}{
		// Position mode
		{prgm: []int{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8}, input: 8, expected: 1},
		{prgm: []int{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8}, input: 7, expected: 0},
		{prgm: []int{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8}, input: 7, expected: 1},
		{prgm: []int{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8}, input: 8, expected: 0},
		// Immediate mode
		{prgm: []int{3, 3, 1108, -1, 8, 3, 4, 3, 99}, input: 8, expected: 1},
		{prgm: []int{3, 3, 1108, -1, 8, 3, 4, 3, 99}, input: 7, expected: 0},
		{prgm: []int{3, 3, 1107, -1, 8, 3, 4, 3, 99}, input: 7, expected: 1},
		{prgm: []int{3, 3, 1107, -1, 8, 3, 4, 3, 99}, input: 8, expected: 0},
		// Jump test position mode
		{prgm: []int{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9}, input: 0, expected: 0},
		{prgm: []int{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9}, input: 1, expected: 1},
		// Jump test immediate mode
		{prgm: []int{3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1}, input: 0, expected: 0},
		{prgm: []int{3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1}, input: 1, expected: 1},
		// Large example
		{prgm: []int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31, 1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104, 999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99}, input: 7, expected: 999},
		{prgm: []int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31, 1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104, 999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99}, input: 8, expected: 1000},
		{prgm: []int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31, 1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104, 999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99}, input: 9, expected: 1001},
	}

	inputchan := make(chan int)
	outputchan := make(chan computer.Output)

	for i, tc := range tcs {
		comp := computer.NewComputer(tc.prgm)
		go comp.Exec(inputchan, outputchan)

		inputchan <- tc.input
		output := <-outputchan

		if output.Error != nil {
			t.Fatal(output.Error)
		}

		if output.Value != tc.expected {
			t.Fatalf("%d. Expected %d; Received %d", i, tc.expected, output.Value)
		}
	}

	close(inputchan)
	close(outputchan)
}
