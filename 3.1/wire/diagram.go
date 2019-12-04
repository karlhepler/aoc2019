package wire

import (
	"github.com/karlhepler/aoc2019/input"
)

func BuildDiagram() (diagram *Diagram, xings []Vector) {
	var layers byte = 0
	diagram = &Diagram{}
	xings = make([]Vector, 0)

	// Build the diagram
	for path := range input.Lines("3.1") {
		layers++
		start := diagram.Origin()

		for move := range MoveAlong(path) {
			start = diagram.RunWire(start, move, layers)
		}
	}

	// Find crossings
	for xing, layer := range *diagram {
		if layer > layers {
			xings = append(xings, xing)
		}
	}

	return
}

type Diagram map[Vector]byte

func (d Diagram) Origin() Vector {
	return Vector{0, 0}
}

// RunWire runs a wire from start (exclusive) to end (inclusive)
func (d *Diagram) RunWire(start, move Vector, layers byte) (end Vector) {
	delta := move.Unit()

	for {
		if move.Empty() {
			break
		}

		start = start.Add(delta)
		move = move.Sub(delta)

		d.SetPoint(start, layers)
	}

	return start
}

func (d *Diagram) SetPoint(point Vector, layers byte) {
	(*d)[point] = ((*d)[point] * layers) + 1
}
