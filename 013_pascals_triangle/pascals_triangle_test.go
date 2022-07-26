package pascals_triangle

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPascalTriangle(t *testing.T) {
	t.Run("Given a number 5 should return a pascal triangle with 5 lines", func(t *testing.T) {
		expected := PascalTriangle(5)
		assert.Equal(t, "1 \n1 1 \n1 2 1 \n1 3 3 1 \n1 4 6 4 1 \n", expected)
	})

	t.Run("Given a number 10 should return a pascal triangle with 10 lines", func(t *testing.T) {
		expected := PascalTriangle(10)
		assert.Equal(t, "1 \n1 1 \n1 2 1 \n1 3 3 1 \n1 4 6 4 1 \n1 5 10 10 5 1 \n1 6 15 20 15 6 1 \n1 7 21 35 35 21 7 1 \n1 8 28 56 70 56 28 8 1 \n1 9 36 84 126 126 84 36 9 1 \n", expected)
	})
}
