package computer

import (
	"fmt"
	"strconv"

	"github.com/karlhepler/aoc2019/2.1/computer"
)

const (
	// OpcodeInput takes a single integer as input and saves it to the position
	// given by its only parameter. For example, the instruction 3,50 would
	// take an input value and store it at address 50.
	OpcodeInput = 3

	// OpcodeOutput outputs the value of its only parameter. For example, the
	// instruction 4,50 would output the value at address 50.
	OpcodeOutput = 4

	// PositionMode causes the parameter to be interpreted as a position -
	// if the parameter is 50, its value is the value stored at address 50 in
	// memory. Until now, all parameters have been in position mode.
	PositionMode = 0

	// ImmediateMode causes a parameter to be interpreted as a value - if
	// the parameter is 50, its value is simply 50.
	ImmediateMode = 1
)

func Exec(prgm []int, inputs ...int) (int, error) {
	inputIndex := 0

	for i, num := 0, len(prgm); i < num; {
		opcode, modes, err := ParseOpcode(prgm[i])
		if err != nil {
			return -1, err
		}

		switch opcode {
		case computer.OpcodeHalt:
			return 0, nil

		case computer.OpcodeAdd:
			params := [3]*int{&prgm[i+1], &prgm[i+2], &prgm[i+3]}
			if err := Add(&prgm, modes, params); err != nil {
				return -1, err
			}
			i += 4

		case computer.OpcodeMult:
			params := [3]*int{&prgm[i+1], &prgm[i+2], &prgm[i+3]}
			if err := Multiply(&prgm, modes, params); err != nil {
				return -1, err
			}
			i += 4

		case OpcodeInput:
			prgm[prgm[i+1]] = inputs[inputIndex]
			inputIndex++
			i += 2

		case OpcodeOutput:
			output, err := Output(&prgm, modes, &prgm[i+1])
			if err != nil {
				return -1, err
			}

			if output != 0 {
				return output, nil
			}
			i += 2

		default:
			return -1, fmt.Errorf("%v is an invalid opcode", prgm[i])
		}
	}

	return -1, nil
}

func ParseOpcode(oc int) (opcode int, modes [3]int, err error) {
	ocstr := fmt.Sprintf("%05d", oc)

	opcode, err = strconv.Atoi(ocstr[3:])
	if err != nil {
		return
	}

	modes[0], err = strconv.Atoi(string(ocstr[2]))
	if err != nil {
		return
	}
	modes[1], err = strconv.Atoi(string(ocstr[1]))
	if err != nil {
		return
	}
	modes[2], err = strconv.Atoi(string(ocstr[0]))

	return
}

func Add(prgm *[]int, modes [3]int, params [3]*int) error {
	vals, err := ParseParams(prgm, modes, params[:])
	if err != nil {
		return err
	}

	*(vals[2]) = *(vals[0]) + *(vals[1])

	return nil
}

func Multiply(prgm *[]int, modes [3]int, params [3]*int) error {
	vals, err := ParseParams(prgm, modes, params[:])
	if err != nil {
		return err
	}

	*(vals[2]) = *(vals[0]) * *(vals[1])

	return nil
}

func Output(prgm *[]int, modes [3]int, param *int) (int, error) {
	params, err := ParseParams(prgm, modes, []*int{param})
	return *params[0], err
}

func ParseParams(prgm *[]int, modes [3]int, params []*int) ([]*int, error) {
	vals := make([]*int, len(params))

	if len(params) > 3 {
		return vals, fmt.Errorf("len(params) must be no greater than 3")
	}

	for i := range params {
		switch modes[i] {
		case ImmediateMode:
			vals[i] = params[i]
		case PositionMode:
			vals[i] = &(*prgm)[*params[i]]
		default:
			return vals, fmt.Errorf("%v is an invalid parameter mode.", modes[i])
		}
	}

	return vals, nil
}
