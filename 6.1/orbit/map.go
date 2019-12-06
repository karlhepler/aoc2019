package orbit

import (
	"strings"
)

// BuildMap builds the map from given lines of input
func BuildMap(m Planeter, lines <-chan string) {
	for line := range lines {
		planetIDs := strings.Split(line, ")")

		p1 := m.Planet(planetIDs[0])
		p2 := m.Planet(planetIDs[1])

		p2.Orbits(p1)
	}
}

// Planeter defines a Planet method that receives a pid and returns a Planet pointer
type Planeter interface {
	Planet(pid interface{}) *Planet
}

// NewMap returns a pointer to a new Map
func NewMap() *Map {
	return &Map{}
}

// Map represents the map of planets and their orbits
type Map map[interface{}]*Planet

// Planet searches the Map for a planet with matching planet ID. If it cannot
// find it, it will create one. Either way, it returns a pointer to it.
func (m *Map) Planet(pid interface{}) *Planet {
	p, ok := (*m)[pid]
	if !ok {
		(*m)[pid] = &Planet{ID: pid, Moons: make([]*Planet, 0)}
		p = (*m)[pid]
	}
	return p
}

// NumOrbits counts and returns the number of orbits in the Map
func (m Map) NumOrbits() (num int) {
	for _, planet := range m {
		num += planet.NumOrbits()
	}

	return
}

// Planet represents a planet in the Map
type Planet struct {
	ID     interface{}
	Moons  []*Planet
	Parent *Planet
}

// Orbits updates the Map to show that planet a orbits planet b
func (a *Planet) Orbits(b *Planet) {
	a.Parent = b
	b.Moons = append(b.Moons, a)
}

// NumOrbits counts the total number of planetary orbits in the Map
func (p Planet) NumOrbits() (num int) {
	num = len(p.Moons)

	for _, moon := range p.Moons {
		num += moon.NumOrbits()
	}

	return
}
