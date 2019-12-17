package arcade

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/karlhepler/aoc2019/13.1/terminator"
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
	tty, err := os.OpenFile("/dev/tty", os.O_RDWR, 0755)
	if err != nil {
		panic(err)
	}
	ui = UserInterface{tty}

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
			dirchan := game.GetDirection()
			direction := 0
		gameloop:
			for {
				select {
				case x := <-output:
					y, tile := <-output, <-output
					if x == -1 && y == 0 {
						ui.Printf("[ SCORE %d ]\n", tile)
					} else {
						game.Grid[Coord{x, y}] = Tile(tile)
					}
				case input <- 0:
					break gameloop
				case direction = <-dirchan:
					input <- direction
					break gameloop
				default:
				}
			}
			game.Render()

			fps := 5
			time.Sleep(time.Duration(int64((1000/fps)*1000)*int64(time.Microsecond) - time.Since(start).Microseconds()))
		}
	}
}

func (game Game) GetDirection() <-chan int {
	direction := make(chan int)

	go func() {
		defer close(direction)

		buf := make([]byte, 1)
		numBytes, err := ui.Read(buf)

		switch {
		case err != nil:
			fatal(err)
		case numBytes < 1:
			direction <- 0
		case buf[0] == 44: // ,
			direction <- -1
		case buf[0] == 46: // .
			direction <- 1
		case buf[0] == 113: // q
			ui.Fatalf("[ QUIT ]")
		default:
			direction <- 0
		}
	}()

	return direction
}

func (game Game) NumTiles(tile Tile) (num int) {
	for _, t := range game.Grid {
		if t == tile {
			num++
		}
	}
	return
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

	// Write the buffer to the screen
	draw(buffer)
}

func draw(buffer []byte) {
	numbytes, err := ui.Write(buffer)

	if err != nil {
		fatal(err)
	}
	if numbytes != ScreenWidth*ScreenHeight {
		fatal(fmt.Errorf("incomplete render: %d/%d bytes", numbytes, ScreenWidth*ScreenHeight))
	}
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
	io.ReadWriter
}

func (ui UserInterface) Println(s interface{}) {
	fmt.Fprintln(ui, s)
}

func (ui UserInterface) Printf(s string, d ...interface{}) {
	fmt.Fprintf(ui, s, d...)
}

func (ui UserInterface) Fatalf(s string, d ...interface{}) {
	ui.Printf(s, d...)
	os.Exit(1)
}

func fatal(err error) {
	ui.Fatalf("[ GAME ERROR ]\nERROR: %s\n", err)
}
