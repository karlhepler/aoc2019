package computer_test

import (
	"testing"

	ocomp "github.com/karlhepler/aoc2019/2.1/computer"
	"github.com/karlhepler/aoc2019/5.1/computer"
)

func TestOrigExec(t *testing.T) {
	t.Run("TestOrigExec", ocomp.CreateTestExec(computer.Exec))
}
