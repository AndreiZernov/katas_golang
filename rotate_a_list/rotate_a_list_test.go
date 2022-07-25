package rotate_a_list

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRotate(t *testing.T) {
	t.Run("Given a list of numbers should rotate at one position", func(t *testing.T) {
		list := []int{1, 2, 3, 4, 5}
		result := Rotate(list, 1)
		assert.Equal(t, []int{2, 3, 4, 5, 1}, result)
	})
	t.Run("Given a list of numbers should rotate at two positions", func(t *testing.T) {
		list := []int{1, 2, 3, 4, 5, 6}
		result := Rotate(list, 2)
		assert.Equal(t, []int{3, 4, 5, 6, 1, 2}, result)
	})
	t.Run("Given a list of numbers should rotate at 8 positions", func(t *testing.T) {
		list := []int{1, 2, 3, 4, 5, 6}
		result := Rotate(list, 8)
		assert.Equal(t, []int{3, 4, 5, 6, 1, 2}, result)
	})
}
