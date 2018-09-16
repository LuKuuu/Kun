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


	//example.TestOfLogisticRegression()
	example.TestOfLinearEquation()



}

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}
