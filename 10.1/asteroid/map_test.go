package asteroid_test

import (
	"reflect"
	"testing"

	"github.com/karlhepler/aoc2019/10.1/asteroid"
)

func TestBuildMap(t *testing.T) {
	tcs := []struct {
		lines []string
		m     asteroid.Map
	}{
		{
			lines: []string{".#..#", ".....", "#####", "....#", "...##"},
			m: asteroid.Map{
				asteroid.Asteroid{asteroid.Vector{1, 0}},
				asteroid.Asteroid{asteroid.Vector{4, 0}},
				asteroid.Asteroid{asteroid.Vector{0, 2}},
				asteroid.Asteroid{asteroid.Vector{1, 2}},
				asteroid.Asteroid{asteroid.Vector{2, 2}},
				asteroid.Asteroid{asteroid.Vector{3, 2}},
				asteroid.Asteroid{asteroid.Vector{4, 2}},
				asteroid.Asteroid{asteroid.Vector{4, 3}},
				asteroid.Asteroid{asteroid.Vector{3, 4}},
				asteroid.Asteroid{asteroid.Vector{4, 4}},
			},
		},
	}

	for i, tc := range tcs {
		lines := make(chan string)

		go func() {
			for _, line := range tc.lines {
				lines <- line
			}
			close(lines)
		}()

		m := asteroid.BuildMap(lines)

		if !reflect.DeepEqual(m, tc.m) {
			t.Errorf("%d. Expected %v; Received %v", i, tc.m, m)
		}
	}
}
