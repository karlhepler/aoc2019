package asteroid

// Vector is a 2 dimensional vector
type Vector [2]int

// Cross returns the cross product of a and b
func (a Vector) Cross(b Vector) int {
	return a[0]*b[1] - a[1]*b[0]
}

// Dot returns the dot product of a and b
func (a Vector) Dot(b Vector) int {
	return a[0]*b[0] + a[1]*b[1]
}

// Sub subtracts b from a, producing a new Vector
func (a Vector) Sub(b Vector) Vector {
	return Vector{
		a[0] - b[0],
		a[1] - b[1],
	}
}

// Determine if the given vector is on a segment
func (c Vector) OnSegment(a Vector, b Vector) bool {
	ab, ac := b.Sub(a), c.Sub(a)

	return ab.Cross(ac) == 0 && // ab and ac are aligned (aka. colinear)
		ab.Dot(ac) > 0 && // ab.ac is positive
		ab.Dot(ac) < ab.Dot(ab) // ab.ac < ab.ab
}
