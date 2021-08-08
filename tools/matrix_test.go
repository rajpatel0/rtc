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

func TestMatrixMultiply(t *testing.T) {
	mat := MakeMatrix(4, 4)
	mat1 := MakeMatrix(4, 4)
	matResult := MakeMatrix(4, 4)
	data := [16]float64{1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 8.0, 7.0, 6.0, 5.0, 4.0, 3.0, 2.0}
	data1 := [16]float64{-2.0, 1.0, 2.0, 3.0, 3.0, 2.0, 1.0, -1.0, 4.0, 3.0, 6.0, 5.0, 1.0, 2.0, 7.0, 8.0}
	dataResult := [16]float64{20.0, 22.0, 50.0, 48.0, 44.0, 54.0, 114.0, 108.0, 40.0, 58.0, 110.0, 102.0, 16.0, 26.0, 46.0, 42.0}
	mat = SetElements(mat, data[:])
	mat1 = SetElements(mat1, data1[:])
	matResult = SetElements(matResult, dataResult[:])
	if !MatrixEqual(matResult, MatrixMultiply(mat, mat1)) {
		t.Errorf("Matrix Multiply Failed")
	}

}

func TestMatrixTupleMultiply(t *testing.T) {
	mat := MakeMatrix(4, 4)
	data := [16]float64{1.0, 2.0, 3.0, 4.0, 2.0, 4.0, 4.0, 2.0, 8.0, 6.0, 4.0, 1.0, 0.0, 0.0, 0.0, 1.0}
	mat = SetElements(mat, data[:])
	tuple := Point(1, 2, 3)

	outPutTuple := Point(18, 24, 33)
	if !EqualTuple(outPutTuple, MatrixTupleMultiply(mat, tuple)) {
		t.Errorf("Matrix Tuple Multiply Failed")
	}
}

func TestMatrixIdentityMultiply(t *testing.T) {
	mat := MakeMatrix(4, 4)
	data := [16]float64{1.0, 2.0, 3.0, 4.0, 2.0, 4.0, 4.0, 2.0, 8.0, 6.0, 4.0, 1.0, 0.0, 0.0, 0.0, 1.0}
	mat = SetElements(mat, data[:])
	idMat := MakeIdentity(4, 4)
	if !MatrixEqual(mat, MatrixMultiply(mat, idMat)) {
		t.Errorf("Identity matrix multiplication failed")
	}
}

func TestMatrixTranspose(t *testing.T) {
	mat := MakeMatrix(4, 4)
	transposeMat := MakeMatrix(4, 4)
	data := [16]float64{0.0, 9.0, 3.0, 0.0, 9.0, 8.0, 0.0, 8.0, 1.0, 8.0, 5.0, 3.0, 0.0, 0.0, 5.0, 8.0}
	transposeData := [16]float64{0.0, 9.0, 1.0, 0.0, 9.0, 8.0, 8.0, 0.0, 3.0, 0.0, 5.0, 5.0, 0.0, 8.0, 3.0, 8.0}
	mat = SetElements(mat, data[:])
	transposeMat = SetElements(transposeMat, transposeData[:])
	idMat := MakeIdentity(4, 4)
	switch {
	case !MatrixEqual(transposeMat, InvertMatrix(mat)):
		t.Errorf("Failed inversion of matrix")
	case !MatrixEqual(idMat, InvertMatrix(idMat)):
		t.Errorf("Failed inversion of identity")
	}
}
