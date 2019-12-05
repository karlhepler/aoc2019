package computer_test

import (
	"log"
	"testing"

	"github.com/karlhepler/aoc2019/5.1/computer"
)

func TestParseOpcode(t *testing.T) {
	tcs := []struct {
		input     int
		expOpcode int
		expModes  [3]int
		expErr    error
	}{
		{
			input:     1002,
			expOpcode: 2,
			expModes:  [3]int{0, 1, 0},
			expErr:    nil,
		},
	}

	for i, tc := range tcs {
		opcode, modes, err := computer.ParseOpcode(tc.input)
		if opcode != tc.expOpcode || modes != tc.expModes || err != tc.expErr {
			t.Errorf(
				"%d. Expected %#v, %#v, %#v; Received %#v, %#v, %#v",
				i,
				tc.expOpcode, tc.expModes, tc.expErr,
				opcode, modes, err,
			)
		}
	}
}

func TestExec(t *testing.T) {
	input := 42
	tcs := []struct {
		input    []int
		expected int
	}{
		{
			input:    []int{1, 0, 0, 0, 99},
			expected: 0,
		},
		{
			input:    []int{2, 3, 0, 3, 99},
			expected: 0,
		},
		{
			input:    []int{2, 4, 4, 5, 99, 0},
			expected: 0,
		},
		{
			input:    []int{1, 1, 1, 4, 99, 5, 6, 0, 99},
			expected: 0,
		},
		{
			input:    []int{1002, 4, 3, 4, 33},
			expected: 0,
		},
		{
			input:    []int{3, 4, 99, 123, 234},
			expected: 0,
		},
	}

	for i, tc := range tcs {
		output, err := computer.Exec(tc.input, input)
		if err != nil {
			log.Fatal(err)
		}

		if tc.expected != output {
			t.Errorf("%d. Expected %v; Received %v", i, tc.expected, output)
		}
	}
}