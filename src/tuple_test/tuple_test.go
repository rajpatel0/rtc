package tuple_test

import (
	"testing"

	tuple "../tuple"
)

func TestPoint(t *testing.T) {
	point := tuple.Tuple{3.3, 2.2, 1.1, 0.0}
	switch {
	case !tuple.Equal(point.X, 3.3):
		t.Errorf("Point Construction mistake")
	case !tuple.Equal(point.Y, 2.2):
		t.Errorf("Point Construction mistake")
	case !tuple.Equal(point.Z, 1.1):
		t.Errorf("Point Construction mistake")
	case !tuple.Equal(point.W, 0.0):
		t.Errorf("Point Construction mistake")
	case !tuple.IsPoint(point):
		t.Errorf("Not a point")
	case tuple.IsVector(point):
		t.Errorf("A vector is not a point")
	}

}

func TestVector(t *testing.T) {
	vector := tuple.Tuple{3.3, 2.2, 1.1, 1.0}
	switch {
	case !tuple.Equal(vector.X, 3.3):
		t.Errorf("Vector Construction mistake")
	case !tuple.Equal(vector.Y, 2.2):
		t.Errorf("Vector Construction mistake")
	case !tuple.Equal(vector.Z, 1.1):
		t.Errorf("Vector Construction mistake")
	case !tuple.Equal(vector.W, 1.0):
		t.Errorf("Vector Construction mistake")
	case !tuple.IsVector(vector):
		t.Errorf("Not a vector")
	case tuple.IsPoint(vector):
		t.Errorf("A point is not a vector")
	}

}
