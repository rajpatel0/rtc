package tools

import "testing"

func TestCanvas(t *testing.T) {

	c := MakeCanvas(10, 20)
	emptyColor := Color(0, 0, 0)
	switch {
	case c.Width != 10:
		t.Errorf("Failed to set width of canvas")
	case c.Height != 20:
		t.Errorf("Failed to set height of canvas")
	}
	for i := range c.Pixels {
		if !EqualTuple(c.Pixels[i], emptyColor) {
			t.Errorf("Not all pixels are zeroed")
		}
	}
}

func TestWritingCanvas(t *testing.T) {

	c := MakeCanvas(10, 20)
	red := Color(1, 0, 0)
	c = WritePixel(c, 2, 3, red)
	if PixelAt(c, 2, 3) != red {
		t.Errorf("Error in setting pixel")
	}
}
