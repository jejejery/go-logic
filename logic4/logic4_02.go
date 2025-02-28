package logic4

func Logic4_02(n int) [][]int{
	_n := n * (n+1) / 2

	mat := CreateEmptyMatrix(_n, _n)
	ctr := 1
	var pivot_row, pivot_col int

	for i := 0; i < n; i++ {
		pivot_row = _n - (i+1)*(i+2)/2
		pivot_col = (i+1)*i/2
		for j := pivot_row; j < pivot_row + i + 1; j++ {
			for k := pivot_col; k < pivot_col + i + 1; k++ {
				if (j-pivot_row) % 2 == 0 {
					mat[j][k] = ctr
				} else {
					mat[j][(i + 2 * pivot_col - k)] = ctr
				}
				ctr+= 2
			}
		}
		ctr = 1 // restore
		
	}

	return mat
}