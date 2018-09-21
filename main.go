package main

import (
	"fmt"
	"github.com/LuKuuu/Kun/LKmath"
	"runtime"
	"time"
)

//testing



func main() {

	fmt.Printf("%v\n", time.Now())


	//example.TestOfNormalEquation()
	//example.TestOfLogisticRegression()
	//example.TestOfLinearEquation()
	//

	//nrlp:=LKmath.NewRandomLayerParameter(true, 2, 3, 5, 8,1,-1)
	//nrlp.Hprint("nrlp:")


	nna:=LKmath.NewNeuralNetworkAttribution(1)

	nna.Cell[0][0] =4; 	nna.Cell[0][1] =5; 	nna.Cell[0][2] =2


	nnn :=LKmath.NewRandomNeuralNetwork(false, nna, 1,-1)
	nnn.LayerParameter[1].NodeParameter[0].W.Cell[0][0]=10
	nnn.LayerParameter[1].NodeParameter[0].W.Cell[0][1]=10
	nnn.LayerParameter[1].NodeParameter[0].W.Cell[0][2]=10
	nnn.LayerParameter[1].NodeParameter[0].W.Cell[0][3]=10
	nnn.LayerParameter[1].NodeParameter[0].W.Cell[0][4]=10
	nnn.LayerParameter[1].NodeParameter[0].B = -20

	nnn.LayerParameter[1].NodeParameter[1].W.Cell[0][0]=-10
	nnn.LayerParameter[1].NodeParameter[1].W.Cell[0][1]=-10
	nnn.LayerParameter[1].NodeParameter[1].W.Cell[0][2]=-10
	nnn.LayerParameter[1].NodeParameter[1].W.Cell[0][3]=-10
	nnn.LayerParameter[1].NodeParameter[1].W.Cell[0][4]=-10
	nnn.LayerParameter[1].NodeParameter[1].B = 20


	nnn.Hprint("nnn")

	inputMatrix :=LKmath.NewRandomMatrix(true, 4, 1,-10,10)
	inputMatrix.Hprint("input matrix")

	result :=nnn.ForwardPropagation(inputMatrix)
	result.Hprint("result of forward propagation")

}

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}
