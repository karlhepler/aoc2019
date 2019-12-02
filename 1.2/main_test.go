package main

import (
	"testing"
)

func TestTotalFuelRequired(t *testing.T) {
	tcs := []struct {
		mass     int
		expected int
	}{{14, 2}, {1969, 966}, {100756, 50346}}

	for _, tc := range tcs {
		if val := totalFuelRequired(tc.mass); val != tc.expected {
			t.Errorf("Expected %v; Received %v", tc.expected, val)
		}
	}
}
