package computer_test

import (
	"log"
	"testing"

	"github.com/karlhepler/aoc2019/2.1/computer"
)

func TestExec(t *testing.T) {
	tcs := []struct {
		input    []int
		expected []int
	}{
		{
			input:    []int{1, 0, 0, 0, 99},
			expected: []int{2, 0, 0, 0, 99},
		},
		{
			input:    []int{2, 3, 0, 3, 99},
			expected: []int{2, 3, 0, 6, 99},
		},
		{
			input:    []int{2, 4, 4, 5, 99, 0},
			expected: []int{2, 4, 4, 5, 99, 9801},
		},
		{
			input:    []int{1, 1, 1, 4, 99, 5, 6, 0, 99},
			expected: []int{30, 1, 1, 4, 2, 5, 6, 0, 99},
		},
	}

	for _, tc := range tcs {
		output, err := computer.Exec(tc.input)
		if err != nil {
			log.Fatal(err)
		}

		if !equal(tc.expected, output) {
			t.Errorf("Expected %v; Received %v", tc.expected, output)
		}
	}
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}
