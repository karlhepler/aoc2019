package sif_test

import (
	"reflect"
	"testing"

	"github.com/karlhepler/aoc2019/8.1/sif"
)

func TestDecode(t *testing.T) {
	tcs := []struct {
		w     int
		h     int
		data  string
		image sif.Image
	}{
		{3, 2, "123456789012", sif.Image{Width: 3, Height: 2, Layers: [][][]byte{[][]byte{[]byte{'1', '2', '3'}, []byte{'4', '5', '6'}}, [][]byte{[]byte{'7', '8', '9'}, []byte{'0', '1', '2'}}}}},
	}

	for i, tc := range tcs {
		image := sif.Decode(tc.w, tc.h, tc.data)
		if !reflect.DeepEqual(image, tc.image) {
			t.Errorf("%d. Expected %v; Received %v", i, tc.image, image)
		}
	}
}

func TestCheck(t *testing.T) {
	tcs := []struct {
		image sif.Image
		code  int
	}{
		{sif.Decode(3, 2, "123456789012"), 1},
		{sif.Decode(3, 2, "120126781002"), 4},
	}

	for i, tc := range tcs {
		code := sif.Check(tc.image)
		if code != tc.code {
			t.Errorf("%d. Expected %d; Received %d", i, tc.code, code)
		}
	}
}
