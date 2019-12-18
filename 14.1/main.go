package main

import (
	"log"

	"github.com/karlhepler/aoc2019/14.1/nano"
	"github.com/karlhepler/aoc2019/input"
)

func main() {
	factory := nano.NewFactory()
	result := factory.OrePerFuel(input.Lines("input/14.1"))
	log.Println(result)
}
