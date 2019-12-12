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
				m = append(m, Vector{float64(x), float64(y)})
			}
		}
		y++
	}

	return
}

// Map is a map of asteroid coordinates.
type Map []Vector

// Translate translates all asteroid positions on the map to a new
// origin and returns the new translated map.
func (m Map) Translate(origin Vector) Map {
	tmap := make(Map, len(m))

	for i, ast := range m {
		tmap[i] = Vector{
			ast[0] - origin[0],
			ast[1] - origin[1],
		}
	}

	return tmap
}

// AsteroidsAtAngle returns a slice of all asteroids that have the
// given clockwise angle.
func (m Map) AsteroidsAtClockwiseAngle(angle float64) ([]Vector, error) {
	asts := make([]Vector, 0)

	for _, ast := range m {
		astang, err := ast.ClockwiseAngle()
		if err != nil {
			return asts, err
		}

		if astang == angle {
			asts = append(asts, ast)
		}
	}

	return asts, nil
}

// byLength is meant to be used with sort.Sort to sort the vectors by their
// lengths.
type byLength []Vector

func (vs byLength) Len() int {
	return len(vs)
}

func (vs byLength) Swap(i, j int) {
	vs[i], vs[j] = vs[j], vs[i]
}

func (vs byLength) Less(i, j int) bool {
	return vs[i].Length() < vs[j].Length()
}

// byClockwiseAngle is meant to be used with sort.Sort to sort the vectors by
// their clockwise angles.
type byClockwiseAngle []Vector

func (vs byClockwiseAngle) Len() int {
	return len(vs)
}

func (vs byClockwiseAngle) Swap(i, j int) {
	vs[i], vs[j] = vs[j], vs[i]
}

func (vs byClockwiseAngle) Less(i, j int) bool {
	ai, err := vs[i].ClockwiseAngle()
	if err != nil {
		panic(err)
	}

	aj, err := vs[j].ClockwiseAngle()
	if err != nil {
		panic(err)
	}

	return ai < aj
}
