package computer

import (
	"fmt"

	comp2 "github.com/karlhepler/aoc2019/2.1/computer"
	comp51 "github.com/karlhepler/aoc2019/5.1/computer"
	comp52 "github.com/karlhepler/aoc2019/5.2/computer"
)

type Output struct {
	Value int
	Error error
}

func NewComputer(prgm []int) Computer {
	return Computer{
		Program: append(prgm[:0:0], prgm...),
	}
}

type Computer struct {
	Program []int
}

func (comp Computer) Exec(input <-chan int, output chan<- *Output) {
	for i, num := 0, len(comp.Program); i < num; {
		opcode, modes, err := comp51.ParseOpcode(comp.Program[i])
		if err != nil {
			output <- &Output{Value: -1, Error: err}
			return
		}

		switch opcode {
		case comp2.OpcodeHalt:
			output <- nil
			return

		case comp2.OpcodeAdd:
			params := [3]*int{&comp.Program[i+1], &comp.Program[i+2], &comp.Program[i+3]}
			if err := comp51.Add(&comp.Program, modes, params); err != nil {
				output <- &Output{Value: -1, Error: err}
				return
			}
			i += 4

		case comp2.OpcodeMult:
			params := [3]*int{&comp.Program[i+1], &comp.Program[i+2], &comp.Program[i+3]}
			if err := comp51.Multiply(&comp.Program, modes, params); err != nil {
				output <- &Output{Value: -1, Error: err}
				return
			}
			i += 4

		case comp51.OpcodeInput:
			comp.Program[comp.Program[i+1]] = <-input
			i += 2

		case comp51.OpcodeOutput:
			val, err := comp51.Output(&comp.Program, modes, &comp.Program[i+1])
			if err != nil {
				output <- &Output{Value: -1, Error: err}
				return
			}
			if val != 0 {
				output <- &Output{Value: val}
				return
			}
			i += 2

		case comp52.OpcodeJumpTrue:
			params := [2]*int{&comp.Program[i+1], &comp.Program[i+2]}
			if err := comp52.JumpIf(true, &i, &comp.Program, modes, params); err != nil {
				output <- &Output{Value: -1, Error: err}
				return
			}

		case comp52.OpcodeJumpFalse:
			params := [2]*int{&comp.Program[i+1], &comp.Program[i+2]}
			if err := comp52.JumpIf(false, &i, &comp.Program, modes, params); err != nil {
				output <- &Output{Value: -1, Error: err}
				return
			}

		case comp52.OpcodeLessThan:
			params := [3]*int{&comp.Program[i+1], &comp.Program[i+2], &comp.Program[i+3]}
			if err := comp52.Compare(comp52.LessThan, &comp.Program, modes, params); err != nil {
				output <- &Output{Value: -1, Error: err}
				return
			}
			i += 4

		case comp52.OpcodeEquals:
			params := [3]*int{&comp.Program[i+1], &comp.Program[i+2], &comp.Program[i+3]}
			if err := comp52.Compare(comp52.Equals, &comp.Program, modes, params); err != nil {
				output <- &Output{Value: -1, Error: err}
				return
			}
			i += 4

		default:
			output <- &Output{Value: -1, Error: fmt.Errorf("%v is an invalid opcode", comp.Program[i])}
			return
		}
	}
}
