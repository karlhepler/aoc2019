package orbit

import "github.com/karlhepler/aoc2019/6.1/orbit"

type Map struct {
	*orbit.Map
}

func NewMap() *Map {
	return &Map{orbit.NewMap()}
}

func BuildMap(m orbit.Planeter, lines <-chan string) {
	orbit.BuildMap(m, lines)
}
