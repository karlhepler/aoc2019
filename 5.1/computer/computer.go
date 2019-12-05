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
		opcode, _, err := ParseOpcode(prgm[i])
		if err != nil {
			return prgm, err
		}

		switch opcode {
		case computer.OpcodeHalt:
			return prgm, nil
		case computer.OpcodeAdd:
			Add(&prgm, prgm[i+1], prgm[i+2], prgm[i+3])
		case computer.OpcodeMult:
			Multiply(&prgm, prgm[i+1], prgm[i+2], prgm[i+3])
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

func Add(prgm *[]int, p1, p2, p3 int) {
	(*prgm)[p3] = (*prgm)[p1] + (*prgm)[p2]
}

func Multiply(prgm *[]int, p1, p2, p3 int) {
	(*prgm)[p3] = (*prgm)[p1] * (*prgm)[p2]
}
