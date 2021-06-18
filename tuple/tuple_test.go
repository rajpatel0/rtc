package tuple

import (
	"testing"
)

func TestPoint(t *testing.T) {
	point := Tuple{3.3, 2.2, 1.1, 0.0}
	switch {
	case !Equal(point.X, 3.3):
		t.Errorf("Point Construction mistake")
	case !Equal(point.Y, 2.2):
		t.Errorf("Point Construction mistake")
	case !Equal(point.Z, 1.1):
		t.Errorf("Point Construction mistake")
	case !Equal(point.W, 0.0):
		t.Errorf("Point Construction mistake")
	case !IsPoint(point):
		t.Errorf("Not a point")
	case IsVector(point):
		t.Errorf("A vector is not a point")
	}
}

func TestVector(t *testing.T) {
	vector := Tuple{3.3, 2.2, 1.1, 1.0}
	switch {
	case !Equal(vector.X, 3.3):
		t.Errorf("Vector Construction mistake")
	case !Equal(vector.Y, 2.2):
		t.Errorf("Vector Construction mistake")
	case !Equal(vector.Z, 1.1):
		t.Errorf("Vector Construction mistake")
	case !Equal(vector.W, 1.0):
		t.Errorf("Vector Construction mistake")
	case !IsVector(vector):
		t.Errorf("Not a vector")
	case IsPoint(vector):
		t.Errorf("A point is not a vector")
	}
}

func TestPointCreator(t *testing.T) {
	point := Point(3.3, 2.2, 1.1)
	switch {
	case !Equal(point.X, 3.3):
		t.Errorf("Point Construction mistake")
	case !Equal(point.Y, 2.2):
		t.Errorf("Point Construction mistake")
	case !Equal(point.Z, 1.1):
		t.Errorf("Point Construction mistake")
	case !Equal(point.W, 0.0):
		t.Errorf("Point Construction mistake")
	case !IsPoint(point):
		t.Errorf("Not a point")
	case IsVector(point):
		t.Errorf("A vector is not a point")
	}
}

func TestVectorCreator(t *testing.T) {
	vector := Vector(3.3, 2.2, 1.1)
	switch {
	case !Equal(vector.X, 3.3):
		t.Errorf("Vector Construction mistake")
	case !Equal(vector.Y, 2.2):
		t.Errorf("Vector Construction mistake")
	case !Equal(vector.Z, 1.1):
		t.Errorf("Vector Construction mistake")
	case !Equal(vector.W, 1.0):
		t.Errorf("Vector Construction mistake")
	case !IsVector(vector):
		t.Errorf("Not a vector")
	case IsPoint(vector):
		t.Errorf("A point is not a vector")
	}
}
