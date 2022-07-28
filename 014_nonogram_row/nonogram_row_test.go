package nonogram_row

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNonogramRow(t *testing.T) {
	nonogramRowTest := []struct {
		Name   string
		Array  []int
		Result []int
	}{
		{
			Name:   "Given an array with no elements should return an empty array",
			Array:  []int{},
			Result: []int{},
		},
		{
			Name:   "Given an array with only 0s should return an empty array",
			Array:  []int{0, 0, 0, 0, 0},
			Result: []int{},
		},
		{
			Name:   "Given an array with only 1s should return an array with number of 1s",
			Array:  []int{1, 1, 1, 1, 1},
			Result: []int{5},
		},
		{
			Name:   "Given and array with 0s and 1s should return an array with number of consecutive 1s, length 2",
			Array:  []int{0, 1, 1, 1, 1, 1, 0, 1, 1, 1, 1},
			Result: []int{5, 4},
		},
		{
			Name:   "Given the array with 0s and 1s should return an array with number of consecutive 1s, length 3",
			Array:  []int{1, 1, 0, 1, 0, 0, 1, 1, 1, 0, 0},
			Result: []int{2, 1, 3},
		},
		{
			Name:   "Given the array with 0s and 1s, started with 0s should return an array with number of consecutive 1s, length 3",
			Array:  []int{0, 0, 0, 0, 1, 1, 0, 0, 1, 0, 1, 1, 1},
			Result: []int{2, 1, 3},
		},
		{
			Name:   "Given the array with 0s and 1s should return an array with number of consecutive 1s, length 8",
			Array:  []int{1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1},
			Result: []int{1, 1, 1, 1, 1, 1, 1, 1},
		},
	}

	for _, tt := range nonogramRowTest {
		t.Run(tt.Name, func(t *testing.T) {
			got := NonogramRow(tt.Array)

			assert.Equal(t, tt.Result, got)
		})
	}
}
