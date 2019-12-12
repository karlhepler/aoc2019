package asteroid

import (
	"math"
)

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
				m = append(m, Coord{float64(x), float64(y)})
			}
		}
		y++
	}

	return
}

// Map is a map of asteroid coordinates.
type Map []Coord

// Translate translates all asteroid positions on the map to a new
// origin and returns the new translated map.
func (m Map) Translate(origin Coord) Map {
	tmap := make(Map, len(m))

	for i, ast := range m {
		tmap[i] = Coord{
			ast[0] - origin[0],
			ast[1] - origin[1],
		}
	}

	return tmap
}

// AsteroidsAtAngle returns a slice of all asteroids that have the
// given clockwise angle.
func (m Map) AsteroidsAtClockwiseAngle(angle float64) []Coord {
	asts := make([]Coord, 0)

	for _, ast := range m {
		if ast.ClockwiseAngle() == angle {
			asts = append(asts, ast)
		}
	}

	return asts
}

// Without returns a copy of Map without the given coordinate.
func (m Map) Without(filter Coord) Map {
	filtered := make(Map, 0)

	for _, c := range m {
		if c != filter {
			filtered = append(filtered, c)
		}
	}

	return filtered
}

func (m Map) IndexOf(c Coord) int {
	for i := range m {
		if m[i] == c {
			return i
		}
	}

	return -1
}

// byDistance is meant to be used with sort.Sort to sort the coordinates by their
// lengths.
type byDistance []Coord

func (cs byDistance) Len() int {
	return len(cs)
}

func (cs byDistance) Less(i, j int) bool {
	return cs[i].Distance() < cs[j].Distance()
}

func (cs byDistance) Swap(i, j int) {
	cs[i], cs[j] = cs[j], cs[i]
}

// byClockwiseAngle is meant to be used with sort.Sort to sort the coordinates by
// their clockwise angles.
type byClockwiseAngle []Coord

func (cs byClockwiseAngle) Len() int {
	return len(cs)
}

func (cs byClockwiseAngle) Less(i, j int) bool {
	iang, jang := cs[i].ClockwiseAngle(), cs[j].ClockwiseAngle()
	return iang < jang || math.IsNaN(iang) && !math.IsNaN(jang)
}

func (cs byClockwiseAngle) Swap(i, j int) {
	cs[i], cs[j] = cs[j], cs[i]
}
