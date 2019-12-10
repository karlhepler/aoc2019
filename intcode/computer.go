package intcode

import (
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
		return
	}

	addr := 0
	for {
		switch comp.Memory[addr] {
		case OpcodeHalt:
			return

		case OpcodeInput:
			params := comp.move(&addr, 1)
			comp.Memory[params[0]] = <-inputs

		case OpcodeOutput:
			outputs <- Output{Value: comp.Memory[comp.move(&addr, 1)[0]]}

		case OpcodeAdd:
			params := comp.move(&addr, 3)
			comp.Memory[params[2]] = comp.Memory[params[1]] + comp.Memory[params[0]]

		case OpcodeMultiply:
			params := comp.move(&addr, 3)
			comp.Memory[params[2]] = comp.Memory[params[1]] * comp.Memory[params[0]]

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
