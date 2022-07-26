package santas_elevator_part_one

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSantaElevatorPartOne(t *testing.T) {
	t.Run("Given a string with instructions should return correct 0 floor to stop", func(t *testing.T) {
		result := SantaElevatorPartOne("(())")
		assert.Equal(t, 0, result)
	})

	t.Run("Given a string with instructions should return correct 3rd floor to stop", func(t *testing.T) {
		result := SantaElevatorPartOne("()(((")
		assert.Equal(t, 3, result)
	})

	t.Run("Given a longer string with instructions should return correct 4th floor to stop", func(t *testing.T) {
		result := SantaElevatorPartOne("(()(()()()(()()((((()((()))((()((((()())))))))))")
		assert.Equal(t, 4, result)
	})

	t.Run("Given a longer string with instructions should return correct -1 floor to stop", func(t *testing.T) {
		result := SantaElevatorPartOne("()(()()()(()()((((()((()))((()((((()())))))))))))))")
		assert.Equal(t, -1, result)
	})
}
