package computer_test

import (
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
