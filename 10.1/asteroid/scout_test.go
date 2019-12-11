package asteroid_test

import (
	"testing"

	"github.com/karlhepler/aoc2019/10.1/asteroid"
)

func TestScoutSearch(t *testing.T) {
	astmap := asteroid.Map{
		asteroid.Asteroid{asteroid.Vector{1, 0}},
		asteroid.Asteroid{asteroid.Vector{4, 0}},
		asteroid.Asteroid{asteroid.Vector{0, 2}},
		asteroid.Asteroid{asteroid.Vector{1, 2}},
		asteroid.Asteroid{asteroid.Vector{2, 2}},
		asteroid.Asteroid{asteroid.Vector{3, 2}},
		asteroid.Asteroid{asteroid.Vector{4, 2}},
		asteroid.Asteroid{asteroid.Vector{4, 3}},
		asteroid.Asteroid{asteroid.Vector{3, 4}},
		asteroid.Asteroid{asteroid.Vector{4, 4}},
	}

	messages := make(chan asteroid.Message)
	defer close(messages)

	tcs := []struct {
		scout      asteroid.Scout
		numVisible int
	}{
		{asteroid.NewScout(astmap.Copy(), asteroid.Vector{1, 0}, messages), 7},
		{asteroid.NewScout(astmap.Copy(), asteroid.Vector{4, 0}, messages), 7},
		{asteroid.NewScout(astmap.Copy(), asteroid.Vector{0, 2}, messages), 6},
		{asteroid.NewScout(astmap.Copy(), asteroid.Vector{1, 2}, messages), 7},
		{asteroid.NewScout(astmap.Copy(), asteroid.Vector{2, 2}, messages), 7},
		{asteroid.NewScout(astmap.Copy(), asteroid.Vector{3, 2}, messages), 7},
		{asteroid.NewScout(astmap.Copy(), asteroid.Vector{4, 2}, messages), 5},
		{asteroid.NewScout(astmap.Copy(), asteroid.Vector{4, 3}, messages), 7},
		{asteroid.NewScout(astmap.Copy(), asteroid.Vector{3, 4}, messages), 8},
		{asteroid.NewScout(astmap.Copy(), asteroid.Vector{4, 4}, messages), 7},
	}

	for i, tc := range tcs {
		go tc.scout.SearchAndReport()
		msg := <-messages

		if msg.NumVisibleAsteroids != tc.numVisible {
			t.Errorf("%d. Expected %d; Received %d", i, tc.numVisible, msg.NumVisibleAsteroids)
		}
	}
}
