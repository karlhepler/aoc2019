package moon

import (
	"regexp"
	"strconv"
)

func NewVector(def string) (vec Vector) {
	// Get matches
	re := regexp.MustCompile(`x=([-0-9]+)|y=([-0-9]+)|z=([-0-9]+)`)
	matches := re.FindAllStringSubmatch(def, -1)
	if matches == nil {
		return
	}

	// Define the vector
	var err error
	for i := 0; i < 3; i++ {
		vec[i], err = strconv.Atoi(matches[i][i+1])
		if err != nil {
			vec[i] = 0
		}
	}

	return
}

type Vector [3]int

func (a Vector) Add(b Vector) Vector {
	return Vector{
		a[0] + b[0],
		a[1] + b[1],
		a[2] + b[2],
	}
}
