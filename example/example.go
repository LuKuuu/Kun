package example

import (
	"github.com/LuKuuu/Kun/LKmath"
)

func TestOfNormalEquation(){

	/*-------------------------------------------test of normal equation--------------------------------------------*/

	//a real problem : find the relationship between features of house and house price

	//features of houses
	X :=LKmath.NewEmptyMatrix(5,4)
	//-----x0-------------size----------bedrooms-----------age-------
	X.Cell[0][0]=1;	X.Cell[0][1]=2104; X.Cell[0][2]=5;X.Cell[0][3]=45
	X.Cell[1][0]=1; X.Cell[1][1]=1416; X.Cell[1][2]=3;X.Cell[1][3]=40
	X.Cell[2][0]=1; X.Cell[2][1]=1534; X.Cell[2][2]=3;X.Cell[2][3]=30
	X.Cell[3][0]=1; X.Cell[3][1]= 850; X.Cell[3][2]=2;X.Cell[3][3]=36
	X.Cell[4][0]=1; X.Cell[4][1]=1300; X.Cell[4][2]=4;X.Cell[4][3]=50
	X.Hprint("features of houses (X):")


	//price of houses
	y :=LKmath.NewEmptyMatrix(5,1)
	y.Cell[0][0]=460
	y.Cell[1][0]=232
	y.Cell[2][0]=315
	y.Cell[3][0]=178
	y.Cell[4][0]=220
	y.Hprint("price of each house are")



	example :=LKmath.NewEmptyMatrix(1, 4)
	example.Cell[0][0] = 1; example.Cell[0][1] = 1000; example.Cell[0][2] = 4; example.Cell[0][3] = 0
	example.Hprint("now I have a house with 1000 square feet, 4 bedrooms and is a brand new house")


	//result:
	result := LKmath.NormalEquation(X, y)
	result.Hprint("the result is :")
	priceOfExampleHouse :=LKmath.MatrixMultiplication(example, result)
	priceOfExampleHouse.Hprint("predicted price of example house is :")


	//regularized result:
	regularizedResult := LKmath.RegularizedNormalEquation(X, y, 0.001)
	regularizedResult.Hprint("the regularized result is :")
	priceOfRegularizedExampleHouse :=LKmath.MatrixMultiplication(example, regularizedResult)
	priceOfRegularizedExampleHouse.Hprint("regularized predicted price of example house is :")

}

func TestOfLinearEquation(){
	/*-------------------------------------------test of normal equation--------------------------------------------*/

	//a real problem : find the relationship between features of house and house price

	//features of houses
	X :=LKmath.NewEmptyMatrix(5,4)
	//-----x0-------------size----------bedrooms-----------age-------
	X.Cell[0][0]=1;	X.Cell[0][1]=2104; X.Cell[0][2]=5;X.Cell[0][3]=45
	X.Cell[1][0]=1; X.Cell[1][1]=1416; X.Cell[1][2]=3;X.Cell[1][3]=40
	X.Cell[2][0]=1; X.Cell[2][1]=1534; X.Cell[2][2]=3;X.Cell[2][3]=30
	X.Cell[3][0]=1; X.Cell[3][1]= 850; X.Cell[3][2]=2;X.Cell[3][3]=36
	X.Cell[4][0]=1; X.Cell[4][1]=1300; X.Cell[4][2]=4;X.Cell[4][3]=50
	X.Hprint("features of houses (X):")


	//--price of houses
	y :=LKmath.NewEmptyMatrix(5,1)
	y.Cell[0][0]=460
	y.Cell[1][0]=232
	y.Cell[2][0]=315
	y.Cell[3][0]=178
	y.Cell[4][0]=220
	y.Hprint("price of each house are")



	example :=LKmath.NewEmptyMatrix(1, 4)
	example.Cell[0][0] = 1; example.Cell[0][1] = 1000; example.Cell[0][2] = 4; example.Cell[0][3] = 0
	example.Hprint("now I have a house with 1000 square feet, 4 bedrooms and is a brand new house")


	startParameter := LKmath.NewEmptyMatrix(4, 1)
	startParameter.Cell[0][0] = 0
	startParameter.Cell[1][0] = 0
	startParameter.Cell[2][0] = 0
	startParameter.Cell[3][0] = 0
	parameter :=LKmath.LinearRegressionGradientDecent(X, y, 0.0000005,startParameter,0.001, 10000000000)
	parameter.Hprint("final result is: ")


	priceOfExampleHouse :=LKmath.MatrixMultiplication(example, parameter)
	priceOfExampleHouse.Hprint("regularized predicted price of example house is :")




}

func ExampleOfScaringLinearRegression(){
	//todo : fix this part

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

}

func TestOfLogisticRegression(){

	//create data with n examples
	n := 1000
	m := 6

	X:=LKmath.NewRandomMatrix(true,m,n, 0, 1)
	y:=LKmath.NewEmptyMatrix(1,n)

	for i :=0; i <n; i++{
		if 1*X.Cell[0][i] + 60* X.Cell[1][i] + 4*X.Cell[2][i] + 9* X.Cell[3][i] +3*X.Cell[4][i]+20*X.Cell[5][i] >51{
			y.Cell[0][i]=1
		}
	}


	//X.Hprint("X is: ")
	//y.Hprint("y is: ")

	w :=LKmath.NewEmptyMatrix(1, m)
	w.Cell[0][0]=0.108186; w.Cell[0][1]=20.228739; w.Cell[0][2]=0.993515; w.Cell[0][3]=2.894106; w.Cell[0][4]=0.353419; w.Cell[0][5]=6.341988
	Parameter :=LKmath.NodeParameter{W:w,B: -16.276119 }


	LKmath.LogisticRegressionGradientDecent(X, y, 0.0001,Parameter, 1000000 )

}