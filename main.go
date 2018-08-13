package main

import (
	"fmt"
	"time"
	"github.com/LuKuuu/Kun/LKmath"
)

//testing

func main() {

	fmt.Printf("%v\n", time.Now())

/*----------------------------------------test of cost/loss function--------------------------------------------------*/

	a,_ :=LKmath.NewEmptyMatrix(10,1)
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


/*-------------------------------------test of matrix multiplication and transpose-------------------------------------*/


	p, _ :=LKmath.NewEmptyMatrix(3,2)
	p.Cell[0][0]=1;	p.Cell[0][1]=2
	p.Cell[1][0]=3; p.Cell[1][1]=4
	p.Cell[2][0]=5; p.Cell[2][1]=6


	q, _ :=LKmath.NewEmptyMatrix(2,2)
	q.Cell[0][0]=5;	q.Cell[0][1]=6
	q.Cell[1][0]=7; q.Cell[1][1]=8

	k ,err:=LKmath.MatrixMultiplication(p, q)
	if err != nil {
		fmt.Printf("%v", err)
	}else{
		k.Hprint()
	}

	z:=LKmath.TransposedMatrix(p)
	p.Hprint()
	z.Hprint()


}
