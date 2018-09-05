package LKmath

import (
	"math"
	"fmt"
)
/*--------------------------------------------------example--------------------------------------------------*/

/*
notice: logistic regression is not only used for classification, but also used for neural network

for convenient,
1. the features matrix X should has n(features) rows and m(number of examples) columns
2. the result matrix y   should has 1 row and m columns
3. there should be one parameter b for each logistical regression as bias
4. features parameter is noted as W which is a n by 1 matrix

here are the relationship between them


 */

/*--------------------------------------------------logical regression--------------------------------------------------*/


type NodeParameter struct {
	W Matrix
	B float64
}

func (this *NodeParameter)Update(np NodeParameter){
	this.W.Update(np.W)
	this.B=np.B
}

func (this *NodeParameter)Hprint(info string){
	fmt.Printf(info+"\n")
	fmt.Printf("W is :\n")
	for i := 0; i < this.W.Row; i++ {
		s := ""
		for j := 0; j < this.W.Column; j++ {
			s = s + fmt.Sprintf("%f ",this.W.Cell[i][j])
		}
		fmt.Printf("%s\n", s)

	}

	fmt.Printf("B is : %f \n", this.B)

	fmt.Println()
}

//Sigmoid Functions

func SigmoidFunction(x float64)float64{
	return 1/(1+ math.Exp(-x))
}

func SigmoidFunctionForMatrix(m Matrix)Matrix{
	resultMatrix := NewCopyMatrix(m)
	for i := 0; i < m.Row; i++{
		for j := 0; j < m.Column; j ++{
			resultMatrix.Cell[i][j] = SigmoidFunction(m.Cell[i][j])
		}
	}
	return resultMatrix
}

func DerivativeOfSigmoidFunction(x float64)float64{
	return SigmoidFunction(x) * (1- SigmoidFunction(x))
}

func DerivativeOfSigmoidFunctionForMatrix(m Matrix)Matrix{
	resultMatrix := NewCopyMatrix(m)
	for i := 0; i < m.Row; i++{
		for j := 0; j < m.Column; j ++{
			resultMatrix.Cell[i][j] = DerivativeOfSigmoidFunction(m.Cell[i][j])
		}
	}
	return resultMatrix
}


func YHat(X Matrix, parameter NodeParameter)Matrix{
	bMatrix := NewValuedMatrix(1, X.Column, parameter.B)
	//yHat = sigmoid(transpose(W)*X + b)
	return SigmoidFunctionForMatrix(MatrixAdd(MatrixMultiplication(TransposedMatrix(parameter.W), X), bMatrix))
}

//loss/cost functions

func LogisticRegressionLossFunction(yHat float64, y float64)float64{
	return -(y * math.Log(yHat))+ (1-y)*(math.Log(1-yHat))
}

func DerivativeOfLogisticRegressionLossFunction(yHat float64, y float64)float64{
	return -(y/yHat)+((1-y)/1-yHat)
}

func DerivativeOfLogisticRegressionLossFunctionForMatrix(yHatMatrix Matrix, yMatrix Matrix)Matrix{
	result :=NewEmptyMatrix(1, yHatMatrix.Column)
	for i := 0; i < yHatMatrix.Column; i ++{
		result.Cell[0][i] = DerivativeOfLogisticRegressionLossFunction(yHatMatrix.Cell[0][i], yMatrix.Cell[0][i])
	}
	return result
}

func LogisticRegressionCostFunction(yHatMatrix *Matrix, yMatrix *Matrix)float64{
	if yHatMatrix.Row !=1 || yMatrix.Row != 1 || yHatMatrix.Column != yMatrix.Column{
		panic("LogisticRegressionCostFunction : format error")
	}

	size := yMatrix.Column
	result := 0.0
	for i := 0; i < size; i++{
		result +=LogisticRegressionLossFunction(yHatMatrix.Cell[0][i],yMatrix.Cell[0][i])
	}
	return result/float64(size)

}



func derivativeOfLogisticalRegressionCostFunction(X Matrix, yMatrix Matrix, parameter NodeParameter)NodeParameter{

	bMatrix :=NewValuedMatrix(1, yMatrix.Column, parameter.B) //1*m matrix

	//z here is an 1 * m Matrix equals to Wt * X + b
	z := MatrixAdd(MatrixMultiplication(TransposedMatrix(parameter.W), X), bMatrix)
	yHatMatrix := SigmoidFunctionForMatrix(z)

	//da here is the derivative part in loss function
	da := DerivativeOfLogisticRegressionLossFunctionForMatrix(yHatMatrix, yMatrix)

	//dl here is the derivative part in sigmoid function
	dl := DerivativeOfSigmoidFunctionForMatrix(z)

	//derivative here is an 1 * m Matrix before ...
	derivative := NewEmptyMatrix(1, yHatMatrix.Column)

	for i := 0; i < derivative.Column; i ++{
		derivative.Cell[0][i] = da.Cell[0][i] * dl.Cell[0][i]
	}

	//for each parameter their derivative is derivative * Xt

	parameterDerivative := NodeParameter{}

	parameterDerivative.B = Average(derivative)
	parameterDerivative.W = MatrixTimesRealNumber(MatrixMultiplication(derivative, TransposedMatrix(X)), 1/float64(yHatMatrix.Column))
	parameterDerivative.Hprint("parameterDerivative")

	return parameterDerivative


}

func LogisticRegressionGradientDecent(X Matrix, y Matrix, alpha float64, startParameter NodeParameter, learningTimes int)NodeParameter{

	//X: data( if we have m examples and n features, X should be an n * m matrix)
	//y: result (m examples means we should have m results so y should be a 1 * m matrix)
	//startParameter (	an 1 * n matrix with parameters
	//			  		and a	parameter b we want to start the gradient decent
	//alpha : learning rate (it should be carefully chose)

	if X.Row != startParameter.W.Row || X.Column != y.Column{
		panic("format error")
	}


	parameterW := NewCopyMatrix(startParameter.W)
	ParameterB := startParameter.B
	parameter := NodeParameter{parameterW, ParameterB}
	derivative := NodeParameter{NewEmptyMatrix(1, parameterW.Column), 0}

	times := 0

	for{
		times ++
		derivative.Update(derivativeOfLogisticalRegressionCostFunction(X, y,parameter))
		if times%5 == 0{
			fmt.Printf("progress : %d%%\n", times/learningTimes)
			derivative.Hprint("current derivative is :")
			parameter.Hprint("and the parameter is: ")

		}

		//pass := true
		//for i := 0; i < derivative.W.Row; i++{
		//	if derivative.W.Cell[i][0] > 0.01 || derivative.W.Cell[i][0] < (-0.01){
		//		pass = false
		//		break
		//	}
		//}
		//if derivative.B > 0.01 || derivative.B < (-0.01){
		//	pass = false
		//	break
		//}
		//
		//if pass == true{
		//	break
		//}

		if times > learningTimes{
			break
		}
		parameter.Update(NodeParameter{W :MatrixSubtraction(parameter.W,MatrixTimesRealNumber(derivative.W,alpha)), B:(parameter.B - alpha * derivative.B)})
		//parameter.Hprint(strconv.Itoa(times)+" times parameter is :")

	}

	return parameter
}
