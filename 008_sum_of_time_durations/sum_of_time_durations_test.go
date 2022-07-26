package sum_of_time_durations

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSumOfTimeDurations(t *testing.T) {
	t.Run("Given a slice with no time durations should return 0", func(t *testing.T) {
		result, err := SumOfTimeDurations([]string{})
		assert.Equal(t, true, err == nil)
		assert.Equal(t, "0:0:0", result)
	})

	t.Run("Given a slice with few durations should return correctly calculated time", func(t *testing.T) {
		result, err := SumOfTimeDurations([]string{"12:32", "34:01", "15:23", "9:27", "55:22", "25:56"})
		assert.Equal(t, true, err == nil)
		assert.Equal(t, "2:32:41", result)
	})

	t.Run("Given a slice with wrong value in time should return calculated time and ignore the wrong value", func(t *testing.T) {
		result, err := SumOfTimeDurations([]string{"12:32", "34:01", "15:23", "9:27", "55:22", "25:56", "wrong"})
		assert.Equal(t, true, err == nil)
		assert.Equal(t, "2:32:41", result)
	})

	t.Run("Given a slice with wrong durations should return error", func(t *testing.T) {
		result, err := SumOfTimeDurations([]string{"12:32", "34:01", "15:23", "9:27", "55:22", "25:56", "wrong:wrong"})
		assert.Equal(t, false, err == nil)
		assert.Equal(t, "0:0:0", result)
	})
}
