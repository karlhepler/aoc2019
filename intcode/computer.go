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

	// OpcodeRelativeBaseOffset adjusts the relative base by the value of its
	// only parameter. The relative base increases (or decreases, if the
	// value is negative) by the value of the parameter.
	OpcodeRelativeBaseOffset = 9
)

const (
	// PositionMode causes the parameter to be interpreted as a position -
	// if the parameter is 50, its value is the value stored at address 50 in
	// memory. Until now, all parameters have been in position mode.
	PositionMode = 0

	// ImmediateMode causes a parameter to be interpreted as a value - if
	// the parameter is 50, its value is simply 50.
	ImmediateMode = 1

	// RelativeMode: Parameters in mode 2, relative mode, behave very similarly
	// to parameters in position mode: the parameter is interpreted as a
	// position. Like position mode, parameters in relative mode can be read
	// from or written to.
	//
	// The important difference is that relative mode parameters don't count
	// from address 0. Instead, they count from a value called the relative
	// base. The relative base starts at 0.
	//
	// The address a relative mode parameter refers to is itself plus the
	// current relative base. When the relative base is 0, relative mode
	// parameters and position mode parameters with the same value refer to the
	// same address.
	RelativeMode = 2
)

// NewComputer returns a pointer to a new Computer instance
func NewComputer() *Computer {
	return &Computer{
		Memory: make([]int, 2048),
	}
}

// Computer is the intcode Computer struct
type Computer struct {
	Memory       []int
	RelativeBase int
}

// Load loads a program into memory
func (comp *Computer) Load(prgm string) error {
	instructions := strings.Split(prgm, ",")

	if comp.Memory == nil || len(comp.Memory) < len(instructions) {
		return fmt.Errorf("NOT ENOUGH MEMORY. NEED >= %v", len(instructions))
	}

	for i, instruction := range instructions {
		comp.Memory[i] = func(s string) int {
			i, err := strconv.Atoi(strings.Trim(s, " "))
			if err != nil {
				log.Fatal(err)
			}
			return i
		}(instruction)
	}

	return nil
}

// Exec executes the program loaded into memory. It receives inputs on
// the given input channel and returns an output channel and a done
// channel, through which errors are passed.
func (comp *Computer) Exec(input <-chan int) (<-chan int, <-chan error) {
	output, done := make(chan int), make(chan error)
	go comp.exec(input, output, done)
	return output, done
}

func (comp *Computer) exec(input <-chan int, output chan<- int, done chan<- error) {
	defer close(done)
	defer close(output)

	if comp.Memory == nil {
		done <- errors.New("NO MEMORY")
		return
	}

	var addr int

	for {
		opcode, modes, err := decode(comp.Memory[addr])
		if err != nil {
			done <- err
			return
		}

		switch opcode {
		case OpcodeHalt:
			return

		case OpcodeInput:
			vals := comp.values(comp.move(&addr, 1), modes)
			*vals[0] = <-input

		case OpcodeOutput:
			vals := comp.values(comp.move(&addr, 1), modes)
			output <- *vals[0]

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

		case OpcodeRelativeBaseOffset:
			vals := comp.values(comp.move(&addr, 1), modes)
			comp.RelativeBase += *vals[0]

		default:
			done <- fmt.Errorf("INVALID OPCODE: %d", opcode)
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
		case RelativeMode:
			vals[i] = &comp.Memory[comp.RelativeBase+params[i]]
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
