package main

import (
	"fmt"
	"log"
	"time"

	sif1 "github.com/karlhepler/aoc2019/8.1/sif"
	"github.com/karlhepler/aoc2019/8.2/sif"
	"github.com/karlhepler/aoc2019/input"
)

func main() {
	start := time.Now()

	image := sif1.Decode(25, 6, <-input.Lines("8.1"))

	sif.Render(log.Writer(), image)

	fmt.Printf("Time: %v\n", time.Since(start))
}
