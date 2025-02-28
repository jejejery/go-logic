package logic4


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

func CreateEmptyStrMatrix(n, m int) [][]string {
	var mat [][]string
	for i := 0; i < n; i++ {
		var row []string
		for j:= 0; j < n; j++{
			row = append(row, "." )	
		}
		mat = append(mat, row)

	}
	return mat
}

func Btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}
