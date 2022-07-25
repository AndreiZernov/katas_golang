package rotate_a_list

func Rotate(list []int, number int) []int {
	for i := 0; i < number; i++ {
		list = append(list[1:], list[0])
	}
	return list
}
