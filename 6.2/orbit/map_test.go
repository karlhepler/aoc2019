package orbit_test

import (
	"testing"

	"github.com/karlhepler/aoc2019/6.2/orbit"
)

func TestCountOrbits(t *testing.T) {
	tcs := []struct {
		lines    []string
		expected int
	}{
		{
			lines:    []string{"COM)B", "B)C", "C)D", "D)E", "E)F", "B)G", "G)H", "D)I", "E)J", "J)K", "K)L"},
			expected: 42,
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

		if numOrbits := m.NumOrbits(); numOrbits != tc.expected {
			t.Errorf("%d. Expected: %d; Received: %d", i, tc.expected, numOrbits)
		}
	}
}
