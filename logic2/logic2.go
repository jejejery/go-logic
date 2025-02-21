package logic2

func CreateEmptyMatrix(n, m int) [][]int {
	var mat [][]int
	for i := 0; i < n; i++ {
		var row []int
		for j := 0; j < m; j++ {
			row = append(row, 0)
		}
		mat = append(mat, row)
	}
	return mat
}

func Logic2_1(n int) [][]int {
	mat := CreateEmptyMatrix(n,n)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			mat[i][j] = 2*j + 1
		}
	}
	return mat
}

func Logic2_2(n int) [][]int {
	mat := CreateEmptyMatrix(n,n)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			mat[i][j] = 2*(j+1) 
		}
	}
	return mat
}

func Logic2_3(n int) [][]int {
	mat := CreateEmptyMatrix(n,n)
	counter := 1
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			mat[i][j] = counter
			counter += 2
		}
	}
	return mat
}

func Logic2_4(n int) [][]int {
	mat := CreateEmptyMatrix(n,n)
	counter := 1
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			mat[i][j] = counter
			counter += 3
		}
	}
	return mat
}

func Logic2_5(n int) [][]int {
	mat := CreateEmptyMatrix(n,n)
	counter := 1
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if(i % 2 == 0){
			mat[i][j] = counter
			} else{
				mat[i][n - 1 - j] = counter
			}
			counter += 2
		}
	}
	return mat
}

func Logic2_6(n int) [][]int {
mat := CreateEmptyMatrix(n, n)
counter := 1
for i := 0; i < n; i++ {
	for j := 0; j < n; j++{
		if i % 2 == 0 { // forward
			mat[i][j] = counter
			counter += 3
		} else{ // backward
			mat[i][n - 1 - j] = counter
			counter += 2
		}
	}
}
return mat
}

func Logic2_7(n int) [][]int {
	mat := CreateEmptyMatrix(n, n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				mat[i][j] = 2*i + 1
			}
		}
	}
	
	return mat
}

func Logic2_8(n int) [][]int {
	mat := CreateEmptyMatrix(n, n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if (n - 1 - i) == j {
				mat[i][j] = 2*i + 1
			}
		}
	}
	
	return mat
}
func Logic2_9(n int) [][] int{
	mat := CreateEmptyMatrix(n, n)

	for i := 0; i < n; i++ {

		for j := 0; j < n; j++ {
			// make sure between the bounds
			if i == j || i + j == n - 1 {
				mat[i][j] = 2*j + 1
			}
		}
	}
	
	return mat
}



func Logic2_10(n int) [][]int {
	mat := CreateEmptyMatrix(n, n)
	row_counter := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if(j <= row_counter){
				mat[i][j] = 2*j + 1
			}
		}
		row_counter++
	}

	return mat
}

func Logic2_11(n int) [][]int {
	mat := CreateEmptyMatrix(n, n)
	row_counter := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if(j >= row_counter){
				mat[i][j] = 2*j + 1
			}
		}
		row_counter++
	}

	return mat
}


func Logic2_12(n int) [][] int {
	mat := CreateEmptyMatrix(n, n)
	lower_bound := -1
	upper_bound := n
	mid := n / 2
	for i := 0; i < n; i++ {
		if(i <= mid){
			lower_bound++
			upper_bound--
		} else {
			lower_bound--
			upper_bound++
		}
		for j := 0; j < n; j++ {
			// make sure between the bounds
			if j <= lower_bound || j >= upper_bound {
				mat[i][j] = 2*j + 1
			}
		}
	}
	
	return mat
}

func Logic2_13(n int) [][]int {
	mat := CreateEmptyMatrix(n, n)
	lower_bound := -1
	upper_bound := n
	mid := n / 2
	for i := 0; i < n; i++ {
		if(i <= mid){
			lower_bound++
			upper_bound--
		} else {
			lower_bound--
			upper_bound++
		}
		for j := 0; j < n; j++ {
			// make sure between the bounds
			if j >= lower_bound && j <= upper_bound {
				mat[i][j] = 2*j + 1
			}
		}
	}
	
	return mat
}