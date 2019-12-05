package input

import (
	"log"
	"strconv"
	"strings"

	"github.com/karlhepler/aoc2019/input"
)

func Program() (prgm []int) {
	strs := strings.Split(<-input.Lines("5.1"), ",")
	prgm = make([]int, len(strs))

	for i, s := range strs {
		prgm[i] = parseint(s)
	}

	return
}

func parseint(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return i
}
