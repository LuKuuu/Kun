package LKmath

import (
	"fmt"
)



/*---------------------------------------------cost/loss function----------------------------------------------------*/

func LinearRegressionLossFunction(yHat float64, y float64)float64{
	return (yHat - y) * (yHat - y)
}

func LinearRegressionCostFunction(yHatMatrix *Matrix, yMatrix *Matrix)(float64, error){
	if yHatMatrix.Column !=1 || yMatrix.Column !=1 || yHatMatrix.Row != yMatrix.Row {
		panic("format error")
	}

	m := yHatMatrix.Row
	if m == 0{
		panic("void matrix")
	}

	var result float64
	for i :=0; i <yMatrix.Row; i++{
		result +=LinearRegressionLossFunction(yHatMatrix.Cell[i][0], yMatrix.Cell[i][0])
	}

	return result/(2*float64(m)), nil

}



/*---------------------------------------------linear regression----------------------------------------------------*/


func LinearRegressionGradientDecent(X Matrix, y Matrix, alpha float64, startParameter Matrix,passDerivative float64, learningTimes int)Matrix{
	//X: data( if we have m examples and n features, X should be a m * (n+1) matrix, with X.data[*][0]=1)
	//y: result (m examples means we should have m results y should be a m * 1 matrix)
	//startParameter (an n * 1 matrix with parameters we want to start the gradient decent)
	//alpha : learning rate (it should be carefully chose)
	if X.Column != startParameter.Row || X.Row != y.Row{
		panic("format error")
	}

	parameter := NewCopyMatrix(startParameter)
	derivative := NewEmptyMatrix(parameter.Row, 1)
	times := 0

	for{
		times ++
		derivative.Update(derivativeOfLinearRegressionCostFunction(X, y,parameter))
		if times%500000 == 0{
			fmt.Printf("progress : %f%%\n", float64(times)/float64(learningTimes))
			derivative.Hprint("current derivative is :")
			parameter.Hprint("and the parameter is: ")

		}

		pass := true
		for i := 0; i < derivative.Row; i++{
			if derivative.Cell[i][0] > passDerivative || derivative.Cell[i][0] < (-passDerivative){
				pass = false
				break
			}
		}

		if pass == true{
			break
		}

		if times > learningTimes{
			break
		}
		parameter.Update(MatrixSubtraction(parameter, ScalarMatrix(derivative,alpha) ))

	}

	return parameter
}


func derivativeOfLinearRegressionCostFunction(X Matrix, y Matrix,parameter Matrix)Matrix{

	//(Xt * ((X * parameter) - y))/m
	return ScalarMatrix(MatrixMultiplication(TransposeMatrix(X), MatrixSubtraction(MatrixMultiplication(X, parameter),y)),1/float64(X.Row))

}


/*-----------------------------------------scaring linear regression--------------------------------------------------*/

//todo : fix bugs in scaring linear regression

//todo add a supporting matrix
const(
	NotChange  = 0
	UseAverage = 1
	UseMax	   = 2
	UseMin	   = 3	//useless

)

//by using feature scaring, the speed of calculating can be improved
func ScaringGradientDecent(X Matrix, y Matrix, alpha float64, supportMatrix Matrix, startParameter Matrix, learningTimes int)Matrix{

	if supportMatrix.Column != X.Column{
		panic("support Matrix format error")
	}
	recoveryMatrix := NewEmptyMatrix(1, X.Column)
	absX := AbsMatrix(X)
	squeezedAverageX := SqueezedAverageColumnMatrix(absX)
	squeezedMaxX := SqueezedMaxColumnMatrix(absX)
	squeezedMinX := SqueezedMinColumnMatrix(absX)

	for i:=0; i < X.Column; i ++{
		switch supportMatrix.Cell[0][i] {
		case NotChange:
			recoveryMatrix.Cell[0][i] = 1
		case UseAverage:
			recoveryMatrix.Cell[0][i] = findDigit(squeezedAverageX.Cell[0][i])
		case UseMax:
			recoveryMatrix.Cell[0][i] = findDigit(squeezedMaxX.Cell[0][i])
		case UseMin:
			recoveryMatrix.Cell[0][i] = findDigit(squeezedMinX.Cell[0][i])
		default:
			fmt.Println("support Matrix value error, column with error value will not perform scaling")
			recoveryMatrix.Cell[0][i] = 1

		}

	}

	improvedX := MatrixAfterFeatureScaling(X, recoveryMatrix)
	improvedX.Hprint("here is the improved matrix")

	improvedStartParameter := StartParameterAfterFeatureScaling(startParameter, recoveryMatrix)



	rowResult := LinearRegressionGradientDecent(improvedX, y, alpha,improvedStartParameter, 0.005,learningTimes)
	rowResult.Hprint("row result is: ")
	return ResultRecovering(rowResult, recoveryMatrix)
}

func MatrixAfterFeatureScaling(X Matrix, recovery Matrix)Matrix{
	scaledX := NewCopyMatrix(X)

	for i := 0; i < scaledX.Column; i ++{
		for j := 0; j < scaledX.Row; j ++{
			scaledX.Cell[j][i] /= recovery.Cell[0][i]
		}
	}
	scaledX.Hprint("scaled matrix")
	return scaledX
}

func StartParameterAfterFeatureScaling(startParameter Matrix, recovery Matrix)Matrix{
	scaledParameter :=NewCopyMatrix(startParameter)
	for i := 0; i < scaledParameter.Column; i ++{
		scaledParameter.Cell[i][0] *= recovery.Cell[0][i]
	}
	return scaledParameter
}

func ResultRecovering(rowResult Matrix, recovery Matrix)Matrix{
	result :=NewCopyMatrix(rowResult)
	for i := 0; i < result.Row; i ++{
		result.Cell[i][0] /= recovery.Cell[0][i]
	}
	return result
}


func findDigit(value float64)float64{

	digit := 1.0
	for{
		if 0.1 <= value && value <= 1{
			return digit
		} else if value > 1{
			digit *= 10
			value /= 10
		}else if value < 0.1{
			digit /= 10
			value *= 10
		}
	}
}


