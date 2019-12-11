package sif

type Image struct {
	Width  int
	Height int
	Layers [][][]byte
}

func Decode(w, h int, data string) Image {
	image := Image{
		Width:  w,
		Height: h,
		Layers: make([][][]byte, 0),
	}

	for d := 0; d < len(data); {
		layer := make([][]byte, image.Height)
		for y := 0; y < image.Height; y++ {
			line := make([]byte, image.Width)
			for x := 0; x < image.Width; x++ {
				line[x] = data[d]
				d++
			}
			layer[y] = line
		}
		image.Layers = append(image.Layers, layer)
	}

	return image
}

func Check(image Image) (code int) {
	var minNum int
	var minLayer [][]byte

	// Find the layer with the fewest 0's
	for i, layer := range image.Layers {
		num := countlayer('0', layer)

		if i == 0 || num == min(num, minNum) {
			minNum = num
			minLayer = layer
		}
	}

	return countlayer('1', minLayer) * countlayer('2', minLayer)
}

func countlayer(b byte, layer [][]byte) (total int) {
	for _, line := range layer {
		total += countline(b, line)
	}
	return
}

func countline(b byte, line []byte) (total int) {
	for _, a := range line {
		if a == b {
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
