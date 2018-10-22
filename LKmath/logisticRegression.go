package LKmath

import (
	"fmt"
	"math"
	"math/rand"
	"time"
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

/*--------------------------------------------------NodeParameter---------------------------------------------------------*/



type NodeParameter struct {
	W Matrix
	B float64
	DW Matrix
	DB float64
}

func NewEmptyNode(featureNum int)NodeParameter{
	return NodeParameter{NewValuedMatrix(1,featureNum, 0), 0,NewValuedMatrix(1,featureNum, 0), 0}
}

func NewValuedNode(featureNum int, value float64)NodeParameter{
	return NodeParameter{NewValuedMatrix(1,featureNum, value), value,NewValuedMatrix(1,featureNum, 0), 0}
}

func NewRandomNode(initialize bool, featureNum int, max float64, min float64)NodeParameter {
	if initialize {
		rand.Seed(time.Now().Unix())
	}
	np :=NodeParameter{NewRandomMatrix(initialize, 1, featureNum, max, min), 0,NewValuedMatrix(1,featureNum, 0), 0}
	np.B=-Sum(np.W)*0.5
	return np
}

func (this *NodeParameter)Update(np NodeParameter){

	this.W.Update(np.W)
	this.B=np.B
	this.DW.Update(np.DW)
	this.DB=np.DB
}

func (this *NodeParameter)Hprint(info string){
	fmt.Printf(info+"\n")

	fmt.Printf("W:\n")
	for i := 0; i < this.W.Row; i++ {
		s := ""
		for j := 0; j < this.W.Column; j++ {
			s = s + fmt.Sprintf("%f\t",this.W.Cell[i][j])
		}
		fmt.Printf("%s\n", s)

	}

	fmt.Printf("B: %f\n", this.B)

	fmt.Printf("dW:\n")
	for i := 0; i < this.DW.Row; i++ {
		s := ""
		for j := 0; j < this.DW.Column; j++ {
			s = s + fmt.Sprintf("%f\t",this.DW.Cell[i][j])
		}
		fmt.Printf("%s\n", s)

	}
	fmt.Printf("dB: %f\n", this.DB)
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

func DerivativeOfSigmoidFunctionForMatrix(yHatMatrix Matrix)Matrix{
	resultMatrix := NewCopyMatrix(yHatMatrix)
	for i := 0; i < yHatMatrix.Row; i++{
		for j := 0; j < yHatMatrix.Column; j ++{
			resultMatrix.Cell[i][j] = (yHatMatrix.Cell[i][j]) * (1 - yHatMatrix.Cell[i][j])
		}
	}
	return resultMatrix
}


func YHat(X Matrix, parameter NodeParameter)Matrix{
	bMatrix := NewValuedMatrix(1, X.Column, parameter.B)
	//yHat = sigmoid(transpose(W)*X + b)
	return SigmoidFunctionForMatrix(MatrixAddition(MatrixMultiplication(parameter.W, X), bMatrix))
}


//loss/cost functions

func LogisticRegressionLossFunction(yHat float64, y float64)float64{

	//return -(y * math.Log(yHat)+ (1-y)*(math.Log(1-yHat)))
	// this sometime will make a stupid return "Not a Number" because it will calculate 0 * inf first

	if y == 0{
		return - math.Log(1-yHat)
	}else{
		return - math.Log(yHat)
	}
}
func LogisticRegressionLossFunctionForMatrix(yHatMatrix Matrix, yMatrix Matrix)Matrix{
	result :=NewEmptyMatrix(1, yHatMatrix.Column)
	for i := 0; i < yHatMatrix.Column; i ++{
		result.Cell[0][i] = LogisticRegressionLossFunction(yHatMatrix.Cell[0][i], yMatrix.Cell[0][i])
	}
	return result
}


func DerivativeOfLogisticRegressionLossFunction(yHat float64, y float64)float64{


	//return -(y/yHat)+((1-y)/(1-yHat))
	//this will result Not a Number sometime because it will calculate 0/0 instead of ignoring the part with numerator == 0

	if y == 0{
		return 1/(1-yHat)
	}else{
		return -1/yHat
	}
}

func DerivativeOfLogisticRegressionLossFunctionForMatrix(yHatMatrix Matrix, yMatrix Matrix)Matrix{
	result :=NewEmptyMatrix(1, yHatMatrix.Column)
	for i := 0; i < yHatMatrix.Column; i ++{
		result.Cell[0][i] = DerivativeOfLogisticRegressionLossFunction(yHatMatrix.Cell[0][i], yMatrix.Cell[0][i])
	}
	return result
}

const maxValue = 0
func FinalDerivativeOfLogisticRegressionForMatrix(yHatMatrix Matrix, yMatrix Matrix)Matrix {

	if yHatMatrix.Column != yMatrix.Column{
		panic("FinalDerivativeOfLogisticRegressionForMatrix: format error")
	}
	da := DerivativeOfLogisticRegressionLossFunctionForMatrix(yHatMatrix, yMatrix)
	dl := DerivativeOfSigmoidFunctionForMatrix(yHatMatrix)


	result :=NewEmptyMatrix(1, da.Column)
	for i:=0; i< da.Column; i++{
		result.Cell[0][i] = da.Cell[0][i] * dl.Cell[0][i]
		//to prevent the problem of not a number
		if math.IsNaN(result.Cell[0][i]){
			result.Cell[0][i] = math.MaxFloat64
		}
	}

	return result


}

func LogisticRegressionCostFunction(yHatMatrix Matrix, yMatrix Matrix)float64{
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



func DerivativeOfLogisticalRegressionCostFunction(X Matrix, yMatrix Matrix, parameter NodeParameter)(NodeParameter, Matrix){


	yHatMatrix :=YHat(X, parameter)

	finalDerivative :=FinalDerivativeOfLogisticRegressionForMatrix(yHatMatrix, yMatrix)

	dW := ScalarMatrix(MatrixMultiplication(finalDerivative, TransposeMatrix(X)), 1/float64(X.Column))

	dB :=Average(finalDerivative)

	return NodeParameter{parameter.W, parameter.B, dW, dB}, finalDerivative






}

/*---------------------------------gradient decent------------------------------*/

func LogisticRegressionGradientDecent(X Matrix, y Matrix, alpha float64, startParameter NodeParameter, learningTimes int)NodeParameter{

	//X: data( if we have m examples and n features, X should be an n * m matrix)
	//y: result (m examples means we should have m results so y should be a 1 * m matrix)
	//startParameter (	an 1 * n matrix with parameters
	//			  		and a	parameter b we want to start the gradient decent
	//alpha : learning rate (it should be carefully chose)

	if X.Row != startParameter.W.Column || X.Column != y.Column{
		panic("format error")
	}

	fmt.Printf("start logistic regression\n")


	parameterW := NewCopyMatrix(startParameter.W)
	ParameterB := startParameter.B
	parameter := NodeParameter{parameterW, ParameterB, NewEmptyMatrix(parameterW.Row, parameterW.Column), 0}

	times := 0

	for{
		times ++
		d, _:=DerivativeOfLogisticalRegressionCostFunction(X, y,parameter)
		parameter.Update(d)
		if times%5000 == 0{

			parameter.Hprint(fmt.Sprintf("\nprogress : %f", float64(times*100)/float64(learningTimes))+"%%")

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

		parameter.W = MatrixSubtraction(parameter.W, ScalarMatrix(parameter.DW,alpha))
		parameter.B -= alpha * parameter.DB

	}

	return parameter
}
