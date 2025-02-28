package logic3

func Logic3_1(n int) [][]int {

	mat := CreateEmptyMatrix(n,n)
	upper_bound_idx := 0
	counter := 1
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if j >= upper_bound_idx {
				mat[i][j] = counter
				counter += 2
			}
		}
		upper_bound_idx += 1
	}
	return mat

}
