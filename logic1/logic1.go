package logic1

import(
	"strconv"
)

func Logic1_1(n int) []int {
	var arr []int
	for i := 0; i < n; i++ {
		arr = append(arr, (2*i + 1))
	}
	return arr
}

func Logic1_2(n int) []int {
	var arr []int
	for i := 0; i < n; i++ {
		arr = append(arr, (2* (i + 1)))
	}
	return arr
}

func Logic1_3(n int) []int {
	var arr []int
	for i := 0; i < n; i++ {
		arr = append(arr, (3 * (i + 1)))
	}
	return arr
}

func Logic1_4(n int) []int {
	var arr []int
	for i := 0; i < n; i++ {
		arr = append(arr, (2 * (n-1 - i) + 1))
	}
	return arr
}

func Logic1_5(n int) []int {
	var arr []int
	for i := 0; i < n; i++ {
		arr = append(arr, (2 * (n - i)))
	}
	return arr
}

func Logic1_6(n int) []int {
	var arr []int
	for i := 0; i < n; i++ {
		arr = append(arr, (3 * (n - i)))
	}
	return arr
}

func Logic1_7(n int) []int {
	var arr []int
	if (n % 2 == 0) {
	for i := 0; i < n/2; i++ {
		arr = append(arr, (2 * i + 1))
	}
	for i := n/2 ; i < n; i++ {
		arr = append(arr, (2 * (n - i) - 1))
	}
	} else {
		for i := 0; i <= n/2; i++ {
			arr = append(arr, (2 * i + 1))
		}
		for i := (n/2 + 1) ; i < n; i++ {
			arr = append(arr, (2 * (n - i) - 1))
		}
	}
	return arr
}

func Logic1_8(n int) []int {
	var arr []int
	if (n % 2 == 0) {
		for i := 0; i < n/2; i++ {
			arr = append(arr, (2 * (i + 1)))
		}
		for i := n/2 ; i < n; i++ {
			arr = append(arr, (2 * (n - i)))
		}
	} else {
		for i := 0; i <= n/2; i++ {
			arr = append(arr, (2 * (i + 1)))
		}
		for i := (n/2 + 1) ; i < n; i++ {
			arr = append(arr, (2 * (n - i)))
		}
	}
	return arr
}

func Logic1_9(n int) []int {
	var arr []int
	if (n % 2 == 0) {
		for i := 0; i < n/2; i++ {
			arr = append(arr, (3 * (i + 1)))
		}
		for i := n/2 ; i < n; i++ {
			arr = append(arr, (3 * (n - i)))
		}
	} else {
		for i := 0; i <= n/2; i++ {
			arr = append(arr, (3 * (i + 1)))
		}
		for i := (n/2 + 1) ; i < n; i++ {
			arr = append(arr, (3 * (n - i)))
		}
	}
	return arr
}

func Logic1_10(n int) []string{
	var arr []string
	value := 2
	for i := 0; i < n; i++ {
		if (i % 2 == 0) {
			arr = append(arr, strconv.Itoa(value))
			value *= 2
		} else {
			arr = append(arr, "fizz")
		}
	}
	return arr
}

func Logic1_11(n int) []string{
	var arr []string
	value := 1
	first := true
	for i := 0; i < n; i++ {
		if (i % 2 == 0) {
			arr = append(arr,"buzz")
		} else {
			arr = append(arr, strconv.Itoa(value))
			if (first) {
				value += 2
			} else {
				value *= 2
			}
		}
	}
	return arr
}

func Logic1_12(n int) []int{
	var arr []int
	for i := 0; i < n; i++ {
		arr = append(arr, (2 * (i % 4) + 1))
	}
	return arr
}