package _02_rotate_a_list

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRotateAList(t *testing.T) {
	adderTest := []struct {
		Name     string
		List     []int
		Number   int
		Expected []int
	}{
		{
			Name:     "Given a list of numbers should rotate at one position",
			List:     []int{1, 2, 3, 4, 5, 6},
			Number:   1,
			Expected: []int{2, 3, 4, 5, 6, 1},
		},
		{
			Name:     "Given a list of numbers should rotate at two positions",
			List:     []int{1, 2, 3, 4, 5, 6},
			Number:   2,
			Expected: []int{3, 4, 5, 6, 1, 2},
		},
		{
			Name:     "Given a list of numbers should rotate at 8 positions",
			List:     []int{1, 2, 3, 4, 5, 6},
			Number:   8,
			Expected: []int{3, 4, 5, 6, 1, 2},
		},
		{
			Name:     "Given an empty list of numbers should return empty list",
			List:     []int{},
			Number:   1,
			Expected: []int{},
		},
	}

	for _, tt := range adderTest {
		t.Run(tt.Name, func(t *testing.T) {
			list := tt.List
			result := RotateAList(list, tt.Number)
			assert.Equal(t, tt.Expected, result)
		})
	}
}
