package rotate_a_list

func RotateAList(list []int, number int) []int {
	if len(list) == 0 {
		return list
	}
	for i := 0; i < number; i++ {
		list = append(list[1:], list[0])
	}
	return list
}
