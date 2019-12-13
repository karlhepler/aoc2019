package intcode_test

import (
	"testing"

	"github.com/karlhepler/aoc2019/input"
	"github.com/karlhepler/aoc2019/intcode"
)

// func TestNewComputer(t *testing.T) {
// 	comp := intcode.NewComputer()

// 	if len(comp.Memory) != 2048 {
// 		t.Errorf("Expected memory size to be 2048. Received %d\n", len(comp.Memory))
// 	}
// }

// func TestComputerLoad(t *testing.T) {
// 	tcs := []struct {
// 		prgm string
// 		exp  []int
// 	}{
// 		{"1,2,3,4,5", []int{1, 2, 3, 4, 5}},
// 		{"1,0,0,0,99", []int{1, 0, 0, 0, 99}},
// 		{"1, 1, 1, 4, 99, 5, 6, 0, 99", []int{1, 1, 1, 4, 99, 5, 6, 0, 99}},
// 	}

// 	for i, tc := range tcs {
// 		comp := &intcode.Computer{
// 			Memory: make([]int, len(tc.exp)),
// 		}

// 		if err := comp.Load(tc.prgm); err != nil {
// 			t.Fatal(err)
// 		}

// 		if !reflect.DeepEqual(tc.exp, comp.Memory) {
// 			t.Errorf("%d. Expected %#v; Received %#v", i, tc.exp, comp.Memory)
// 		}
// 	}
// }

// func TestComputerExecMemory(t *testing.T) {
// 	tcs := []struct {
// 		initialState []int
// 		finalState   []int
// 	}{
// 		{
// 			initialState: []int{1, 0, 0, 0, 99},
// 			finalState:   []int{2, 0, 0, 0, 99},
// 		},
// 		{
// 			initialState: []int{2, 3, 0, 3, 99},
// 			finalState:   []int{2, 3, 0, 6, 99},
// 		},
// 		{
// 			initialState: []int{2, 4, 4, 5, 99, 0},
// 			finalState:   []int{2, 4, 4, 5, 99, 9801},
// 		},
// 		{
// 			initialState: []int{1, 1, 1, 4, 99, 5, 6, 0, 99},
// 			finalState:   []int{30, 1, 1, 4, 2, 5, 6, 0, 99},
// 		},
// 		{
// 			initialState: []int{1002, 4, 3, 4, 33},
// 			finalState:   []int{1002, 4, 3, 4, 99},
// 		},
// 		{
// 			initialState: []int{1101, 100, -1, 4, 0},
// 			finalState:   []int{1101, 100, -1, 4, 99},
// 		},
// 	}

// 	comp := intcode.NewComputer()
// 	for i, tc := range tcs {
// 		inputs := make(chan int)

// 		comp.Memory = tc.initialState

// 		_, halt := comp.Exec(inputs)
// 		err := <-halt
// 		close(inputs)
// 		if err != nil {
// 			t.Fatal(err)
// 		}

// 		if !reflect.DeepEqual(tc.finalState, comp.Memory) {
// 			t.Errorf("%d. Expected %#v; Received %#v", i, tc.finalState, comp.Memory)
// 		}
// 	}
// }

// func TestComputerOutput(t *testing.T) {
// 	tcs := []struct {
// 		prgm   string
// 		input  int
// 		output int
// 	}{
// 		// Position mode
// 		{prgm: "3,9,8,9,10,9,4,9,99,-1,8", input: 8, output: 1},
// 		{prgm: "3,9,8,9,10,9,4,9,99,-1,8", input: 7, output: 0},
// 		{prgm: "3,9,7,9,10,9,4,9,99,-1,8", input: 7, output: 1},
// 		{prgm: "3,9,7,9,10,9,4,9,99,-1,8", input: 8, output: 0},
// 		// Immediate mode
// 		{prgm: "3,3,1108,-1,8,3,4,3,99", input: 8, output: 1},
// 		{prgm: "3,3,1108,-1,8,3,4,3,99", input: 7, output: 0},
// 		{prgm: "3,3,1107,-1,8,3,4,3,99", input: 7, output: 1},
// 		{prgm: "3,3,1107,-1,8,3,4,3,99", input: 8, output: 0},
// 		// Jump test position mode
// 		{prgm: "3,12,6,12,15,1,13,14,13,4,13,99,-1,0,1,9", input: 0, output: 0},
// 		{prgm: "3,12,6,12,15,1,13,14,13,4,13,99,-1,0,1,9", input: 1, output: 1},
// 		// Jump test immediate mode
// 		{prgm: "3,3,1105,-1,9,1101,0,0,12,4,12,99,1", input: 0, output: 0},
// 		{prgm: "3,3,1105,-1,9,1101,0,0,12,4,12,99,1", input: 1, output: 1},
// 		// Large example
// 		{prgm: "3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31,1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,104,999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99", input: 7, output: 999},
// 		{prgm: "3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31,1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,104,999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99", input: 8, output: 1000},
// 		{prgm: "3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31,1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,104,999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99", input: 9, output: 1001},
// 		// Relative base
// 		{prgm: "104,1125899906842624,99", input: -1, output: 1125899906842624},
// 		{prgm: "1102,34915192,34915192,7,4,7,99,0", input: -1, output: 1219070632396864},
// 		{prgm: "109,1,204,-1,1001,100,1,100,1008,100,16,101,1006,101,0,99", input: -1, output: 109},
// 	}

// 	comp := intcode.NewComputer()
// 	for i, tc := range tcs {
// 		inputs := make(chan int)

// 		if err := comp.Load(tc.prgm); err != nil {
// 			t.Fatal(err)
// 		}

// 		if tc.input >= 0 {
// 			go func() {
// 				inputs <- tc.input
// 			}()
// 		}

// 		outputs, halt := comp.Exec(inputs)

// 		select {
// 		case output := <-outputs:
// 			if output != tc.output {
// 				t.Errorf("%d. Expected %#v; Received %#v", i, tc.output, output)
// 			}
// 		case err := <-halt:
// 			close(inputs)
// 			t.Fatal(err)
// 		}
// 	}
// }

func TestDay5Part1(t *testing.T) {
	comp := intcode.NewComputer()
	if err := comp.Load(<-input.Lines("5.1", true)); err != nil {
		t.Fatal(err)
	}

	var code int
	inputs := make(chan int)
	codes, halt := comp.Exec(inputs)

	go func() {
		defer close(inputs)
		inputs <- 1
		for code = range codes {
			//
		}
	}()

	if err := <-halt; err != nil {
		t.Fatal(err)
	}

	if code != 5074395 {
		t.Errorf("Expected %d; Received %d\n", 5074395, code)
	}
}
