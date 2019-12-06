package orbit

import "github.com/karlhepler/aoc2019/6.1/orbit"

// NewMap returns a pointer to a new instance of Map
func NewMap() *Map {
	return &Map{orbit.NewMap()}
}

// BuildMap simply wraps orbit.BuildMap defined in 6.1
func BuildMap(m orbit.Planeter, lines <-chan string) {
	orbit.BuildMap(m, lines)
}

// Map decorates orbit.Map defined in 6.1
type Map struct {
	*orbit.Map
}

// NumOrbitalTransfers returns the minimum number of orbital transfers required
// to move from pid a to pid b
func (m *Map) NumOrbitalTransfers(a, b string) (num int) {
	return
}
