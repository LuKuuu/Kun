package LKmath

import (
	"math/rand"
	"time"
	"fmt"
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
	NodeDerivative  		[]NodeParameter

}

func NewRandomLayerParameter(initialize bool, PreviousLayerNo int, NextLayerNo int, PreviousLayerNum int, NextLayerNum int, max float64, min float64) LayerParameter {

	if initialize{
		rand.Seed(time.Now().Unix())
	}

	nd :=[]NodeParameter{}
	n :=[]NodeParameter{}

	for i:=0; i< NextLayerNum; i++{
		newNode :=NewRandomNode(initialize, PreviousLayerNum,max, min)
		newEmptyDerivative :=NewEmptyNode(PreviousLayerNum)
		n =append(n, newNode)
		nd = append(nd, newEmptyDerivative)

	}

	layer := LayerParameter{
		PreviousLayerNo:  		PreviousLayerNo,
		NextLayerNo:			NextLayerNo,
		PreviousLayerNum: 		PreviousLayerNum,
		NextLayerNum:    		NextLayerNum,
		NodeParameter:   		n,
		NodeDerivative: 	 	nd,

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


func (this *NeuralNetwork)ForwardPropagation(inputMatrix Matrix)(outputMatrix Matrix){

	temp :=[]Matrix{}
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

		temp[i].Hprint(fmt.Sprintf("temp %d", i))


	}

	return NewCopyMatrix(temp[this.HiddenLayerNum])
}


