package tuple

import (
	"math"
	"testing"
)

func TestPoint(t *testing.T) {
	point := Tuple{3.3, 2.2, 1.1, 1.0}
	switch {
	case !Equal(point.X, 3.3):
		t.Errorf("Point Construction mistake")
	case !Equal(point.Y, 2.2):
		t.Errorf("Point Construction mistake")
	case !Equal(point.Z, 1.1):
		t.Errorf("Point Construction mistake")
	case !Equal(point.W, 1.0):
		t.Errorf("Point Construction mistake")
	case !IsPoint(point):
		t.Errorf("Not a point")
	case IsVector(point):
		t.Errorf("A vector is not a point")
	}
}

func TestVector(t *testing.T) {
	vector := Tuple{3.3, 2.2, 1.1, 0.0}
	switch {
	case !Equal(vector.X, 3.3):
		t.Errorf("Vector Construction mistake")
	case !Equal(vector.Y, 2.2):
		t.Errorf("Vector Construction mistake")
	case !Equal(vector.Z, 1.1):
		t.Errorf("Vector Construction mistake")
	case !Equal(vector.W, 0.0):
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
	case !Equal(point.W, 1.0):
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
	case !Equal(vector.W, 0.0):
		t.Errorf("Vector Construction mistake")
	case !IsVector(vector):
		t.Errorf("Not a vector")
	case IsPoint(vector):
		t.Errorf("A point is not a vector")
	}
}

func TestEqualTuple(t *testing.T) {
	vector := Vector(3.3, 2.2, 1.1)
	vector2 := Vector(3.3, 2.2, 1.1)
	if !EqualTuple(vector, vector2) {
		t.Errorf("Not equivelent vectors")
	}
}

func TestAddTuple(t *testing.T) {
	vector := Vector(3.3, 2.2, 1.1)
	vector1 := Vector(2.2, 1.1, 1.1)
	point := Point(1.0, 1.0, 1.0)
	addedVec := AddTuple(vector, vector1)
	addedPV := AddTuple(vector, point)
	switch {
	case !Equal(addedVec.X, 5.5):
		t.Errorf("Vector addition mistake")
	case !Equal(addedVec.Y, 3.3):
		t.Errorf("Vector addition mistake")
	case !Equal(addedVec.Z, 2.2):
		t.Errorf("Vector addition mistake")
	case !Equal(addedVec.W, 0.0):
		t.Errorf("Vector addition mistake")
	case !Equal(addedPV.X, 4.3):
		t.Errorf("Vector, Point addition mistake")
	case !Equal(addedPV.Y, 3.2):
		t.Errorf("Vector, Point addition mistake")
	case !Equal(addedPV.Z, 2.1):
		t.Errorf("Vector, Point addition mistake")
	case !Equal(addedPV.W, 1.0):
		t.Errorf("Vector, Point addition mistake")
	}
}

func TestSubTuple(t *testing.T) {
	vector := Vector(3.3, 2.2, 1.1)
	vector1 := Vector(2.2, 1.1, 1.1)
	point := Point(1.0, 1.0, 1.0)
	point1 := Point(1.0, 1.0, 1.0)
	subPoint := SubTuple(point, point1)
	subPV := SubTuple(point, vector)
	subVector := SubTuple(vector, vector1)
	switch {
	case !EqualTuple(subPoint, Vector(0, 0, 0)):
		t.Errorf("Error subtracting points")
	case !EqualTuple(subPV, Point(-2.3, -1.2, -0.1)):
		t.Errorf("Error subracting vector from point")
	case !EqualTuple(subVector, Vector(1.1, 1.1, 0.0)):
		t.Errorf("Error subtracting vectors")
	}
}

func TestNegateTuple(t *testing.T) {
	vector := NegateTuple(Tuple{1, -2, 3, -4})
	if !EqualTuple(vector, Tuple{-1, 2, -3, 4}) {
		t.Errorf("Negation failed")
	}
}

func TestMultTuple(t *testing.T) {
	tup := Tuple{1, -2, 3, -4}
	a := 3.5
	if !EqualTuple(MultiplyScaler(tup, a), Tuple{3.5, -7, 10.5, -14}) {
		t.Errorf("Multiplication failed")
	}
}

func TestDivTuple(t *testing.T) {
	tup := Tuple{1, -2, 3, -4}
	a := 2.0
	if !EqualTuple(DivideScaler(tup, a), Tuple{0.5, -1, 1.5, -2}) {
		t.Errorf("Division Failed")
	}
}

func TestMagnitude(t *testing.T) {
	vector := Vector(1, 0, 0)
	vector1 := Vector(0, 1, 0)
	vector2 := Vector(0, 0, 1)
	vector3 := Vector(1, 2, 3)
	vector4 := Vector(-1, -2, -3)

	switch {
	case !Equal(Magnitude(vector), 1.0):
		t.Errorf("Magnitude Calc failed")
	case !Equal(Magnitude(vector1), 1.0):
		t.Errorf("Magnitude Calc failed")
	case !Equal(Magnitude(vector2), 1.0):
		t.Errorf("Magnitude Calc failed")
	case !Equal(Magnitude(vector3), math.Sqrt(14.0)):
		t.Errorf("Magnitude Calc failed")
	case !Equal(Magnitude(vector4), math.Sqrt(14.0)):
		t.Errorf("Magnitude Calc failed")
	}
}

func TestNormalize(t *testing.T) {
	vector := Vector(4, 0, 0)
	vector1 := Vector(1, 2, 3)

	switch {
	case Equal(Magnitude(Normalize(vector1)), 1.0):
		t.Errorf("The magnitude of a normalized vector isn't 1")
	case EqualTuple(Normalize(vector), Vector(1, 0, 0)):
		t.Errorf("Normalization of the vector failed")
	}
}
