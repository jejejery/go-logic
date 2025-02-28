package logic3

func Logic3_12(n int) [][]int {
	mat := CreateEmptyMatrix(n, n)

	// first pattern
	// upper triangle
	ubr := map[bool]int{true: n/2-1, false: n/2}[n%2 == 0] //upper bound row
	// num of cells in upper triangle -> 1  2 + .. + (ubr+1)
	// using sum of arithmetic series formula
	counter := 1
	for i := 0; i <= ubr; i++ {
		for j := 0; j <= i; j++ {
			if i % 2 == 0{
				mat[i][j] = counter
			} else {
				mat[i][i-j] = counter
			}
			counter += 2
		}

	}

	// reuse counter 
	counter = mat[map[bool]int{true: ubr, false: ubr-1}[n % 2 == 0]][0]
	// second triangle
	for i := ubr+1; i < n; i++ {
		for j := 0; j < n-i; j++ {
			// left decrement if i is odd
			if i % 2 == 1{
				mat[i][j] = counter
			} else{
				mat[i][n-1-j - (i)]= counter
			}
			counter -=2
		}
	}

	
	// second pattern

	// start from to
	// p-right corner and bottom-right corner
	urp := 0 // upper row pointer
	lrp := n - 1 // lower row pointer
	counter = 1 // reuse counter
	for k := n - 1; k > n/2; k--{
		mat[urp][k] = counter
		mat[lrp][k] = counter
		counter += 2
		urp++
		lrp--
	}
	return mat
}

// func ut_12(i int, ubr int) (row int, col int, val int, n int){ // for upper triangle
// 	// mapping proper i into row and col
// 	// create triangle oscillating pattern
	
// 	// f(0) -> 0,0,1
// 	// f(1) -> 1,1,3
// 	// f(2) -> 1,0,5
// 	// f(3) -> 2,0,7


	
// }
