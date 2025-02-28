package logic4

import "fmt"

func Logic4_03(n int) [][]int {
	_n := n * n // 1 + 2 + .. + (n-1) + n + (n-1) + ... + 2 + 1
	mat := CreateEmptyMatrix(_n, _n)
	start_idx := 0
	var _len int
	var tlr, tlc, trr, trc int
	// top left
	tlr = start_idx
	tlc = tlr
	// top right
	trr = 0
	trc = _n - 1
	// bottom left
	blr := _n - 1
	blc := 0
	// bottom right
	brr := _n - 1
	brc := _n - 1
		
	
	// interate through the matrix

	for k := 0; k < n-1; k++ {
		_len = k + 1 // length of submatrix
		
		
		
		// bl = _n - 1 - (start_idx + _len - 1)
		// br = _n - 1 - (start_idx + _len - 1)
		// fill submatrix
		fill_submatrix(mat, tlr, tlc, _len)
		fill_submatrix(mat, trr, trc, _len)
		fill_submatrix(mat, blr, blc, _len)
		fill_submatrix(mat, brr, brc, _len)


		// update start_idx
		start_idx += _len

		// update top left
		tlr = start_idx
		tlc = tlr
		// update top right
		trr = trr + _len
		trc = trc - (_len + 1)
		// update bottom left
		blr = blr - (_len + 1)
		blc = blc + _len
		// update bottom right
		brr = brr - (_len + 1)
		brc = brc - (_len + 1)

	}

	// finally, fill the center submatrix
	_len = n
	fill_submatrix2(mat, start_idx, _len)

	return mat
}

func fill_submatrix(mat [][]int, row_idx int, col_idx int, _len int) {
	counter := 0 // always reset counter
	for i := row_idx; i < row_idx+_len; i++ {
		for j := col_idx; j < col_idx+_len; j++ {
			if (i-row_idx)%2 == 0 {
				mat[i][j] = 1 + 2*counter
			} else {
				mat[i][col_idx+_len-1-(j-col_idx)] = 1 + 2*counter
			}
			counter += 1

		}
	}
}

func fill_submatrix2(mat [][]int, start_idx int, _len int) {
	counter := 0 // always reset counter
	var startRow int
	if start_idx > len(mat)/2 {
		// top right
		startRow = len(mat) - 1 - start_idx
		fmt.Println("startIndex", start_idx)
		fmt.Println("startRow", startRow)
	} else {
		// bottom left

		startRow = start_idx
	}
	for i := startRow; i < startRow+_len; i++ {
		for j := start_idx; j < start_idx+_len; j++ {
			if (i-start_idx)%2 == 0 {
				mat[i][j] = 1 + 2*counter
			} else {
				mat[i][start_idx+_len-1-(j-start_idx)] = 1 + 2*counter
			}
			counter += 1

		}
	}
}