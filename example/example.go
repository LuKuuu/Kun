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
	n := 5
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

	Parameter := LKmath.NewEmptyNode(m)



	LKmath.LogisticRegressionGradientDecent(X, y, 0.01,Parameter, 100000000 )



//progress : 44.940000%
//	W:
//	1.565060	108.032650	6.853914	16.520635	3.933365	36.684508
//B: -91.193275
//dW:
//	-0.000001	-0.000079	-0.000005	-0.000012	-0.000003	-0.000027
//dB: 0.000067


}


func TestOfNeuralNetwork(){
	//nrlp:=LKmath.NewRandomLayerParameter(true, 2, 3, 5, 8,1,-1)
	//nrlp.Hprint("nrlp:")


	n :=100

	nna:=LKmath.NewNeuralNetworkAttribution(1)

	nna.Cell[0][0] = 5; nna.Cell[0][1] =2; 	nna.Cell[0][2] =2

	X :=LKmath.NewRandomMatrix(true, 5, n, 0,1)
	X.Hprint("X")
	Y :=LKmath.NewEmptyMatrix(2, n)

	enn :=LKmath.NewRandomNeuralNetwork(false, nna, 1,0)
	//enn.Hprint("enn:")

	enn.LayerParameter[0].NodeParameter[0].W.Cell[0][0]=1
	enn.LayerParameter[0].NodeParameter[0].W.Cell[0][1]=1
	enn.LayerParameter[0].NodeParameter[0].W.Cell[0][2]=4
	enn.LayerParameter[0].NodeParameter[0].W.Cell[0][3]=1
	enn.LayerParameter[0].NodeParameter[0].W.Cell[0][4]=-4
	enn.LayerParameter[0].NodeParameter[0].B = -1

	enn.LayerParameter[0].NodeParameter[1].W.Cell[0][0]=2
	enn.LayerParameter[0].NodeParameter[1].W.Cell[0][1]=8
	enn.LayerParameter[0].NodeParameter[1].W.Cell[0][2]=-10
	enn.LayerParameter[0].NodeParameter[1].W.Cell[0][3]=3
	enn.LayerParameter[0].NodeParameter[1].W.Cell[0][4]=6
	enn.LayerParameter[0].NodeParameter[1].B = -25


	enn.LayerParameter[1].NodeParameter[0].W.Cell[0][0]=100
	enn.LayerParameter[1].NodeParameter[0].W.Cell[0][1]=-30
	enn.LayerParameter[1].NodeParameter[0].B=-60

	enn.LayerParameter[1].NodeParameter[1].W.Cell[0][0]=-100
	enn.LayerParameter[1].NodeParameter[1].W.Cell[0][1]=30
	enn.LayerParameter[1].NodeParameter[1].B=60

	neuralNetworkData :=LKmath.NewNeuralNetworkData()
	neuralNetworkData.ConnectToDatabase("mysql", "root:cjkj@tcp(127.0.0.1:3306)/neural_network")

	//LKmath.SaveToJson("enn.json",&enn)

	enn = LKmath.ReadFromJson("./data/neural_network_data/","enn.json")

	//neuralNetworkData.Insert("enn",enn)

	Y, _ =enn.ForwardPropagation(X)

	Y = LKmath.CleanY(Y)

	Y.Hprint("Y")

	//copyedNN :=LKmath.NewRandomNeuralNetwork(false, nna, 100, -100)
	//copyedNN ,_= neuralNetworkData.ReadFromDatabase("enn", copyedNN)
	//copyedNN.Hprint("copyedNN")
	//
	//
	//neuralNetworkData.Insert("copiedNN",copyedNN)


	//nnn :=LKmath.NewRandomNeuralNetwork(false, nna, 1,0)

	//nnn,_ = neuralNetworkData.ReadFromDatabase("nnn", nnn)
	//nnn.LayerParameter[0].NodeParameter[0].W.Cell[0][0]=0.66
	//nnn.LayerParameter[0].NodeParameter[0].W.Cell[0][1]=1.68
	//nnn.LayerParameter[0].NodeParameter[0].W.Cell[0][2]=6.1
	//nnn.LayerParameter[0].NodeParameter[0].W.Cell[0][3]=2.1
	//nnn.LayerParameter[0].NodeParameter[0].W.Cell[0][4]=-1.8
	//nnn.LayerParameter[0].NodeParameter[0].B = -1.9
	//
	//nnn.LayerParameter[0].NodeParameter[1].W.Cell[0][0]=2.98
	//nnn.LayerParameter[0].NodeParameter[1].W.Cell[0][1]=3.97
	//nnn.LayerParameter[0].NodeParameter[1].W.Cell[0][2]=14.97
	//nnn.LayerParameter[0].NodeParameter[1].W.Cell[0][3]=2.98
	//nnn.LayerParameter[0].NodeParameter[1].W.Cell[0][4]=7.98
	//nnn.LayerParameter[0].NodeParameter[1].B = -10.12
	//
	//
	//nnn.LayerParameter[1].NodeParameter[0].W.Cell[0][0]=33
	//nnn.LayerParameter[1].NodeParameter[0].W.Cell[0][1]=-37
	//nnn.LayerParameter[1].NodeParameter[0].B=-4.9




	//nnn.Hprint("nnn before gradient decent")
	//
	//nnn =LKmath.NeuralNetworkGradientDecent("nnn", X,Y, 0.001, nnn,10000000000)
	//
	//nnn.Hprint("nnn after first gradient decent")


	//nna.Cell[0][1] =6
	//ntnn :=LKmath.NewRandomNeuralNetwork(false, nna, 1,0)
	//ntnn.Hprint("ntnn before gradient decent")


	//yHat, _ :=ntnn.ForwardPropagation(X)
	//yHat.Hprint("yHat with ntnn")
	//yHat = LKmath.CleanY(yHat)
	//yHat.Hprint("cleaned yHat with ntnn")


	//neuralNetworkData.Insert("ntnn", ntnn)

	NoBugNN :=LKmath.NewRandomNeuralNetwork(false, nna, 1, 0)
	//neuralNetworkData.Insert("dNoBugNN",NoBugNN)

	NoBugNN =LKmath.NeuralNetworkGradientDecent("dNoBugNN", X,Y, 0.001, NoBugNN,10000000000)




}