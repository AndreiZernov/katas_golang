package square_or_root

import "math"

func SquareOrRoot(list []int) []int {
	var result []int
	for _, number := range list {
		if math.Sqrt(float64(number)) == math.Trunc(math.Sqrt(float64(number))) {
			result = append(result, int(math.Sqrt(float64(number))))
		} else {
			result = append(result, number*number)
		}
	}
	return result
}
