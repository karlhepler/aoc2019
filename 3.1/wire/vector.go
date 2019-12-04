package wire

import "math"

type Vector [2]float64

func (v Vector) Add(a Vector) Vector {
	return Vector{
		v[0] + a[0],
		v[1] + a[1],
	}
}

func (v Vector) Sub(a Vector) Vector {
	return Vector{
		v[0] - a[0],
		v[1] - a[1],
	}
}

func (v Vector) Empty() bool {
	return v == Vector{}
}

func (v Vector) Manhattan(a Vector) float64 {
	return math.Abs(a[0]-v[0]) + math.Abs(a[1]-v[1])
}

func (v Vector) Unit() Vector {
	mag := v.Magnitude()
	if mag == 0 {
		return Vector{}
	}

	return Vector{
		v[0] / mag,
		v[1] / mag,
	}
}

func (v Vector) Magnitude() float64 {
	return math.Sqrt(v[0]*v[0] + v[1]*v[1])
}
