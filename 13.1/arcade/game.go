package arcade

import (
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"

	"github.com/karlhepler/aoc2019/13.2/terminator"
	"github.com/karlhepler/aoc2019/intcode"
	"golang.org/x/crypto/ssh/terminal"
)

const (
	ScreenWidth  = 300
	ScreenHeight = 300
)

var ui UserInterface
var term *terminal.Terminal
var computer *intcode.Computer

func PowerOn() (func() error, error) {
	reset, err := terminator.RawMode()
	if err != nil {
		return reset, err
	}

	ui = UserInterface{os.Stdin, os.Stdout}
	term = terminal.NewTerminal(ui, "")
	computer = intcode.NewComputer()

	return reset, nil
}

func LoadGame(prgm string) (*Game, error) {
	game := &Game{make(map[Coord]Tile)}
	computer.UpgradeMemory(len(prgm))
	return game, computer.Load(prgm)
}

type Game struct {
	Grid map[Coord]Tile
}

func (game *Game) InsertQuarters() {
	// Memory address 0 represents the number of quarters that have
	// been inserted; set it to 2 to play for free.
	computer.Memory[0] = 2
}

func (game *Game) Play() error {
	input := make(chan int)
	defer close(input)

	output, done := computer.Exec(input)

	for {
		select {
		case err := <-done:
			return err
		default:
			if err := game.Update(output, input); err != nil {
				return err
			}
			if err := game.Render(); err != nil {
				return err
			}
		}
	}
}

func (game Game) NumTiles(tile Tile) (num int) {
	for _, t := range game.Grid {
		if t == tile {
			num++
		}
	}
	return
}

func (game *Game) Update(req <-chan int, res chan<- int) error {
	x, y, tile := <-req, <-req, Tile(<-req)
	game.Grid[Coord{x, y}] = tile

	position := make([]byte, 1)
	_, err := ui.Joystick.Read(position)
	if err != nil {
		return err
	}

	switch position[0] {
	case 80: // left arrow
		res <- -1
	case 79: // right arrow
		res <- 1
	case 41: // escape
		return errors.New("you pressed the power button")
	default:
		res <- 0
	}

	return nil
}

func (game Game) Render() error {
	clear := exec.Command("clear")
	clear.Stdout = ui.Screen
	clear.Run()

	screenWidth, screenHeight := 300, 300
	buffer := make([]byte, screenWidth*screenHeight)

	for y := 0; y < scrnHeight; y++ {
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

type Coord [2]int

type Tile int

func (tile Tile) Byte() byte {
	switch tile {
	case EmptyTile:
		return ' '
	case WallTile:
		return '+'
	case BlockTile:
		return '#'
	case PaddleTile:
		return '='
	case BallTile:
		return 'o'
	}
	return 255
}

const (
	EmptyTile Tile = iota
	WallTile
	BlockTile
	PaddleTile
	BallTile
)

type UserInterface struct {
	Joystick io.Reader
	Screen   io.Writer
}

func (ui UserInterface) Read(p []byte) (n int, err error) {
	return ui.Joystick.Read(p)
}

func (ui UserInterface) Write(p []byte) (n int, err error) {
	return ui.Screen.Write(p)
}
