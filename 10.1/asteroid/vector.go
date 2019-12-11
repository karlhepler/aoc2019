package asteroid

type Vector [2]int

// Cross calculates the cross product of two Vectors
func (a Vector) Cross(b Vector) int {
	return a[0]*b[1] - a[1]*b[0]
}

func (a Vector) Dot(b Vector) int {
	return a[0]*b[0] + a[1]*b[1]
}

func (a Vector) Sub(b Vector) Vector {
	return Vector{
		b[0] - a[0],
		b[1] - a[1],
	}
}

func Colinear(a, b Vector) bool {
	return a.Cross(b) == 0
}
