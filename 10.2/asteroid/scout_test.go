package asteroid_test

import (
	"testing"

	"github.com/karlhepler/aoc2019/10.2/asteroid"
)

func TestMonitoringStation(t *testing.T) {
	lines := make(chan string)

	go func() {
		defer close(lines)
		input := []string{
			".#..#",
			".....",
			"#####",
			"....#",
			"...##",
		}

		for _, line := range input {
			lines <- line
		}
	}()

	estation := asteroid.Coord{3, 4}
	evisible := 8

	m := asteroid.BuildMap(lines)
	station, visible := asteroid.MonitoringStation(m)

	if station != estation {
		t.Errorf("[STATION] Expected %v; Received %v", estation, station)
	}
	if len(visible) != evisible {
		t.Errorf("[VISIBLE] Expected %d/%d; Received %d/%d", evisible, len(m), len(visible), len(m))
	}
}
