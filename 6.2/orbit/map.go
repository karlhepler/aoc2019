package orbit

import (
	"errors"
	"fmt"

	"github.com/karlhepler/aoc2019/6.1/orbit"
)

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
func (m *Map) NumOrbitalTransfers(pids ...string) (xfers int, err error) {
	var allAncestors [][]*orbit.Planet

	allAncestors, err = m.AllAncestors(pids...)
	if err != nil {
		return
	}

	i := firstNoMatchIndex(allAncestors)
	if i == -1 {
		err = errors.New("No connecting orbits found")
		return
	}

	for _, ancestors := range allAncestors {
		xfers += len(ancestors[i:])
	}

	return
}

// AllAncestors finds the ancestors for each pid and returns slice of slices of planets
func (m *Map) AllAncestors(pids ...string) ([][]*orbit.Planet, error) {
	errchan := make(chan error)
	ancestorschan := make(chan []*orbit.Planet, len(pids))

	for _, pid := range pids {
		go func(pid string) {
			as, err := m.Ancestors(pid)
			if err != nil {
				errchan <- err
				return
			}
			ancestorschan <- as
		}(pid)
	}

	ancestors := make([][]*orbit.Planet, 0)

	for i := 0; i < len(pids); i++ {
		select {
		case err := <-errchan:
			return nil, err
		case as := <-ancestorschan:
			ancestors = append(ancestors, as)
		}
	}

	return ancestors, nil
}

// Ancestors finds all planetary ancestors for the given pid
func (m *Map) Ancestors(pid string) ([]*orbit.Planet, error) {
	planet, ok := (*m.Map)[pid]
	if !ok {
		return nil, fmt.Errorf("Planet ID %s does not exist", pid)
	}

	ancestors := make([]*orbit.Planet, 0)

	planet = planet.Parent
	for planet != nil {
		ancestors = append([]*orbit.Planet{planet}, ancestors...)
		planet = planet.Parent
	}

	return ancestors, nil
}

// Find the first index at which the ancestors do NOT match. Return -1 if no
// match is found.
func firstNoMatchIndex(allAncestors [][]*orbit.Planet) int {
	for i := range allAncestors[0] {
		for _, ancestors := range allAncestors {
			if ancestors[i] != allAncestors[0][i] {
				return i
			}
		}
	}

	return -1
}
