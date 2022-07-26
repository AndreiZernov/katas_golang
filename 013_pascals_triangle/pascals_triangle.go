package pascals_triangle

import "fmt"

func PascalTriangle(n int) string {
	var result string
	for i := 1; i <= n; i++ {
		result += PascalRow(i) + "\n"
	}
	return result
}

func PascalRow(n int) string {
	var result string
	for i := 1; i <= n; i++ {
		result += fmt.Sprintf("%d ", Pascal(n, i))
	}
	return result
}

func Pascal(n, i int) int {
	if i == 1 || i == n {
		return 1
	}
	return Pascal(n-1, i-1) + Pascal(n-1, i)
}
