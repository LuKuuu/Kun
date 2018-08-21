package main

import (
	"fmt"
	"github.com/LuKuuu/Kun/LKmath"
	"time"
)

//testing

func main() {

	fmt.Printf("%v\n", time.Now())

/*----------------------------------------test of cost/loss function--------------------------------------------------*/

	//a :=LKmath.NewEmptyMatrix(10,1)
	//b :=LKmath.NewCopyMatrix(a)
	//LKmath.MatrixRandom(b,100,-100)
	//
	//
	//a.Hprint()
	//b.Hprint()
	//
	//
	//OldCost, err := LKmath.OldCostFunction(&b, &a)
	//if err != nil{
	//	fmt.Printf("%v", err)
	//}else{
	//	fmt.Printf("the old cost function result between a and b is %v\n",OldCost)
	//
	//}
	//
	//a.MatrixSigmoid()
	//b.MatrixSigmoid()
	//
	//Cost, err := LKmath.CostFunction(&a, &b)
	//if err != nil{
	//	fmt.Printf("%v", err)
	//}else{
	//	fmt.Printf("the cost function result between sigmoid a and sigmiod b is %v\n",Cost)
	//
	//}


/*-------------------------------------test of matrix multiplication and transpose-------------------------------------*/
	//
	//
	//p :=LKmath.NewEmptyMatrix(3,2)
	//p.Cell[0][0]=1;	p.Cell[0][1]=2
	//p.Cell[1][0]=3; p.Cell[1][1]=4
	//p.Cell[2][0]=5; p.Cell[2][1]=6
	//
	//
	//q :=LKmath.NewEmptyMatrix(2,2)
	//q.Cell[0][0]=5;	q.Cell[0][1]=6
	//q.Cell[1][0]=7; q.Cell[1][1]=8
	//
	//k :=LKmath.MatrixMultiplication(p, q)
	//k.Hprint("K is:")
	//
	//z:=LKmath.TransposedMatrix(p)
	//p.Hprint("p is:")
	//z.Hprint("z is:")


/*-------------------------------------------test of normal equation--------------------------------------------*/

	//a real problem : find the relationship between features of house and house price

	X :=LKmath.NewEmptyMatrix(5,4)
	//features of houses
	//-----x0-------------size----------bedrooms-----------age-------
	X.Cell[0][0]=1;	X.Cell[0][1]=2104; X.Cell[0][2]=5;X.Cell[0][3]=45
	X.Cell[1][0]=1; X.Cell[1][1]=1416; X.Cell[1][2]=3;X.Cell[1][3]=40
	X.Cell[2][0]=1; X.Cell[2][1]=1534; X.Cell[2][2]=3;X.Cell[2][3]=30
	X.Cell[3][0]=1; X.Cell[3][1]= 850; X.Cell[3][2]=2;X.Cell[3][3]=36
	X.Cell[4][0]=1; X.Cell[4][1]=1300; X.Cell[4][2]=4;X.Cell[4][3]=50



	X.Hprint("features of houses (X):")

	y :=LKmath.NewEmptyMatrix(5,1)
	//--price of houses
	y.Cell[0][0]=460
	y.Cell[1][0]=232
	y.Cell[2][0]=315
	y.Cell[3][0]=178
	y.Cell[4][0]=220
	y.Hprint("price of each house are")



	result := LKmath.NormalEquation(X, y)
	result.Hprint("the result is :")


	example :=LKmath.NewEmptyMatrix(1, 4)
	example.Cell[0][0] = 1; example.Cell[0][1] = 1000; example.Cell[0][2] = 4; example.Cell[0][3] = 0
	example.Hprint("now I have a house with 1000 square feet, 4 bedrooms and is a brand new house")
	priceOfExampleHouse :=LKmath.MatrixMultiplication(example, result)
	priceOfExampleHouse.Hprint("predicted price of example house is :")



























}
