package main

import (
	"encoding/json"
	"fmt"
	"github.com/LuKuuu/Kun/LKmath"
	"math/rand"
	"runtime"
	"time"
)



func main() {

	fmt.Printf("%v\n", time.Now())

	rand.Seed(time.Now().UnixNano())

	m:=LKmath.NewEmptyMatrix(3,4)
	s,_ :=json.Marshal(m)
	fmt.Printf("%s\n",s)

	b:=LKmath.NewEmptyMatrix(3,4)
	json.Unmarshal(s, &b)

	fmt.Printf("%d",cap(b.Cell))
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

	//example.Handwriting_test()
	//example.Test()

	//example.TestOfLogisticRegression()








}

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}
