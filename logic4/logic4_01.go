package logic4

func Logic4_01(n int) [][]int{
	_n := n * (n+1) / 2

	mat := CreateEmptyMatrix(_n, _n)
	ctr := 1
	var pivot int
	for i := 0; i < n; i++ {
		pivot = (i+1)*i/2
		for j := 0; j < i + 1; j++ {
			for k := 0; k < i+1; k++ {
				if j % 2 == 0 {
					mat[pivot + j][pivot + k] = ctr
				} else {
					mat[pivot + j][pivot + i - k] = ctr
				}
				ctr+= 2
			}
		}
		ctr = 1 // restore
	}

	return mat
}