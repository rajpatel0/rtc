package tools

import "testing"

func TestMatrixFourByFour(t *testing.T) {
	mat := MakeMatrix(4, 4)
	data := [16]float64{1.0, 2.0, 3.0, 4.0, 5.5, 6.5, 7.5, 8.5, 9.0, 10.0, 11.0, 12.0, 13.5, 14.5, 15.5, 16.5}
	mat = SetElements(mat, data[:])
	switch {
	case !Equal(GetElem(mat, 0, 0), 1.0):
		t.Errorf("Expected 1 got %f", GetElem(mat, 0, 0))
	case !Equal(GetElem(mat, 0, 3), 4.0):
		t.Errorf("Expected 4 got %f", GetElem(mat, 0, 3))
	case !Equal(GetElem(mat, 1, 0), 5.5):
		t.Errorf("Expected 5.5 got %f", GetElem(mat, 1, 0))
	case !Equal(GetElem(mat, 1, 2), 7.5):
		t.Errorf("Expected 7.5 got %f", GetElem(mat, 1, 2))
	case !Equal(GetElem(mat, 2, 2), 11.0):
		t.Errorf("Expected 11 got %f", GetElem(mat, 2, 2))
	case !Equal(GetElem(mat, 3, 0), 13.5):
		t.Errorf("Expected 13.5 got %f", GetElem(mat, 3, 0))
	case !Equal(GetElem(mat, 3, 2), 15.5):
		t.Errorf("Expected 15.5 got %f", GetElem(mat, 3, 2))

	}
}

func TestMatrixTwoByTwo(t *testing.T) {
	mat := MakeMatrix(2, 2)
	data := [4]float64{-3.0, 5.0, 1.0, -2.0}
	mat = SetElements(mat, data[:])
	switch {
	case !Equal(GetElem(mat, 0, 0), -3.0):
		t.Errorf("Expected -3.0 got %f", GetElem(mat, 0, 0))
	case !Equal(GetElem(mat, 0, 1), 5.0):
		t.Errorf("Expected 5.0 got %f", GetElem(mat, 0, 1))
	case !Equal(GetElem(mat, 1, 0), 1.0):
		t.Errorf("Expected 1 got %f", GetElem(mat, 1, 0))
	case !Equal(GetElem(mat, 1, 1), -2.0):
		t.Errorf("Expected -2.0 got %f", GetElem(mat, 1, 1))
	}
}

func TestMatrixEquality(t *testing.T) {
	mat := MakeMatrix(4, 4)
	mat1 := MakeMatrix(4, 4)
	data := [16]float64{1.0, 2.0, 3.0, 4.0, 5.5, 6.5, 7.5, 8.5, 9.0, 10.0, 11.0, 12.0, 13.5, 14.5, 15.5, 16.5}
	mat = SetElements(mat, data[:])
	mat1 = SetElements(mat1, data[:])
	if !MatrixEqual(mat, mat1) {
		t.Errorf("Matrices equal but failed equality test")
	}
}
func TestMatrixInequality(t *testing.T) {
	mat := MakeMatrix(4, 4)
	mat1 := MakeMatrix(4, 4)
	data := [16]float64{1.0, 2.0, 3.0, 4.0, 5.5, 6.5, 7.5, 8.5, 9.0, 10.0, 11.0, 12.0, 13.5, 14.5, 15.5, 16.5}
	data1 := [16]float64{3.0, 2.0, 3.0, 4.0, 5.5, 6.5, 7.5, 8.5, 9.0, 10.0, 11.0, 12.0, 13.5, 14.5, 15.5, 16.5}
	mat = SetElements(mat, data[:])
	mat1 = SetElements(mat1, data1[:])
	if MatrixEqual(mat, mat1) {
		t.Errorf("Matrices not equal but passed equality test")
	}
}
