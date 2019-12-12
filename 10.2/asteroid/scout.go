package asteroid

import (
	"sort"
)

// MonitoringStation returns the map coordinates of the monitoring station.
func MonitoringStation(m Map) (station Coord, visible []Coord) {
	var maxlen int

	for _, ast := range m {
		viz := Visible(m.Translate(ast))
		if maxlen < len(viz) {
			maxlen = len(viz)
			station = ast
			visible = viz
		}
	}

	return
}

// Visible returns a slice of coordinates to visible asteroids
func Visible(m Map) []Coord {
	visible := make([]Coord, 0)

	// Sort the map by clockwise angle
	sort.Sort(byClockwiseAngle(m))

	// Loop through the map by clockwise angle
	for _, ast := range m {
		// Get all of the asteroids at the current clockwise angle, sort them
		// by distance, and choose the first (closest) one.
		matches := m.AsteroidsAtClockwiseAngle(ast.ClockwiseAngle())
		if len(matches) == 0 {
			visible = append(visible, ast)
			continue
		}

		sort.Sort(byDistance(matches))
		visible = append(visible, matches[0])
	}

	return visible
}
