package asteroid

func BuildMap(lines <-chan string) Map {
	m := make(Map)

	y := 0
	for line := range lines {
		for x, b := range line {
			if b == '#' {
				m[Vector{x, y}] = Asteroid{}
			}
		}
		y++
	}

	return m
}

type Map map[Vector]Asteroid

type Asteroid struct {
	//
}
