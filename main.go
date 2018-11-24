package main

import (
	"fmt"
	"github.com/LuKuuu/Kun/example"
	"runtime"
	"time"
)



func main() {

	fmt.Printf("%v\n", time.Now())

	//neuralNetworkData :=io.NewNeuralNetworkData()
	//neuralNetworkData.ConnectToDatabase("mysql", "root:cjkj@tcp(127.0.0.1:3306)/heart")
	//




	//example.TestOfNormalEquation()
	//example.TestOfLogisticRegression()
	//example.TestOfLinearEquation()
	//example.TestOfNeuralNetwork()
	//
	//m :=LKmath.NewRandomMatrix(true, 3,3,0,1)
	//m.Hprint("m")
	//
	//sm :=LKmath.ScalarMatrix(m, 3)
	//sm.Hprint("sm")
	//
	//n :=LKmath.SqueezedAverageRowMatrix(m)
	//n.Hprint("n")

	example.Handwriting_test()
	//example.Test()







}

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}
