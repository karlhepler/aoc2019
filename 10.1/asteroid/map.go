package asteroid

// BuildMap builds a Map from a given channel of lines where each character in
// the line (sitting on a y coordinate) represents an x coordinate on the map.
// The character '.' means the point is NOT an asteroid, while the character
// '#' means there IS an asteroid. Since we only care about the positions of
// asteroids, we ignore '.' and only add to Map if an asteroid is found.
func BuildMap(lines <-chan string) (m Map) {
	y := 0

	for line := range lines {
		for x, b := range line {
			if b == '#' {
				m = append(m, Asteroid{Vector{x, y}})
			}
		}
		y++
	}

	return
}

// Map is a map of asteroids, where the key is the 2D coordinate vector and the
// value is the asteroid itself.
type Map []Asteroid

// Copy returns a copy of the map
func (a Map) Copy() Map {
	b := make(Map, len(a))
	for k, v := range a {
		b[k] = v
	}
	return b
}

// Asteroid represents an asteroid. So far, it has no properties.
type Asteroid struct {
	Pos Vector
}
