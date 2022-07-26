package merge_list

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMergeList(t *testing.T) {
	t.Run("Given two lists should return a sorted merged list", func(t *testing.T) {
		list1 := []int64{1, 2, 3, 4, 5}
		list2 := []int64{6, 7, 8, 9, 10}
		expected := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		result := MergeList(list1, list2)
		assert.Equal(t, expected, result)
	})
	t.Run("Given two mixed lists should return a sorted merged list", func(t *testing.T) {
		list1 := []int64{1, 4, 6}
		list2 := []int64{2, 3, 5}
		expected := []int64{1, 2, 3, 4, 5, 6}
		result := MergeList(list1, list2)
		assert.Equal(t, expected, result)
	})
}
