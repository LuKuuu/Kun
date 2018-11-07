package main

import (
	"fmt"
	"github.com/LuKuuu/Kun/example"
	"runtime"
	"time"
)

/*
important : to avoid the problem of not a number and to for better gradient decent, it would be better to set all parameter greater than 0!
 */


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





}

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}
