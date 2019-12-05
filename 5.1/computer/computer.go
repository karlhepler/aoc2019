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

func Exec(prgm []int) ([]int, error) {
	for i, num := 0, len(prgm); i < num; i += 4 {
		opcode, modes, err := ParseOpcode(prgm[i])
		if err != nil {
			return prgm, err
		}

		switch opcode {
		case computer.OpcodeHalt:
			return prgm, nil
		case computer.OpcodeAdd:
			params := [3]*int{&prgm[i+1], &prgm[i+2], &prgm[i+3]}
			if err := Add(&prgm, modes, params); err != nil {
				return prgm, err
			}
		case computer.OpcodeMult:
			params := [3]*int{&prgm[i+1], &prgm[i+2], &prgm[i+3]}
			if err := Multiply(&prgm, modes, params); err != nil {
				return prgm, err
			}
		case OpcodeInput:
			//
		case OpcodeOutput:
			//
		default:
			return prgm, fmt.Errorf("%v is an invalid opcode", prgm[i])
		}
	}

	return prgm, nil
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
	vals, err := ParseParams(prgm, modes, params)
	if err != nil {
		return err
	}

	*(vals[2]) = *(vals[0]) + *(vals[1])

	return nil
}

func Multiply(prgm *[]int, modes [3]int, params [3]*int) error {
	vals, err := ParseParams(prgm, modes, params)
	if err != nil {
		return err
	}

	*(vals[2]) = *(vals[0]) * *(vals[1])

	return nil
}

func ParseParams(prgm *[]int, modes [3]int, params [3]*int) ([3]*int, error) {
	var vals [3]*int

	for i, m := range modes {
		switch m {
		case ImmediateMode:
			vals[i] = params[i]
			if i == 2 {
				return vals, fmt.Errorf("Unexpected immediate mode for last param.")
			}
		case PositionMode:
			vals[i] = &(*prgm)[*params[i]]
		default:
			return vals, fmt.Errorf("%v is an invalid parameter mode.", m)
		}
	}

	return vals, nil
}
