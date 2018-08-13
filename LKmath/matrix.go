package LKmath

import (
	"sync"
	"math/rand"
	"fmt"
	"time"
	"github.com/pkg/errors"
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


func NewEmptyMatrix(row int, column int)(Matrix, error){

	if row == 0 || column == 0{
		return Matrix{}, errors.New("index error")
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

	return matrix, nil

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
			ma.Cell[i][j] =Sigmoid(ma.Cell[i][j])
		}
	}
	return *ma
}


func (ma *Matrix)Hprint(){
	for i := 0; i < ma.Row; i++ {
		s := ""
		for j := 0; j < ma.Column; j++ {
			s = s + fmt.Sprintf("%f ",ma.Cell[i][j])
		}
		fmt.Printf("%s\n", s)

	}

	fmt.Println()
}

func MatrixMultiplication(a Matrix, b Matrix)(Matrix, error){
	if a.Column != b.Row {
		return Matrix{}, errors.New("a.column unequal to b.row, cannot perform multiplication")
	}
	result,_ := NewEmptyMatrix(a.Row, b.Column)

	for i :=0; i < b.Column; i ++{
		for j := 0; j< a.Row; j++{
			cellSum := 0.0
			for k :=0; k< a.Column; k++{
				cellSum +=a.Cell[j][k] * b.Cell[k][i]
			}
			result.Cell[j][i] = cellSum
		}
	}

	return result, nil


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

	matrix :=Matrix{
		Row:m.Column,
		Column:m.Row,
		Cell:data,
	}
	return matrix
}


