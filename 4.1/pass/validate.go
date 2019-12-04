package pass

import (
	"log"
	"strconv"
)

// OrderBy is used by CheckSort to determine the sort order
type OrderBy byte

const (
	// ASC is order by ascending
	ASC OrderBy = iota + 1
	// DESC is order by descending
	DESC
)

// Validator sets up the validation rules to use with its Valid method
type Validator struct {
	Length      int
	LowerBound  int
	UpperBound  int
	NumAdjacent int
	OrderBy
}

// Valid returns true if the password is valid
func (v Validator) Valid(pw []byte) bool {
	return CheckLength(pw, v.Length) &&
		CheckRange(pw, v.LowerBound, v.UpperBound) &&
		CheckAdjacent(pw, v.NumAdjacent) &&
		CheckSort(pw, v.OrderBy)
}

// CheckLength returns true if bs is of length ln
func CheckLength(bs []byte, ln int) bool {
	return len(bs) == ln
}

// CheckRange returns true if the int value of vs is within the lower and upper bounds
func CheckRange(bs []byte, lower, upper int) bool {
	if len(bs) <= 0 {
		return false
	}

	num, err := strconv.Atoi(string(bs))
	if err != nil {
		log.Fatal(err)
	}

	return num > lower && num < upper
}

// CheckAdjacent returns true if bs have at least one set of adjacent bytes of length n
func CheckAdjacent(bs []byte, n int) bool {
	count := 0

	for i, b := range bs {
		if i == 0 {
			continue
		}

		if b == bs[i-1] {
			count++
		} else {
			count = 0
		}

		if count == n-1 {
			return true
		}
	}

	return false
}

// CheckSort returns true if bs are sorted left-to-right by ob
func CheckSort(bs []byte, ob OrderBy) bool {
	for i, b := range bs {
		if i == 0 {
			continue
		}

		switch {
		case ob == ASC && b < bs[i-1]:
			return false
		case ob == DESC && b > bs[i-1]:
			return false
		}
	}

	return true
}
