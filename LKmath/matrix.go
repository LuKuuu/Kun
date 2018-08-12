package LKmath

import (
	"sync"
	"math/rand"
	"fmt"
)

type Matrix struct{
	Row 		int
	Column 		int
	Data  		[][]float32
	mu			sync.RWMutex
}

const(
	INT  = 0
	FLOAT = 1
)


func NewEmptyMatrix(row int, column int)Matrix{


	var data [][]float32
	for i := 0; i < row; i++ {
		rowData := make([]float32, 0, row)
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

func MatrixRandom(ma Matrix, max float32, min float32)Matrix{
	for i := 0; i < ma.Row; i++ {
		for j := 0; j < ma.Column; j++ {
			ma.Data[i][j] =((max - min) *rand.Float32()) + min
		}
	}
	return ma
}

func Hprint(ma Matrix){
	for i := 0; i < ma.Row; i++ {
		s := ""
		for j := 0; j < ma.Column; j++ {
			s = s + fmt.Sprintf("%f ",ma.Data[i][j])
		}
		fmt.Printf("%s\n", s)

	}
}

func matrixMultipulate(){}



