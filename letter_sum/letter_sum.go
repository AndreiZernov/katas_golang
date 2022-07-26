package letter_sum

import "strings"

func LetterSum(letters string) int {
	var sum int
	for _, letter := range strings.ToLower(letters) {
		sum += int(letter) - 96
	}
	return sum
}
