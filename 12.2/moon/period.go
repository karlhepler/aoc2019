package moon

import (
	"fmt"

	"github.com/karlhepler/aoc2019/12.1/moon"
)

func NewVector(def string) (vec moon.Vector) {
	return moon.NewVector(def)
}

func NewMoon(pos moon.Vector) moon.Moon {
	return moon.NewMoon(pos)
}

type Moons moon.Moons

// Period finds the period of each dimension (the number of steps it
// takes for each dimension to go back to its starting point) and
// returns the least common multiple of them.
func (moons Moons) Period() int {
	init := make(Moons, len(moons))
	copy(init, moons)
	periods := moon.Vector{}

	for step := 1; ; step++ {
		moon.Moons(moons).ApplyGravity()
		moon.Moons(moons).Move()

		for dim := 0; dim < 3; dim++ {
			match := true
			for m := 0; m < len(moons); m++ {
				posmatch := moons[m].Position[dim] == init[m].Position[dim]
				velmatch := moons[m].Velocity[dim] == init[m].Velocity[dim]
				if !posmatch || !velmatch {
					match = false
					break
				}
			}

			if match == true {
				periods[dim] = step
			}
		}

		if periods[0] != 0 && periods[1] != 0 && periods[2] != 0 {
			break
		}
	}

	fmt.Printf("%v\n", periods)

	return LCM(periods[0], periods[1], periods[2])
}
