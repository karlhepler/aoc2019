package intcode

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
)

const (
	// OpcodeHalt means that the program is finished and should immediately halt.
	OpcodeHalt = 99

	// OpcodeAdd adds together numbers read from two positions and stores the
	// result in a third position. The three integers immediately after the
	// opcode tell you these three positions - the first two indicate the
	// positions from which you should read the input values, and the third
	// indicates the position at which the output should be stored.
	OpcodeAdd = 1

	// OpcodeMuliply works exactly like opcode 1, except it multiplies the two
	// inputs instead of adding them. Again, the three integers after the
	// opcode indicate where the inputs and outputs are, not their values.
	OpcodeMultiply = 2

	// OpcodeInput takes a single integer as input and saves it to the position
	// given by its only parameter. For example, the instruction 3,50 would
	// take an input value and store it at address 50.
	OpcodeInput = 3

	// OpcodeOutput outputs the value of its only parameter. For example, the
	// instruction 4,50 would output the value at address 50.
	OpcodeOutput = 4

	// OpcodeJumpTrue: if the first parameter is non-zero, it sets the
	// instruction pointer to the value from the second parameter. Otherwise,
	// it does nothing.
	OpcodeJumpIfTrue = 5

	// OpcodeJumpFalse: if the first parameter is zero, it sets the instruction
	// pointer to the value from the second parameter. Otherwise, it does
	// nothing.
	OpcodeJumpIfFalse = 6

	// OpcodeLessThan: if the first parameter is less than the second
	// parameter, it stores 1 in the position given by the third parameter.
	// Otherwise, it stores 0.
	OpcodeLessThan = 7

	// OpcodeEquals: if the first parameter is equal to the second parameter,
	// it stores 1 in the position given by the third parameter. Otherwise, it
	// stores 0.
	OpcodeEquals = 8
)

const (
	// PositionMode causes the parameter to be interpreted as a position -
	// if the parameter is 50, its value is the value stored at address 50 in
	// memory. Until now, all parameters have been in position mode.
	PositionMode = 0

	// ImmediateMode causes a parameter to be interpreted as a value - if
	// the parameter is 50, its value is simply 50.
	ImmediateMode = 1
)

// NewComputer returns a pointer to a new Computer instance
func NewComputer() *Computer {
	return &Computer{}
}

// Computer is the intcode Computer struct
type Computer struct {
	Memory []int
}

// Load loads a program into memory
func (comp *Computer) Load(prgm string) {
	instructions := strings.Split(prgm, ",")
	comp.Memory = make([]int, len(instructions))

	for i, instruction := range instructions {
		comp.Memory[i] = func(s string) int {
			i, err := strconv.Atoi(strings.Trim(s, " "))
			if err != nil {
				log.Fatal(err)
			}
			return i
		}(instruction)
	}
}

// Output is what the computer outputs
type Output struct {
	Value int
	Error error
}

// Exec executes the program loaded into memory
func (comp *Computer) Exec(inputs <-chan int) <-chan Output {
	outputs := make(chan Output)
	go comp.exec(inputs, outputs)
	return outputs
}

func (comp *Computer) exec(inputs <-chan int, outputs chan<- Output) {
	defer close(outputs)

	if comp.Memory == nil {
		outputs <- Output{Error: errors.New("No program loaded in memory")}
		return
	}

	addr := 0

	for {
		opcode, modes, err := decode(comp.Memory[addr])
		if err != nil {
			outputs <- Output{Error: err}
			return
		}

		switch opcode {
		case OpcodeHalt:
			return

		case OpcodeInput:
			vals := comp.values(comp.move(&addr, 1), modes)
			*vals[0] = <-inputs

		case OpcodeOutput:
			vals := comp.values(comp.move(&addr, 1), modes)
			outputs <- Output{Value: *vals[0]}

		case OpcodeAdd:
			vals := comp.values(comp.move(&addr, 3), modes)
			*vals[2] = *vals[1] + *vals[0]

		case OpcodeMultiply:
			vals := comp.values(comp.move(&addr, 3), modes)
			*vals[2] = *vals[1] * *vals[0]

		case OpcodeJumpIfTrue:
			vals := comp.values(comp.move(&addr, 2), modes)
			if *vals[0] != 0 {
				addr = *vals[1]
			}

		case OpcodeJumpIfFalse:
			vals := comp.values(comp.move(&addr, 2), modes)
			if *vals[0] == 0 {
				addr = *vals[1]
			}

		case OpcodeLessThan:
			vals := comp.values(comp.move(&addr, 3), modes)
			if *vals[0] < *vals[1] {
				*vals[2] = 1
			} else {
				*vals[2] = 0
			}

		case OpcodeEquals:
			vals := comp.values(comp.move(&addr, 3), modes)
			if *vals[0] == *vals[1] {
				*vals[2] = 1
			} else {
				*vals[2] = 0
			}

		default:
			outputs <- Output{Error: fmt.Errorf("%d is an invalid intcode", comp.Memory[0])}
			return
		}
	}
}

func (comp Computer) move(addr *int, num int) []int {
	first, last := *addr+1, *addr+1+num
	*addr = last
	return comp.Memory[first:last]
}

func (comp Computer) values(params []int, modes [3]int) []*int {
	vals := make([]*int, 3)

	for i := range params {
		switch modes[i] {
		case ImmediateMode:
			vals[i] = &params[i]
		case PositionMode:
			vals[i] = &comp.Memory[params[i]]
		}
	}

	return vals
}

func decode(code int) (opcode int, modes [3]int, err error) {
	codestr := fmt.Sprintf("%05d", code)

	opcode, err = strconv.Atoi(codestr[3:])
	if err != nil {
		return
	}

	if modes[0], err = strconv.Atoi(string(codestr[2])); err != nil {
		return
	}
	if modes[1], err = strconv.Atoi(string(codestr[1])); err != nil {
		return
	}
	modes[2], err = strconv.Atoi(string(codestr[0]))

	return
}
