package LKmath

import (
	"fmt"
)


/*-------------------------------------------------example----------------------------------------------------*/

func EXAMPLE(){

	//a real problem : find the relationship between features of house and house price

	X :=NewEmptyMatrix(5,4)
	//features of houses
	//-----x0-------------size----------bedrooms-----------age-------
	X.Cell[0][0]=1;	X.Cell[0][1]=2104; X.Cell[0][2]=5;X.Cell[0][3]=45
	X.Cell[1][0]=1; X.Cell[1][1]=1416; X.Cell[1][2]=3;X.Cell[1][3]=40
	X.Cell[2][0]=1; X.Cell[2][1]=1534; X.Cell[2][2]=3;X.Cell[2][3]=30
	X.Cell[3][0]=1; X.Cell[3][1]= 850; X.Cell[3][2]=2;X.Cell[3][3]=36
	X.Cell[4][0]=1; X.Cell[4][1]=1300; X.Cell[4][2]=4;X.Cell[4][3]=50



	X.Hprint("features of houses (X):")

	a := Determinant(MatrixMultiplication(TransposedMatrix(X), X))
	fmt.Printf("%v\n", a)


	y :=NewEmptyMatrix(5,1)
	//--price of houses
	y.Cell[0][0]=460
	y.Cell[1][0]=232
	y.Cell[2][0]=315
	y.Cell[3][0]=178
	y.Cell[4][0]=220
	y.Hprint("price of each house are")



	example :=NewEmptyMatrix(1, 4)
	example.Cell[0][0] = 1; example.Cell[0][1] = 1000; example.Cell[0][2] = 4; example.Cell[0][3] = 0
	example.Hprint("now I have a house with 1000 square feet, 4 bedrooms and is a brand new house")



	result := NormalEquation(X, y)
	result.Hprint("the result is :")
	priceOfExampleHouse :=MatrixMultiplication(example, result)
	priceOfExampleHouse.Hprint("predicted price of example house is :")



	regularizedResult := RegularizedNormalEquation(X, y, 0.001)
	regularizedResult.Hprint("the regularized result is :")
	priceOfRegularizedExampleHouse :=MatrixMultiplication(example, regularizedResult)
	priceOfRegularizedExampleHouse.Hprint("regularized predicted price of example house is :")

}


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




func LinearRegressionGradientDecent(X Matrix, y Matrix, alpha float64, startParameter Matrix, learningTimes int)Matrix{
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
			fmt.Printf("progress : %d%%\n", times/learningTimes)
			derivative.Hprint("current derivative is :")
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


func derivativeOfLinearRegressionCostFunction(X Matrix, y Matrix,parameter Matrix)Matrix{

	//(Xt * ((X * parameter) - y))/m
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



	rowResult := LinearRegressionGradientDecent(improvedX, y, alpha,improvedStartParameter, learningTimes)
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


