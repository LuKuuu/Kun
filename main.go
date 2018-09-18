package main

import (
	"fmt"
	"time"
	"runtime"
	"github.com/LuKuuu/Kun/example"
)

//testing



func main() {

	fmt.Printf("%v\n", time.Now())


	example.TestOfLogisticRegression()
	//example.TestOfLinearEquation()
	//
	//nna:=LKmath.NewNeuralNetworkAttribution(3)
	//nna.Cell[0][0] =4; 	nna.Cell[0][1] =5; 	nna.Cell[0][2] =3
	//
	//
	//nnn :=LKmath.NewRandomNeuralNetwork(true, nna, 1,0)
	//nnn.Hprint("nnn")

}

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}
