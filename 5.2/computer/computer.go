package computer

import (
	"fmt"

	comp2 "github.com/karlhepler/aoc2019/2.1/computer"
	comp5 "github.com/karlhepler/aoc2019/5.1/computer"
)

type ComparisonOperator byte

const (
	Equals ComparisonOperator = iota + 1
	LessThan
)

const (
	// OpcodeJumpTrue: if the first parameter is non-zero, it sets the
	// instruction pointer to the value from the second parameter. Otherwise,
	// it does nothing.
	OpcodeJumpTrue = 5

	// OpcodeJumpFalse: if the first parameter is zero, it sets the instruction
	// pointer to the value from the second parameter. Otherwise, it does
	// nothing.
	OpcodeJumpFalse = 6

	// OpcodeLessThan: if the first parameter is less than the second
	// parameter, it stores 1 in the position given by the third parameter.
	// Otherwise, it stores 0.
	OpcodeLessThan = 7

	// OpcodeEquals: if the first parameter is equal to the second parameter,
	// it stores 1 in the position given by the third parameter. Otherwise, it
	// stores 0.
	OpcodeEquals = 8
)

func Exec(prgm []int, input int) (int, error) {
	for i, num := 0, len(prgm); i < num; {
		opcode, modes, err := comp5.ParseOpcode(prgm[i])
		if err != nil {
			return -1, err
		}

		switch opcode {
		case comp2.OpcodeHalt:
			return 0, nil

		case comp2.OpcodeAdd:
			params := [3]*int{&prgm[i+1], &prgm[i+2], &prgm[i+3]}
			if err := comp5.Add(&prgm, modes, params); err != nil {
				return -1, err
			}
			i += 4

		case comp2.OpcodeMult:
			params := [3]*int{&prgm[i+1], &prgm[i+2], &prgm[i+3]}
			if err := comp5.Multiply(&prgm, modes, params); err != nil {
				return -1, err
			}
			i += 4

		case comp5.OpcodeInput:
			prgm[prgm[i+1]] = input
			i += 2

		case comp5.OpcodeOutput:
			output, err := comp5.Output(&prgm, modes, &prgm[i+1])
			if err != nil {
				return -1, err
			}

			if output != 0 {
				return output, nil
			}
			i += 2

		case OpcodeJumpTrue:
			params := [2]*int{&prgm[i+1], &prgm[i+2]}
			if err := JumpIf(true, &i, &prgm, modes, params); err != nil {
				return -1, err
			}

		case OpcodeJumpFalse:
			params := [2]*int{&prgm[i+1], &prgm[i+2]}
			if err := JumpIf(false, &i, &prgm, modes, params); err != nil {
				return -1, err
			}

		case OpcodeLessThan:
			params := [3]*int{&prgm[i+1], &prgm[i+2], &prgm[i+3]}
			if err := Compare(LessThan, &prgm, modes, params); err != nil {
				return -1, err
			}
			i += 4

		case OpcodeEquals:
			params := [3]*int{&prgm[i+1], &prgm[i+2], &prgm[i+3]}
			if err := Compare(Equals, &prgm, modes, params); err != nil {
				return -1, err
			}
			i += 4

		default:
			return -1, fmt.Errorf("%v is an invalid opcode", prgm[i])
		}
	}

	return -1, nil
}

func JumpIf(cond bool, i *int, prgm *[]int, modes [3]int, params [2]*int) error {
	vals, err := comp5.ParseParams(prgm, modes, params[:])
	if err != nil {
		return err
	}

	switch {
	case cond == true && *(vals[0]) != 0:
		*i = *(vals[1])
	case cond == false && *(vals[0]) == 0:
		*i = *(vals[1])
	default:
		*i += 3
	}

	return nil
}

func Compare(op ComparisonOperator, prgm *[]int, modes [3]int, params [3]*int) error {
	vals, err := comp5.ParseParams(prgm, modes, params[:])
	if err != nil {
		return err
	}

	var val = 0

	switch {
	case op == Equals:
		if *(vals[0]) == *(vals[1]) {
			val = 1
		}
	case op == LessThan:
		if *(vals[0]) < *(vals[1]) {
			val = 1
		}
	default:
		return fmt.Errorf("%v is an invalid comparison operator", op)
	}

	*(vals[2]) = val

	return nil
}
