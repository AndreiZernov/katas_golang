package nonogram_row

func NonogramRow(array []int) []int {
	var (
		result  = []int{}
		current int
	)

	for _, value := range array {
		switch {
		case value == 1:
			current++
		case current > 0:
			result = append(result, current)
			current = 0
		}
	}
	if current > 0 {
		result = append(result, current)
	}
	return result
}
