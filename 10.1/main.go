package main

import (
	"fmt"
	"log"
	"time"

	"github.com/karlhepler/aoc2019/10.1/asteroid"
	"github.com/karlhepler/aoc2019/input"
)

func main() {
	start := time.Now()

	dispatcher := asteroid.NewDispatcher(
		asteroid.BuildMap(input.Lines("input/10.1")),
	)

	var most int
	var numReported int
	messages := dispatcher.SurveyAndListen()
	for numReported < len(dispatcher.Map) {
		msg := <-messages
		numReported++
		most = max(most, msg.NumVisibleAsteroids)
	}

	log.Println(most)

	fmt.Printf("Time: %v\n", time.Since(start))
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
