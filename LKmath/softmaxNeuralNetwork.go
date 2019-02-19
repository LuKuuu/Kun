package LKmath

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)


/*
SoftMax Neural network is contributed by layer parameters
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
type SMNeuralNetwork struct {
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
func NewSMNeuralNetworkAttribution(hiddenLayerNo int) Matrix {
	return NewEmptyMatrix(1, hiddenLayerNo+2)
}

func NewRandomSMNeuralNetwork(attribution Matrix, max float64, min float64) SMNeuralNetwork {

	inputLayerNum := int(attribution.Cell[0][0])
	outputLayerNum := int(attribution.Cell[0][attribution.Column-1])

	Layers := []LayerParameter{}

	for i := 0; i < attribution.Column-1; i++ {
		Layers = append(Layers, NewRandomLayer(int(attribution.Cell[0][i]), int(attribution.Cell[0][i+1]), max, min))
	}

	nn := SMNeuralNetwork{
		Name:           time.Now().Format("2006-01-02 15:04:05"),
		InputLayerNum:  inputLayerNum,
		OutputLayerNum: outputLayerNum,
		Attribution:    attribution,
		HiddenLayerNum: attribution.Column - 2,
		LayerParameter: Layers,
	}

	return nn

}

func (this *SMNeuralNetwork) Hprint(info string) {
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

func (this *SMNeuralNetwork) ForwardPropagation(inputMatrix Matrix) (outputMatrix Matrix, temp []Matrix) {

	temp = append(temp, this.LayerParameter[0].YHat(inputMatrix))
	for i := 0; i < this.HiddenLayerNum; i++ {
		if i==this.HiddenLayerNum-1{
			temp = append(temp, this.LayerParameter[i+1].SoftMaxYHat(temp[i]))

		}else{
			temp = append(temp, this.LayerParameter[i+1].YHat(temp[i]))

		}
	}



	return NewCopyMatrix(temp[this.HiddenLayerNum]), temp
}

/*--------------------------------backward propagation ------------------------------------------*/

func (this *SMNeuralNetwork) UpdateDerivative(X Matrix, y Matrix) {



	yHat, forwardTemp := this.ForwardPropagation(X)

	backwardTempMatrix := Matrix{}
	for i := this.HiddenLayerNum; i >= 0; i-- {

		if i == this.HiddenLayerNum {
			backwardTempMatrix = this.LayerParameter[i].UpdateDerivativeForSoftMax(yHat,forwardTemp[i-1], y)
		} else if i == 0 {
			this.LayerParameter[i].UpdateDerivativeWithDerivative(X, backwardTempMatrix)
		} else {
			backwardTempMatrix = this.LayerParameter[i].UpdateDerivativeWithDerivative(forwardTemp[i-1], backwardTempMatrix)
		}

	}

}

func (this *SMNeuralNetwork) getCost(X Matrix, y Matrix) (cost float64) {

	yHatMatrix, _ := this.ForwardPropagation(X)
	return SoftMaxCostFunctionForMatrix(yHatMatrix, y)

}

/*-----------------------------------------------------gradient decent--------------------------------------------------*/

//X: data( if we have m examples and n features, X should be a (number of input layer) * m matrix)
//y: result (m examples means we should have m results so y should be a (number of output layer) * m matrix)
//SMNeuralNetwork a neural network
//alpha : learning rate (it should be carefully chose) [according to Ng, A, 0.01 is a good choice)
func (this *SMNeuralNetwork) GradientDecent(X Matrix, y Matrix, alpha float64, learningTimes int) float64{

	if X.Row != this.InputLayerNum || y.Row != this.OutputLayerNum || X.Column != y.Column {
		fmt.Printf("%v", this.Attribution)
		panic("format error")
	}

	//SMNeuralNetworkData :=NewSMNeuralNetworkData()
	//SMNeuralNetworkData.ConnectToDatabase("mysql", "root:cjkj@tcp(127.0.0.1:3306)/neural_network")
	//
	//fmt.Printf("start reading data from databse...\n")
	//SMNeuralNetwork,_ = SMNeuralNetworkData.ReadFromDatabase(SMNeuralNetworkName, SMNeuralNetwork)


	fmt.Printf("start gradient decent of the neural network\n")
	t := time.Now()
	d := time.Now().Sub(t)

	cost := this.getCost(X, y)
	oldCost :=cost


	for times := 0; times < learningTimes; times++ {


		this.UpdateDerivative(X, y)
		for i := 0; i <= this.HiddenLayerNum; i++ {
			this.LayerParameter[i].W.Update(MatrixSubtraction(this.LayerParameter[i].W, ScalarMatrix(this.LayerParameter[i].dW, alpha)))
			this.LayerParameter[i].B.Update(MatrixSubtraction(this.LayerParameter[i].B, ScalarMatrix(this.LayerParameter[i].dB, alpha)))
		}

		cost = this.getCost(X,y)

		d = time.Now().Sub(t)
		t = time.Now()
		fmt.Printf(fmt.Sprintf("\nprogress : %f", float64((times+1)*100)/float64(learningTimes)) + "%% \t\t ")
		fmt.Printf("learning rate is %f\n",alpha)
		fmt.Printf("the dueration for each update is around %s\n", d.String())
		fmt.Printf("expect to get the next result on %s\n", t.Add(d).String())
		fmt.Printf("old cost is: %v,new cost is %v\n", oldCost, cost)

		if cost <= oldCost {
			alpha *= 1.2
			oldCost=cost
		} else {
			fmt.Printf("Warning:cost is becoming bigger\n")
			times--
			for i := 0; i <= this.HiddenLayerNum; i++ {
				this.LayerParameter[i].W.Update(MatrixAddition(this.LayerParameter[i].W, ScalarMatrix(this.LayerParameter[i].dW, alpha)))
				this.LayerParameter[i].B.Update(MatrixAddition(this.LayerParameter[i].B, ScalarMatrix(this.LayerParameter[i].dB, alpha)))
			}

			alpha *= 0.1
			fmt.Printf("learning rate will be changed to %f\n", alpha)
		}



		this.UpdateDerivative(X, y)




	}

	this.SaveToJson(DefaultSMNeuralNetworkDirection)

	return alpha


}




func (this *SMNeuralNetwork)SaveToJson(dir string){

	this.saveToJsonAsTemp(dir)

	for ;;{
		err:=os.Remove(dir +this.Name+".json")
		if err!=nil{
			fmt.Printf("error:%v",err)
			time.Sleep(time.Second)
			break
		}else{
			break
		}

	}

	for ;;{
		err:=os.Rename(dir+this.Name+"(temp).json", dir+this.Name+".json")
		if err!=nil{
			fmt.Printf("error:%v",err)
			time.Sleep(time.Second)
		}else{
			break
		}

	}


}

const DefaultSMNeuralNetworkDirection ="data/SM_neural_network_data/"

func (this *SMNeuralNetwork) saveToJsonAsTemp(dir string){

	JSON, MarshalErr := json.MarshalIndent(&this, "", "\t")
	if MarshalErr !=nil{
		panic(MarshalErr)
	}

	fileAddress := dir +this.Name+"(temp).json"
	outputFile, outputError := os.OpenFile(fileAddress, os.O_CREATE, 0666)
	if outputError != nil {
		panic(outputError)
	}
	defer outputFile.Close()
	outputWriter := bufio.NewWriter(outputFile)
	outputWriter.WriteString(string(JSON))
	outputWriter.Flush()

	fmt.Printf("successfully save SMNeuralNetwork to %s\n",fileAddress)
}


func (this *SMNeuralNetwork)ReadFromJson(dir string,SMNeuralNetworkName string){

	fileAddress := dir+SMNeuralNetworkName+".json"
	data, err := ioutil.ReadFile(fileAddress)
	if err != nil {
		panic(err)
	} else {
		fmt.Printf("Loading data from %s\n", fileAddress)
	}

	unmarshalErr := json.Unmarshal(data, &this)
	if unmarshalErr != nil {
		panic(unmarshalErr)
	}

	fmt.Printf("successfully read neural network from %s\n",fileAddress)

}
