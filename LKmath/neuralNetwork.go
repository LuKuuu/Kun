package LKmath

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

/*
Neural network is contributed by layer parameters
a n hidden layers neural network should have n+1 layer parameters
the LayerParameter[0] is the parameter between input layer and first hidden layer
the LayerParameter[n} is the parameter between hidden layer n and output layer

in each layer
layerParameter are parameters for nodes on the next layer
LayerParameter[x], node parameter[y] is the parameter for the {y+1)th node of the next layer
*/

//layer parameter is the parameters between previous layer and the next layer

/*--------------------------------------------------LayerParameter---------------------------------------------------------*/


type LayerParameter struct{
	PreviousLayerNo			int					//0 means input layer
	NextLayerNo				int

	PreviousLayerNum		int
	NextLayerNum			int

	NodeParameter  			[]NodeParameter

}

func NewRandomLayerParameter(initialize bool, PreviousLayerNo int, NextLayerNo int, PreviousLayerNum int, NextLayerNum int, max float64, min float64) LayerParameter {

	if initialize{
		rand.Seed(time.Now().Unix())
	}

	n :=[]NodeParameter{}

	for i:=0; i< NextLayerNum; i++{
		newNode :=NewRandomNode(initialize, PreviousLayerNum,max, min)
		n =append(n, newNode)

	}

	layer := LayerParameter{
		PreviousLayerNo:  		PreviousLayerNo,
		NextLayerNo:			NextLayerNo,
		PreviousLayerNum: 		PreviousLayerNum,
		NextLayerNum:    		NextLayerNum,
		NodeParameter:   		n,

	}

	return layer
}

func (this *LayerParameter)Hprint(info string){
	fmt.Printf(info+"\n" )
	for i,v:=range this.NodeParameter {
		info :=fmt.Sprintf("parameter of node %d on layer %d has attribution of ", i+1, this.NextLayerNo)
		v.Hprint(info)
	}
}

/*--------------------------------------------------neural network---------------------------------------------------------*/



/*
layerParameter[0] is the parameter between input layer and first hidden layer
layerParameter[x] is the parameter between hidden layer x and hidden layer x+1
*/
type NeuralNetwork struct{
	InputLayerNum				int
	OutputLayerNum				int
	Attribution					Matrix
	HiddenLayerNum				int
	LayerParameter				[]LayerParameter

}

//new neural network attribution will create a 1 * (hiddenLayerNum+2) matrix
//cell[0][0] is the number of data of input layer
//cell[0][x] is the number of node of (x)th hidden layer
//cell[0][hiddenLayerNo+1] is the number of data of output layer
func NewNeuralNetworkAttribution(hiddenLayerNo int)Matrix{
	return NewEmptyMatrix(1, hiddenLayerNo+2)
}

func NewRandomNeuralNetwork(initialize bool, attribution Matrix, max float64, min float64)NeuralNetwork{
	if initialize{
		rand.Seed(time.Now().Unix())
	}

	inputLayerNum := int(attribution.Cell[0][0])
	outputLayerNum := int(attribution.Cell[0][attribution.Column-1])

	Layers :=[]LayerParameter{}

	for i :=1; i<attribution.Column; i++{
		Layers = append(Layers, NewRandomLayerParameter(initialize,i-1, i,int(attribution.Cell[0][i-1]),int(attribution.Cell[0][i]),max, min ))
	}

	nn :=NeuralNetwork{
		InputLayerNum:		inputLayerNum,
		OutputLayerNum:		outputLayerNum,
		Attribution:		attribution,
		HiddenLayerNum:		attribution.Column-2,
		LayerParameter:		Layers,
	}

	return nn


}

func (this *NeuralNetwork)Hprint(info string){
	fmt.Printf(info+"\n")
	for i, v:=range this.LayerParameter{
		if i == 0{
			v.Hprint(fmt.Sprintf("-----------------parameter of nodes between input layer and hidden layer %d----------------------------", i+1))
		}else if i == this.HiddenLayerNum{
			v.Hprint(fmt.Sprintf("-----------------parameter of nodes between hidden layer %d and output layer---------------------------", i))
		}else{
			v.Hprint(fmt.Sprintf("-----------------parameter of nodes between hidden layer %d and hidden layer %d----------------------------", i, i+1))
		}
	}
}

/*--------------------------------------------------forward propagation---------------------------------------------------------*/


func (this *NeuralNetwork)ForwardPropagation(inputMatrix Matrix)(outputMatrix Matrix, temp []Matrix){

	for i :=0; i<=this.HiddenLayerNum; i++{
		temp =append(temp, NewEmptyMatrix(int(this.Attribution.Cell[0][i+1]), inputMatrix.Column))
	}

	for i :=0; i<=this.HiddenLayerNum; i++{

		for j:=0; j<this.LayerParameter[i].NextLayerNum;j++{

			for exampleNo :=0; exampleNo < inputMatrix.Column; exampleNo++{

				if i == 0{
					temp[i].Cell[j][exampleNo] = YHat(inputMatrix,this.LayerParameter[0].NodeParameter[j]).Cell[0][exampleNo]

				}else{
					temp[i].Cell[j][exampleNo] = YHat(temp[i-1],this.LayerParameter[i].NodeParameter[j]).Cell[0][exampleNo]

				}

			}
		}

		//temp[i].Hprint(fmt.Sprintf("temp %d", i))


	}

	return NewCopyMatrix(temp[this.HiddenLayerNum]), temp
}


/*--------------------------------backward propagation ------------------------------------------*/

//calculate derivative of one single layer

func UpdateFinalLayerDerivative(temp Matrix, y Matrix, nn NeuralNetwork)(NeuralNetwork,Matrix){

	backTemp :=NewEmptyMatrix(nn.OutputLayerNum,y.Column)
	for i :=0; i <  nn.OutputLayerNum; i++{
		cm :=KeepOneRow(y, i)
		d , backTempRow:=DerivativeOfLogisticalRegressionCostFunction(temp,cm , nn.LayerParameter[nn.HiddenLayerNum].NodeParameter[i])
		nn.LayerParameter[nn.HiddenLayerNum].NodeParameter[i].DB = d.DB
		nn.LayerParameter[nn.HiddenLayerNum].NodeParameter[i].DW.Update(d.DW)
		for j:=0;j<backTemp.Column;j++{
			backTemp.Cell[i][j]=backTempRow.Cell[0][j]
		}
	}

	return nn, backTemp
}

func UpdateHiddenLayerDerivative(layerNo int, temp Matrix, backTemp Matrix, nn NeuralNetwork)(NeuralNetwork, Matrix){

	parameterTemp :=NewEmptyMatrix(nn.LayerParameter[layerNo+1].NextLayerNum, nn.LayerParameter[layerNo].NextLayerNum)

	for i:=0; i< parameterTemp.Row; i++ {
		for j :=0; j< parameterTemp.Column; j++{
			parameterTemp.Cell[i][j] = nn.LayerParameter[layerNo+1].NodeParameter[i].W.Cell[0][j]
		}
	}

	parameterTemp = ScalarMatrix(MatrixMultiplication(TransposeMatrix(parameterTemp), backTemp),1/float64(nn.LayerParameter[layerNo+1].NextLayerNum))
	NewBackTemp :=NewEmptyMatrix(nn.LayerParameter[layerNo].NextLayerNum,backTemp.Column )

	for i :=0; i<nn.LayerParameter[layerNo].NextLayerNum; i++ {

		yHatMatrix :=YHat(temp, nn.LayerParameter[layerNo].NodeParameter[i])
		dl := DerivativeOfSigmoidFunctionForMatrix(yHatMatrix)

		finalDerivative :=NewEmptyMatrix(1, dl.Column)
		for j :=0; j < dl.Column; j++{
			finalDerivative.Cell[0][j] = parameterTemp.Cell[i][j] * dl.Cell[0][j]
			NewBackTemp.Cell[i][j] = parameterTemp.Cell[i][j] * dl.Cell[0][j]
			//to prevent the problem of not a number
			if math.IsNaN(finalDerivative.Cell[0][j]){
				finalDerivative.Cell[0][j] = math.MaxFloat64
			}
		}

		nn.LayerParameter[layerNo].NodeParameter[i].DB =Average(finalDerivative)
		nn.LayerParameter[layerNo].NodeParameter[i].DW.Update(ScalarMatrix(MatrixMultiplication(finalDerivative, TransposeMatrix(temp)), 1/float64(temp.Column)))

	}


	return nn, backTemp


}

/*-----------------------------------------------------gradient decent--------------------------------------------------*/


func NeuralNetworkGradientDecent(NeuralNetworkName string, X Matrix, y Matrix, alpha float64, NeuralNetwork NeuralNetwork, learningTimes int)NeuralNetwork{

	//X: data( if we have m examples and n features, X should be a (number of input layer) * m matrix)
	//y: result (m examples means we should have m results so y should be a (number of output layer) * m matrix)
	//NeuralNetwork a neural network
	//alpha : learning rate (it should be carefully chose) [according to Ng, A, 0.01 is a good choice)

	if X.Row != NeuralNetwork.InputLayerNum || y.Row != NeuralNetwork.OutputLayerNum || X.Column != y.Column{
		panic("format error")
	}

	neuralNetworkData :=NewNeuralNetworkData()
	neuralNetworkData.ConnectToDatabase("mysql", "root:cjkj@tcp(127.0.0.1:3306)/neural_network")

	fmt.Printf("start reading data from databse...\n")
	NeuralNetwork,_ = neuralNetworkData.ReadFromDatabase(NeuralNetworkName, NeuralNetwork)

	NeuralNetwork.Hprint(NeuralNetworkName +" before gradient decent")

	fmt.Printf("start gradient decent of the neural network\n")


	times :=0
	cost :=math.Inf(1)
	backTemp :=Matrix{}

	for{
		times ++


		yHat, temp := NeuralNetwork.ForwardPropagation(X)



		if times%5001 == 0{
			NeuralNetwork.Hprint(fmt.Sprintf("\nprogress : %f", float64(times*100)/float64(learningTimes))+"%%")
			yHat.Hprint("yHat")
			fmt.Printf("cost is %v\n", LogisticRegressionCostFunction(yHat,y))

			if cost >= LogisticRegressionCostFunction(yHat,y){
				cost =LogisticRegressionCostFunction(yHat,y)
			}else{
				fmt.Printf("before, cost is %v, while cost at present is %v, cost is becoming bigger\n",cost,LogisticRegressionCostFunction(yHat,y))
			}

			neuralNetworkData.Insert(NeuralNetworkName, NeuralNetwork)
			fmt.Printf("saved to MySQL successfully!\n")


		}

		if times > learningTimes{
			break
		}


		NeuralNetwork,backTemp  = UpdateFinalLayerDerivative(temp[NeuralNetwork.HiddenLayerNum-1], y, NeuralNetwork)

		for j :=NeuralNetwork.HiddenLayerNum-1; j>=0;j--{
			if j ==0{
				NeuralNetwork, backTemp = UpdateHiddenLayerDerivative(0, X, backTemp, NeuralNetwork)
			}else{
				NeuralNetwork, backTemp = UpdateHiddenLayerDerivative(j, temp[j], backTemp, NeuralNetwork)

			}
		}

		for j:=0; j< NeuralNetwork.HiddenLayerNum+1;j++{
			for k :=0; k<NeuralNetwork.LayerParameter[j].NextLayerNum;k++{
				NeuralNetwork.LayerParameter[j].NodeParameter[k].W.Update(MatrixSubtraction(NeuralNetwork.LayerParameter[j].NodeParameter[k].W, ScalarMatrix(NeuralNetwork.LayerParameter[j].NodeParameter[k].DW,alpha)))
				NeuralNetwork.LayerParameter[j].NodeParameter[k].B -= NeuralNetwork.LayerParameter[j].NodeParameter[k].DB * alpha
			}
		}

	}
	return NeuralNetwork
}


func CleanY(y Matrix)Matrix{

	zero :=0
	one :=0
	for i:=0; i<y.Row;i++{
		for j:=0; j<y.Column;j++{
			if 	y.Cell[i][j]>0.5{
				y.Cell[i][j]=1
				one++
			}else{
				y.Cell[i][j]=0
				zero++
			}

		}
	}

	fmt.Printf("there are %d zero and %d one\n",zero, one)
	return y

}