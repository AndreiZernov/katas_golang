package game_of_three

func GameOfThree(number int) []int {
	var result []int
	for number != 1 {
		result = append(result, number)
		switch {
		case number%3 == 0:
			number /= 3
		case number%3 == 1:
			number--
		case number%3 == 2:
			number++
		}
	}
	return append(result, 1)
}
