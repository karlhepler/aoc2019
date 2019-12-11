package asteroid

type Vector [2]int

func Cross(a, b Vector) int {
	return a[0]*b[1] - a[1]*b[0]
}
