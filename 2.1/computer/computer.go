package computer

import "fmt"

const (
	OpcodeHalt = 99
	OpcodeAdd  = 1
	OpcodeMult = 2
)

func Exec(prgm []int) ([]int, error) {
	for i, num := 0, len(prgm); i < num; i += 4 {
		switch prgm[i] {
		case OpcodeHalt:
			return prgm, nil
		case OpcodeAdd:
			prgm[prgm[i+3]] = prgm[prgm[i+1]] + prgm[prgm[i+2]]
		case OpcodeMult:
			prgm[prgm[i+3]] = prgm[prgm[i+1]] * prgm[prgm[i+2]]
		default:
			return prgm, fmt.Errorf("%v is an invalid opcode", prgm[i])
		}
	}

	return prgm, nil
}
