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
	ScreenWidth  = 640
	ScreenHeight = 480
)

var ui UserInterface
var term *terminal.Terminal
var computer *intcode.Computer

func init() {
	ui = UserInterface{os.Stdin, os.Stdout}
}

func PowerOn() (func() error, error) {
	ui.Println("[ POWER ON ]")

	reset, err := terminator.RawMode()
	if err != nil {
		return reset, err
	}

	term = terminal.NewTerminal(ui, "")
	computer = intcode.NewComputer()

	ui.Println("[ READY ]")

	return reset, nil
}

func LoadGame(prgm string) (*Game, error) {
	ui.Println("[ LOAD GAME ]")

	game := &Game{make(map[Coord]Tile)}
	computer.UpgradeMemory(len(prgm))
	err := computer.Load(prgm)

	ui.Println("[ READY ]")

	return game, err
}

type Game struct {
	Grid map[Coord]Tile
}

func (game *Game) InsertQuarters() {
	// Memory address 0 represents the number of quarters that have
	// been inserted; set it to 2 to play for free.
	computer.Memory[0] = 2
}

func (game *Game) Play() {
	input := make(chan int)
	defer close(input)

	output, done := computer.Exec(input)

	for {
		select {
		case err := <-done:
			if err != nil {
				ui.Fatalf("[ GAME ERROR ]\nERROR: %s\n", err)
			}

			ui.Println("[ GAME OVER ]")
			return
		default:
			if err := game.Update(output, input); err != nil {
				ui.Fatalf("[ GAME ERROR ]\nERROR: %s\n", err)
			}
			if err := game.Render(); err != nil {
				ui.Fatalf("[ GAME ERROR ]\nERROR: %s\n", err)
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
	if err := clear.Run(); err != nil {
		return err
	}

	// Build the buffer
	buffer := make([]byte, ScreenWidth*ScreenHeight)
	for y := 0; y < ScreenHeight; y++ {
		// Generate the byte slice to render
		for x := 0; x < ScreenWidth-1; x++ {
			buffer[x*y+x] = game.Grid[Coord{x, y}].Byte()
		}
		buffer[ScreenWidth] = '\n'
	}

	// Write the buffer to the screen
	numbytes, err := ui.Screen.Write(buffer)
	if err != nil {
		return err
	}
	if numbytes != ScreenWidth*ScreenHeight {
		return fmt.Errorf("incomplete render: %d/%d bytes", numbytes, ScreenWidth*ScreenHeight)
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

func (ui UserInterface) Println(s string) {
	fmt.Fprintln(ui.Screen, s)
}

func (ui UserInterface) Fatalf(s string, d ...interface{}) {
	fmt.Fprintf(ui.Screen, s, d...)
	os.Exit(1)
}
