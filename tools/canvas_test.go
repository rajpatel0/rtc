package tools

import "testing"

func TestCanvas(t *testing.T) {

	c := MakeCanvas(10, 20)
	emptyColor := Color(0, 0, 0)
	switch {
	case c.width != 10:
		t.Errorf("Failed to set width of canvas")
	case c.height != 20:
		t.Errorf("Failed to set height of canvas")
	}
	for i := range c.pixels {
		if !EqualTuple(c.pixels[i], emptyColor) {
			t.Errorf("Not all pixels are zeroed")
		}
	}
}
