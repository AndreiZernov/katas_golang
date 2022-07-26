package santas_elevator_part_one

func SantaElevatorPartOne(instructions string) int {
	var floor int
	for _, stairs := range instructions {
		if stairs == '(' {
			floor++
		} else if stairs == ')' {
			floor--
		}
	}
	return floor
}
