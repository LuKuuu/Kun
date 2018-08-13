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
	Data  		[][]float64
	mu			sync.RWMutex
}

const(
	INT  = 0
	FLOAT = 1
)


func NewEmptyMatrix(row int, column int)Matrix{
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
		Data:data,
	}

	return matrix

}


func NewCopyMatrix(m Matrix)Matrix{
	var data [][]float64
	for i := 0; i < m.Row; i++ {
		rowData := make([]float64, 0, m.Row)
		for j := 0; j < m.Column; j++ {
			rowData = append(rowData,0)
		}
		data = append(data,rowData)
	}

	matrix :=Matrix{
		Row:m.Row,
		Column:m.Column,
		Data:data,
	}
	return matrix
}



func MatrixRandom(m Matrix, max float64, min float64)Matrix{
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < m.Row; i++ {
		for j := 0; j < m.Column; j++ {
			m.Data[i][j] =((max - min) *rand.Float64()) + min
		}
	}
	return m
}

func (ma *Matrix)MatrixSigmoid()Matrix{
	for i := 0; i < ma.Row; i++ {
		for j := 0; j < ma.Column; j++ {
			ma.Data[i][j] =Sigmoid(ma.Data[i][j])
		}
	}
	return *ma
}


func (ma *Matrix)Hprint(){
	for i := 0; i < ma.Row; i++ {
		s := ""
		for j := 0; j < ma.Column; j++ {
			s = s + fmt.Sprintf("%f ",ma.Data[i][j])
		}
		fmt.Printf("%s\n", s)

	}

	fmt.Println()
}

func matrixMultipulate(){}



