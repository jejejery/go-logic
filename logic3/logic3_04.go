package logic3

func Logic3_04(n int) [][]int {

	mat := CreateEmptyMatrix(n, n)
	ctr := 1
	var _col int
	for i := 0; i < n; i++ {
		for j := n-1-i; j < n; j++ {
			if i % 2 == 0 {
				mat[i][j] = ctr
			} else {
				_col = 2*n-j-i-2
				mat[i][_col] = ctr
			}
			
			ctr+=2
		}
	}

	return mat

}