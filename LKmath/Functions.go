package LKmath

import (
	"math"
	"errors"
)

//for linear regression

func LinearRegressionLossFunction(yHat float64, y float64)float64{
	return (yHat - y) * (yHat - y)
}

func LinearRegressionCostFunction(yHatMatrix *Matrix, yMatrix *Matrix)(float64, error){
	if yHatMatrix.Column !=1 || yMatrix.Column !=1 || yHatMatrix.Row != yMatrix.Row {
		return 0, errors.New("format error")
	}

	number := yHatMatrix.Row
	if number == 0{
		return 0, errors.New("void matrix")
	}

	var result float64
	for i :=0; i <yMatrix.Row; i++{
		result +=LinearRegressionLossFunction(yHatMatrix.Cell[i][0], yMatrix.Cell[i][0])
	}

	return result/(2*float64(number)), nil

}


func NormalEquation(X Matrix, y Matrix)Matrix{
	if y.Row != X.Row || y.Column != 1{
		panic("value format error")
	}
	XT := TransposedMatrix(X)
	return MatrixMultiplication(MatrixMultiplication(InverseMatrix(MatrixMultiplication(XT, X)), XT), y)


}




//for logical regression

func SigmoidFunction(x float64)float64{
	return 1/(1+ math.Exp(x))
}

func LogicalRegressionLossFunction(yHat float64, y float64)float64{
	return -(y * math.Log(yHat))+ (1-y)*(math.Log(1-yHat))
}

func LogicalRegressionCostFunction(yHatMatrix *Matrix, yMatrix *Matrix)(float64, error){
	if yHatMatrix.Column !=1 || yMatrix.Column !=1 || yHatMatrix.Row != yMatrix.Row {
		return 0, errors.New("format error")
	}

	number := yHatMatrix.Row
	if number == 0{
		return 0, errors.New("void matrix")
	}

	var result float64
	for i :=0; i <yMatrix.Row; i++{
		result +=LogicalRegressionLossFunction(yHatMatrix.Cell[i][0], yMatrix.Cell[i][0])
	}

	return result/float64(number), nil

}

