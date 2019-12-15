package arcade

import (
	"io"

	"github.com/karlhepler/aoc2019/intcode"
)

func LoadGame(prgm string) (*Game, error) {
	game := &Game{
		Grid:     make(map[Coord]Tile),
		Computer: intcode.NewComputer(),
	}
	game.Computer.UpgradeMemory(len(prgm))
	return game, game.Computer.Load(prgm)
}

type Game struct {
	Grid     map[Coord]Tile
	Computer *intcode.Computer
}

func (game *Game) Play() error {
	input := make(chan int)
	defer close(input)

	output, done := game.Computer.Exec(input)

	for {
		select {
		case err := <-done:
			return err
		default:
			x, y, tile := <-output, <-output, Tile(<-output)
			game.Grid[Coord{x, y}] = tile
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

func (game Game) Render(w io.Writer) error {
	return nil
}

type Coord [2]int

type Tile int

const (
	EmptyTile Tile = iota
	WallTile
	BlockTile
	PaddleTile
	BallTile
)
