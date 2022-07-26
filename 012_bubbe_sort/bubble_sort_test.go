package bubbe_sort

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	t.Run("Given an empty array should return an empty array", func(t *testing.T) {
		result := BubbleSort([]int{})
		assert.Equal(t, []int{}, result)
	})

	t.Run("Given an array with one element should return the same array", func(t *testing.T) {
		result := BubbleSort([]int{1})
		assert.Equal(t, []int{1}, result)
	})

	t.Run("Given an array with three elements should return the sorted array", func(t *testing.T) {
		result := BubbleSort([]int{3, 2, 1, 0, 4, 5, 6, 7, 8, 9})
		assert.Equal(t, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, result)
	})

	t.Run("Given an array with negative elements should return the sorted array", func(t *testing.T) {
		result := BubbleSort([]int{-3, 2, 3, 1, -2, -1})
		assert.Equal(t, []int{-3, -2, -1, 1, 2, 3}, result)
	})
}
