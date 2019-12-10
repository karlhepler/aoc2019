package computer_test

// import (
// 	"testing"

// 	"github.com/karlhepler/aoc2019/7.2/computer"
// )

// func TestAmplifierChain(t *testing.T) {
// 	tcs := []struct {
// 		input         int
// 		prgm          []int
// 		phaseSettings [5]int
// 		expected      int
// 	}{
// 		{0, []int{3, 26, 1001, 26, -4, 26, 3, 27, 1002, 27, 2, 27, 1, 27, 26, 27, 4, 27, 1001, 28, -1, 28, 1005, 28, 6, 99, 0, 0, 5}, [5]int{9, 8, 7, 6, 5}, 139629729},
// 		{0, []int{3, 52, 1001, 52, -5, 52, 3, 53, 1, 52, 56, 54, 1007, 54, 5, 55, 1005, 55, 26, 1001, 54, -5, 54, 1105, 1, 12, 1, 53, 54, 53, 1008, 54, 0, 55, 1001, 55, 1, 55, 2, 53, 55, 53, 4, 53, 1001, 56, -1, 56, 1005, 56, 6, 99, 0, 0, 0, 0, 10}, [5]int{9, 7, 8, 5, 6}, 18216},
// 	}

// 	for i, tc := range tcs {
// 		chain := computer.NewAmplifierChain(tc.prgm, tc.phaseSettings)

// 		output, err := chain.Exec(tc.input)
// 		if err != nil {
// 			t.Fatalf("%d. %s", i, err)
// 		}

// 		if output != tc.expected {
// 			t.Errorf("%d. Expected %d; Received %d", i, tc.expected, output)
// 		}
// 	}
// }
