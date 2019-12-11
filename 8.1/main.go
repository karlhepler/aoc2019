package main

import (
	"fmt"
	"time"

	"github.com/karlhepler/aoc2019/8.1/sif"
	"github.com/karlhepler/aoc2019/input"
)

func main() {
	start := time.Now()

	image := sif.Decode(25, 6, <-input.Lines("8.1"))
	code := sif.Check(image)

	fmt.Printf("Check Code: %d\n", code)

	fmt.Printf("Time: %v\n", time.Since(start))
}
