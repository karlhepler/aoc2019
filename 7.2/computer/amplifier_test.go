package computer_test

import (
	"testing"

	"github.com/karlhepler/aoc2019/7.2/computer"
)

func TestAmplifierChain(t *testing.T) {
	tcs := []struct {
		input         int
		prgm          []int
		phaseSettings []int
		expected      int
	}{
		{0, []int{3, 15, 3, 16, 1002, 16, 10, 16, 1, 16, 15, 15, 4, 15, 99, 0, 0}, []int{4, 3, 2, 1, 0}, 43210},
		{0, []int{3, 23, 3, 24, 1002, 24, 10, 24, 1002, 23, -1, 23, 101, 5, 23, 23, 1, 24, 23, 23, 4, 23, 99, 0, 0}, []int{0, 1, 2, 3, 4}, 54321},
		{0, []int{3, 31, 3, 32, 1002, 32, 10, 32, 1001, 31, -2, 31, 1007, 31, 0, 33, 1002, 33, 7, 33, 1, 33, 31, 31, 1, 32, 31, 31, 4, 31, 99, 0, 0, 0}, []int{1, 0, 4, 3, 2}, 65210},
	}

	for i, tc := range tcs {
		chain := computer.NewAmplifierChain(tc.prgm, tc.phaseSettings)

		output, err := chain.Exec(tc.input)
		if err != nil {
			t.Fatalf("%d. %s", i, err)
		}

		if output != tc.expected {
			t.Errorf("%d. Expected %d; Received %d", i, tc.expected, output)
		}
	}
}
