package logic3

func Logic3_03(n int) [][]int{
	mat := CreateEmptyMatrix(n, n)
	counter := 2
	for i := 0; i < n; i++{
		for j := 0; j < n - i; j++{
			// handle special case
			if i == 0 && j == 0{
				mat[i][j] = counter
				counter += 2
			} else if i == 0{ // special case for first row
				mat[i][j] = counter
				counter += 3
				//
			} else if i % 2 == 0 {// general case, increment by 2 for even row
				mat[i][j] = counter
				counter += 2
			} else { // general case, increment by 3 for odd row
				mat[i][n-1-j-i] = counter
				counter += 3
			}
			
		}
		// after the end of the row, should restore counter
		if i == 0{
			// counter = counter - 3 + 3
		} else if i % 2 == 0{ // switch between even and odd row
			counter = counter - 2 + 3
		} else {
			counter = counter - 3 + 2
		}
	}

	return mat
}