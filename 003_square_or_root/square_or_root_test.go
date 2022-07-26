package _03_square_or_root

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSquareOrRoot(t *testing.T) {
	t.Run("Given a list of numbers should return correct calculation", func(t *testing.T) {
		list := []int{4, 3, 9, 7, 2, 1}
		result := SquareOrRoot(list)
		assert.Equal(t, []int{2, 9, 3, 49, 4, 1}, result)
	})
	t.Run("Given an another list of numbers should return correct calculation", func(t *testing.T) {
		list := []int{1, 4, 9, 5, 4, 1}
		result := SquareOrRoot(list)
		assert.Equal(t, []int{1, 2, 3, 25, 2, 1}, result)
	})
}
