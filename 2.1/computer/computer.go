package computer

import (
	"fmt"
	"log"
	"testing"
)

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

func CreateTestExec(exec func([]int) ([]int, error)) func(*testing.T) {
	return func(t *testing.T) {
		tcs := []struct {
			input    []int
			expected []int
		}{
			{
				input:    []int{1, 0, 0, 0, 99},
				expected: []int{2, 0, 0, 0, 99},
			},
			{
				input:    []int{2, 3, 0, 3, 99},
				expected: []int{2, 3, 0, 6, 99},
			},
			{
				input:    []int{2, 4, 4, 5, 99, 0},
				expected: []int{2, 4, 4, 5, 99, 9801},
			},
			{
				input:    []int{1, 1, 1, 4, 99, 5, 6, 0, 99},
				expected: []int{30, 1, 1, 4, 2, 5, 6, 0, 99},
			},
		}

		for _, tc := range tcs {
			output, err := exec(tc.input)
			if err != nil {
				log.Fatal(err)
			}

			if !equal(tc.expected, output) {
				t.Errorf("Expected %v; Received %v", tc.expected, output)
			}
		}
	}
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}
