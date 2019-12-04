package main

import (
	"log"
	"strconv"
	"strings"

	"github.com/karlhepler/aoc2019/4.2/pass"
	"github.com/karlhepler/aoc2019/input"
)

func main() {
	line := <-input.Lines("4.1")
	upper, lower := bounds(line)

	v := pass.Validator{
		Length:      6,
		LowerBound:  lower,
		UpperBound:  upper,
		NumAdjacent: 2,
		OrderBy:     pass.ASC,
	}

	valid := make([]int, 0)
	for i := lower + 1; i < upper; i++ {
		if v.Valid([]byte(strconv.Itoa(i))) {
			valid = append(valid, i)
		}
	}

	log.Printf("Num Valid: %v", len(valid))
}

func bounds(s string) (upper, lower int) {
	var err error
	bs := strings.Split(s, "-")

	lower, err = strconv.Atoi(bs[0])
	if err != nil {
		log.Fatal(err)
	}

	upper, err = strconv.Atoi(bs[1])
	if err != nil {
		log.Fatal(err)
	}

	return
}
