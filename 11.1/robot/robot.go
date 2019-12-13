package robot

import (
	"errors"
	"fmt"
	"math"

	"github.com/karlhepler/aoc2019/intcode"
)

type Color int

func (c Color) String() string {
	switch c {
	case Black:
		return "□"
	case White:
		return "■"
	default:
		return fmt.Sprintf("%d", c)
	}
}

const (
	Black Color = 0
	White Color = 1
)

type Direction float64

func (d Direction) String() string {
	switch d {
	case Left:
		return "←"
	case Down:
		return "↓"
	case Right:
		return "→"
	case Up:
		return "↑"
	default:
		return fmt.Sprintf("%f", d)
	}
}

const (
	Left  Direction = 0
	Down  Direction = 0.5
	Right Direction = 1
	Up    Direction = 1.5
)

type Coord [2]int

func (c Coord) String() string {
	return fmt.Sprintf("(%d,%d)", c[0], c[1])
}

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
	outputs, halt := rob.Computer.Exec(inputs)

	go func() {
		defer close(inputs)

		for rob.Computer.Running {
			inputs <- int(rob.Camera())
			color, turn := <-outputs, <-outputs

			rob.Paint(Color(color))
			rob.Turn(Direction(turn))
			rob.Move()
		}
	}()

	err = <-halt

	return len(rob.PaintedPanels), err
}

func (rob *Robot) Turn(dir Direction) (err error) {
	switch dir {
	case Left:
		rob.Direction = Direction(math.Mod(float64(rob.Direction)+0.5, 2))
	case Right:
		rob.Direction = Direction(math.Mod(float64(rob.Direction)-0.5, 2))
	default:
		err = errors.New("Robot can only turn Left or Right")
	}

	// Fix 360deg right turn
	if rob.Direction == -0.5 {
		rob.Direction = 1.5
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
	} else if color == Black {
		fmt.Printf("[COLOR] %v\n", color)
	}
	return color
}
