package asteroid

import (
	"math"
)

type Coord [2]float64

func (c Coord) ClockwiseAngle() float64 {
	angle := c.Angle()

	switch c.Quadrant() {
	case 1:
		return 90 - angle
	case 2:
		return 90 + angle
	case 3:
		return 270 - angle
	default:
		return 270 + angle
	}
}

func (c Coord) Angle() float64 {
	return math.Abs(math.Atan(c[1] / c[0]))
}

func (c Coord) Distance() float64 {
	return math.Sqrt(c[0]*c[0] + c[1]*c[1])
}

func (c Coord) Quadrant() int {
	switch {
	case c[0] > 0 && c[1] > 0:
		return 1
	case c[0] > 0 && c[1] < 0:
		return 2
	case c[0] < 0 && c[1] < 0:
		return 3
	default:
		return 4
	}
}
