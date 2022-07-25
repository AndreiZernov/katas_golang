package game_of_three

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGameOfThree(t *testing.T) {
	t.Run("Given a number 3 should return correct calculation", func(t *testing.T) {
		number := 3
		result := GameOfThree(number)
		assert.Equal(t, []int{3, 1}, result)
	})
	t.Run("Given a number 100 should return correct calculation", func(t *testing.T) {
		number := 100
		result := GameOfThree(number)
		assert.Equal(t, []int{100, 99, 33, 11, 12, 4, 3, 1}, result)
	})
}
