package fuel

import "github.com/karlhepler/aoc2019/1.1/fuel"

func Required(mass int) int {
	total := 0

	mass = fuel.Required(mass)
	for mass > 0 {
		total += mass
		mass = fuel.Required(mass)
	}

	return total
}
