package robot_test

import (
	"testing"

	"github.com/karlhepler/aoc2019/11.1/robot"
)

func TestRobotTurn(t *testing.T) {
	tcs := []struct {
		turn      robot.Direction
		direction robot.Direction
	}{
		{robot.Left, robot.Left},
		{robot.Left, robot.Down},
		{robot.Left, robot.Right},
		{robot.Left, robot.Up},
		{robot.Right, robot.Right},
		{robot.Right, robot.Down},
		{robot.Right, robot.Left},
		{robot.Right, robot.Up},
	}

	rob := robot.New()
	if rob.Direction != robot.Up {
		t.Fatal("Robot MUST start pointing Up!")
	}

	for i, tc := range tcs {
		rob.Turn(tc.turn)
		if rob.Direction != tc.direction {
			t.Errorf("%d. Expected %v; Received %v\n", i, tc.direction, rob.Direction)
		}
	}
}
