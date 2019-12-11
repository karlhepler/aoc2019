package sif

import (
	"fmt"
	"io"

	"github.com/karlhepler/aoc2019/8.1/sif"
)

const (
	black = '0'
	white = '1'
	trans = '2'
)

func Render(w io.Writer, image sif.Image) error {
	for y := 0; y < image.Height; y++ {
		// Get the real pixels for the line
		line := make([]byte, image.Width+1)
		for x := 0; x < image.Width; x++ {
			line[x] = pixel(x, y, image.Layers)
		}
		line[image.Width] = '\n'

		// Render the line
		numbytes, err := w.Write(line)
		if err != nil {
			return err
		}
		if numbytes != image.Width+1 {
			return fmt.Errorf("Incomplete render: %d/%d bytes", numbytes, image.Width+1)
		}
	}

	return nil
}

func pixel(x, y int, layers [][][]byte) byte {
	for _, layer := range layers {
		switch layer[y][x] {
		case white:
			return 219 // '█'
		case black:
			return 32 // ' '
		}
	}

	return trans
}
