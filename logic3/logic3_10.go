package logic3

func Logic3_10(n int) [][]int {
	if n < 2 {
		panic("n must be greater than 1")
	}
	createdMatrix := CreateEmptyMatrix(n, n)
	lowerBound := n/2 - 1 * Btoi(n%2 == 0)
	upperBound := n/2
	counter := n
	var _length int
	for i := 0; i < n; i++ {
		_length = upperBound - lowerBound + 1
		for j := 0; j < n; j++ {
			if j >= lowerBound && j <= upperBound {
				createdMatrix[i][j] = counter
				counter = counter - 2 * Btoi(j < (lowerBound + _length/2)) + 2 * Btoi(j >= (lowerBound + _length/2) || (_length / 2 + lowerBound - 1 == j && _length % 2 == 0))
			}
		}
		//update lowerBound and upperBound
		lowerBound = lowerBound - 1 * Btoi(i < n/2) + 1 * Btoi(i >= n/2 || (i + 1 == n/2 && n % 2 == 0)) 
		upperBound = upperBound + 1 * Btoi(i < n/2) - 1 * Btoi(i >= n/2 || (i + 1 == n/2 && n % 2 == 0)) 
		// restore the counter
		counter = n
	}
	return createdMatrix
}
