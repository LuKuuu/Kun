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


/*-------------------------------------test of matrix multiplication and transpose-------------------------------------*/


	p :=LKmath.NewEmptyMatrix(3,2)
	p.Cell[0][0]=1;	p.Cell[0][1]=2
	p.Cell[1][0]=3; p.Cell[1][1]=4
	p.Cell[2][0]=5; p.Cell[2][1]=6


	q :=LKmath.NewEmptyMatrix(2,2)
	q.Cell[0][0]=5;	q.Cell[0][1]=6
	q.Cell[1][0]=7; q.Cell[1][1]=8

	k :=LKmath.MatrixMultiplication(p, q)
	if err != nil {
		fmt.Printf("%v", err)
	}else{
		k.Hprint()
	}

	z:=LKmath.TransposedMatrix(p)
	p.Hprint()
	z.Hprint()

	X :=LKmath.NewEmptyMatrix(4,5)
	//features of houses
	//-----x0-------------size----------bedrooms-----------floors--------age----
	X.Cell[0][0]=1;	X.Cell[0][1]=2104; X.Cell[0][2]=5;X.Cell[0][3]=1;X.Cell[0][4]=45
	X.Cell[1][0]=1; X.Cell[1][1]=1416; X.Cell[1][2]=3;X.Cell[1][3]=2;X.Cell[1][4]=40
	X.Cell[2][0]=1; X.Cell[2][1]=1534; X.Cell[2][2]=3;X.Cell[2][3]=2;X.Cell[2][4]=30
	X.Cell[3][0]=1; X.Cell[3][1]= 852; X.Cell[3][2]=2;X.Cell[3][3]=1;X.Cell[3][4]=36

	X.Hprint()

	Xt:=LKmath.TransposedMatrix(X)

	Xt.Hprint()

	XtimesXt := LKmath.MatrixMultiplication(X, Xt)
	if err != nil{
		fmt.Printf("%v", err)
	}else{
		XtimesXt.Hprint()
	}


	y :=LKmath.NewEmptyMatrix(4,1)
	//--price of houses
	y.Cell[0][0]=460
	y.Cell[1][0]=232
	y.Cell[2][0]=315
	y.Cell[3][0]=178

	Xcut := LKmath.CutMatrix(X, 1, 3, 1, 4)
	Xcut.Hprint()








}
