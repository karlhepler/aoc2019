package sif_test

import (
	"bytes"
	"testing"

	sif1 "github.com/karlhepler/aoc2019/8.1/sif"
	"github.com/karlhepler/aoc2019/8.2/sif"
)

func TestRender(t *testing.T) {
	image := sif1.Decode(2, 2, "0222112222120000")

	var b bytes.Buffer
	if err := sif.Render(&b, image); err != nil {
		t.Fatal(err)
	}

	expected := string([]byte{32, 219, 10, 219, 32, 10})
	if b.String() != expected {
		t.Errorf("Expected %s; Received %s", expected, b.String())
	}
}
