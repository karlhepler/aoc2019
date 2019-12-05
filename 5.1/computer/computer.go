package computer

import "github.com/karlhepler/aoc2019/2.1/computer"

const (
	OpcodeInput = 3
	OpcodeOuput = 4
)

func Exec(prgm []int) ([]int, error) {
	var err error

	for i, num := 0, len(prgm); i < num; {
		prgm, err = computer.ExecLine(i, prgm)

		if err == nil {
			i += 4
			continue
		}

		switch err.(computer.Error).Code {
		case computer.ErrorHalt:
			return prgm, nil
		case computer.ErrorOpcode:
			switch prgm[i] {
			case OpcodeInput:
				// Opcode 3 takes a single integer as input and saves it to the
				// position given by its only parameter. For example, the
				// instruction 3,50 would take an input value and store it at
				// address 50.
				i += 2
			case OpcodeOuput:
				// Opcode 4 outputs the value of its only parameter. For
				// example, the instruction 4,50 would output the value at
				// address 50.
				i += 2
			default:
				return prgm, err
			}
		}
	}

	return prgm, nil
}
