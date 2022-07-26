package square_or_root

import "math"

func SquareOrRoot(list []int) []int {
	result := make([]int, len(list))
	for i, number := range list {
		square := math.Sqrt(float64(number))
		if square == math.Trunc(square) {
			result[i] = int(square)
		} else {
			result[i] = number * number
		}
	}
	return result
}
