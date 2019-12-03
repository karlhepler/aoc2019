package fuel

import "math"

func Required(mass int) int {
	return int(math.Trunc(float64(mass/3))) - 2
}
