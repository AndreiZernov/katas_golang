package align

import (
	"errors"
	"math"
	"strings"
)

var InvalidPosition = errors.New("invalid position")
var InvalidLength = errors.New("invalid length")

func Align(word string, length int, position string) (string, error) {
	if len(word) > length {
		return "", InvalidLength
	}

	switch position {
	case "L":
		leftSide := strings.Repeat(".", length-len(word))
		return word + leftSide, nil
	case "R":
		rightSide := strings.Repeat(".", length-len(word))
		return rightSide + word, nil
	case "C":
		leftSideReps := math.Ceil(float64(length-len(word)) / 2)
		leftSide := strings.Repeat(".", int(leftSideReps))

		rightSieReps := length - len(word) - int(leftSideReps)
		rightSide := strings.Repeat(".", rightSieReps)

		return leftSide + word + rightSide, nil
	default:
		return "", InvalidPosition
	}
}
