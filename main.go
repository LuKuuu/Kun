package main

import (
	"fmt"
	"time"
	"github.com/LuKuuu/Kun/LKmath"
)

//testing

func main() {

	fmt.Printf("%v\n", time.Now())

/*----------------------------------------test of cost/loss function--------------------------------------------------*/

	//a :=LKmath.NewEmptyMatrix(10,1)
	//b :=LKmath.NewCopyMatrix(a)
	//LKmath.MatrixRandom(b,100,-100)
	//
	//
	//a.Hprint()
	//b.Hprint()
	//
	//
	//OldCost, err := LKmath.OldCostFunction(&b, &a)
	//if err != nil{
	//	fmt.Printf("%v", err)
	//}else{
	//	fmt.Printf("the old cost function result between a and b is %v\n",OldCost)
	//
	//}
	//
	//a.MatrixSigmoid()
	//b.MatrixSigmoid()
	//
	//Cost, err := LKmath.CostFunction(&a, &b)
	//if err != nil{
	//	fmt.Printf("%v", err)
	//}else{
	//	fmt.Printf("the cost function result between sigmoid a and sigmiod b is %v\n",Cost)
	//
	//}


/*-------------------------------------test of matrix multiplication and transpose-------------------------------------*/
	//
	//
	//p :=LKmath.NewEmptyMatrix(3,2)
	//p.Cell[0][0]=1;	p.Cell[0][1]=2
	//p.Cell[1][0]=3; p.Cell[1][1]=4
	//p.Cell[2][0]=5; p.Cell[2][1]=6
	//
	//
	//q :=LKmath.NewEmptyMatrix(2,2)
	//q.Cell[0][0]=5;	q.Cell[0][1]=6
	//q.Cell[1][0]=7; q.Cell[1][1]=8
	//
	//k :=LKmath.MatrixMultiplication(p, q)
	//k.Hprint("K is:")
	//
	//z:=LKmath.TransposedMatrix(p)
	//p.Hprint("p is:")
	//z.Hprint("z is:")


/*-------------------------------------------test of normal equation--------------------------------------------*/

	//a real problem : find the relationship between features of house and house price

	//X :=LKmath.NewEmptyMatrix(5,4)
	////features of houses
	////-----x0-------------size----------bedrooms-----------age-------
	//X.Cell[0][0]=1;	X.Cell[0][1]=2104; X.Cell[0][2]=5;X.Cell[0][3]=45
	//X.Cell[1][0]=1; X.Cell[1][1]=1416; X.Cell[1][2]=3;X.Cell[1][3]=40
	//X.Cell[2][0]=1; X.Cell[2][1]=1534; X.Cell[2][2]=3;X.Cell[2][3]=30
	//X.Cell[3][0]=1; X.Cell[3][1]= 850; X.Cell[3][2]=2;X.Cell[3][3]=36
	//X.Cell[4][0]=1; X.Cell[4][1]=1300; X.Cell[4][2]=4;X.Cell[4][3]=50
	//
	//
	//
	//X.Hprint("features of houses (X):")
	//
	//a := LKmath.Determinant(LKmath.MatrixMultiplication(LKmath.TransposedMatrix(X), X))
	//fmt.Printf("%v\n", a)
	//
	//
	//y :=LKmath.NewEmptyMatrix(5,1)
	////--price of houses
	//y.Cell[0][0]=460
	//y.Cell[1][0]=232
	//y.Cell[2][0]=315
	//y.Cell[3][0]=178
	//y.Cell[4][0]=220
	//y.Hprint("price of each house are")
	//
	//
	//
	//example :=LKmath.NewEmptyMatrix(1, 4)
	//example.Cell[0][0] = 1; example.Cell[0][1] = 1000; example.Cell[0][2] = 4; example.Cell[0][3] = 0
	//example.Hprint("now I have a house with 1000 square feet, 4 bedrooms and is a brand new house")
	//
	//
	//
	//result := LKmath.NormalEquation(X, y)
	//result.Hprint("the result is :")
	//priceOfExampleHouse :=LKmath.MatrixMultiplication(example, result)
	//priceOfExampleHouse.Hprint("predicted price of example house is :")
	//
	//
	//
	//regularizedResult := LKmath.RegularizedNormalEquation(X, y, 0.001)
	//regularizedResult.Hprint("the regularized result is :")
	//priceOfRegularizedExampleHouse :=LKmath.MatrixMultiplication(example, regularizedResult)
	//priceOfRegularizedExampleHouse.Hprint("regularized predicted price of example house is :")


	//A := LKmath.NewEmptyMatrix(3, 4)
	//A.Hprint("original A")
	//A.Update(LKmath.NewIdentityMatrix(3))
	//A.Hprint("new A")
	//A.Update(LKmath.NewRandomMatrix(3,4,31,32))
	//A.Hprint("random A")
	//A.Update(LKmath.MatrixRandom(A, -16, 16))
	//A.Hprint("random A using matrixRandom")


/*----------------------------------------------test of gradient decent-----------------------------------------------*/
	//
	//startParameter := LKmath.NewEmptyMatrix(4, 1)
	//startParameter.Cell[0][0] = 0
	//startParameter.Cell[1][0] = 0
	//startParameter.Cell[2][0] = 0
	//startParameter.Cell[3][0] = 0
	//parameter :=LKmath.LinearRegressionGradientDecent(X, y, 0.0000005,startParameter, 10000000000)
	//parameter.Hprint("final result is: ")



	//supportMatrix := LKmath.NewEmptyMatrix(1, 4)
	//supportMatrix.Cell[0][0] = LKmath.NotChange; supportMatrix.Cell[0][1] = LKmath.UseMin; supportMatrix.Cell[0][2] = LKmath.UseAverage; supportMatrix.Cell[0][3] = LKmath.UseAverage
	////NormalStartParameter := LKmath.NewValuedMatrix(4, 1, 100)
	//
	//newParameter := LKmath.ScariedGradientDecent(X, y, 0.000005,supportMatrix, startParameter, 5000000 )
	//newParameter.Hprint("new parameter")
	//
	//
	//gradientDecentResult := LKmath.NewEmptyMatrix(4,1)
	//gradientDecentResult.Cell[0][0] = 169.782315
	//gradientDecentResult.Cell[1][0] =0.000109
	//gradientDecentResult.Cell[2][0] =7.068752
	//gradientDecentResult.Cell[3][0] =-0.071219
	//
	//testResult :=LKmath.MatrixMultiplication(X, parameter)
	//testResult.Hprint("test result is: ")

	//

/*-------------------------------------------test of logistic regression --------------------------------------------*/


//create data


	n := 10

	X:=LKmath.NewRandomMatrix(2,n, 0, 10)
	y:=LKmath.NewEmptyMatrix(1,n)

	for i :=0; i <n; i++{
		if 1*X.Cell[0][i] + 6* X.Cell[1][i] >35{
			y.Cell[0][i]=1
		}
	}

	X.Cell[0][0]=5
	X.Cell[1][0]=9.84

	X.Hprint("X is: ")
	y.Hprint("y is: ")

	w :=LKmath.NewEmptyMatrix(2, 1)
	w.Cell[0][0]=1
	w.Cell[1][0]=6

	Parameter :=LKmath.NodeParameter{W:w,B: -20}
	yHat :=LKmath.YHat(X, Parameter)
	yHat.Hprint("yHat is: ")

	lossFunction:=LKmath.LogisticRegressionLossFunctionForMatrix(yHat, y)
	lossFunction.Hprint("loss function result for each example are: ")

	lossFunctionDerivative :=LKmath.DerivativeOfLogisticRegressionLossFunctionForMatrix(yHat, y)
	lossFunctionDerivative.Hprint("loss function derivatives are: ")


	finalDerivative :=LKmath.FinalDerivativeOfLogisticRegressionForMatrix(yHat, y)
	finalDerivative.Hprint("final derivative is: ")

	fmt.Printf("cost function result is: %v\n", LKmath.Average(lossFunction) )
	fmt.Printf("average final derivative is :%v\n ", LKmath.Average(finalDerivative))


	//fmt.Printf("%v", math.Log(1))
	//
	//parameter :=LKmath.LogisticRegressionGradientDecent(X, y, 0.000005,startParameter,100000000)
	//parameter.Hprint("final parameter is ")

	//fmt.Printf("%v",LKmath.LogisticRegressionLossFunction(0,0))


















}
