package main

import (
	"fmt"
	"log"
	"time"

	"github.com/karlhepler/aoc2019/10.2/asteroid"
	"github.com/karlhepler/aoc2019/input"
)

func main() {
	start := time.Now()

	m := asteroid.BuildMap(input.Lines("10.1"))
	station, visible := asteroid.MonitoringStation(m)

	vaporized := 0

	for len(m) > 1 {
		for _, ast := range visible {
			i := m.IndexOf(ast)
			m = append(m[:i], m[i+1:]...)
			vaporized++

			if vaporized == 200 {
				log.Printf("[SOLUTION] %v\n", ast[0]*100+ast[1])
				break
			}
		}

		vres := make(chan asteroid.VisibleResponse)
		go asteroid.Visible(m, station, vres)
		res := <-vres
		visible = res.Visible
	}

	fmt.Printf("Time: %v\n", time.Since(start))
}
