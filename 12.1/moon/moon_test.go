package moon_test

import (
	"testing"

	"github.com/karlhepler/aoc2019/12.1/moon"
)

func TestSimulate(t *testing.T) {
	tcs := []struct {
		moons  moon.Moons
		steps  int
		energy int
	}{
		{
			moon.Moons{
				moon.NewMoon(moon.NewVector("<x=-1, y=0, z=2>")),
				moon.NewMoon(moon.NewVector("<x=2, y=-10, z=-7>")),
				moon.NewMoon(moon.NewVector("<x=4, y=-8, z=8>")),
				moon.NewMoon(moon.NewVector("<x=3, y=5, z=-1>")),
			}, 10, 179,
		},
		{
			moon.Moons{
				moon.NewMoon(moon.NewVector("<x=-8, y=-10, z=0>")),
				moon.NewMoon(moon.NewVector("<x=5, y=5, z=10>")),
				moon.NewMoon(moon.NewVector("<x=2, y=-7, z=3>")),
				moon.NewMoon(moon.NewVector("<x=9, y=-8, z=-3>")),
			}, 100, 1940,
		},
	}

	for i, tc := range tcs {
		tc.moons.Simulate(tc.steps)
		if energy := tc.moons.Energy(); energy != tc.energy {
			t.Errorf("%d. Expected %d; Received %d", i, tc.energy, energy)
		}
	}
}
