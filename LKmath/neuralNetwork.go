package LKmath

import (
	"fmt"
	"math"
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

/*--------------------------------------------------neural network---------------------------------------------------------*/

/*
layerParameter[0] is the parameter between input layer and first hidden layer
layerParameter[x] is the parameter between hidden layer x and hidden layer x+1
*/
type NeuralNetwork struct {
	InputLayerNum  int
	OutputLayerNum int
	Attribution    Matrix
	HiddenLayerNum int
	LayerParameter []LayerParameter
	Name           string
}

//new neural network attribution will create a 1 * (hiddenLayerNum+2) matrix
//cell[0][0] is the number of data of input layer
//cell[0][x] is the number of node of (x)th hidden layer
//cell[0][hiddenLayerNo+1] is the number of data of output layer
func NewNeuralNetworkAttribution(hiddenLayerNo int) Matrix {
	return NewEmptyMatrix(1, hiddenLayerNo+2)
}

func NewRandomNeuralNetwork(attribution Matrix, max float64, min float64) NeuralNetwork {

	inputLayerNum := int(attribution.Cell[0][0])
	outputLayerNum := int(attribution.Cell[0][attribution.Column-1])

	Layers := []LayerParameter{}

	for i := 0; i < attribution.Column-1; i++ {
		Layers = append(Layers, NewRandomLayer(int(attribution.Cell[0][i]), int(attribution.Cell[0][i+1]), max, min))
	}

	nn := NeuralNetwork{
		Name:           time.Now().Format("2006-01-02 15:04:05"),
		InputLayerNum:  inputLayerNum,
		OutputLayerNum: outputLayerNum,
		Attribution:    attribution,
		HiddenLayerNum: attribution.Column - 2,
		LayerParameter: Layers,
	}

	return nn

}

func (this *NeuralNetwork) Hprint(info string) {
	fmt.Printf(info + "\n")
	for i, v := range this.LayerParameter {
		if i == 0 {
			v.Hprint(fmt.Sprintf("-----------------parameter between input layer and hidden layer %d----------------------------", i+1))
		} else if i == this.HiddenLayerNum {
			v.Hprint(fmt.Sprintf("-----------------parameter between hidden layer %d and output layer---------------------------", i))
		} else {
			v.Hprint(fmt.Sprintf("-----------------parameter between hidden layer %d and hidden layer %d----------------------------", i, i+1))
		}
	}
}

/*--------------------------------------------------forward propagation---------------------------------------------------------*/

func (this *NeuralNetwork) ForwardPropagation(inputMatrix Matrix) (outputMatrix Matrix, temp []Matrix) {

	temp = append(temp, this.LayerParameter[0].YHat(inputMatrix))
	for i := 0; i < this.HiddenLayerNum; i++ {
		temp = append(temp, this.LayerParameter[i+1].YHat(temp[i]))
	}

	return NewCopyMatrix(temp[this.HiddenLayerNum]), temp
}

/*--------------------------------backward propagation ------------------------------------------*/

func (this *NeuralNetwork) UpdateDerivative(X Matrix, y Matrix) (cost float64) {

	yHatMatrix, forwardTemp := this.ForwardPropagation(X)
	cost = LogisticRegressionCostFunction(yHatMatrix, y)

	backwardTempMatrix := Matrix{}
	for i := this.HiddenLayerNum; i >= 0; i-- {

		if i == this.HiddenLayerNum {
			backwardTempMatrix = this.LayerParameter[i].UpdateDerivative(forwardTemp[i-1], y)
		} else if i == 0 {
			this.LayerParameter[i].UpdateDerivativeWithDerivative(X, backwardTempMatrix)
		} else {
			backwardTempMatrix = this.LayerParameter[i].UpdateDerivativeWithDerivative(forwardTemp[i-1], backwardTempMatrix)
		}

	}

	return cost
}

/*-----------------------------------------------------gradient decent--------------------------------------------------*/

//X: data( if we have m examples and n features, X should be a (number of input layer) * m matrix)
//y: result (m examples means we should have m results so y should be a (number of output layer) * m matrix)
//NeuralNetwork a neural network
//alpha : learning rate (it should be carefully chose) [according to Ng, A, 0.01 is a good choice)
func (this *NeuralNetwork) GradientDecent(X Matrix, y Matrix, alpha float64, learningTimes int) {

	if X.Row != this.InputLayerNum || y.Row != this.OutputLayerNum || X.Column != y.Column {
		fmt.Printf("%v", this.Attribution)
		panic("format error")
	}

	//neuralNetworkData :=NewNeuralNetworkData()
	//neuralNetworkData.ConnectToDatabase("mysql", "root:cjkj@tcp(127.0.0.1:3306)/neural_network")
	//
	//fmt.Printf("start reading data from databse...\n")
	//NeuralNetwork,_ = neuralNetworkData.ReadFromDatabase(NeuralNetworkName, NeuralNetwork)

	//fmt.Printf("start reading data from JSON file...\n")
	//NeuralNetwork =ReadFromJson("./data/neural_network_data/",NeuralNetworkName+".json")

	fmt.Printf("start gradient decent of the neural network\n")
	t := time.Now()
	d := time.Now().Sub(t)
	oldCost := math.Inf(1)
	cost := math.Inf(1)

	for times := 0; times < learningTimes; times++ {
		cost = this.UpdateDerivative(X, y)
		for i := 0; i <= this.HiddenLayerNum; i++ {
			this.LayerParameter[i].W.Update(MatrixSubtraction(this.LayerParameter[i].W, ScalarMatrix(this.LayerParameter[i].DW, alpha)))
			this.LayerParameter[i].B.Update(MatrixSubtraction(this.LayerParameter[i].B, ScalarMatrix(this.LayerParameter[i].DB, alpha)))
		}

		if times%1==0{
			d = time.Now().Sub(t)
			t = time.Now()
			fmt.Printf(fmt.Sprintf("\nprogress : %f", float64(times*100)/float64(learningTimes)) + "%%\n")
			fmt.Printf("the dueration for each update is around %s\n", d.String())
			fmt.Printf("expect to get the next result on %s\n", t.Add(d).String())
			fmt.Printf("old cost is: %v,new cost is %v\n", oldCost, cost)

			if cost < oldCost {
				alpha *= 1.1
				this.SaveToJson(DefaultNeuralNetworkDirection)
			} else {
				fmt.Printf("Warning:cost is becoming bigger\n")
				fmt.Printf("learning rate will be changed to %f\n", alpha)
				return
			}
			oldCost = cost

			//this.Hprint("nn")

		}
		}



}

func CleanY(y Matrix) Matrix {

	zero := 0
	one := 0
	for i := 0; i < y.Row; i++ {
		for j := 0; j < y.Column; j++ {
			if y.Cell[i][j] > 0.5 {
				y.Cell[i][j] = 1
				one++
			} else {
				y.Cell[i][j] = 0
				zero++
			}

		}
	}

	fmt.Printf("there are %d zero and %d one\n", zero, one)
	return y

}
