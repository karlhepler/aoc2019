package asteroid

import (
	"math"
)

type Coord [2]float64

func (c Coord) ClockwiseAngle() float64 {
	switch {
	case c[0] == 0 && c[1] < 0:
		return 0
	case c[0] > 0 && c[1] == 0:
		return 90
	case c[0] == 0 && c[1] > 0:
		return 180
	case c[0] < 0 && c[1] == 0:
		return 270
	}

	angle := c.Angle()
	switch c.Quadrant() {
	case 1:
		return 90 - angle
	case 2:
		return 90 + angle
	case 3:
		return 270 - angle
	case 4:
		return 270 + angle
	}

	return 360
}

// Angle returns the acute angle of the coordinate in degrees.
func (c Coord) Angle() float64 {
	return math.Abs(math.Atan(c[1]/c[0])) * (180 / math.Pi)
}

func (c Coord) Distance() float64 {
	return math.Sqrt(c[0]*c[0] + c[1]*c[1])
}

func (c Coord) Quadrant() int {
	switch {
	case c[0] > 0 && c[1] < 0:
		return 1
	case c[0] > 0 && c[1] > 0:
		return 2
	case c[0] < 0 && c[1] > 0:
		return 3
	case c[0] < 0 && c[1] < 0:
		return 4
	default:
		return 0
	}
}

func (c Coord) Negative() Coord {
	return Coord{-c[0], -c[1]}
}
