package intcode_test

import (
	"reflect"
	"testing"

	"github.com/karlhepler/aoc2019/intcode"
)

func TestNewComputer(t *testing.T) {
	comp := intcode.NewComputer()

	if comp.Memory != nil {
		t.Errorf("Expected memory to be nil. Received %#v\n", comp.Memory)
	}
}

func TestComputerLoad(t *testing.T) {
	tcs := []struct {
		prgm string
		exp  []int
	}{
		{"1,2,3,4,5", []int{1, 2, 3, 4, 5}},
		{"1,0,0,0,99", []int{1, 0, 0, 0, 99}},
		{"1, 1, 1, 4, 99, 5, 6, 0, 99", []int{1, 1, 1, 4, 99, 5, 6, 0, 99}},
	}

	comp := intcode.NewComputer()
	for i, tc := range tcs {
		comp.Load(tc.prgm)
		if !reflect.DeepEqual(tc.exp, comp.Memory) {
			t.Errorf("%d. Expected %#v; Received %#v", i, tc.exp, comp.Memory)
		}
	}
}

func TestComputerExec21(t *testing.T) {
	tcs := []struct {
		initialState []int
		finalState   []int
	}{
		{initialState: nil, finalState: nil},
		{
			initialState: []int{1, 0, 0, 0, 99},
			finalState:   []int{2, 0, 0, 0, 99},
		},
		{
			initialState: []int{2, 3, 0, 3, 99},
			finalState:   []int{2, 3, 0, 6, 99},
		},
		{
			initialState: []int{2, 4, 4, 5, 99, 0},
			finalState:   []int{2, 4, 4, 5, 99, 9801},
		},
		{
			initialState: []int{1, 1, 1, 4, 99, 5, 6, 0, 99},
			finalState:   []int{30, 1, 1, 4, 2, 5, 6, 0, 99},
		},
	}

	comp := intcode.NewComputer()
	for i, tc := range tcs {
		inputs := make(chan int)

		comp.Memory = tc.initialState

		output := <-comp.Exec(inputs)
		if output.Error != nil {
			t.Fatal(output.Error)
		}

		if !reflect.DeepEqual(tc.finalState, comp.Memory) {
			t.Errorf("%d. Expected %#v; Received %#v", i, tc.finalState, comp.Memory)
		}

		close(inputs)
	}
}

func TestComputer51(t *testing.T) {
	//
}
