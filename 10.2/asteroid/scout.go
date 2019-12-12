package asteroid

import (
	"sort"
)

// MonitoringStation returns the map coordinates of the monitoring station.
func MonitoringStation(m Map) (station Coord, visible []Coord) {
	var maxlen int

	for _, ast := range m {
		tm := m.Translate(ast)
		vis := Visible(tm)
		if maxlen < len(vis) {
			maxlen = len(vis)
			station = ast
			visible = vis
		}
	}

	return
}

// Visible returns a slice of coordinates to visible asteroids
func Visible(m Map) []Coord {
	visible := make([]Coord, 0)

	// Remove the origin
	m = m.Without(Coord{0, 0})

	// Sort the map by clockwise angle
	sort.Sort(byClockwiseAngle(m))

	// Loop through the map by clockwise angle
	for _, ast := range m {
		// Get all of the asteroids at the current clockwise angle, sort them
		// by distance, and choose the first (closest) one.
		matches := m.AsteroidsAtClockwiseAngle(ast.ClockwiseAngle())
		sort.Sort(byDistance(matches))

		// Append if it doesn't already exist
		exists := false
		for _, v := range visible {
			if v == matches[0] {
				exists = true
				break
			}
		}
		if !exists {
			visible = append(visible, matches[0])
		}
	}

	return visible
}
