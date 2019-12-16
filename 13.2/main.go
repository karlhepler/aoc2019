package main

import (
	"fmt"
	"os"

	"github.com/karlhepler/aoc2019/13.1/arcade"
	"github.com/karlhepler/aoc2019/input"
)

func main() {
	// Power on the arcade and defer power off
	fmt.Println("He flips the power switch on...")
	powerOff, err := arcade.PowerOn()
	if err != nil {
		fmt.Printf("[ GAME OVER ]\nERROR: %s\n", err)
		os.Exit(1)
	}
	defer powerOff()

	// Load the game
	fmt.Println("He inserts his brand new Breakout! game cartridge...")
	breakout, err := arcade.LoadGame(<-input.Lines("input/13.1"))
	if err != nil {
		fmt.Printf("[ GAME OVER ]\nERROR: %s\n", err)
		os.Exit(1)
	}

	// Gotta pay to play
	fmt.Println("He inserts precisely π quarters...")
	breakout.InsertQuarters()

	// Play the game
	breakout.Play()
}
