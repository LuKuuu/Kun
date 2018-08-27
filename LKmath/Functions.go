package LKmath

import (
	"math"
	"fmt"
	"strconv"
)

/*-------------------------------------------------linear regression----------------------------------------------------*/

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


func NormalEquation(X Matrix, y Matrix)Matrix{
	//note: when the normal equation meets the problem of non-invertible matrix
	//		regularized normal equation can fix it ;-)
	//		just put a really small number like o.ooo1 into lambda

	if y.Row != X.Row || y.Column != 1{
		panic("value format error")
	}
	XT := TransposedMatrix(X)
	return MatrixMultiplication(MatrixMultiplication(InverseMatrix(MatrixMultiplication(XT, X)), XT), y)


}


//regularized normal equation here is used to prevent over-fitting
//by adding the parameter lambda, the function can be more smooth
func RegularizedNormalEquation(X Matrix, y Matrix, lambda float64)Matrix{

	//lambda here is usually a really small number

	if y.Row != X.Row || y.Column != 1{
		panic("value format error")
	}
	XT := TransposedMatrix(X)
	
	penaltyMatrix := NewIdentityMatrix(X.Column)
	penaltyMatrix.Cell[0][0] = 0

	penalty := MatrixTimesRealNumber(penaltyMatrix, lambda)

	
	return MatrixMultiplication(MatrixMultiplication(InverseMatrix(MatrixAdd(MatrixMultiplication(XT, X),penalty)), XT), y)

}




func NormalGradientDecent(X Matrix, y Matrix, alpha float64, startParameter Matrix, learningTimes int)Matrix{
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
		derivative.Update(derivativeOfCostFunction(X, y,parameter))
		if times%100 == 0{
			derivative.Hprint(strconv.Itoa(times)+" times learning derivative is :")
			parameter.Hprint("and the parameter is: ")

		}

		pass := true
		for i := 0; i < derivative.Row; i++{
			if derivative.Cell[i][0] > 0.01 || derivative.Cell[i][0] < (-0.01){
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
		parameter.Update(MatrixSubtraction(parameter, MatrixTimesRealNumber(derivative,alpha) ))
		//parameter.Hprint(strconv.Itoa(times)+" times parameter is :")

	}

	return parameter
}


func derivativeOfCostFunction(X Matrix, y Matrix,parameter Matrix)Matrix{

	//Xt * ((X * parameter) - y)
	return MatrixTimesRealNumber(MatrixMultiplication(TransposedMatrix(X), MatrixSubtraction(MatrixMultiplication(X, parameter),y)),1/float64(X.Row))

}

//todo add a supporting matrix
const(
	NotChange  = 0
	UseAverage = 1
	UseMax	   = 2
	UseMin	   = 3	//useless

)

//by using feature scaring, the result can be improved
func ScariedGradientDecent(X Matrix, y Matrix, alpha float64, supportMatrix Matrix, startParameter Matrix, learningTimes int)Matrix{

	if supportMatrix.Column != X.Column{
		panic("support Matrix format error")
	}
	recoveryMatrix := NewEmptyMatrix(1, X.Column)
	absX := AbsMatrix(X)
	squeezedAverageX := SqueezedAverageRowMatrix(absX)
	squeezedMaxX := SqueezedMaxRowMatrix(absX)
	squeezedMinX := SqueezedMinRowMatrix(absX)

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



	rowResult :=NormalGradientDecent(improvedX, y, alpha,improvedStartParameter, learningTimes)
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




/*--------------------------------------------------logical regression--------------------------------------------------*/

func SigmoidFunction(x float64)float64{
	return 1/(1+ math.Exp(x))
}

func LogicalRegressionLossFunction(yHat float64, y float64)float64{
	return -(y * math.Log(yHat))+ (1-y)*(math.Log(1-yHat))
}

func LogicalRegressionCostFunction(yHatMatrix *Matrix, yMatrix *Matrix)float64{
	if yHatMatrix.Column !=1 || yMatrix.Column !=1 || yHatMatrix.Row != yMatrix.Row {
		panic("format error")
	}

	number := yHatMatrix.Row
	if number == 0{
		panic("void matrix")
	}

	var result float64
	for i :=0; i <yMatrix.Row; i++{
		result +=LogicalRegressionLossFunction(yHatMatrix.Cell[i][0], yMatrix.Cell[i][0])
	}

	return result/float64(number)

}

