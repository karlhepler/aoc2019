package pass

import (
	"github.com/karlhepler/aoc2019/4.1/pass"
)

const (
	ASC  = pass.ASC
	DESC = pass.DESC
)

type Validator pass.Validator

func (v Validator) Valid(pw []byte) bool {
	return pass.CheckLength(pw, v.Length) &&
		pass.CheckRange(pw, v.LowerBound, v.UpperBound) &&
		pass.CheckAdjacent(pw, v.NumAdjacent) &&
		pass.CheckSort(pw, v.OrderBy) &&
		CheckGroup(pw, v.NumAdjacent)
}

// CheckGroup returns true if bs have at least one set of adjacent bytes of
// length n which are not a part of a larger group
func CheckGroup(bs []byte, n int) bool {
	count := 0

	for i, b := range bs {
		if i == 0 {
			continue
		}

		if b == bs[i-1] {
			count++
		} else {
			if count == n-1 {
				return true
			}
			count = 0
		}
	}

	return count == n-1
}
