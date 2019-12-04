package pass_test

import (
	"testing"

	"github.com/karlhepler/aoc2019/4.2/pass"
)

func TestValidatorValid(t *testing.T) {
	tcs := []struct {
		bs       []byte
		expected bool
	}{
		{[]byte("112233"), true},
		{[]byte("123444"), false},
		{[]byte("111122"), true},
		{[]byte("111111"), false},
	}

	v := pass.Validator{
		Length:      6,
		LowerBound:  0,
		UpperBound:  123445,
		NumAdjacent: 2,
		OrderBy:     pass.ASC,
	}

	for i, tc := range tcs {
		result := v.Valid(tc.bs)
		if result != tc.expected {
			t.Errorf("%d. Expected %v; Received %v", i, tc.expected, result)
		}
	}
}
