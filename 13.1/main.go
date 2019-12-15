package main

import (
	"fmt"
	"log"
	"time"

	"github.com/karlhepler/aoc2019/13.1/arcade"
	"github.com/karlhepler/aoc2019/input"
)

func main() {
	start := time.Now()

	game, err := arcade.LoadGame(<-input.Lines("input/13.1"))
	if err != nil {
		log.Fatal(err)
	}

	err = game.Play()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Number of Block Tiles: %d\n", game.NumTiles(arcade.BlockTile))

	fmt.Printf("Time: %v\n", time.Since(start))
}
