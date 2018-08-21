package LKmath

import (
	"sync"
	"math/rand"
	"fmt"
	"time"
)

type Matrix struct{
	Row 		int
	Column 		int
	Cell  		[][]float64
	mu			sync.RWMutex
}

const(
	INT  = 0
	FLOAT = 1
)


func NewEmptyMatrix(row int, column int)Matrix{

	if row == 0 || column == 0{
		panic("index error")

		return Matrix{}
	}
	var data [][]float64
	for i := 0; i < row; i++ {
		rowData := make([]float64, 0, row)
		for j := 0; j < column; j++ {
			rowData = append(rowData,0)
		}
		data = append(data,rowData)
	}

	matrix :=Matrix{
		Row:row,
		Column:column,
		Cell:data,
	}

	return matrix

}


func NewCopyMatrix(m Matrix)Matrix{
	var data [][]float64
	for i := 0; i < m.Row; i++ {
		rowData := make([]float64, 0, m.Row)
		for j := 0; j < m.Column; j++ {
			rowData = append(rowData,m.Cell[i][j])
		}
		data = append(data,rowData)
	}

	matrix :=Matrix{
		Row:m.Row,
		Column:m.Column,
		Cell:data,
	}
	return matrix
}



func MatrixRandom(m Matrix, max float64, min float64)Matrix{
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < m.Row; i++ {
		for j := 0; j < m.Column; j++ {
			m.Cell[i][j] =((max - min) *rand.Float64()) + min
		}
	}
	return m
}

func (ma *Matrix)MatrixSigmoid()Matrix{
	for i := 0; i < ma.Row; i++ {
		for j := 0; j < ma.Column; j++ {
			ma.Cell[i][j] =SigmoidFunction(ma.Cell[i][j])
		}
	}
	return *ma
}


func (ma *Matrix)Hprint(info string){
	fmt.Printf(info+"\n")
	for i := 0; i < ma.Row; i++ {
		s := ""
		for j := 0; j < ma.Column; j++ {
			s = s + fmt.Sprintf("%f ",ma.Cell[i][j])
		}
		fmt.Printf("%s\n", s)

	}

	fmt.Println()
}

func MatrixMultiplication(a Matrix, b Matrix)Matrix{
	if a.Column != b.Row {
		panic("a.column unequal to b.row, cannot perform multiplication")
		return Matrix{}
	}
	result:= NewEmptyMatrix(a.Row, b.Column)

	for i :=0; i < b.Column; i ++{
		for j := 0; j< a.Row; j++{
			cellSum := 0.0
			for k :=0; k< a.Column; k++{
				cellSum +=a.Cell[j][k] * b.Cell[k][i]
			}
			result.Cell[j][i] = cellSum
		}
	}

	return result


}

func TransposedMatrix(m Matrix)Matrix{
	var data [][]float64
	for i := 0; i < m.Column; i++ {
		rowData := make([]float64, 0, m.Column)
		for j := 0; j < m.Row; j++ {
			rowData = append(rowData,m.Cell[j][i])
		}
		data = append(data,rowData)
	}

	mT :=Matrix{
		Row:m.Column,
		Column:m.Row,
		Cell:data,
	}
	return mT
}

func InverseMatrix(m Matrix)Matrix{
	if m.Column != m.Row {
		panic("non-square matrix. cannot perform inverse")
		return Matrix{}
	}

	det :=Determinant(m)

	if det == 0{
		fmt.Printf("The determinant is 0, the matrix is not invertible\n")
		return MoorePenroseInverse(m)

	}


	return MatrixTimesRealNumber(AdjugateMatrix(m), 1/det)

}


func MoorePenroseInverse(m Matrix)Matrix{
	//todo: add Moore-Penrose Inverse in case the matrix is invertible
	panic("moore penrose inverse is not available right now")
	return Matrix{}
}






//cut matrix into a new matrix (include both begin value and end value)
func CutMatrix(m Matrix, rowBegin int, rowEnd int, columnBegin int, columnEnd int)Matrix{

	if rowBegin<0 || rowEnd >= m.Row || columnBegin<0 || columnEnd >= m.Column{
		panic("index out of range")
	}

	row := rowEnd - rowBegin +1
	column := columnEnd - columnBegin +1


	var data [][]float64
	for i := rowBegin; i <= rowEnd; i++ {
		rowData := make([]float64, 0, row)
		for j := columnBegin; j <= columnEnd; j++ {
			rowData = append(rowData,m.Cell[i][j])
		}
		data = append(data,rowData)
	}

	matrix :=Matrix{
		Row:row,
		Column:column,
		Cell:data,
	}
	return matrix
}



func RemoveRow(m Matrix, rowIndex int)Matrix{

	if rowIndex >= m.Row{
		panic("index out of range")
	}

	row := m.Row - 1
	column := m.Column


	var data [][]float64
	for i := 0; i < m.Row; i++ {
		if i == rowIndex{continue}
		rowData := make([]float64, 0, row)
		for j := 0; j < m.Column; j++ {
			rowData = append(rowData,m.Cell[i][j])
		}
		data = append(data,rowData)
	}

	matrix :=Matrix{
		Row:row,
		Column:column,
		Cell:data,
	}
	return matrix
}


func RemoveColumn(m Matrix, columnIndex int)Matrix{

	if columnIndex >= m.Column{
		panic("RemoveColumn index out of range")
	}

	row := m.Row
	column := m.Column - 1


	var data [][]float64
	for i := 0; i < m.Row; i++ {
		rowData := make([]float64, 0, row)
		for j := 0; j < m.Column; j++ {
			if j == columnIndex{continue}
			rowData = append(rowData,m.Cell[i][j])
		}
		data = append(data,rowData)
	}

	matrix :=Matrix{
		Row:row,
		Column:column,
		Cell:data,
	}
	return matrix
}


func RemoveRowAndColumn(m Matrix, rowIndex int, columnIndex int)Matrix{

	if rowIndex >= m.Row || columnIndex >= m.Column{
		panic("index out of range")
	}

	row := m.Row - 1
	column := m.Column - 1


	var data [][]float64
	for i := 0; i < m.Row; i++ {
		if i == rowIndex{continue}
		rowData := make([]float64, 0, row)
		for j := 0; j < m.Column; j++ {
			if j == columnIndex{continue}
			rowData = append(rowData,m.Cell[i][j])
		}
		data = append(data,rowData)
	}

	matrix :=Matrix{
		Row:row,
		Column:column,
		Cell:data,
	}
	return matrix
}


func Determinant(m Matrix)float64{

	if m.Row != m.Column{
		panic("non-square matrix. can't calculate determinant")
		return 0
	}

	if m.Row == 2{
		return m.Cell[0][0] * m.Cell[1][1] - m.Cell[0][1] * m.Cell[1][0]
	}else{
		var result float64
		for i := 0; i < m.Row; i ++{
			sign := 0.0
			if i ==0 || i % 2 == 0{
				sign = 1
			}else{
				sign = -1
			}

			subMatrixResult :=sign * m.Cell[0][i] * Determinant(RemoveRowAndColumn(m, 0, i))
			result = result + subMatrixResult


		}

		return result



	}



}


func MatrixTimesRealNumber(m Matrix, rn float64)Matrix{
	var data [][]float64
	for i := 0; i < m.Row; i++ {
		rowData := make([]float64, 0, m.Row)
		for j := 0; j < m.Column; j++ {
			rowData = append(rowData,m.Cell[i][j] * rn)
		}
		data = append(data,rowData)
	}

	matrix :=Matrix{
		Row:m.Row,
		Column:m.Column,
		Cell:data,
	}
	return matrix
}

func AdjugateMatrix(m Matrix)Matrix{
	var data [][]float64
	for i := 0; i < m.Row; i++ {
		rowData := make([]float64, 0, m.Row)
		for j := 0; j < m.Column; j++ {
			rowData = append(rowData, adjugatedCell(m, j, i))
		}
		data = append(data,rowData)
	}

	matrix :=Matrix{
		Row:m.Row,
		Column:m.Column,
		Cell:data,
	}
	return matrix
}

func adjugatedCell(m Matrix, rowIndex int, columnIndex int)float64{

	sign := 0.0
	if rowIndex + columnIndex == 0 || (rowIndex + columnIndex)%2 == 0{
		sign = 1
	}else{
		sign = -1
	}

	return sign * Determinant(RemoveRowAndColumn(m, rowIndex, columnIndex))
}