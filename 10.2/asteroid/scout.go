package asteroid

import (
	"sort"
)

// MonitoringStation returns the map coordinates of the monitoring station and
// a slice of asteroids that are visible from it.
func MonitoringStation(m Map) (station Coord, visible []Coord) {
	res := make(chan data)

	for _, ast := range m {
		go Visible(ast, m.Translate(ast), res)
	}

	i := 0
	maxlen := 0
	for d := range res {
		if maxlen < len(d.visible) {
			maxlen = len(d.visible)
			station = d.station
			visible = d.visible
		}

		i++
		if i == len(m) {
			close(res)
		}
	}

	return
}

type data struct {
	station Coord
	visible []Coord
}

// Visible returns a slice of coordinates to visible asteroids
func Visible(station Coord, omap Map, res chan<- data) {
	visible := make([]Coord, 0)

	// Remove the origin
	omap = omap.Without(Coord{0, 0})

	// Sort the map by clockwise angle
	sort.Sort(byClockwiseAngle(omap))

	// Loop through the map by clockwise angle
	for _, ast := range omap {
		// Get all of the asteroids at the current clockwise angle, sort them
		// by distance, and choose the first (closest) one.
		matches := omap.AsteroidsAtClockwiseAngle(ast.ClockwiseAngle())
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

	res <- data{station, visible}
}
