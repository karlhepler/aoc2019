package sif

import (
	"log"
	"strconv"
)

type Image struct {
	Width  int
	Height int
	Layers [][][]int
}

func Decode(w, h int, data string) (image Image, err error) {
	image = Image{
		Width:  w,
		Height: h,
		Layers: make([][][]int, 0),
	}

	for d := 0; d < len(data); {
		layer := make([][]int, image.Height)
		for y := 0; y < image.Height; y++ {
			line := make([]int, image.Width)
			for x := 0; x < image.Width; x++ {
				line[x], err = strconv.Atoi(string(data[d]))
				if err != nil {
					return
				}
				d++
			}
			layer[y] = line
		}
		image.Layers = append(image.Layers, layer)
	}

	return
}

func MustDecode(w, h int, data string) Image {
	image, err := Decode(w, h, data)
	if err != nil {
		log.Fatal(err)
	}
	return image
}

func Check(image Image) (code int) {
	var minNum int
	var minLayer [][]int

	// Find the layer with the fewest 0's
	for i, layer := range image.Layers {
		num := countlayer(0, layer)

		if i == 0 || num == min(num, minNum) {
			minNum = num
			minLayer = layer
		}
	}

	return countlayer(1, minLayer) * countlayer(2, minLayer)
}

func countlayer(c int, layer [][]int) (total int) {
	for _, line := range layer {
		total += countline(c, line)
	}
	return
}

func countline(c int, line []int) (total int) {
	for _, n := range line {
		if n == c {
			total++
		}
	}
	return
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
