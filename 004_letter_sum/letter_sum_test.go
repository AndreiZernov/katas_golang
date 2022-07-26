package _04_letter_sum

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLetterSum(t *testing.T) {
	testLettersTest := []struct {
		Name     string
		Letters  string
		Expected int
	}{
		{Name: "Given a string with no letters should return 0", Letters: "", Expected: 0},
		{Name: "Given a string with one letter should return 1", Letters: "a", Expected: 1},
		{Name: "Given a string with two letters should return 3", Letters: "ab", Expected: 3},
		{Name: "GIven a string with upper case letters should 3", Letters: "AB", Expected: 3},
		{Name: "Given a string microspectrophotometries should return 317", Letters: "microspectrophotometries", Expected: 317},
	}

	for _, tt := range testLettersTest {
		t.Run(tt.Name, func(t *testing.T) {
			result := LetterSum(tt.Letters)
			assert.Equal(t, tt.Expected, result)
		})
	}
}
