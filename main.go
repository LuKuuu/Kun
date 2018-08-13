package main

import (
	"fmt"
	"github.com/LuKuuu/Kun/LKmath"
	"time"
)

//testing

func main() {

	fmt.Printf("%v\n", time.Now())

	a :=LKmath.NewEmptyMatrix(10,1)
	b :=LKmath.NewCopyMatrix(a)
	LKmath.MatrixRandom(b,100,-100)


	a.Hprint()
	b.Hprint()


	OldCost, err := LKmath.OldCostFunction(&b, &a)
	if err != nil{
		fmt.Printf("%v", err)
	}else{
		fmt.Printf("the old cost function result between a and b is %v\n",OldCost)

	}

	a.MatrixSigmoid()
	b.MatrixSigmoid()

	Cost, err := LKmath.CostFunction(&a, &b)
	if err != nil{
		fmt.Printf("%v", err)
	}else{
		fmt.Printf("the cost function result between sigmoid a and sigmiod b is %v\n",Cost)

	}


}
