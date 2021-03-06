package LKmath

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
	"time"
)

// this packet includes functions for matrix creation and calculation

type Matrix struct {
	Row    int
	Column int
	Cell   [][]float64
	mu     sync.RWMutex
}

/*-----------------------------------------functions for creating matrix----------------------------------------------*/
//New Empty Matrix will create a Matrix filled with 0
func NewEmptyMatrix(row int, column int) Matrix {

	if row == 0 || column == 0 {
		panic("index error")
	}

	var data [][]float64
	for i := 0; i < row; i++ {
		rowData := make([]float64, 0, row)
		for j := 0; j < column; j++ {
			rowData = append(rowData, 0)
		}
		data = append(data, rowData)
	}

	matrix := Matrix{
		Row:    row,
		Column: column,
		Cell:   data,
	}

	return matrix

}

//New Valued Matrix will create a Matrix filled with a certain value
func NewValuedMatrix(row int, column int, value float64) Matrix {

	if row == 0 || column == 0 {
		panic("index error")

	}
	var data [][]float64
	for i := 0; i < row; i++ {
		rowData := make([]float64, 0, row)
		for j := 0; j < column; j++ {
			rowData = append(rowData, value)
		}
		data = append(data, rowData)
	}

	matrix := Matrix{
		Row:    row,
		Column: column,
		Cell:   data,
	}

	return matrix

}

//New Identity Matrix will create a size by size Identity Matrix
func NewIdentityMatrix(size int) Matrix {
	var data [][]float64
	for i := 0; i < size; i++ {
		rowData := make([]float64, 0, size)
		for j := 0; j < size; j++ {
			if i == j {
				rowData = append(rowData, 1)
			} else {
				rowData = append(rowData, 0)
			}
		}
		data = append(data, rowData)
	}

	matrix := Matrix{
		Row:    size,
		Column: size,
		Cell:   data,
	}
	return matrix
}

//New Copy Matrix will copy a Matrix and create a new one
func NewCopyMatrix(m Matrix) Matrix {
	var data [][]float64
	for i := 0; i < m.Row; i++ {
		rowData := make([]float64, 0, m.Row)
		for j := 0; j < m.Column; j++ {
			rowData = append(rowData, m.Cell[i][j])
		}
		data = append(data, rowData)
	}

	matrix := Matrix{
		Row:    m.Row,
		Column: m.Column,
		Cell:   data,
	}
	return matrix
}

//New Random Matrix will create a random matrix
//if initialize is false, the matrix will always be the same
func NewRandomMatrix(row int, column int, min float64, max float64) Matrix {


	if row == 0 || column == 0 {
		panic("index error")

	}
	var data [][]float64
	for i := 0; i < row; i++ {
		rowData := make([]float64, 0, row)
		for j := 0; j < column; j++ {
			rowData = append(rowData, ((max-min)*rand.Float64())+min)
		}
		data = append(data, rowData)
	}

	matrix := Matrix{
		Row:    row,
		Column: column,
		Cell:   data,
	}

	return matrix

}

/*---------------------------------------functions for recreating matrix----------------------------------------------*/

//Matrix Random will make an existing matrix become a random matrix
func RandomMatrix(m Matrix, min float64, max float64) Matrix {
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < m.Row; i++ {
		for j := 0; j < m.Column; j++ {
			m.Cell[i][j] = ((max - min) * rand.Float64()) + min
		}
	}
	return m
}

//return the abstract value of the matrix
func AbsMatrix(m Matrix) Matrix {
	var data [][]float64
	for i := 0; i < m.Row; i++ {
		rowData := make([]float64, 0, m.Row)
		for j := 0; j < m.Column; j++ {
			rowData = append(rowData, math.Abs(m.Cell[i][j]))
		}
		data = append(data, rowData)
	}

	matrix := Matrix{
		Row:    m.Row,
		Column: m.Column,
		Cell:   data,
	}
	return matrix
}

//update the matrix
func (this *Matrix) Update(newMatrix Matrix) {
	if this.Row != newMatrix.Row || this.Column != newMatrix.Column {
		panic("format will be changed\n")
		this.Column = newMatrix.Column
		this.Row = newMatrix.Row
	}

	var data [][]float64
	for i := 0; i < newMatrix.Row; i++ {
		rowData := make([]float64, 0, newMatrix.Row)
		for j := 0; j < newMatrix.Column; j++ {
			rowData = append(rowData, newMatrix.Cell[i][j])
		}
		data = append(data, rowData)
	}

	this.Cell = data

}

/*---------------------------------------functions for reshaping matrix----------------------------------------------*/

//cut matrix into a new matrix (include both begin value and end value)
func CutMatrix(m Matrix, rowBegin int, rowEnd int, columnBegin int, columnEnd int) Matrix {

	if rowBegin < 0 || rowEnd >= m.Row || columnBegin < 0 || columnEnd >= m.Column {
		panic("index out of range")
	}

	row := rowEnd - rowBegin + 1
	column := columnEnd - columnBegin + 1

	var data [][]float64
	for i := rowBegin; i <= rowEnd; i++ {
		rowData := make([]float64, 0, row)
		for j := columnBegin; j <= columnEnd; j++ {
			rowData = append(rowData, m.Cell[i][j])
		}
		data = append(data, rowData)
	}

	matrix := Matrix{
		Row:    row,
		Column: column,
		Cell:   data,
	}
	return matrix
}

//remove a row
func RemoveRow(m Matrix, rowIndex int) Matrix {

	if rowIndex >= m.Row {
		panic("index out of range")
	}

	row := m.Row - 1
	column := m.Column

	var data [][]float64
	for i := 0; i < m.Row; i++ {
		if i == rowIndex {
			continue
		}
		rowData := make([]float64, 0, row)
		for j := 0; j < m.Column; j++ {
			rowData = append(rowData, m.Cell[i][j])
		}
		data = append(data, rowData)
	}

	matrix := Matrix{
		Row:    row,
		Column: column,
		Cell:   data,
	}
	return matrix
}

//remove a column
func RemoveColumn(m Matrix, columnIndex int) Matrix {

	if columnIndex >= m.Column {
		panic("RemoveColumn index out of range")
	}

	row := m.Row
	column := m.Column - 1

	var data [][]float64
	for i := 0; i < m.Row; i++ {
		rowData := make([]float64, 0, row)
		for j := 0; j < m.Column; j++ {
			if j == columnIndex {
				continue
			}
			rowData = append(rowData, m.Cell[i][j])
		}
		data = append(data, rowData)
	}

	matrix := Matrix{
		Row:    row,
		Column: column,
		Cell:   data,
	}
	return matrix
}

//remove row and column at the same time
func RemoveRowAndColumn(m Matrix, rowIndex int, columnIndex int) Matrix {

	if rowIndex >= m.Row || columnIndex >= m.Column {
		panic("index out of range")
	}

	row := m.Row - 1
	column := m.Column - 1

	var data [][]float64
	for i := 0; i < m.Row; i++ {
		if i == rowIndex {
			continue
		}
		rowData := make([]float64, 0, row)
		for j := 0; j < m.Column; j++ {
			if j == columnIndex {
				continue
			}
			rowData = append(rowData, m.Cell[i][j])
		}
		data = append(data, rowData)
	}

	matrix := Matrix{
		Row:    row,
		Column: column,
		Cell:   data,
	}
	return matrix
}

func KeepOneRow(m Matrix, rowIndex int) Matrix {

	if rowIndex >= m.Row {
		panic("index out of range")
	}

	row := 1
	column := m.Column

	var data [][]float64
	for i := 0; i < m.Row; i++ {
		if i != rowIndex {
			continue
		}
		rowData := make([]float64, 0, row)
		for j := 0; j < m.Column; j++ {
			rowData = append(rowData, m.Cell[i][j])
		}
		data = append(data, rowData)
	}

	matrix := Matrix{
		Row:    row,
		Column: column,
		Cell:   data,
	}
	return matrix
}

/*---------------------------------------functions for printing matrix----------------------------------------------*/

// print matrix in a human-readable way together with an info you want to print
func (ma *Matrix) Hprint(info string) {
	fmt.Printf(info + "\n")
	for i := 0; i < ma.Row; i++ {
		s := ""
		for j := 0; j < ma.Column; j++ {
			s = s + fmt.Sprintf("%f\t", ma.Cell[i][j])
		}
		fmt.Printf("%s\n", s)

	}

	fmt.Println()
}

//print matrix in detail for better testing
func (ma *Matrix) Dprint(info string) {
	fmt.Printf(info + "\n")
	for i := 0; i < ma.Row; i++ {
		s := ""
		for j := 0; j < ma.Column; j++ {
			s = s + fmt.Sprintf("%v\t", ma.Cell[i][j])
		}
		fmt.Printf("%s\n", s)

	}

	fmt.Println()
}

/*---------------------------------------functions for calculating matrix----------------------------------------------*/

//multiply Matrix a and Matrix b
func MatrixMultiplication(a Matrix, b Matrix) Matrix {
	if a.Column != b.Row {
		fmt.Printf("a.columu is %d, b.row is %d\n", a.Column, b.Row)
		panic("a.column unequal to b.row, cannot perform multiplication")
	}
	result := NewEmptyMatrix(a.Row, b.Column)

	for i := 0; i < b.Column; i ++ {
		for j := 0; j < a.Row; j++ {
			cellSum := 0.0
			for k := 0; k < a.Column; k++ {
				cellSum += a.Cell[j][k] * b.Cell[k][i]
			}
			result.Cell[j][i] = cellSum
		}
	}

	return result

}

func DotProduct(a Matrix, b Matrix) Matrix {
	if a.Row != b.Row || a.Column != b.Column {
		panic("DotProduct: format error")
	}
	result := NewEmptyMatrix(a.Row, a.Column)
	for i := 0; i < a.Row; i++ {
		for j := 0; j < a.Column; j++ {
			result.Cell[i][j] = a.Cell[i][j] * b.Cell[i][j]
			if math.IsNaN(result.Cell[i][j]){
				result.Cell[i][j]=0
			}
		}
	}

	return result
}

//return the transposed matrix of m
func TransposeMatrix(m Matrix) Matrix {
	var data [][]float64
	for i := 0; i < m.Column; i++ {
		rowData := make([]float64, 0, m.Column)
		for j := 0; j < m.Row; j++ {
			rowData = append(rowData, m.Cell[j][i])
		}
		data = append(data, rowData)
	}

	mT := Matrix{
		Row:    m.Column,
		Column: m.Row,
		Cell:   data,
	}
	return mT
}

//return the inverse matrix of m
func InverseMatrix(m Matrix) Matrix {
	if m.Column != m.Row {
		panic("non-square matrix. cannot perform inverse")
	}

	det := Determinant(m)

	if det == 0 {
		fmt.Printf("The determinant is 0, the matrix is not invertible\n")
		return MoorePenroseInverse(m)

	}

	return ScalarMatrix(AdjugateMatrix(m), 1/det)

}

//return the Moore-Penrose inverse matrix of m
func MoorePenroseInverse(m Matrix) Matrix {
	//todo: add Moore-Penrose Inverse in case the matrix is invertible
	panic("moore penrose inverse is not available right now")
}

//calculate the matrix times a real number
func ScalarMatrix(m Matrix, rn float64) Matrix {
	var data [][]float64
	for i := 0; i < m.Row; i++ {
		rowData := make([]float64, 0, m.Row)
		for j := 0; j < m.Column; j++ {
			rowData = append(rowData, m.Cell[i][j]*rn)
		}
		data = append(data, rowData)
	}

	matrix := Matrix{
		Row:    m.Row,
		Column: m.Column,
		Cell:   data,
	}
	return matrix
}

//calculate the adjugate of a matrix
func AdjugateMatrix(m Matrix) Matrix {
	var data [][]float64
	for i := 0; i < m.Row; i++ {
		rowData := make([]float64, 0, m.Row)
		for j := 0; j < m.Column; j++ {
			rowData = append(rowData, adjugatedCell(m, j, i))
		}
		data = append(data, rowData)
	}

	matrix := Matrix{
		Row:    m.Row,
		Column: m.Column,
		Cell:   data,
	}
	return matrix
}

func adjugatedCell(m Matrix, rowIndex int, columnIndex int) float64 {

	sign := 0.0
	if rowIndex+columnIndex == 0 || (rowIndex+columnIndex)%2 == 0 {
		sign = 1
	} else {
		sign = -1
	}

	return sign * Determinant(RemoveRowAndColumn(m, rowIndex, columnIndex))
}

//add two matrices
func MatrixAddition(a Matrix, b Matrix) Matrix {
	if a.Row != b.Row || a.Column != b.Column {
		a.Hprint("a")
		b.Hprint("b:")
		panic("matrix doesn't fit each other, can't perform adding")
	}
	var data [][]float64
	for i := 0; i < a.Row; i++ {
		rowData := make([]float64, 0, a.Row)
		for j := 0; j < a.Column; j++ {
			rowData = append(rowData, a.Cell[i][j]+b.Cell[i][j])
		}
		data = append(data, rowData)
	}

	matrix := Matrix{
		Row:    a.Row,
		Column: a.Column,
		Cell:   data,
	}
	return matrix
}

//subtract two matrices
func MatrixSubtraction(a Matrix, b Matrix) Matrix {
	if a.Row != b.Row || a.Column != b.Column {
		panic("matrix doesn't fit each other, can't perform subtraction")
	}
	var data [][]float64
	for i := 0; i < a.Row; i++ {
		rowData := make([]float64, 0, a.Row)
		for j := 0; j < a.Column; j++ {
			rowData = append(rowData, a.Cell[i][j]-b.Cell[i][j])
		}
		data = append(data, rowData)
	}

	matrix := Matrix{
		Row:    a.Row,
		Column: a.Column,
		Cell:   data,
	}
	return matrix
}

/*-----------------------------functions for calculate the attribution of matrix--------------------------------------*/

//calculate the Determinant of a matrix
func Determinant(m Matrix) float64 {

	if m.Row != m.Column {
		panic("non-square matrix. can't calculate determinant")
		return 0
	}

	if m.Row == 2 {
		return m.Cell[0][0]*m.Cell[1][1] - m.Cell[0][1]*m.Cell[1][0]
	} else {
		var result float64
		for i := 0; i < m.Row; i ++ {
			sign := 0.0
			if i == 0 || i%2 == 0 {
				sign = 1
			} else {
				sign = -1
			}

			subMatrixResult := sign * m.Cell[0][i] * Determinant(RemoveRowAndColumn(m, 0, i))
			result = result + subMatrixResult

		}

		return result

	}

}

//return the sum of the whole matrix
func Sum(m Matrix) float64 {
	sum := 0.0
	for i := 0; i < m.Row; i++ {
		for j := 0; j < m.Column; j++ {
			sum += m.Cell[i][j]
		}
	}
	return sum
}

//return the average value of the whole matrix
func Average(m Matrix) float64 {
	return Sum(m) / float64(m.Row*m.Column)
}

/*----------------------------------------------------squeeze matrix--------------------------------------------------*/

//return the sum of each column
func SqueezedSumColumnMatrix(m Matrix) Matrix {

	var data [][]float64
	rowData := make([]float64, 0, m.Row)
	for i := 0; i < m.Column; i ++ {
		cellResult := 0.0
		for j := 0; j < m.Row; j++ {
			cellResult += m.Cell[j][i]
		}
		rowData = append(rowData, cellResult)
	}
	data = append(data, rowData)

	matrix := Matrix{
		Row:    1,
		Column: m.Column,
		Cell:   data,
	}

	return matrix
}

//return the average of each column
func SqueezedAverageColumnMatrix(m Matrix) Matrix {

	return ScalarMatrix(SqueezedSumColumnMatrix(m), 1/float64(m.Row))
}

//return the sum of each row
func SqueezedSumRowMatrix(m Matrix) Matrix {

	var data [][]float64
	for i := 0; i < m.Row; i++ {
		rowData := make([]float64, 0, m.Row)
		cellResult := 0.0
		for j := 0; j < m.Column; j++ {
			cellResult += m.Cell[i][j]
		}

		rowData = append(rowData, cellResult)
		data = append(data, rowData)

	}

	matrix := Matrix{
		Row:    m.Row,
		Column: 1,
		Cell:   data,
	}

	return matrix
}

//return the average of each row
func SqueezedAverageRowMatrix(m Matrix) Matrix {

	temp := SqueezedSumRowMatrix(m)
	return ScalarMatrix(temp, 1/float64(m.Column))
}

//return the maximum value of each column
func SqueezedMaxColumnMatrix(m Matrix) Matrix {

	var data [][]float64
	rowData := make([]float64, 0, m.Row)
	for i := 0; i < m.Column; i ++ {
		Max := m.Cell[0][i]
		for j := 1; j < m.Row; j++ {
			if m.Cell[j][i] > Max {
				Max = m.Cell[j][i]
			}
		}
		rowData = append(rowData, Max)
	}
	data = append(data, rowData)

	matrix := Matrix{
		Row:    1,
		Column: m.Column,
		Cell:   data,
	}

	return matrix
}

//return the minimum value of each column
func SqueezedMinColumnMatrix(m Matrix) Matrix {

	var data [][]float64
	rowData := make([]float64, 0, m.Row)
	for i := 0; i < m.Column; i ++ {
		Min := m.Cell[0][i]
		for j := 1; j < m.Row; j++ {
			if m.Cell[j][i] < Min {
				Min = m.Cell[j][i]
			}
		}
		rowData = append(rowData, Min)
	}
	data = append(data, rowData)

	matrix := Matrix{
		Row:    1,
		Column: m.Column,
		Cell:   data,
	}

	return matrix
}
