package asteroid

import (
	"errors"
	"math"
)

type Vector [2]float64

func (v Vector) ClockwiseAngle() (float64, error) {
	angle, err := v.Angle()
	if err != nil {
		return 0, err
	}

	switch v.Quadrant() {
	case 1:
		return 90 - angle, nil
	case 2:
		return 90 + angle, nil
	case 3:
		return 270 - angle, nil
	default:
		return 270 + angle, nil
	}
}

func (v Vector) Angle() (float64, error) {
	if v[0] == 0 {
		return 0, errors.New("division by zero")
	}

	return math.Abs(math.Atan(v[1] / v[0])), nil
}

func (v Vector) Length() float64 {
	return math.Sqrt(v[0]*v[0] + v[1]*v[1])
}

func (v Vector) Quadrant() int {
	switch {
	case v[0] > 0 && v[1] > 0:
		return 1
	case v[0] > 0 && v[1] < 0:
		return 2
	case v[0] < 0 && v[1] < 0:
		return 3
	default:
		return 4
	}
}
