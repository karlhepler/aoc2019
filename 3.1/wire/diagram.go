package wire

type Diagram map[Vector]byte

func (d Diagram) AtOrigin(p Vector) bool {
	o := d.Origin()
	return p[0] == o[0] && p[1] == o[1]
}

func (d Diagram) Origin() Vector {
	return Vector{0, 0}
}

// RunWire runs a wire from start (exclusive) to end (inclusive)
func (d *Diagram) RunWire(start, move Vector, i byte) (end Vector) {
	delta := move.Unit()

	for {
		if move.Empty() {
			break
		}

		start = start.Add(delta)
		move = move.Sub(delta)

		d.SetPoint(start, i)
	}

	return start
}

func (d *Diagram) SetPoint(point Vector, b byte) {
	(*d)[point] = ((*d)[point] * b) + 1
}
