package fuel_test

import (
	"testing"

	"github.com/karlhepler/aoc2019/1.2/fuel"
)

func TestFuelRequired(t *testing.T) {
	tcs := []struct {
		mass     int
		expected int
	}{{14, 2}, {1969, 966}, {100756, 50346}}

	for _, tc := range tcs {
		if val := fuel.Required(tc.mass); val != tc.expected {
			t.Errorf("Expected %v; Received %v", tc.expected, val)
		}
	}
}
