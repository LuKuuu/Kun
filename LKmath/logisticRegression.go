package LKmath

import (
	"fmt"
	"math"
)

/*--------------------------------------------------example--------------------------------------------------*/

/*
notice: logistic regression is not only used for classification, but also used for neural network

for convenient,
1. the features matrix X should has n(features) rows and m(number of examples) columns
2. the result matrix y   should has rn(result features) row and m columns
3. there should be one parameter b for each logistical regression as bias; therefore, there should be a basis matrix
	it should has rn rows and 1 column
4. features parameter is noted as W which is a n by rn matrix

here are the relationship between them

yHat=sigmoid(W*X+ b)


 */

/*--------------------------------------------------LayerParameter---------------------------------------------------------*/

type LayerParameter struct {
	FeatureNum int
	OutputNum  int
	W          Matrix
	B          Matrix
	dW         Matrix
	dB         Matrix
}

func NewEmptyNode(featureNum int, outputNum int) LayerParameter {
	return LayerParameter{
		featureNum, outputNum,
		NewValuedMatrix(outputNum, featureNum, 0),
		NewValuedMatrix(outputNum, 1, 0),
		NewValuedMatrix(outputNum, featureNum, 0),
		NewValuedMatrix(outputNum, 1, 0),
	}
}

func NewValuedNode(featureNum int, outputNum int, value float64) LayerParameter {
	return LayerParameter{
		featureNum, outputNum,
		NewValuedMatrix(outputNum, featureNum, value),
		NewValuedMatrix(outputNum, 1, value),
		NewValuedMatrix(outputNum, featureNum, 0),
		NewValuedMatrix(outputNum, 1, 0),
	}
}

func NewRandomNode(featureNum int, outputNum int, max float64, min float64) LayerParameter {

	return LayerParameter{
		featureNum, outputNum,
		NewRandomMatrix(outputNum, featureNum, max, min),
		NewRandomMatrix(outputNum, 1, max, min),
		NewValuedMatrix(outputNum, featureNum, 0),
		NewValuedMatrix(outputNum, 1, 0),
	}
}

func (this *LayerParameter) Update(np LayerParameter) {

	if this.FeatureNum != np.FeatureNum || this.OutputNum != np.OutputNum {
		panic("format error")
	}

	this.W.Update(np.W)
	this.B.Update(np.B)
	this.dW.Update(np.dW)
	this.dB.Update(np.dB)
}

func (this *LayerParameter) Hprint(info string) {
	fmt.Printf(info + "\n")

	fmt.Printf("parameter:\n")

	for i := 0; i < this.W.Row; i++ {
		s := fmt.Sprintf("B: %f \t\t|W:\t", this.B.Cell[i][0])
		for j := 0; j < this.W.Column; j++ {
			s = s + fmt.Sprintf("%f \t", this.W.Cell[i][j])
		}
		fmt.Printf("%s\n", s)

	}

	fmt.Printf("derivative of parameter:\n")
	for i := 0; i < this.W.Row; i++ {
		s := fmt.Sprintf("dB: %f \t\t|dW:\t", this.dB.Cell[i][0])
		for j := 0; j < this.W.Column; j++ {
			s = s + fmt.Sprintf("%f \t", this.dW.Cell[i][j])
		}
		fmt.Printf("%s\n", s)

	}

	fmt.Println()
}

//Sigmoid Functions

func SigmoidFunction(x float64) float64 {
	return 1 / (1 + math.Exp(-x))
}

func SigmoidFunctionForMatrix(m Matrix) Matrix {
	resultMatrix := NewCopyMatrix(m)
	for i := 0; i < m.Row; i++ {
		for j := 0; j < m.Column; j ++ {
			resultMatrix.Cell[i][j] = SigmoidFunction(m.Cell[i][j])
		}
	}
	return resultMatrix
}

func DerivativeOfSigmoidFunction(x float64) float64 {
	return SigmoidFunction(x) * (1 - SigmoidFunction(x))
}

func DerivativeOfSigmoidFunctionForMatrix(yHatMatrix Matrix) Matrix {
	resultMatrix := NewCopyMatrix(yHatMatrix)
	for i := 0; i < yHatMatrix.Row; i++ {
		for j := 0; j < yHatMatrix.Column; j ++ {
			resultMatrix.Cell[i][j] = DerivativeOfSigmoidFunction(yHatMatrix.Cell[i][j])
		}
	}
	return resultMatrix
}

func YHat(X Matrix, parameter LayerParameter) Matrix {
	bMatrix := parameter.ExpandedB(X.Column)
	//yHat = sigmoid(transpose(W)*X + b)
	return SigmoidFunctionForMatrix(MatrixAddition(MatrixMultiplication(parameter.W, X), bMatrix))
}

func (parameter *LayerParameter)YHat(X Matrix)Matrix {
	bMatrix := parameter.ExpandedB(X.Column)
	//yHat = sigmoid(transpose(W)*X + b)
	return SigmoidFunctionForMatrix(MatrixAddition(MatrixMultiplication(parameter.W, X), bMatrix))
}

func (this *LayerParameter) ExpandedB(n int) Matrix {
	result := NewEmptyMatrix(this.B.Row, n)
	for i := 0; i < result.Row; i++ {
		for j := 0; j < result.Column; j++ {
			result.Cell[i][j] = this.B.Cell[i][0]
		}
	}
	return result
}

//loss/cost functions

func LogisticRegressionLossFunction(yHat float64, y float64) float64 {

	//return -(y * math.Log(yHat)+ (1-y)*(math.Log(1-yHat)))
	// this sometime will make a stupid return "Not a Number" because it will calculate 0 * inf first

	if y < 0.5 {
		return - math.Log(1 - yHat)
	} else {
		return - math.Log(yHat)
	}
}
func LogisticRegressionLossFunctionForMatrix(yHatMatrix Matrix, yMatrix Matrix) Matrix {

	if (yHatMatrix.Column != yMatrix.Column) || (yHatMatrix.Row != yMatrix.Row) {
		panic("LogisticRegressionCostFunction : format error")
	}
	result := NewEmptyMatrix( yHatMatrix.Row, yHatMatrix.Column)

	for i := 0; i < result.Row; i++ {
		for j := 0; j < result.Column; j ++ {
			result.Cell[i][j] = LogisticRegressionLossFunction(yHatMatrix.Cell[i][j], yMatrix.Cell[i][j])
		}
	}

	return result
}

func LogisticRegressionCostFunction(yHatMatrix Matrix, yMatrix Matrix) float64 {

	return Average(LogisticRegressionLossFunctionForMatrix(yHatMatrix, yMatrix))

}

func DerivativeOfLogisticRegressionLossFunction(yHat float64, y float64) float64 {

	//return -(y/yHat)+((1-y)/(1-yHat))
	//this will result Not a Number sometime because it will calculate 0/0 instead of ignoring the part with numerator == 0

	if y == 0 {
		return 1 / (1 - yHat)
	} else {
		return -1 / yHat
	}
}

func DerivativeOfLogisticRegressionLossFunctionForMatrix(yHatMatrix Matrix, yMatrix Matrix) Matrix {
	result := NewEmptyMatrix(yHatMatrix.Row, yHatMatrix.Column)

	for i := 0; i < result.Row; i++ {
		for j := 0; j < result.Column; j ++ {
			result.Cell[i][j] = DerivativeOfLogisticRegressionLossFunction(yHatMatrix.Cell[i][j], yMatrix.Cell[i][j])
		}
	}

	return result
}

func FinalDerivativeOfLogisticRegressionForMatrix(LastDerivativeMatrix Matrix, yHatMatrix Matrix) Matrix {

	if yHatMatrix.Column != LastDerivativeMatrix.Column || yHatMatrix.Row != LastDerivativeMatrix.Row {
		panic("FinalDerivativeOfLogisticRegressionForMatrix: format error")
	}
	dl := DerivativeOfSigmoidFunctionForMatrix(yHatMatrix)

	return DotProduct(LastDerivativeMatrix, dl)

}

func (this *LayerParameter) UpdateDerivative(X Matrix, y Matrix) (derivativeMatrix Matrix) {

	yHatMatrix := this.YHat(X)

	da := DerivativeOfLogisticRegressionLossFunctionForMatrix(yHatMatrix, y)

	finalDerivative := FinalDerivativeOfLogisticRegressionForMatrix(da, yHatMatrix)

	this.dW= ScalarMatrix(MatrixMultiplication(finalDerivative, TransposeMatrix(X)), 1/float64(X.Column))

	this.dB= SqueezedAverageRowMatrix(finalDerivative)

	return  finalDerivative

}


func (this *LayerParameter) UpdateDerivativeWithDerivative(X Matrix, lastDerivativeMatrix Matrix) (derivativeMatrix Matrix) {

	yHatMatrix := this.YHat(X)

	finalDerivative := FinalDerivativeOfLogisticRegressionForMatrix(lastDerivativeMatrix, yHatMatrix)

	this.dW= ScalarMatrix(MatrixMultiplication(finalDerivative, TransposeMatrix(X)), 1/float64(X.Column))

	this.dB= SqueezedAverageRowMatrix(finalDerivative)

	return  finalDerivative

}
/*---------------------------------gradient decent------------------------------*/

func LogisticRegressionGradientDecent(X Matrix, y Matrix, alpha float64, startParameter LayerParameter, learningTimes int) LayerParameter {

	//X: data( if we have m examples and n features, X should be an n * m matrix)
	//y: result (m examples means we should have m results so y should be a 1 * m matrix)
	//startParameter (	an 1 * n matrix with parameters
	//			  		and a	startParameter b we want to start the gradient decent
	//alpha : learning rate (it should be carefully chose)

	if X.Row != startParameter.FeatureNum || startParameter.OutputNum != y.Row || X.Column != y.Column {
		panic("format error")
	}

	fmt.Printf("start logistic regression\n")

	parameter := NewEmptyNode(startParameter.FeatureNum, startParameter.OutputNum)
	parameter.Update(startParameter)
	for i := 0; i < learningTimes; i++ {
		parameter.UpdateDerivative(X,y)
		parameter.W.Update(MatrixSubtraction(parameter.W, ScalarMatrix(parameter.dW, alpha)))
		parameter.B.Update(MatrixSubtraction(parameter.B, ScalarMatrix(parameter.dB, alpha)))

		if i%50000== 0 {

			fmt.Printf("cost is :%f\n",LogisticRegressionCostFunction(parameter.YHat(X),y))
			parameter.Hprint(fmt.Sprintf("\nprogress : %f", float64(i*100)/float64(learningTimes)) + "%%")

		}

	}

	return startParameter
}
