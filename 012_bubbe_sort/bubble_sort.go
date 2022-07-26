package bubbe_sort

func BubbleSort(numbers []int) []int {
	var (
		sorted  bool
		counter int
	)
	for !sorted {
		sorted = true
		for i := 0; i < len(numbers)-1; i++ {
			if numbers[i] > numbers[i+1] {
				numbers[i], numbers[i+1] = numbers[i+1], numbers[i]
				sorted = false
			}
		}
		counter++
	}
	return numbers
}
