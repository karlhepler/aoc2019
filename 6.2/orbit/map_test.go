package orbit_test

import (
	"testing"

	"github.com/karlhepler/aoc2019/6.2/orbit"
)

func TestNumOrbitalTransfers(t *testing.T) {
	tcs := []struct {
		lines    []string
		expected int
	}{
		{
			lines:    []string{"COM)B", "B)C", "C)D", "D)E", "E)F", "B)G", "G)H", "D)I", "E)J", "J)K", "K)L", "K)YOU", "I)SAN"},
			expected: 4,
		},
	}

	for i, tc := range tcs {
		lineschan := make(chan string)

		go func(lines []string) {
			for _, line := range lines {
				lineschan <- line
			}
			close(lineschan)
		}(tc.lines)

		m := orbit.NewMap()
		orbit.BuildMap(m, lineschan)

		xfers, err := m.NumOrbitalTransfers("YOU", "SAN")
		if err != nil {
			t.Fatal(err)
		}
		if xfers != tc.expected {
			t.Errorf("%d. Expected: %d; Received: %d", i, tc.expected, xfers)
		}
	}
}
