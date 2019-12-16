package arcade

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/karlhepler/aoc2019/13.2/terminator"
	"github.com/karlhepler/aoc2019/intcode"
	tdim "github.com/wayneashleyberry/terminal-dimensions"
	"golang.org/x/crypto/ssh/terminal"
)

var (
	ScreenWidth  int
	ScreenHeight int
)

var ui UserInterface
var term *terminal.Terminal
var computer *intcode.Computer

func init() {
	ui = UserInterface{os.Stdin, os.Stdout}

	w, err := tdim.Width()
	if err != nil {
		fatal(err)
	}
	ScreenWidth = int(w) + 1
	h, err := tdim.Height()
	if err != nil {
		fatal(err)
	}
	ScreenHeight = int(h) + 1
}

func PowerOn() func() error {
	ui.Println("[ POWER ON ]")
	ui.Printf("[ DISPLAY SIZE %dx%d ]\n", ScreenWidth, ScreenHeight)

	reset, err := terminator.RawMode()
	if err != nil {
		fatal(err)
	}

	term = terminal.NewTerminal(ui, "")
	term.SetSize(ScreenWidth, ScreenHeight)
	computer = intcode.NewComputer()

	ui.Println("[ READY ]")

	return reset
}

func LoadGame(prgm string) *Game {
	ui.Println("[ LOAD GAME ]")

	game := &Game{make(map[Coord]Tile)}
	computer.UpgradeMemory(len(prgm))
	if err := computer.Load(prgm); err != nil {
		fatal(err)
	}

	ui.Println("[ READY ]")

	return game
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
	ui.Println("[ PLAY GAME ]")

	input := make(chan int)
	defer close(input)

	output, done := computer.Exec(input)

	for {
		start := time.Now()

		select {
		case err := <-done:
			if err != nil {
				fatal(err)
			}

			ui.Println("[ GAME OVER ]")
			return
		default:
			game.UpdateState(output)
			game.Render()
			// game.ProcessInput(input)

			time.Sleep(time.Duration(16600*int64(time.Microsecond) - time.Since(start).Microseconds()))
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

func (game *Game) UpdateState(state <-chan int) {
	var i, x, y, tile int
	for {
		switch i % 3 {
		case 0:
			x = <-state
		case 1:
			y = <-state
		case 2:
			tile = <-state
		}
		i++

		// Show the score in the segment display in this case
		if x == -1 && y == 0 {
			ui.Printf("[ SCORE %d ]\n", tile)
			return
		}

		game.Grid[Coord{x, y}] = Tile(tile)
	}
}

func (game Game) Render() {
	// Build the buffer
	buffer := make([]byte, ScreenWidth*ScreenHeight)
	for y := 0; y < ScreenHeight; y++ {
		// Generate the byte slice to render
		for x := 0; x < ScreenWidth; x++ {
			buffer[x+y*(ScreenWidth-1)] = game.Grid[Coord{x, y}].Byte()
		}
	}

	// Clear the screen
	draw(make([]byte, ScreenWidth*ScreenHeight))

	// Write the buffer to the screen
	draw(buffer)
}

func draw(buffer []byte) {
	numbytes, err := ui.Screen.Write(buffer)

	if err != nil {
		fatal(err)
	}
	if numbytes != ScreenWidth*ScreenHeight {
		fatal(fmt.Errorf("incomplete render: %d/%d bytes", numbytes, ScreenWidth*ScreenHeight))
	}
}

func (game Game) ProcessInput(input chan<- int) {
	// TODO how to read input in the terminal?

	// switch /* INPUT */ {
	// case 80: // left arrow
	// 	input <- -1
	// case 79: // right arrow
	// 	input <- 1
	// case 41: // escape
	// 	ui.Fatalf("[ ESCAPE ]")
	// default:
	// 	input <- 0
	// }
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

func (ui UserInterface) Printf(s string, d ...interface{}) {
	fmt.Fprintf(ui.Screen, s, d...)
}

func (ui UserInterface) Fatalf(s string, d ...interface{}) {
	ui.Printf(s, d...)
	os.Exit(1)
}

func fatal(err error) {
	ui.Fatalf("[ GAME ERROR ]\nERROR: %s\n", err)
}
