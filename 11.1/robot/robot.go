package robot

import (
	"errors"

	"github.com/karlhepler/aoc2019/intcode"
)

type Color int

const (
	Black Color = 0
	White Color = 1
)

type Direction float64

const (
	Left  Direction = 0
	Right Direction = 1
	Up    Direction = 2
	Down  Direction = 3
)

type Coord [2]int

func New() *Robot {
	return &Robot{
		Direction:     Up,
		PaintedPanels: make(map[Coord]Color),
		Computer:      intcode.NewComputer(),
	}
}

type Robot struct {
	Direction
	Position      Coord
	PaintedPanels map[Coord]Color
	Computer      *intcode.Computer
}

func (rob *Robot) Activate() (numPaintedPanels int, err error) {
	inputs := make(chan int)
	outputs := rob.Computer.Exec(inputs)

	go func() {
		defer close(inputs)

		for rob.Computer.Running {
			inputs <- int(rob.Camera())

			color := <-outputs
			if err = color.Error; err != nil {
				return
			}

			turn := <-outputs
			if err = turn.Error; err != nil {
				return
			}

			rob.Paint(Color(color.Value))
			if err = rob.Turn(Direction(turn.Value)); err != nil {
				return
			}

			rob.Move()
		}
	}()

	<-rob.Computer.Halt

	return len(rob.PaintedPanels), err
}

func (rob *Robot) Turn(dir Direction) (err error) {
	switch {
	// Left
	case dir == Left && rob.Direction == Left:
		rob.Direction = Down
	case dir == Left && rob.Direction == Down:
		rob.Direction = Right
	case dir == Left && rob.Direction == Right:
		rob.Direction = Up
	case dir == Left && rob.Direction == Up:
		rob.Direction = Left

	// Right
	case dir == Right && rob.Direction == Left:
		rob.Direction = Up
	case dir == Right && rob.Direction == Up:
		rob.Direction = Right
	case dir == Right && rob.Direction == Right:
		rob.Direction = Down
	case dir == Right && rob.Direction == Down:
		rob.Direction = Left

	default:
		err = errors.New("Robot can only turn Left or Right")
	}

	return
}

func (rob *Robot) Move() {
	switch rob.Direction {
	case Up:
		rob.Position[1] -= 1
	case Right:
		rob.Position[0] += 1
	case Down:
		rob.Position[1] += 1
	case Left:
		rob.Position[1] -= 1
	}
}

func (rob *Robot) Paint(color Color) {
	rob.PaintedPanels[rob.Position] = color
}

func (rob *Robot) Camera() Color {
	color, ok := rob.PaintedPanels[rob.Position]
	if !ok {
		color = Black
	}
	return color
}
