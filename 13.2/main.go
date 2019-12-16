package main

import (
	"github.com/karlhepler/aoc2019/13.1/arcade"
	"github.com/karlhepler/aoc2019/input"
)

func main() {
	// Power on the arcade and defer power off
	powerOff := arcade.PowerOn()
	defer powerOff()

	// Load the game
	breakout := arcade.LoadGame(<-input.Lines("input/13.1"))

	// Gotta pay to play
	breakout.InsertQuarters()

	// Play the game
	breakout.Play()
}
