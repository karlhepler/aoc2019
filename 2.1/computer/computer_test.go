package computer_test

import (
	"testing"

	"github.com/karlhepler/aoc2019/2.1/computer"
)

func TestExec(t *testing.T) {
	t.Run("TestExec", computer.CreateTestExec(computer.Exec))
}
