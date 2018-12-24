package example

//import (
//	"fmt"
//	"github.com/LuKuuu/Kun/LKmath"
//	"github.com/petar/GoMNIST"
//)
//
//func Handwriting_test()  {
//
//
//
//	m :=0
//	n:=0
//
//
//
//	rows, cols, imgs, err := GoMNIST.ReadImageFile("data/train-images-idx3-ubyte.gz")
//	if err!=nil{
//		panic(err)
//	}
//	//fmt.Printf("%d, %d, %d\n",rows, cols, cap(imgs))
//
//	m = cap(imgs)
//	n = rows * cols
//
//	labels, e := GoMNIST.ReadLabelFile("data/train-labels-idx1-ubyte.gz")
//	if e!=nil{
//		panic(e)
//	}
//
//	X :=LKmath.NewEmptyMatrix(n, m)
//
//	for i:=0;i<X.Column;i++{
//		for j:=0;j<X.Row;j++{
//			X.Cell[j][i]=float64(imgs[i][j])/256
//		}
//	}
//
//	//X.Hprint("X")
//
//
//	Y :=LKmath.NewEmptyMatrix(10, m)
//	for i:=0;i<m;i++{
//		Y.Cell[int(labels[i])][i]=1
//	}
//
//	//Y.Hprint("Y")
//
//	nna :=LKmath.NewNeuralNetworkAttribution(4)
//	nna.Cell[0][0]= float64(n)
//	nna.Cell[0][1]=800
//	nna.Cell[0][2]=50
//	nna.Cell[0][3]=100
//	nna.Cell[0][4]=40
//	nna.Cell[0][5]=10
//
//
//
//
//
//
//	hw :=LKmath.NewRandomNeuralNetwork(false, nna, 1, -1)
//	LKmath.SaveToJson("hwTDeep.json", &hw)
//
//
//	for{
//		for batch:=0;batch<59000;batch+=100{
//
//			a:=batch
//			b:=batch+100
//
//
//
//			XCut :=LKmath.CutMatrix(X,0,X.Row-1, a, b)
//			YCut :=LKmath.CutMatrix(Y,0,Y.Row-1, a, b)
//
//
//			//hw=LKmath.ReadFromJson("hw.json")
//
//
//			//hw.Hprint("hw")
//
//
//
//
//
//
//			//hw =LKmath.NeuralNetworkMiniBatchGradientDecent("hwT",X,Y,0.001, hw,1000)
//			hw = LKmath.NeuralNetworkGradientDecent("hwTDeep",XCut,YCut,0.05, hw,5)
//			fmt.Printf("\n")
//
//
//
//		}
//	}
//
//
//
//
//
//
//
//
//}
//
//
//func Test(){
//
//	m :=0
//	n:=0
//	a:=0
//	b:=0
//
//
//	rows, cols, imgs, err := GoMNIST.ReadImageFile("data/train-images-idx3-ubyte.gz")
//	if err!=nil{
//		panic(err)
//	}
//	//fmt.Printf("%d, %d, %d\n",rows, cols, cap(imgs))
//
//	m = cap(imgs)
//	n = rows * cols
//
//	labels, e := GoMNIST.ReadLabelFile("data/train-labels-idx1-ubyte.gz")
//	if e!=nil{
//		panic(e)
//	}
//
//	X :=LKmath.NewEmptyMatrix(n, m)
//
//	for i:=0;i<X.Column;i++{
//		for j:=0;j<X.Row;j++{
//			X.Cell[j][i]=float64(imgs[i][j])/256
//		}
//	}
//
//	//X.Hprint("X")
//
//
//	Y :=LKmath.NewEmptyMatrix(10, m)
//	for i:=0;i<m;i++{
//		Y.Cell[int(labels[i])][i]=1
//	}
//
//
//	X =LKmath.CutMatrix(X,0,X.Row-1, a, b)
//	Y =LKmath.CutMatrix(Y,0,Y.Row-1, a, b)
//
//	X.Hprint("X")
//
//	Y.Hprint("Y")
//
//
//	nna :=LKmath.NewNeuralNetworkAttribution(1)
//	nna.Cell[0][0]= float64(n)
//	nna.Cell[0][1]=100
//	nna.Cell[0][2]=10
//
//	hw :=LKmath.NewRandomNeuralNetwork(false, nna, 1, 0)
//	hw=LKmath.ReadFromJson("./data/neural_network_data/","hwT(2).json")
//
//
//	yHat, _ :=hw.ForwardPropagation(X)
//	yHat.Hprint("yHat")
//
//
//
//
//
//
//
//
//
//}
//
//
//
