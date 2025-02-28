package logic3

import(
	// integer to string package
	"strconv"
)

func Logic3_13(n int) [][]string {
	if(n % 2 == 0){
		panic("n must be odd")
	}
	mat := CreateEmptyStrMatrix(n, n)
	var dec int
	var multiplier int
	var val string
	for i := 0; i < n; i++ {
		
		if i < n/2 {
			dec = i
		} else{
			dec = n - i - 1
		}
		start := n/2 - (dec)
		for j := start; j <= start + 2 * dec; j++ {
			if j < n/2 {
				multiplier = j
			} else{
				multiplier = n - j - 1
			}
			// check condition if row is odd/even
			if (i % 2 == 0 && j % 2 == 0) || (i % 2 == 1 && j % 2 == 1){
				val = strconv.Itoa(2*multiplier + 1)
			} else if (i % 2 == 0 && j % 2 == 1) || (i % 2 == 1 && j % 2 == 0) {
				val = "#"
			}

			mat[i][j] = val
		}
	}

	return mat
}
