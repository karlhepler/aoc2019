package main

import (
	"fmt"
	"os"

	"github.com/karlhepler/aoc2019/13.1/arcade"
	"github.com/karlhepler/aoc2019/input"
)

func main() {
	// Power on the arcade and defer power off
	defer arcade.PowerOn()

	// Load the game
	breakout, err := arcade.LoadGame(<-input.Lines("input/13.1"))
	if err != nil {
		fmt.Printf("[ GAME OVER ]\nERROR: %s\n", err)
		os.Exit(1)
	}

	// Gotta pay to play
	breakout.InsertQuarters()

	// Play the game
	if err := breakout.Play(); err != nil {
		fmt.Printf("[ GAME OVER ]\nERROR: %s\n", err)
		os.Exit(1)
	}

	fmt.Println("[ GAME OVER ]")
}
