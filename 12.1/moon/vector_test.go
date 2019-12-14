package moon_test

import (
	"testing"

	"github.com/karlhepler/aoc2019/12.1/moon"
)

func TestNewVector(t *testing.T) {
	tcs := []struct {
		def string
		vec moon.Vector
	}{
		{"<x=-1, y=0, z=2>", moon.Vector{-1, 0, 2}},
		{"<x=2, y=-10, z=-7>", moon.Vector{2, -10, -7}},
		{"<x=4, y=-8, z=8>", moon.Vector{4, -8, 8}},
		{"<x=3, y=5, z=-1>", moon.Vector{3, 5, -1}},
		{"<x=-8, y=-10, z=0>", moon.Vector{-8, -10, 0}},
		{"<x=5, y=5, z=10>", moon.Vector{5, 5, 10}},
		{"<x=2, y=-7, z=3>", moon.Vector{2, -7, 3}},
		{"<x=9, y=-8, z=-3>", moon.Vector{9, -8, -3}},
	}

	for i, tc := range tcs {
		if vec := moon.NewVector(tc.def); vec != tc.vec {
			t.Errorf("%d. Expected %v; Received %v", i, tc.vec, vec)
		}
	}
}
