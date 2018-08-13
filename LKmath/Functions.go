package LKmath

import (
	"math"
	"errors"
)

func Sigmoid(x float64)float64{
	return 1/(1+ math.Exp(x))
}

func OldLossFunction(yHat float64, y float64)float64{
	return (yHat - y) * (yHat - y)
}

func OldCostFunction(yHatMatrix *Matrix, yMatrix *Matrix)(float64, error){
	if yHatMatrix.Column !=1 || yMatrix.Column !=1 || yHatMatrix.Row != yMatrix.Row {
		return 0, errors.New("format error")
	}

	number := yHatMatrix.Row
	if number == 0{
		return 0, errors.New("void matrix")
	}

	var result float64
	for i :=0; i <yMatrix.Row; i++{
		result +=OldLossFunction(yHatMatrix.Data[i][0], yMatrix.Data[i][0])
	}

	return result/(2*float64(number)), nil

}


func LossFunction(yHat float64, y float64)float64{
	return -(y * math.Log(yHat))+ (1-y)*(math.Log(1-yHat))
}

func CostFunction(yHatMatrix *Matrix, yMatrix *Matrix)(float64, error){
	if yHatMatrix.Column !=1 || yMatrix.Column !=1 || yHatMatrix.Row != yMatrix.Row {
		return 0, errors.New("format error")
	}

	number := yHatMatrix.Row
	if number == 0{
		return 0, errors.New("void matrix")
	}

	var result float64
	for i :=0; i <yMatrix.Row; i++{
		result +=LossFunction(yHatMatrix.Data[i][0], yMatrix.Data[i][0])
	}

	return result/float64(number), nil

}