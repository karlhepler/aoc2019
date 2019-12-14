package robot

import (
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"os/exec"
	"time"

	"github.com/karlhepler/aoc2019/intcode"
)

type Color int

func (c Color) String() string {
	b := c.Byte()
	if b == 0 {
		return fmt.Sprintf("%d", c)
	}
	return string(b)
}

func (c Color) Byte() byte {
	switch c {
	case Black:
		return ' '
	case White:
		return '#'
	default:
		return 0
	}
}

const (
	Black Color = 0
	White Color = 1
)

type Direction float64

func (d Direction) Byte() byte {
	switch d {
	case Left:
		return '<'
	case Down:
		return 'v'
	case Right:
		return '>'
	case Up:
		return '^'
	default:
		return 0
	}
}

func (d Direction) String() string {
	b := d.Byte()
	if b == 0 {
		return fmt.Sprintf("%f", d)
	}
	return string(b)
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
	Position       Coord
	HullDimensions Coord
	PaintedPanels  map[Coord]Color
	Computer       *intcode.Computer
}

func (rob *Robot) Activate() (numPaintedPanels int, err error) {
	input := make(chan int)
	defer close(input)

	output, done := rob.Computer.Exec(input)

	rob.Render(log.Writer())

	for {
		select {
		case input <- int(rob.Camera()):
			rob.Paint(Color(<-output))
			if err = rob.Turn(Direction(<-output)); err != nil {
				return
			}
			rob.Move()

		case err = <-done:
			return len(rob.PaintedPanels), err

		default:
			time.Sleep(100 * time.Millisecond)
			if err = rob.Render(log.Writer()); err != nil {
				return
			}
		}
	}
}

func (rob *Robot) Render(w io.Writer) error {
	clear := exec.Command("clear")
	clear.Stdout = os.Stdout
	clear.Run()

	for y := 0; y < rob.HullDimensions[1]; y++ {
		// Generate the byte slice to render
		line := make([]byte, rob.HullDimensions[0]+1)
		for x := 0; x < rob.HullDimensions[0]; x++ {
			pos := Coord{x, y}
			if pos == rob.Position {
				line[x] = rob.Direction.Byte()
			} else {
				line[x] = rob.PaintedPanels[Coord{x, y}].Byte()
			}
		}
		line[rob.HullDimensions[0]] = '\n'

		// Render the byte slice
		numbytes, err := w.Write(line)
		if err != nil {
			return err
		}
		if numbytes != rob.HullDimensions[0]+1 {
			return fmt.Errorf("incomplete render: %d/%d bytes", numbytes, rob.HullDimensions[0]+1)
		}
	}

	return nil
}

func (rob *Robot) Turn(dir Direction) (err error) {
	switch dir {
	case Left:
		rob.Direction = Direction(math.Mod(float64(rob.Direction)+0.5, 2))
	case Right:
		rob.Direction = Direction(math.Mod(float64(rob.Direction)-0.5, 2))
	default:
		err = fmt.Errorf("cannot turn %v", dir)
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
		rob.Position[1]--
	case Right:
		rob.Position[0]++
	case Down:
		rob.Position[1]++
	case Left:
		rob.Position[0]--
	}

	// Update hull dimensions
	rob.HullDimensions[0] = max(rob.HullDimensions[0], abs(rob.Position[0]*2))
	rob.HullDimensions[1] = max(rob.HullDimensions[1], abs(rob.Position[1]+1))
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

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
