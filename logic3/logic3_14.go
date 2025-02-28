package logic3

func Logic3_14(n int) [][]int {
	// throw error if n is less than 2
	if n < 2 {
		panic("n must be greater than 1")
	}
	mat := CreateEmptyMatrix(n,n)
	// base case
	mat[0][0] = 1
	mat[1][0] = 1


	for i := 2; i < n*n; i++ {
		pr, pc := helper14(i-1, n)
		ppr, ppc := helper14(i-2, n)
		cr, cc := helper14(i, n)

		// for 2 ... n*n - 3
		if(i < n*n - 2){
		// fibonaci rules
			mat[cr][cc] = mat[pr][pc] + mat[ppr][ppc]
		} else {
			// for n*n - 2 and n*n - 1
			mat[cr][cc] = mat[pr][pc] + 2
		}
	}

	return mat

}

func helper14(i int, n int) (row int, col int){
	// mapping proper i into row and col
	// which i in range 0 to n*n

	// col is straightforward
	_col := i / n
	// how to find row??

	// find the modulo of i % 2*n
	var _row int
	modulo := i % (2*n)
	_row = map[bool]int{true: modulo, false: 2*n - modulo - 1}[modulo < n]

	return _row, _col

}