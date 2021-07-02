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
