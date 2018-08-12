package main

import (
	"fmt"
	"github.com/LuKuuu/Kun/LKmath"
	"time"
)

//testing

func main() {
	a :=LKmath.NewEmptyMatrix(7,7)
	LKmath.Hprint(a)

	fmt.Println()

	b :=LKmath.MatrixRandom(a,9,-9)
	LKmath.Hprint(b)

	fmt.Printf("%v", time.Now())

}
