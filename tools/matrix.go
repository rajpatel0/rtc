package tools

type Matrix struct {
	Row    int
	Column int
	Data   []float64
}

func MakeMatrix(row, column int) *Matrix {
	return &Matrix{Row: row, Column: column, Data: make([]float64, row*column)}
}

func SetElements(mat *Matrix, elems []float64) *Matrix {
	rows := mat.Row
	columns := mat.Column
	elements := rows * columns
	for i := 0; i < elements; i++ {
		mat.Data[i] = elems[i]
	}
	return mat
}

func SetElement(mat *Matrix, row, column int, elem float64) *Matrix {
	if row < mat.Row && column < mat.Column {
		mat.Data[row*mat.Column+column] = elem
	}
	return mat
}

func GetElem(mat *Matrix, row, column int) float64 {
	return mat.Data[row*mat.Column+column]
}

func MatrixEqual(mat1, mat2 *Matrix) bool {
	if mat1.Column != mat2.Column || mat1.Row != mat2.Row {
		return false
	}
	elems := mat1.Column * mat1.Row
	for i := 0; i < elems; i++ {
		if !Equal(mat1.Data[i], mat2.Data[i]) {
			return false
		}
	}
	return true

}

func MatrixMultiply(mat1, mat2 *Matrix) *Matrix {
	//if mat1.Column != mat2.Row{
	//	return errors
	//} For when I add error checking
	retMat := MakeMatrix(mat1.Row, mat2.Column)
	for i := 0; i < mat1.Row; i++ {
		for j := 0; j < mat2.Column; j++ {
			sum := 0.0
			for k := 0; k < mat1.Column; k++ {
				sum += mat1.Data[i*mat1.Column+k] * mat2.Data[j+k*mat2.Column]
			}
			retMat.Data[i*retMat.Column+j] = sum
		}
	}
	return retMat
}

func MatrixTupleMultiply(mat *Matrix, tup Tuple) Tuple {
	//Note output can only be of size 4
	//if mat.Column != 4{
	//	return errors
	//} For when I add error checking
	tupMat := MakeMatrix(4, 1)
	data := [4]float64{tup.X, tup.Y, tup.Z, tup.W}
	tupMat = SetElements(tupMat, data[:])
	outPutTuple := MatrixMultiply(mat, tupMat)
	return Tuple{X: outPutTuple.Data[0], Y: outPutTuple.Data[1], Z: outPutTuple.Data[2], W: outPutTuple.Data[3]}
}
