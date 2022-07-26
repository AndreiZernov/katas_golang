package santas_elevator_part_one

func SantaElevatorPartOne(input string) int {
	var floor int
	for _, stairs := range input {
		if stairs == '(' {
			floor++
		} else if stairs == ')' {
			floor--
		}
	}
	return floor
}
