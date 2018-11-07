package example

import (
	"fmt"
	"github.com/LuKuuu/Kun/LKmath"
	"github.com/petar/GoMNIST"
)

func Handwriting_test()  {

	m :=0
	n:=0



	rows, cols, imgs, err := GoMNIST.ReadImageFile("data/train-images-idx3-ubyte.gz")
	if err!=nil{
		panic(err)
	}
	fmt.Printf("%d, %d, %d\n",rows, cols, cap(imgs))

	m = cap(imgs)/100
	n = rows * cols

	labels, e := GoMNIST.ReadLabelFile("data/train-labels-idx1-ubyte.gz")
	if e!=nil{
		panic(e)
	}

	X :=LKmath.NewEmptyMatrix(n, m)

	for i:=0;i<X.Column;i++{
		for j:=0;j<X.Row;j++{
			X.Cell[j][i]=float64(imgs[i][j])/256
		}
	}

	//X.Hprint("X")


	Y :=LKmath.NewEmptyMatrix(10, m)
	for i:=0;i<m;i++{
		Y.Cell[int(labels[i])][i]=1
	}

	//Y.Hprint("Y")


	nna :=LKmath.NewNeuralNetworkAttribution(1)
	nna.Cell[0][0]= float64(n)
	nna.Cell[0][1]=100
	nna.Cell[0][2]=10
	//nna.Cell[0][3]=30
	//nna.Cell[0][4]=20
	//nna.Cell[0][5]=20
	//nna.Cell[0][6]=10

	hw :=LKmath.NewRandomNeuralNetwork(false, nna, 1, 0)

	//hw.Hprint("hw")

	//LKmath.SaveToJson("hw.json", &hw)






	neuralNetworkData :=LKmath.NewNeuralNetworkData()
	neuralNetworkData.ConnectToDatabase("mysql", "root:cjkj@tcp(127.0.0.1:3306)/neural_network")
	//neuralNetworkData.Insert("hw",hw)

	hw = LKmath.NeuralNetworkGradientDecent("hw",X,Y,0.1, hw,10000000)
	neuralNetworkData.Insert("hw", hw)




}