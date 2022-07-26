package align

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAlign(t *testing.T) {
	t.Run("Given a string, a length and the position L should return the string aligned to the left", func(t *testing.T) {
		result, err := Align("SaltPay", 11, "L")
		assert.NoError(t, err)
		assert.Equal(t, "SaltPay....", result)
	})

	t.Run("Given a string, a length and the position R should return the string aligned to the right", func(t *testing.T) {
		result, err := Align("SaltPay", 11, "R")
		assert.NoError(t, err)
		assert.Equal(t, "....SaltPay", result)
	})

	t.Run("Given a string, a length 11 and the position C should return the string aligned to the center", func(t *testing.T) {
		result, err := Align("SaltPay", 11, "C")
		assert.NoError(t, err)
		assert.Equal(t, "..SaltPay..", result)
	})

	t.Run("Given a string, a length 10 and the position C should return the string aligned to the center", func(t *testing.T) {
		result, err := Align("SaltPay", 10, "C")
		assert.NoError(t, err)
		assert.Equal(t, "..SaltPay.", result)
	})

	t.Run("Given a string with length bigger then the length should return an error", func(t *testing.T) {
		result, err := Align("SaltPay", 0, "L")
		assert.Error(t, err)
		assert.Equal(t, "", result)
	})

	t.Run("Given a wrong position should return an error", func(t *testing.T) {
		result, err := Align("SaltPay", 0, "L")
		assert.Error(t, err)
		assert.Equal(t, "", result)
	})
}
