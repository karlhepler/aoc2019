package computer

import "fmt"

const (
	OpcodeHalt = 99
	OpcodeAdd  = 1
	OpcodeMult = 2
)

type ErrorCode int

const (
	ErrorHalt ErrorCode = iota + 1
	ErrorOpcode
)

type Error struct {
	Code    ErrorCode
	Message string
}

func (err Error) Error() string {
	return err.Message
}

func Exec(prgm []int) ([]int, error) {
	var err error

	for i, num := 0, len(prgm); i < num; i += 4 {
		prgm, err = ExecLine(i, prgm)

		if err == nil {
			continue
		}

		switch err.(Error).Code {
		case ErrorHalt:
			return prgm, nil
		case ErrorOpcode:
			return prgm, err
		}
	}

	return prgm, nil
}

func ExecLine(i int, prgm []int) ([]int, error) {
	switch prgm[i] {
	case OpcodeHalt:
		return prgm, Error{Code: ErrorHalt}
	case OpcodeAdd:
		prgm[prgm[i+3]] = prgm[prgm[i+1]] + prgm[prgm[i+2]]
	case OpcodeMult:
		prgm[prgm[i+3]] = prgm[prgm[i+1]] * prgm[prgm[i+2]]
	default:
		msg := fmt.Sprintf("%v is an invalid opcode", prgm[i])
		return prgm, Error{Code: ErrorOpcode, Message: msg}
	}

	return prgm, nil
}
