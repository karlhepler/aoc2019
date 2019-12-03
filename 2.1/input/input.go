package input

import (
	"log"
	"strconv"
	"strings"

	"github.com/karlhepler/aoc2019/input"
)

var cache []int

func Program(noun, verb int) []int {
	if cache == nil {
		strs := strings.Split(<-input.Lines("2.1"), ",")
		cache = make([]int, len(strs))

		for i, s := range strs {
			cache[i] = parseint(s)
		}
	}

	prgm := append(cache[:0:0], cache...)

	prgm[1] = noun
	prgm[2] = verb

	return prgm
}

func parseint(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return i
}
