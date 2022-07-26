package die_roll

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDieRoll(t *testing.T) {
	t.Run("Given a die with 6 sides should return a slice with 6 elements", func(t *testing.T) {
		result := DieRoll(6, 1000000)
		assert.Equal(t, 6, len(result))
	})
	t.Run("Given a die with 100 sides should return a slice with 100 elements", func(t *testing.T) {
		result := DieRoll(100, 1000000)
		assert.Equal(t, 100, len(result))
	})
}
