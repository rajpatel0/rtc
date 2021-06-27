package tools

type Canvas struct {
	width  int
	height int
	pixels []Tuple
}

func MakeCanvas(width, height int) *Canvas {
	return ZeroPixels(&Canvas{width: width, height: height, pixels: make([]Tuple, height*width)})
}

func ZeroPixels(canvas *Canvas) *Canvas {
	for i := range canvas.pixels {
		canvas.pixels[i] = Color(0, 0, 0)
	}
	return canvas
}
