package LKmath

import (
	"math/rand"
	"time"
	"fmt"
)

type Layer struct{
	FeatureNum			int
	NodeNum				int
	Node				[]NodeParameter
	NodeDerivative		[]NodeParameter

}

func NewRandomLayer(initialize bool, featureNum int, NodeNum int, max float64, min float64)Layer{

	if initialize{
		rand.Seed(time.Now().Unix())
	}

	n :=[]NodeParameter{}
	nd :=[]NodeParameter{}

	for i:=0; i< NodeNum; i++{
		newNode :=NewRandomNode(initialize,featureNum,max, min)
		newEmptyDerivative :=NewEmptyNode(featureNum)
		n =append(n, newNode)
		nd = append(nd, newEmptyDerivative)

	}

	layer := Layer{
		FeatureNum:featureNum,
		NodeNum:NodeNum,
		Node:n,
		NodeDerivative:nd,

	}

	return layer
}

func (this *Layer)Hprint(LayerNo int, info string){
	fmt.Printf(info+"\n" )
	for i,v:=range this.Node{
		info :=fmt.Sprintf("layer %d, node %d has attribution of ", LayerNo, i)
		v.Hprint(info)
	}
}



type NeuralNetwork struct{
	Attribution		Matrix
	LayerNum		int
	Layer			[]Layer

}

func NewNeuralNetworkAttribution(layerNo int)Matrix{
	return NewEmptyMatrix(1, layerNo)
}

func NewRandomNeuralNetwork(initialize bool, attribution Matrix, max float64, min float64)NeuralNetwork{
	if initialize{
		rand.Seed(time.Now().Unix())
	}
	Layers :=[]Layer{}

	for i :=1; i<attribution.Column; i++{
		Layers = append(Layers, NewRandomLayer(initialize,int(attribution.Cell[0][i-1]),int(attribution.Cell[0][i]),max, min ))
	}

	nn :=NeuralNetwork{
		Attribution:attribution,
		LayerNum:attribution.Column,
		Layer:Layers,
	}

	return nn


}

func (this *NeuralNetwork)Hprint(info string){
	fmt.Printf(info+"\n")
	for i, v:=range this.Layer{
		v.Hprint(i+1,"-----------------a new layer----------------------------")
	}
}




