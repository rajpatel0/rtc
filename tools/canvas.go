package tools

import (
	"log"
	"os"
	"strconv"
)

type Canvas struct {
	Width  int
	Height int
	Pixels []Tuple
}

func MakeCanvas(width, height int) *Canvas {
	return ZeroPixels(&Canvas{Width: width, Height: height, Pixels: make([]Tuple, height*width)})
}

func ZeroPixels(canvas *Canvas) *Canvas {
	for i := range canvas.Pixels {
		canvas.Pixels[i] = Color(0, 0, 0)
	}
	return canvas
}

func WritePixel(canvas *Canvas, x, y int, color Tuple) *Canvas {
	if x < 0 || x > canvas.Width || y < 0 || y > canvas.Height {
		return canvas
	}
	idx := x + canvas.Width*y
	canvas.Pixels[idx] = color
	return canvas
}

func PixelAt(canvas *Canvas, x, y int) Tuple {
	idx := x + canvas.Width*y
	return canvas.Pixels[idx]
}

func CanvasToPPM(canvas *Canvas) string {
	maxLineWidth := 70
	maxColorValue := 255
	maxPixelWidth := 12 //if all three pixels are 3 digits will be 12 characters
	numChars := 0
	ppmString := "P3\n" + strconv.Itoa(canvas.Width) + " " + strconv.Itoa(canvas.Height) + "\n" + strconv.Itoa(maxColorValue)
	for i := range canvas.Pixels {
		if i%canvas.Width == 0 || (numChars+maxPixelWidth > maxLineWidth) {
			ppmString += "\n"
			numChars = 0
		}
		color := canvas.Pixels[i]
		r := clamp(color.X, maxColorValue)
		b := clamp(color.Y, maxColorValue)
		g := clamp(color.Z, maxColorValue)
		addedString := strconv.Itoa(r) + " " + strconv.Itoa(b) + " " + strconv.Itoa(g) + " "
		numChars += len(addedString)
		ppmString += addedString
	}
	return ppmString
}

func clamp(colorVal float64, maxColorValue int) int {
	colorVal = colorVal * 255
	if colorVal <= 0.0 {
		return 0
	}
	if colorVal >= 255 {
		return 255
	}
	return int(colorVal)
}

func PrintCanvas(ppmString string) {
	file, err := os.Create("testing.ppm")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	file.WriteString(ppmString + "\n")
}
