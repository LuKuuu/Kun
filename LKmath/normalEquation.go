package LKmath

/*
normal equation is suitable for calculating with less examples and features
it is much more faster than linear regression
 */

func NormalEquation(X Matrix, y Matrix)Matrix{
	//note: when the normal equation meets the problem of non-invertible matrix
	//		regularized normal equation can fix it ;-)
	//		just put a really small number like o.ooo1 into lambda

	if y.Row != X.Row || y.Column != 1{
		panic("value format error")
	}
	XT := TransposeMatrix(X)
	return MatrixMultiplication(MatrixMultiplication(InverseMatrix(MatrixMultiplication(XT, X)), XT), y)


}


//regularized normal equation here is used to prevent over-fitting
//by adding the parameter lambda, the function can be more smooth
func RegularizedNormalEquation(X Matrix, y Matrix, lambda float64)Matrix{

	//lambda here is usually a really small number

	if y.Row != X.Row || y.Column != 1{
		panic("value format error")
	}
	XT := TransposeMatrix(X)

	penaltyMatrix := NewIdentityMatrix(X.Column)
	penaltyMatrix.Cell[0][0] = 0

	penalty := ScalarMatrix(penaltyMatrix, lambda)


	return MatrixMultiplication(MatrixMultiplication(InverseMatrix(MatrixAddition(MatrixMultiplication(XT, X),penalty)), XT), y)

}

