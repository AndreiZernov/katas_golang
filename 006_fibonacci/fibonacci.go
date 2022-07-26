package fibonacci

import (
	"errors"
	"math"
)

var (
	OverflowErrorMessage          = errors.New("integer overflow")
	OutsideOfSequenceErrorMessage = errors.New("position outside of fibonacci sequence")
)

func Fibonacci(position int64) (int64, error) {
	switch {
	case position < 0:
		return 0, OutsideOfSequenceErrorMessage
	case position == 0:
		return 0, nil
	case position >= 93:
		return 0, OverflowErrorMessage
	default:
		firstTerm := math.Pow(math.Phi, float64(position))
		secondTerm := math.Pow(math.Phi-1, float64(position))
		result := math.Round((firstTerm + secondTerm) / math.Sqrt(5))

		return int64(result), nil
	}
}
