package arcade

import (
	"fmt"
	"io"
	"os"
	"os/exec"

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
		ui.Fatalf("[ GAME ERROR ]\nERROR: %s\n", err)
	}
	ScreenWidth = int(w)
	h, err := tdim.Height()
	if err != nil {
		ui.Fatalf("[ GAME ERROR ]\nERROR: %s\n", err)
	}
	ScreenHeight = int(h)
}

func PowerOn() func() error {
	ui.Println("[ POWER ON ]")

	reset, err := terminator.RawMode()
	if err != nil {
		ui.Fatalf("[ GAME ERROR ]\nERROR: %s\n", err)
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
		ui.Fatalf("[ GAME ERROR ]\nERROR: %s\n", err)
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
		select {
		case err := <-done:
			if err != nil {
				ui.Fatalf("[ GAME ERROR ]\nERROR: %s\n", err)
			}

			ui.Println("[ GAME OVER ]")
			return
		default:
			ui.Println("[ DEBUG ]")
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
	x, y, tile := <-req, <-req, <-req

	// Show the score in the segment display in this case
	if x == -1 && y == 0 {
		ui.Printf("[ SCORE %d ]\n", tile)
		return nil
	}

	game.Grid[Coord{x, y}] = Tile(tile)

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
		ui.Fatalf("[ POWER OFF ]")
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

func (ui UserInterface) Printf(s string, d ...interface{}) {
	fmt.Fprintf(ui.Screen, s, d...)
}

func (ui UserInterface) Fatalf(s string, d ...interface{}) {
	ui.Printf(s, d...)
	os.Exit(1)
}
