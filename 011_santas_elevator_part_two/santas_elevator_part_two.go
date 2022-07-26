package santas_elevator_part_two

import "errors"

var NeverInterredBasement = errors.New("never entered basement")

func SantaElevatorPartTwo(instructions string) (int, error) {
	var floor int
	for i := 0; i < len(instructions); i++ {
		switch {
		case instructions[i] == '(':
			floor++
		case instructions[i] == ')':
			floor--
		}
		if floor == -1 {
			return i + 1, nil

		}
	}
	return floor, NeverInterredBasement
}
