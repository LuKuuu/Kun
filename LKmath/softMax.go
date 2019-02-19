package LKmath

import "math"

func SoftMaxYHat(X Matrix, lastLayer LayerParameter) (yHat Matrix) {
	bMatrix := lastLayer.ExpandedB(X.Column)
	//yHat = (e^(transpose(W)*X + b))/sum(e)
	yHat = ExpMatrix(MatrixAddition(MatrixMultiplication(lastLayer.W, X), bMatrix))

	sumMatrix := SqueezedSumColumnMatrix(yHat)

	for i := 0; i < yHat.Column; i++ {
		for j := 0; j < yHat.Row; j++ {
			yHat.Cell[j][i] /= sumMatrix.Cell[0][i]
		}

	}

	return yHat


}

func (lastLayer *LayerParameter)SoftMaxYHat(X Matrix) (yHat Matrix) {
	bMatrix := lastLayer.ExpandedB(X.Column)
	//yHat = (e^(transpose(W)*X + b))/sum(e)
	yHat = ExpMatrix(MatrixAddition(MatrixMultiplication(lastLayer.W, X), bMatrix))

	sumMatrix := SqueezedSumColumnMatrix(yHat)

	for i := 0; i < yHat.Column; i++ {
		for j := 0; j < yHat.Row; j++ {
			yHat.Cell[j][i] /= sumMatrix.Cell[0][i]
		}

	}

	return yHat


}


func ExpMatrix(x Matrix) Matrix {
	result := NewEmptyMatrix(x.Row, x.Column)
	for i := 0; i < result.Column; i++ {
		for j := 0; j < result.Row; j++ {
			result.Cell[j][i] = math.Exp(x.Cell[j][i])
		}

	}
	return result
}

func softMaxCostFunction(y float64,yHat float64)float64{
	if y<=0.5{
		return 0
	}else{
		return -math.Log(yHat)
	}
}

func DerivativeOfSoftMaxCostFunction(y float64,yHat float64)float64{
	return yHat-y
}


func DerivativeOfSoftMaxCostFunctionForMatrix(yHat Matrix, y Matrix)Matrix{
	result := NewEmptyMatrix(y.Row, y.Column)
	for i := 0; i < result.Column; i++ {
		for j := 0; j < result.Row; j++ {
			result.Cell[j][i] = DerivativeOfSoftMaxCostFunction(y.Cell[j][i],yHat.Cell[j][i])
		}

	}

	return result
}


func SoftMaxCostFunctionForMatrix(yHat Matrix, y Matrix)float64{
	result := NewEmptyMatrix(y.Row, y.Column)
	for i := 0; i < result.Column; i++ {
		for j := 0; j < result.Row; j++ {
			result.Cell[j][i] = softMaxCostFunction(y.Cell[j][i],yHat.Cell[j][i])
		}

	}
	return Average(result)
}
