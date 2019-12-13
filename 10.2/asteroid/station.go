package asteroid

import (
	"sort"
)

// MonitoringStation returns the map coordinates of the monitoring station and
// a slice of asteroids that are visible from it.
func MonitoringStation(m Map) (station Coord, visible []Coord) {
	vres := make(chan VisibleResponse, len(m))

	for _, ast := range m {
		go Visible(m, ast, vres)
	}

	i := 0
	maxlen := 0
	for res := range vres {
		if maxlen < len(res.Visible) {
			maxlen = len(res.Visible)
			station = res.Station
			visible = res.Visible
		}

		if i++; i == len(m) {
			close(vres)
		}
	}

	return
}

type VisibleResponse struct {
	Station Coord
	Visible []Coord
}

// Visible returns a slice of coordinates to visible asteroids
func Visible(m Map, station Coord, res chan<- VisibleResponse) {
	visible := make([]Coord, 0)
	omap := m.Translate(station)

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

	res <- VisibleResponse{station, Map(visible).Translate(station.Negative())}
}
